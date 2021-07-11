// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// BribeABI is the input ABI used to generate the binding from.
const BribeABI = "[{\"inputs\":[],\"name\":\"bribe\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]"

// BribeBin is the compiled bytecode used for deploying new contracts.
var BribeBin = "0x6080604052348015600f57600080fd5b50608d8061001e6000396000f3fe608060405260043610601c5760003560e01c806337d0208c146021575b600080fd5b60276029565b005b60405141903480156108fc02916000818181858888f193505050501580156054573d6000803e3d6000fd5b5056fea26469706673582212202e04417bb0be216ba8f4d64d5668ea24b458b0e474814d008dd7819e28c1288364736f6c63430008040033"

// DeployBribe deploys a new Ethereum contract, binding an instance of Bribe to it.
func DeployBribe(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Bribe, error) {
	parsed, err := abi.JSON(strings.NewReader(BribeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BribeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Bribe{BribeCaller: BribeCaller{contract: contract}, BribeTransactor: BribeTransactor{contract: contract}, BribeFilterer: BribeFilterer{contract: contract}}, nil
}

// Bribe is an auto generated Go binding around an Ethereum contract.
type Bribe struct {
	BribeCaller     // Read-only binding to the contract
	BribeTransactor // Write-only binding to the contract
	BribeFilterer   // Log filterer for contract events
}

// BribeCaller is an auto generated read-only Go binding around an Ethereum contract.
type BribeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BribeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BribeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BribeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BribeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BribeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BribeSession struct {
	Contract     *Bribe            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BribeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BribeCallerSession struct {
	Contract *BribeCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BribeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BribeTransactorSession struct {
	Contract     *BribeTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BribeRaw is an auto generated low-level Go binding around an Ethereum contract.
type BribeRaw struct {
	Contract *Bribe // Generic contract binding to access the raw methods on
}

// BribeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BribeCallerRaw struct {
	Contract *BribeCaller // Generic read-only contract binding to access the raw methods on
}

// BribeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BribeTransactorRaw struct {
	Contract *BribeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBribe creates a new instance of Bribe, bound to a specific deployed contract.
func NewBribe(address common.Address, backend bind.ContractBackend) (*Bribe, error) {
	contract, err := bindBribe(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Bribe{BribeCaller: BribeCaller{contract: contract}, BribeTransactor: BribeTransactor{contract: contract}, BribeFilterer: BribeFilterer{contract: contract}}, nil
}

// NewBribeCaller creates a new read-only instance of Bribe, bound to a specific deployed contract.
func NewBribeCaller(address common.Address, caller bind.ContractCaller) (*BribeCaller, error) {
	contract, err := bindBribe(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BribeCaller{contract: contract}, nil
}

// NewBribeTransactor creates a new write-only instance of Bribe, bound to a specific deployed contract.
func NewBribeTransactor(address common.Address, transactor bind.ContractTransactor) (*BribeTransactor, error) {
	contract, err := bindBribe(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BribeTransactor{contract: contract}, nil
}

// NewBribeFilterer creates a new log filterer instance of Bribe, bound to a specific deployed contract.
func NewBribeFilterer(address common.Address, filterer bind.ContractFilterer) (*BribeFilterer, error) {
	contract, err := bindBribe(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BribeFilterer{contract: contract}, nil
}

// bindBribe binds a generic wrapper to an already deployed contract.
func bindBribe(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BribeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bribe *BribeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bribe.Contract.BribeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bribe *BribeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bribe.Contract.BribeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bribe *BribeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bribe.Contract.BribeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bribe *BribeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bribe.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bribe *BribeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bribe.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bribe *BribeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bribe.Contract.contract.Transact(opts, method, params...)
}

// Bribe is a paid mutator transaction binding the contract method 0x37d0208c.
//
// Solidity: function bribe() payable returns()
func (_Bribe *BribeTransactor) Bribe(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bribe.contract.Transact(opts, "bribe")
}

// Bribe is a paid mutator transaction binding the contract method 0x37d0208c.
//
// Solidity: function bribe() payable returns()
func (_Bribe *BribeSession) Bribe() (*types.Transaction, error) {
	return _Bribe.Contract.Bribe(&_Bribe.TransactOpts)
}

// Bribe is a paid mutator transaction binding the contract method 0x37d0208c.
//
// Solidity: function bribe() payable returns()
func (_Bribe *BribeTransactorSession) Bribe() (*types.Transaction, error) {
	return _Bribe.Contract.Bribe(&_Bribe.TransactOpts)
}
