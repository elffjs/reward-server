package main

import (
	"bytes"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/big"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/elffjs/reward-server/contract"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
	"github.com/gofiber/fiber/v2"
)

type RewardData struct {
	VehicleID *big.Int       `json:"vehicleId"`
	Owner     common.Address `json:"owner"`
	Week      *big.Int       `json:"week"`
	Points    *big.Int       `json:"points"`
}

type Input struct {
	Data      RewardData `json:"data"`
	Signature string     `json:"signature"`
}

var typ = []apitypes.Type{
	{Name: "vehicleId", Type: "uint256"},
	{Name: "owner", Type: "address"},
	{Name: "week", Type: "uint256"},
	{Name: "points", Type: "uint256"},
}

type AuthResp struct {
	State     string `json:"state"`
	Challenge string `json:"challenge"`
}

type ChalResp struct {
	AccessToken string `json:"access_token"`
}

type TEResp struct {
	Token string `json:"token"`
}

type TokenReq struct {
	NFTContractAddress common.Address `json:"nftContractAddress"`
	TokenID            *big.Int       `json:"tokenId"`
	Privileges         []*big.Int     `json:"privileges"`
}

var licenseAddr = common.HexToAddress("0x9bfC50a5fC024f36635CA8D92B0202D5B5b51fb2")

type VCResp struct {
	VCQuery string `json:"vcQuery"`
}

