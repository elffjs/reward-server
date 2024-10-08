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

// SACDMetaData contains all meta data concerning the SACD contract.
var SACDMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"grantee\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"permissionIndex\",\"type\":\"uint8\"}],\"name\":\"hasPermission\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// SACDABI is the input ABI used to generate the binding from.
// Deprecated: Use SACDMetaData.ABI instead.
var SACDABI = SACDMetaData.ABI

// SACD is an auto generated Go binding around an Ethereum contract.
type SACD struct {
	SACDCaller     // Read-only binding to the contract
	SACDTransactor // Write-only binding to the contract
	SACDFilterer   // Log filterer for contract events
}

// SACDCaller is an auto generated read-only Go binding around an Ethereum contract.
type SACDCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SACDTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SACDTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SACDFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SACDFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SACDSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SACDSession struct {
	Contract     *SACD             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SACDCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SACDCallerSession struct {
	Contract *SACDCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// SACDTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SACDTransactorSession struct {
	Contract     *SACDTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SACDRaw is an auto generated low-level Go binding around an Ethereum contract.
type SACDRaw struct {
	Contract *SACD // Generic contract binding to access the raw methods on
}

// SACDCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SACDCallerRaw struct {
	Contract *SACDCaller // Generic read-only contract binding to access the raw methods on
}

// SACDTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SACDTransactorRaw struct {
	Contract *SACDTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSACD creates a new instance of SACD, bound to a specific deployed contract.
func NewSACD(address common.Address, backend bind.ContractBackend) (*SACD, error) {
	contract, err := bindSACD(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SACD{SACDCaller: SACDCaller{contract: contract}, SACDTransactor: SACDTransactor{contract: contract}, SACDFilterer: SACDFilterer{contract: contract}}, nil
}

// NewSACDCaller creates a new read-only instance of SACD, bound to a specific deployed contract.
func NewSACDCaller(address common.Address, caller bind.ContractCaller) (*SACDCaller, error) {
	contract, err := bindSACD(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SACDCaller{contract: contract}, nil
}

// NewSACDTransactor creates a new write-only instance of SACD, bound to a specific deployed contract.
func NewSACDTransactor(address common.Address, transactor bind.ContractTransactor) (*SACDTransactor, error) {
	contract, err := bindSACD(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SACDTransactor{contract: contract}, nil
}

// NewSACDFilterer creates a new log filterer instance of SACD, bound to a specific deployed contract.
func NewSACDFilterer(address common.Address, filterer bind.ContractFilterer) (*SACDFilterer, error) {
	contract, err := bindSACD(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SACDFilterer{contract: contract}, nil
}

// bindSACD binds a generic wrapper to an already deployed contract.
func bindSACD(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SACDMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SACD *SACDRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SACD.Contract.SACDCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SACD *SACDRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SACD.Contract.SACDTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SACD *SACDRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SACD.Contract.SACDTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SACD *SACDCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SACD.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SACD *SACDTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SACD.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SACD *SACDTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SACD.Contract.contract.Transact(opts, method, params...)
}

// HasPermission is a free data retrieval call binding the contract method 0x48eb48f5.
//
// Solidity: function hasPermission(address asset, uint256 tokenId, address grantee, uint8 permissionIndex) view returns(bool)
func (_SACD *SACDCaller) HasPermission(opts *bind.CallOpts, asset common.Address, tokenId *big.Int, grantee common.Address, permissionIndex uint8) (bool, error) {
	var out []interface{}
	err := _SACD.contract.Call(opts, &out, "hasPermission", asset, tokenId, grantee, permissionIndex)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasPermission is a free data retrieval call binding the contract method 0x48eb48f5.
//
// Solidity: function hasPermission(address asset, uint256 tokenId, address grantee, uint8 permissionIndex) view returns(bool)
func (_SACD *SACDSession) HasPermission(asset common.Address, tokenId *big.Int, grantee common.Address, permissionIndex uint8) (bool, error) {
	return _SACD.Contract.HasPermission(&_SACD.CallOpts, asset, tokenId, grantee, permissionIndex)
}

// HasPermission is a free data retrieval call binding the contract method 0x48eb48f5.
//
// Solidity: function hasPermission(address asset, uint256 tokenId, address grantee, uint8 permissionIndex) view returns(bool)
func (_SACD *SACDCallerSession) HasPermission(asset common.Address, tokenId *big.Int, grantee common.Address, permissionIndex uint8) (bool, error) {
	return _SACD.Contract.HasPermission(&_SACD.CallOpts, asset, tokenId, grantee, permissionIndex)
}
