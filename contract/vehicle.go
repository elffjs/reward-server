// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// VehicleMetaData contains all meta data concerning the Vehicle contract.
var VehicleMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"exists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// VehicleABI is the input ABI used to generate the binding from.
// Deprecated: Use VehicleMetaData.ABI instead.
var VehicleABI = VehicleMetaData.ABI

// Vehicle is an auto generated Go binding around an Ethereum contract.
type Vehicle struct {
	VehicleCaller     // Read-only binding to the contract
	VehicleTransactor // Write-only binding to the contract
	VehicleFilterer   // Log filterer for contract events
}

// VehicleCaller is an auto generated read-only Go binding around an Ethereum contract.
type VehicleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VehicleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VehicleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VehicleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VehicleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VehicleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VehicleSession struct {
	Contract     *Vehicle          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VehicleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VehicleCallerSession struct {
	Contract *VehicleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// VehicleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VehicleTransactorSession struct {
	Contract     *VehicleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// VehicleRaw is an auto generated low-level Go binding around an Ethereum contract.
type VehicleRaw struct {
	Contract *Vehicle // Generic contract binding to access the raw methods on
}

// VehicleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VehicleCallerRaw struct {
	Contract *VehicleCaller // Generic read-only contract binding to access the raw methods on
}

// VehicleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VehicleTransactorRaw struct {
	Contract *VehicleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVehicle creates a new instance of Vehicle, bound to a specific deployed contract.
func NewVehicle(address common.Address, backend bind.ContractBackend) (*Vehicle, error) {
	contract, err := bindVehicle(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Vehicle{VehicleCaller: VehicleCaller{contract: contract}, VehicleTransactor: VehicleTransactor{contract: contract}, VehicleFilterer: VehicleFilterer{contract: contract}}, nil
}

// NewVehicleCaller creates a new read-only instance of Vehicle, bound to a specific deployed contract.
func NewVehicleCaller(address common.Address, caller bind.ContractCaller) (*VehicleCaller, error) {
	contract, err := bindVehicle(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VehicleCaller{contract: contract}, nil
}

// NewVehicleTransactor creates a new write-only instance of Vehicle, bound to a specific deployed contract.
func NewVehicleTransactor(address common.Address, transactor bind.ContractTransactor) (*VehicleTransactor, error) {
	contract, err := bindVehicle(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VehicleTransactor{contract: contract}, nil
}

// NewVehicleFilterer creates a new log filterer instance of Vehicle, bound to a specific deployed contract.
func NewVehicleFilterer(address common.Address, filterer bind.ContractFilterer) (*VehicleFilterer, error) {
	contract, err := bindVehicle(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VehicleFilterer{contract: contract}, nil
}

// bindVehicle binds a generic wrapper to an already deployed contract.
func bindVehicle(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VehicleMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Vehicle *VehicleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Vehicle.Contract.VehicleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Vehicle *VehicleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vehicle.Contract.VehicleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Vehicle *VehicleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Vehicle.Contract.VehicleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Vehicle *VehicleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Vehicle.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Vehicle *VehicleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vehicle.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Vehicle *VehicleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Vehicle.Contract.contract.Transact(opts, method, params...)
}

// Exists is a free data retrieval call binding the contract method 0x4f558e79.
//
// Solidity: function exists(uint256 tokenId) view returns(bool)
func (_Vehicle *VehicleCaller) Exists(opts *bind.CallOpts, tokenId *big.Int) (bool, error) {
	var out []interface{}
	err := _Vehicle.contract.Call(opts, &out, "exists", tokenId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Exists is a free data retrieval call binding the contract method 0x4f558e79.
//
// Solidity: function exists(uint256 tokenId) view returns(bool)
func (_Vehicle *VehicleSession) Exists(tokenId *big.Int) (bool, error) {
	return _Vehicle.Contract.Exists(&_Vehicle.CallOpts, tokenId)
}

// Exists is a free data retrieval call binding the contract method 0x4f558e79.
//
// Solidity: function exists(uint256 tokenId) view returns(bool)
func (_Vehicle *VehicleCallerSession) Exists(tokenId *big.Int) (bool, error) {
	return _Vehicle.Contract.Exists(&_Vehicle.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Vehicle *VehicleCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Vehicle.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Vehicle *VehicleSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Vehicle.Contract.OwnerOf(&_Vehicle.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Vehicle *VehicleCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Vehicle.Contract.OwnerOf(&_Vehicle.CallOpts, tokenId)
}