func main() {
	privateKey, err := crypto.HexToECDSA(os.Getenv("PRIVATE_KEY"))
	if err != nil {
		log.Fatal(err)
	}

	publicKeyECDSA := privateKey.Public().(*ecdsa.PublicKey)

	ourAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	fmt.Printf("Our address: %s\n", ourAddress)

	vehicleAddr := common.HexToAddress("0x45fbCD3ef7361d156e8b16F5538AE36DEdf61Da8")

	client, err := ethclient.Dial("https://rpc-amoy.polygon.technology/")
	if err != nil {
		log.Fatal(err)
	}

	sacd, err := contract.NewSACD(common.HexToAddress("0x4E5F9320b1c7cB3DE5ebDD760aD67375B66cF8a3"), client)
	if err != nil {
		log.Fatal(err)
	}

	vehicle, err := contract.NewVehicle(vehicleAddr, client)
	if err != nil {
		log.Fatal(err)
	}

	registry, err := contract.NewRegistry(common.HexToAddress("0x5eAA326fB2fc97fAcCe6A79A304876daD0F2e96c"), client)
	if err != nil {
		log.Fatal(err)
	}

	app := fiber.New()
	app.Post("/", func(c *fiber.Ctx) error {
		u, err := url.ParseRequestURI("https://auth.dev.dimo.zone/auth/web3/generate_challenge")
		if err != nil {
			return err
		}

		v := url.Values{}
		v.Set("client_id", "dimo-frontend")
		v.Set("domain", "https://app.dev.dimo.zone/auth/callback")
		v.Set("scope", "openid email")
		v.Set("response_type", "code")
		v.Set("address", licenseAddr.Hex())

		u.RawQuery = v.Encode()

		resp, err := http.Post(u.String(), "", nil)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		b, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		if resp.StatusCode != 200 {
			return fmt.Errorf("status code %d: %s", resp.StatusCode, string(b))
		}

		var ar AuthResp
		if err := json.Unmarshal(b, &ar); err != nil {
			return err
		}

		fmt.Println("Challenge: ", ar.Challenge)

		arH := accounts.TextHash([]byte(ar.Challenge))

		sig, err := crypto.Sign(arH, privateKey)
		if err != nil {
			return err
		}

		sig[crypto.RecoveryIDOffset] += 27

		v = url.Values{}
		v.Set("client_id", "dimo-frontend")
		v.Set("domain", "https://app.dev.dimo.zone/auth/callback")
		v.Set("grant_type", "authorization_code")
		v.Set("state", ar.State)
		v.Set("signature", hexutil.Encode(sig))

		resp, err = http.Post("https://auth.dev.dimo.zone/auth/web3/submit_challenge", "application/x-www-form-urlencoded", strings.NewReader(v.Encode()))
		if err != nil {
			return err
		}

		defer resp.Body.Close()

		b, err = io.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		if resp.StatusCode != 200 {
			return fmt.Errorf("status code %d: %s", resp.StatusCode, string(b))
		}

		var cr ChalResp
		if err := json.Unmarshal(b, &cr); err != nil {
			return err
		}

		fmt.Println("Access token", cr.AccessToken)

		var body Input
		if err := c.BodyParser(&body); err != nil {
			return err
		}

		if exists, err := vehicle.Exists(nil, body.Data.VehicleID); err != nil {
			return err
		} else if !exists {
			return fiber.NewError(fiber.StatusBadRequest, fmt.Sprintf("Vehicle %d does not exist.", body.Data.VehicleID))
		}

		if owner, err := vehicle.OwnerOf(nil, body.Data.VehicleID); err != nil {
			return err
		} else if owner != body.Data.Owner {
			return fiber.NewError(fiber.StatusBadRequest, fmt.Sprintf("Vehicle %d owned by %s.", body.Data.VehicleID, owner))
		}

		if has, err := sacd.HasPermission(nil, vehicleAddr, body.Data.VehicleID, ourAddress, 5); err != nil {
			return err
		} else if !has {
			return fiber.NewError(fiber.StatusBadRequest, fmt.Sprintf("Please grant permission 5 (VIN credential) to %s.", ourAddress))
		}

		tr := TokenReq{
			NFTContractAddress: vehicleAddr,
			TokenID:            body.Data.VehicleID,
			Privileges:         []*big.Int{big.NewInt(5)},
		}

		trRaw, err := json.Marshal(tr)
		if err != nil {
			return err
		}

		req, err := http.NewRequestWithContext(c.Context(), "POST", "https://token-exchange-api.dev.dimo.zone/v1/tokens/exchange", bytes.NewReader(trRaw))
		if err != nil {
			return err
		}

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer "+cr.AccessToken)

		resp, err = http.DefaultClient.Do(req)
		if err != nil {
			return err
		}

		defer resp.Body.Close()

		b, err = io.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		if resp.StatusCode != 200 {
			return fmt.Errorf("exchange failure: status code %d: %s", resp.StatusCode, string(b))
		}

		fmt.Println("XDDICANT", string(b))

		var teResp TEResp

		if err := json.Unmarshal(b, &teResp); err != nil {
			return err
		}

		req, err = http.NewRequestWithContext(c.Context(), "POST", "https://attestation-api.dev.dimo.zone/v1/vc/vin/"+body.Data.VehicleID.String(), nil)
		if err != nil {
			return err
		}

		req.Header.Set("Authorization", "Bearer "+teResp.Token)

		resp, err = http.DefaultClient.Do(req)
		if err != nil {
			return fmt.Errorf("WHAT %w", err)
		}

		defer resp.Body.Close()

		b, err = io.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		if resp.StatusCode != 200 {
			return fmt.Errorf("exchange failure: status code %d: %s", resp.StatusCode, string(b))
		}

		var vcRes VCResp
		if err := json.Unmarshal(b, &vcRes); err != nil {
			return err
		}

		type QueryReq struct {
			Query string `json:"query"`
		}

		b, err = json.Marshal(QueryReq{Query: vcRes.VCQuery})
		if err != nil {
			return err
		}

		req, err = http.NewRequestWithContext(c.Context(), "POST", "https://telemetry-api.dev.dimo.zone/query", bytes.NewBuffer(b))
		if err != nil {
			return err
		}

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer "+teResp.Token)

		resp, err = http.DefaultClient.Do(req)
		if err != nil {
			return err
		}

		defer resp.Body.Close()

		b, err = io.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		if resp.StatusCode != 200 {
			return fmt.Errorf("status code %d: %s", resp.StatusCode, string(b))
		}

		var vtr VINTelResp

		if err := json.Unmarshal(b, &vtr); err != nil {
			return err
		}

		rawVC := vtr.Data.VINVCLatest.RawVC

		var vc VINVC
		if err := json.Unmarshal([]byte(rawVC), &vc); err != nil {
			return err
		}

		dr := DecodeReq{
			VIN: vc.CredentialSubject.VehicleIdentificationNumber,
		}

		b, err = json.Marshal(dr)
		if err != nil {
			return err
		}

		req, err = http.NewRequestWithContext(c.Context(), "POST", "https://device-definitions-api.dev.dimo.zone/device-definitions/decode-vin", bytes.NewReader(b))
		if err != nil {
			return err
		}

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer "+cr.AccessToken)

		resp, err = http.DefaultClient.Do(req)
		if err != nil {
			return err
		}

		defer resp.Body.Close()

		b, err = io.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		if resp.StatusCode != 200 {
			return fmt.Errorf("status code %d: %s", resp.StatusCode, string(b))
		}

		var decodeResp DecodeResp
		if err := json.Unmarshal(b, &decodeResp); err != nil {
			return err
		}

		dd, err := registry.GetDeviceDefinitionIdByVehicleId(nil, body.Data.VehicleID)
		if err != nil {
			return err
		}

		if decodeResp.DeviceDefinitionID != dd {
			return fiber.NewError(fiber.StatusBadRequest, fmt.Sprintf("Vehicle VIN decoded to %s, but token has %s", decodeResp.DeviceDefinitionID, dd))
		}

		td := apitypes.TypedData{
			Types: apitypes.Types{
				"EIP712Domain": []apitypes.Type{
					{Name: "name", Type: "string"},
					{Name: "version", Type: "string"},
					{Name: "chainId", Type: "uint256"},
					{Name: "verifyingContract", Type: "address"},
				},
				"BaselineRequest": []apitypes.Type{
					{Name: "vehicleId", Type: "uint256"},
					{Name: "owner", Type: "address"},
					{Name: "week", Type: "uint256"},
					{Name: "points", Type: "uint256"},
				},
			},
			PrimaryType: "BaselineRequest",
			Domain: apitypes.TypedDataDomain{
				Name:              "Baseline",
				Version:           "1",
				ChainId:           math.NewHexOrDecimal256(80002),
				VerifyingContract: ourAddress.Hex(),
			},
			Message: apitypes.TypedDataMessage{
				"vehicleId": body.Data.VehicleID,
				"owner":     body.Data.Owner.Hex(),
				"week":      body.Data.Week,
				"points":    body.Data.Points,
			},
		}

		b, _, err = apitypes.TypedDataAndHash(td)
		if err != nil {
			return fmt.Errorf("couldn't hash typed data: %w", err)
		}

		sig, err = crypto.Sign(b, privateKey)
		if err != nil {
			return err
		}

		sig[crypto.RecoveryIDOffset] += 27

		uint256Type, _ := abi.NewType("uint256", "", nil)
		addrType, _ := abi.NewType("address", "", nil)
		bytesType, _ := abi.NewType("bytes", "", nil)

		args := abi.Arguments{
			{Type: uint256Type},
			{Type: addrType},
			{Type: uint256Type},
			{Type: uint256Type},
			{Type: bytesType},
		}

		out, err := args.Pack(body.Data.VehicleID, body.Data.Owner, body.Data.Week, body.Data.Points, sig)
		if err != nil {
			return fmt.Errorf("packing error: %w", err)
		}

		fmt.Printf("%s\n", hexutil.Encode(out))

		return c.JSON(Response{
			Signature: hexutil.Encode(sig),
		})
	})

	fmt.Println("WHAT")

	if err := app.Listen(":3001"); err != nil {
		log.Fatal(err)
	}
}

type Response struct {
	Signature string `json:"signature"`
}

type VINVC struct {
	CredentialSubject struct {
		VehicleIdentificationNumber string `json:"vehicleIdentificationNumber"`
	} `json:"credentialSubject"`
}

type VINTelResp struct {
	Data struct {
		VINVCLatest struct {
			RawVC string `json:"rawVC"`
		} `json:"vinVCLatest"`
	} `json:"data"`
}

type DecodeReq struct {
	VIN string `json:"vin"`
}

type DecodeResp struct {
	DeviceDefinitionID string `json:"deviceDefinitionId"`
}
