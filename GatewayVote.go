// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ethwatcher

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

// GatewayVoteABI is the input ABI used to generate the binding from.
const GatewayVoteABI = "[{\"stateMutability\":\"view\",\"payable\":false,\"name\":\"hasConfirmed\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"type\":\"function\",\"inputs\":[{\"name\":\"operation\",\"type\":\"uint256\"},{\"name\":\"voter\",\"type\":\"address\"}],\"constant\":true},{\"stateMutability\":\"view\",\"payable\":false,\"name\":\"isChainCode\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"type\":\"function\",\"inputs\":[{\"name\":\"code\",\"type\":\"uint32\"}],\"constant\":true},{\"stateMutability\":\"nonpayable\",\"payable\":false,\"name\":\"revoke\",\"outputs\":[],\"type\":\"function\",\"inputs\":[{\"name\":\"operation\",\"type\":\"uint256\"}],\"constant\":false},{\"stateMutability\":\"nonpayable\",\"inputs\":[{\"type\":\"uint64\",\"name\":\"wad\"},{\"type\":\"string\",\"name\":\"dstDescribe\"},{\"type\":\"uint64\",\"name\":\"fee\"}],\"name\":\"burn\",\"outputs\":[],\"type\":\"function\",\"constant\":false,\"payable\":false},{\"stateMutability\":\"nonpayable\",\"payable\":false,\"name\":\"start\",\"outputs\":[],\"type\":\"function\",\"inputs\":[{\"name\":\"proposal\",\"type\":\"string\"}],\"constant\":false},{\"stateMutability\":\"nonpayable\",\"payable\":false,\"name\":\"addApp\",\"outputs\":[],\"type\":\"function\",\"inputs\":[{\"name\":\"app\",\"type\":\"address\"},{\"name\":\"chain\",\"type\":\"uint32\"},{\"name\":\"token\",\"type\":\"uint32\"},{\"name\":\"proposal\",\"type\":\"string\"}],\"constant\":false},{\"stateMutability\":\"view\",\"payable\":false,\"name\":\"mMaxAppCode\",\"outputs\":[{\"name\":\"\",\"type\":\"uint32\"}],\"type\":\"function\",\"inputs\":[],\"constant\":true},{\"stateMutability\":\"view\",\"payable\":false,\"name\":\"getChainName\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"type\":\"function\",\"inputs\":[{\"name\":\"code\",\"type\":\"uint32\"}],\"constant\":true},{\"stateMutability\":\"view\",\"payable\":false,\"name\":\"getAppCode\",\"outputs\":[{\"name\":\"\",\"type\":\"uint32\"}],\"type\":\"function\",\"inputs\":[{\"name\":\"app\",\"type\":\"address\"}],\"constant\":true},{\"stateMutability\":\"view\",\"payable\":false,\"name\":\"mMaxChainCode\",\"outputs\":[{\"name\":\"\",\"type\":\"uint32\"}],\"type\":\"function\",\"inputs\":[],\"constant\":true},{\"stateMutability\":\"nonpayable\",\"payable\":false,\"name\":\"changeGatewayAddr\",\"outputs\":[],\"type\":\"function\",\"inputs\":[{\"name\":\"appCode\",\"type\":\"uint32\"},{\"name\":\"newer\",\"type\":\"address\"},{\"name\":\"proposal\",\"type\":\"string\"}],\"constant\":false},{\"stateMutability\":\"nonpayable\",\"payable\":false,\"name\":\"stop\",\"outputs\":[],\"type\":\"function\",\"inputs\":[{\"name\":\"proposal\",\"type\":\"string\"}],\"constant\":false},{\"stateMutability\":\"view\",\"payable\":false,\"name\":\"getAppAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"type\":\"function\",\"inputs\":[{\"name\":\"code\",\"type\":\"uint32\"}],\"constant\":true},{\"stateMutability\":\"view\",\"payable\":false,\"name\":\"isAppCode\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"type\":\"function\",\"inputs\":[{\"name\":\"code\",\"type\":\"uint32\"}],\"constant\":true},{\"stateMutability\":\"view\",\"payable\":false,\"name\":\"isCaller\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"type\":\"function\",\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"constant\":true},{\"stateMutability\":\"view\",\"payable\":false,\"name\":\"getAppTokenCode\",\"outputs\":[{\"name\":\"\",\"type\":\"uint32\"}],\"type\":\"function\",\"inputs\":[{\"name\":\"code\",\"type\":\"uint32\"}],\"constant\":true},{\"stateMutability\":\"nonpayable\",\"payable\":false,\"name\":\"removeApp\",\"outputs\":[],\"type\":\"function\",\"inputs\":[{\"name\":\"code\",\"type\":\"uint32\"},{\"name\":\"proposal\",\"type\":\"string\"}],\"constant\":false},{\"stateMutability\":\"view\",\"payable\":false,\"name\":\"getAppChainCode\",\"outputs\":[{\"name\":\"\",\"type\":\"uint32\"}],\"type\":\"function\",\"inputs\":[{\"name\":\"code\",\"type\":\"uint32\"}],\"constant\":true},{\"stateMutability\":\"view\",\"payable\":false,\"name\":\"getAppInfo\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint32\"},{\"name\":\"\",\"type\":\"uint32\"}],\"type\":\"function\",\"inputs\":[{\"name\":\"code\",\"type\":\"uint32\"}],\"constant\":true},{\"stateMutability\":\"view\",\"payable\":false,\"name\":\"isVoter\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"type\":\"function\",\"inputs\":[{\"name\":\"voter\",\"type\":\"address\"}],\"constant\":true},{\"stateMutability\":\"nonpayable\",\"payable\":false,\"name\":\"addChain\",\"outputs\":[],\"type\":\"function\",\"inputs\":[{\"name\":\"chain\",\"type\":\"string\"},{\"name\":\"proposal\",\"type\":\"string\"}],\"constant\":false},{\"stateMutability\":\"view\",\"payable\":false,\"name\":\"mStopped\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"type\":\"function\",\"inputs\":[],\"constant\":true},{\"stateMutability\":\"nonpayable\",\"payable\":false,\"name\":\"mintByGateway\",\"outputs\":[],\"type\":\"function\",\"inputs\":[{\"name\":\"appCode\",\"type\":\"uint32\"},{\"name\":\"wad\",\"type\":\"uint64\"},{\"name\":\"receiver\",\"type\":\"address\"},{\"name\":\"proposal\",\"type\":\"string\"}],\"constant\":false},{\"stateMutability\":\"view\",\"payable\":false,\"name\":\"isApper\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"type\":\"function\",\"inputs\":[{\"name\":\"app\",\"type\":\"address\"}],\"constant\":true},{\"stateMutability\":\"view\",\"payable\":false,\"name\":\"getChainCode\",\"outputs\":[{\"name\":\"\",\"type\":\"uint32\"}],\"type\":\"function\",\"inputs\":[{\"name\":\"chain\",\"type\":\"string\"}],\"constant\":true},{\"stateMutability\":\"view\",\"payable\":false,\"name\":\"mNumVoters\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"type\":\"function\",\"inputs\":[],\"constant\":true},{\"stateMutability\":\"nonpayable\",\"payable\":false,\"name\":\"changeVoter\",\"outputs\":[],\"type\":\"function\",\"inputs\":[{\"name\":\"older\",\"type\":\"address\"},{\"name\":\"newer\",\"type\":\"address\"},{\"name\":\"proposal\",\"type\":\"string\"}],\"constant\":false},{\"stateMutability\":\"nonpayable\",\"payable\":false,\"name\":\"removeVoter\",\"outputs\":[],\"type\":\"function\",\"inputs\":[{\"name\":\"older\",\"type\":\"address\"},{\"name\":\"proposal\",\"type\":\"string\"}],\"constant\":false},{\"stateMutability\":\"nonpayable\",\"payable\":false,\"name\":\"addVoter\",\"outputs\":[],\"type\":\"function\",\"inputs\":[{\"name\":\"newer\",\"type\":\"address\"},{\"name\":\"proposal\",\"type\":\"string\"}],\"constant\":false},{\"stateMutability\":\"nonpayable\",\"payable\":false,\"name\":\"burnForGateway\",\"outputs\":[],\"type\":\"function\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"dstDescribe\",\"type\":\"string\"},{\"name\":\"wad\",\"type\":\"uint64\"},{\"name\":\"fee\",\"type\":\"uint64\"}],\"constant\":false},{\"stateMutability\":\"view\",\"payable\":false,\"name\":\"isChain\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"type\":\"function\",\"inputs\":[{\"name\":\"chain\",\"type\":\"string\"}],\"constant\":true},{\"stateMutability\":\"nonpayable\",\"type\":\"constructor\",\"inputs\":[{\"name\":\"voters\",\"type\":\"address[]\"}],\"payable\":false},{\"anonymous\":false,\"name\":\"Stopped\",\"type\":\"event\",\"inputs\":[{\"indexed\":true,\"name\":\"operation\",\"type\":\"uint256\"}]},{\"anonymous\":false,\"name\":\"Started\",\"type\":\"event\",\"inputs\":[{\"indexed\":true,\"name\":\"operation\",\"type\":\"uint256\"}]},{\"anonymous\":false,\"name\":\"Confirmation\",\"type\":\"event\",\"inputs\":[{\"indexed\":false,\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"operation\",\"type\":\"uint256\"}]},{\"anonymous\":false,\"name\":\"OperationDone\",\"type\":\"event\",\"inputs\":[{\"indexed\":false,\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"operation\",\"type\":\"uint256\"}]},{\"anonymous\":false,\"name\":\"Revoke\",\"type\":\"event\",\"inputs\":[{\"indexed\":false,\"name\":\"revoker\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"operation\",\"type\":\"uint256\"}]},{\"anonymous\":false,\"name\":\"VoterChanged\",\"type\":\"event\",\"inputs\":[{\"indexed\":false,\"name\":\"oldVoter\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"newVoter\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"operation\",\"type\":\"uint256\"}]},{\"anonymous\":false,\"name\":\"VoterAdded\",\"type\":\"event\",\"inputs\":[{\"indexed\":false,\"name\":\"newVoter\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"operation\",\"type\":\"uint256\"}]},{\"anonymous\":false,\"name\":\"VoterRemoved\",\"type\":\"event\",\"inputs\":[{\"indexed\":false,\"name\":\"oldVoter\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"operation\",\"type\":\"uint256\"}]},{\"anonymous\":false,\"name\":\"ChainAdded\",\"type\":\"event\",\"inputs\":[{\"indexed\":false,\"name\":\"chain\",\"type\":\"string\"},{\"indexed\":true,\"name\":\"operation\",\"type\":\"uint256\"}]},{\"anonymous\":false,\"name\":\"AppAdded\",\"type\":\"event\",\"inputs\":[{\"indexed\":false,\"name\":\"app\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"chain\",\"type\":\"uint32\"},{\"indexed\":false,\"name\":\"token\",\"type\":\"uint32\"},{\"indexed\":true,\"name\":\"operation\",\"type\":\"uint256\"}]},{\"anonymous\":false,\"name\":\"AppRemoved\",\"type\":\"event\",\"inputs\":[{\"indexed\":false,\"name\":\"code\",\"type\":\"uint32\"},{\"indexed\":true,\"name\":\"operation\",\"type\":\"uint256\"}]},{\"anonymous\":false,\"name\":\"MintByGateway\",\"type\":\"event\",\"inputs\":[{\"indexed\":false,\"name\":\"appCode\",\"type\":\"uint32\"},{\"indexed\":false,\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"wad\",\"type\":\"uint64\"},{\"indexed\":true,\"name\":\"operation\",\"type\":\"uint256\"}]},{\"anonymous\":false,\"name\":\"BurnForGateway\",\"type\":\"event\",\"inputs\":[{\"indexed\":false,\"name\":\"appCode\",\"type\":\"uint32\"},{\"indexed\":false,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"dstDescribe\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"wad\",\"type\":\"uint64\"},{\"indexed\":false,\"name\":\"fee\",\"type\":\"uint64\"}]},{\"anonymous\":false,\"name\":\"GatewayAddrChanged\",\"type\":\"event\",\"inputs\":[{\"indexed\":false,\"name\":\"appCode\",\"type\":\"uint32\"},{\"indexed\":false,\"name\":\"newer\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"operation\",\"type\":\"uint256\"}]}]"

// GatewayVote is an auto generated Go binding around an Ethereum contract.
type GatewayVote struct {
	GatewayVoteCaller     // Read-only binding to the contract
	GatewayVoteTransactor // Write-only binding to the contract
	GatewayVoteFilterer   // Log filterer for contract events
}

// GatewayVoteCaller is an auto generated read-only Go binding around an Ethereum contract.
type GatewayVoteCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GatewayVoteTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GatewayVoteTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GatewayVoteFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GatewayVoteFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GatewayVoteSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GatewayVoteSession struct {
	Contract     *GatewayVote      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GatewayVoteCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GatewayVoteCallerSession struct {
	Contract *GatewayVoteCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// GatewayVoteTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GatewayVoteTransactorSession struct {
	Contract     *GatewayVoteTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// GatewayVoteRaw is an auto generated low-level Go binding around an Ethereum contract.
type GatewayVoteRaw struct {
	Contract *GatewayVote // Generic contract binding to access the raw methods on
}

// GatewayVoteCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GatewayVoteCallerRaw struct {
	Contract *GatewayVoteCaller // Generic read-only contract binding to access the raw methods on
}

// GatewayVoteTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GatewayVoteTransactorRaw struct {
	Contract *GatewayVoteTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGatewayVote creates a new instance of GatewayVote, bound to a specific deployed contract.
func NewGatewayVote(address common.Address, backend bind.ContractBackend) (*GatewayVote, error) {
	contract, err := bindGatewayVote(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GatewayVote{GatewayVoteCaller: GatewayVoteCaller{contract: contract}, GatewayVoteTransactor: GatewayVoteTransactor{contract: contract}, GatewayVoteFilterer: GatewayVoteFilterer{contract: contract}}, nil
}

// NewGatewayVoteCaller creates a new read-only instance of GatewayVote, bound to a specific deployed contract.
func NewGatewayVoteCaller(address common.Address, caller bind.ContractCaller) (*GatewayVoteCaller, error) {
	contract, err := bindGatewayVote(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GatewayVoteCaller{contract: contract}, nil
}

// NewGatewayVoteTransactor creates a new write-only instance of GatewayVote, bound to a specific deployed contract.
func NewGatewayVoteTransactor(address common.Address, transactor bind.ContractTransactor) (*GatewayVoteTransactor, error) {
	contract, err := bindGatewayVote(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GatewayVoteTransactor{contract: contract}, nil
}

// NewGatewayVoteFilterer creates a new log filterer instance of GatewayVote, bound to a specific deployed contract.
func NewGatewayVoteFilterer(address common.Address, filterer bind.ContractFilterer) (*GatewayVoteFilterer, error) {
	contract, err := bindGatewayVote(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GatewayVoteFilterer{contract: contract}, nil
}

// bindGatewayVote binds a generic wrapper to an already deployed contract.
func bindGatewayVote(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GatewayVoteABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GatewayVote *GatewayVoteRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _GatewayVote.Contract.GatewayVoteCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GatewayVote *GatewayVoteRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayVote.Contract.GatewayVoteTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GatewayVote *GatewayVoteRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GatewayVote.Contract.GatewayVoteTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GatewayVote *GatewayVoteCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _GatewayVote.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GatewayVote *GatewayVoteTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayVote.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GatewayVote *GatewayVoteTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GatewayVote.Contract.contract.Transact(opts, method, params...)
}

// GetAppAddress is a free data retrieval call binding the contract method 0x7466ee48.
//
// Solidity: function getAppAddress(code uint32) constant returns(address)
func (_GatewayVote *GatewayVoteCaller) GetAppAddress(opts *bind.CallOpts, code uint32) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _GatewayVote.contract.Call(opts, out, "getAppAddress", code)
	return *ret0, err
}

// GetAppAddress is a free data retrieval call binding the contract method 0x7466ee48.
//
// Solidity: function getAppAddress(code uint32) constant returns(address)
func (_GatewayVote *GatewayVoteSession) GetAppAddress(code uint32) (common.Address, error) {
	return _GatewayVote.Contract.GetAppAddress(&_GatewayVote.CallOpts, code)
}

// GetAppAddress is a free data retrieval call binding the contract method 0x7466ee48.
//
// Solidity: function getAppAddress(code uint32) constant returns(address)
func (_GatewayVote *GatewayVoteCallerSession) GetAppAddress(code uint32) (common.Address, error) {
	return _GatewayVote.Contract.GetAppAddress(&_GatewayVote.CallOpts, code)
}

// GetAppChainCode is a free data retrieval call binding the contract method 0xa17ea25b.
//
// Solidity: function getAppChainCode(code uint32) constant returns(uint32)
func (_GatewayVote *GatewayVoteCaller) GetAppChainCode(opts *bind.CallOpts, code uint32) (uint32, error) {
	var (
		ret0 = new(uint32)
	)
	out := ret0
	err := _GatewayVote.contract.Call(opts, out, "getAppChainCode", code)
	return *ret0, err
}

// GetAppChainCode is a free data retrieval call binding the contract method 0xa17ea25b.
//
// Solidity: function getAppChainCode(code uint32) constant returns(uint32)
func (_GatewayVote *GatewayVoteSession) GetAppChainCode(code uint32) (uint32, error) {
	return _GatewayVote.Contract.GetAppChainCode(&_GatewayVote.CallOpts, code)
}

// GetAppChainCode is a free data retrieval call binding the contract method 0xa17ea25b.
//
// Solidity: function getAppChainCode(code uint32) constant returns(uint32)
func (_GatewayVote *GatewayVoteCallerSession) GetAppChainCode(code uint32) (uint32, error) {
	return _GatewayVote.Contract.GetAppChainCode(&_GatewayVote.CallOpts, code)
}

// GetAppCode is a free data retrieval call binding the contract method 0x56bb9b43.
//
// Solidity: function getAppCode(app address) constant returns(uint32)
func (_GatewayVote *GatewayVoteCaller) GetAppCode(opts *bind.CallOpts, app common.Address) (uint32, error) {
	var (
		ret0 = new(uint32)
	)
	out := ret0
	err := _GatewayVote.contract.Call(opts, out, "getAppCode", app)
	return *ret0, err
}

// GetAppCode is a free data retrieval call binding the contract method 0x56bb9b43.
//
// Solidity: function getAppCode(app address) constant returns(uint32)
func (_GatewayVote *GatewayVoteSession) GetAppCode(app common.Address) (uint32, error) {
	return _GatewayVote.Contract.GetAppCode(&_GatewayVote.CallOpts, app)
}

// GetAppCode is a free data retrieval call binding the contract method 0x56bb9b43.
//
// Solidity: function getAppCode(app address) constant returns(uint32)
func (_GatewayVote *GatewayVoteCallerSession) GetAppCode(app common.Address) (uint32, error) {
	return _GatewayVote.Contract.GetAppCode(&_GatewayVote.CallOpts, app)
}

// GetAppInfo is a free data retrieval call binding the contract method 0xa5cae60c.
//
// Solidity: function getAppInfo(code uint32) constant returns(address, uint32, uint32)
func (_GatewayVote *GatewayVoteCaller) GetAppInfo(opts *bind.CallOpts, code uint32) (common.Address, uint32, uint32, error) {
	var (
		ret0 = new(common.Address)
		ret1 = new(uint32)
		ret2 = new(uint32)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
	}
	err := _GatewayVote.contract.Call(opts, out, "getAppInfo", code)
	return *ret0, *ret1, *ret2, err
}

// GetAppInfo is a free data retrieval call binding the contract method 0xa5cae60c.
//
// Solidity: function getAppInfo(code uint32) constant returns(address, uint32, uint32)
func (_GatewayVote *GatewayVoteSession) GetAppInfo(code uint32) (common.Address, uint32, uint32, error) {
	return _GatewayVote.Contract.GetAppInfo(&_GatewayVote.CallOpts, code)
}

// GetAppInfo is a free data retrieval call binding the contract method 0xa5cae60c.
//
// Solidity: function getAppInfo(code uint32) constant returns(address, uint32, uint32)
func (_GatewayVote *GatewayVoteCallerSession) GetAppInfo(code uint32) (common.Address, uint32, uint32, error) {
	return _GatewayVote.Contract.GetAppInfo(&_GatewayVote.CallOpts, code)
}

// GetAppTokenCode is a free data retrieval call binding the contract method 0x8d22b345.
//
// Solidity: function getAppTokenCode(code uint32) constant returns(uint32)
func (_GatewayVote *GatewayVoteCaller) GetAppTokenCode(opts *bind.CallOpts, code uint32) (uint32, error) {
	var (
		ret0 = new(uint32)
	)
	out := ret0
	err := _GatewayVote.contract.Call(opts, out, "getAppTokenCode", code)
	return *ret0, err
}

// GetAppTokenCode is a free data retrieval call binding the contract method 0x8d22b345.
//
// Solidity: function getAppTokenCode(code uint32) constant returns(uint32)
func (_GatewayVote *GatewayVoteSession) GetAppTokenCode(code uint32) (uint32, error) {
	return _GatewayVote.Contract.GetAppTokenCode(&_GatewayVote.CallOpts, code)
}

// GetAppTokenCode is a free data retrieval call binding the contract method 0x8d22b345.
//
// Solidity: function getAppTokenCode(code uint32) constant returns(uint32)
func (_GatewayVote *GatewayVoteCallerSession) GetAppTokenCode(code uint32) (uint32, error) {
	return _GatewayVote.Contract.GetAppTokenCode(&_GatewayVote.CallOpts, code)
}

// GetChainCode is a free data retrieval call binding the contract method 0xccedf3d2.
//
// Solidity: function getChainCode(chain string) constant returns(uint32)
func (_GatewayVote *GatewayVoteCaller) GetChainCode(opts *bind.CallOpts, chain string) (uint32, error) {
	var (
		ret0 = new(uint32)
	)
	out := ret0
	err := _GatewayVote.contract.Call(opts, out, "getChainCode", chain)
	return *ret0, err
}

// GetChainCode is a free data retrieval call binding the contract method 0xccedf3d2.
//
// Solidity: function getChainCode(chain string) constant returns(uint32)
func (_GatewayVote *GatewayVoteSession) GetChainCode(chain string) (uint32, error) {
	return _GatewayVote.Contract.GetChainCode(&_GatewayVote.CallOpts, chain)
}

// GetChainCode is a free data retrieval call binding the contract method 0xccedf3d2.
//
// Solidity: function getChainCode(chain string) constant returns(uint32)
func (_GatewayVote *GatewayVoteCallerSession) GetChainCode(chain string) (uint32, error) {
	return _GatewayVote.Contract.GetChainCode(&_GatewayVote.CallOpts, chain)
}

// GetChainName is a free data retrieval call binding the contract method 0x426ad3f5.
//
// Solidity: function getChainName(code uint32) constant returns(string)
func (_GatewayVote *GatewayVoteCaller) GetChainName(opts *bind.CallOpts, code uint32) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _GatewayVote.contract.Call(opts, out, "getChainName", code)
	return *ret0, err
}

// GetChainName is a free data retrieval call binding the contract method 0x426ad3f5.
//
// Solidity: function getChainName(code uint32) constant returns(string)
func (_GatewayVote *GatewayVoteSession) GetChainName(code uint32) (string, error) {
	return _GatewayVote.Contract.GetChainName(&_GatewayVote.CallOpts, code)
}

// GetChainName is a free data retrieval call binding the contract method 0x426ad3f5.
//
// Solidity: function getChainName(code uint32) constant returns(string)
func (_GatewayVote *GatewayVoteCallerSession) GetChainName(code uint32) (string, error) {
	return _GatewayVote.Contract.GetChainName(&_GatewayVote.CallOpts, code)
}

// HasConfirmed is a free data retrieval call binding the contract method 0x08aff933.
//
// Solidity: function hasConfirmed(operation uint256, voter address) constant returns(bool)
func (_GatewayVote *GatewayVoteCaller) HasConfirmed(opts *bind.CallOpts, operation *big.Int, voter common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _GatewayVote.contract.Call(opts, out, "hasConfirmed", operation, voter)
	return *ret0, err
}

// HasConfirmed is a free data retrieval call binding the contract method 0x08aff933.
//
// Solidity: function hasConfirmed(operation uint256, voter address) constant returns(bool)
func (_GatewayVote *GatewayVoteSession) HasConfirmed(operation *big.Int, voter common.Address) (bool, error) {
	return _GatewayVote.Contract.HasConfirmed(&_GatewayVote.CallOpts, operation, voter)
}

// HasConfirmed is a free data retrieval call binding the contract method 0x08aff933.
//
// Solidity: function hasConfirmed(operation uint256, voter address) constant returns(bool)
func (_GatewayVote *GatewayVoteCallerSession) HasConfirmed(operation *big.Int, voter common.Address) (bool, error) {
	return _GatewayVote.Contract.HasConfirmed(&_GatewayVote.CallOpts, operation, voter)
}

// IsAppCode is a free data retrieval call binding the contract method 0x78a62e9d.
//
// Solidity: function isAppCode(code uint32) constant returns(bool)
func (_GatewayVote *GatewayVoteCaller) IsAppCode(opts *bind.CallOpts, code uint32) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _GatewayVote.contract.Call(opts, out, "isAppCode", code)
	return *ret0, err
}

// IsAppCode is a free data retrieval call binding the contract method 0x78a62e9d.
//
// Solidity: function isAppCode(code uint32) constant returns(bool)
func (_GatewayVote *GatewayVoteSession) IsAppCode(code uint32) (bool, error) {
	return _GatewayVote.Contract.IsAppCode(&_GatewayVote.CallOpts, code)
}

// IsAppCode is a free data retrieval call binding the contract method 0x78a62e9d.
//
// Solidity: function isAppCode(code uint32) constant returns(bool)
func (_GatewayVote *GatewayVoteCallerSession) IsAppCode(code uint32) (bool, error) {
	return _GatewayVote.Contract.IsAppCode(&_GatewayVote.CallOpts, code)
}

// IsApper is a free data retrieval call binding the contract method 0xb9358ba4.
//
// Solidity: function isApper(app address) constant returns(bool)
func (_GatewayVote *GatewayVoteCaller) IsApper(opts *bind.CallOpts, app common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _GatewayVote.contract.Call(opts, out, "isApper", app)
	return *ret0, err
}

// IsApper is a free data retrieval call binding the contract method 0xb9358ba4.
//
// Solidity: function isApper(app address) constant returns(bool)
func (_GatewayVote *GatewayVoteSession) IsApper(app common.Address) (bool, error) {
	return _GatewayVote.Contract.IsApper(&_GatewayVote.CallOpts, app)
}

// IsApper is a free data retrieval call binding the contract method 0xb9358ba4.
//
// Solidity: function isApper(app address) constant returns(bool)
func (_GatewayVote *GatewayVoteCallerSession) IsApper(app common.Address) (bool, error) {
	return _GatewayVote.Contract.IsApper(&_GatewayVote.CallOpts, app)
}

// IsCaller is a free data retrieval call binding the contract method 0x7ac07dcc.
//
// Solidity: function isCaller(addr address) constant returns(bool)
func (_GatewayVote *GatewayVoteCaller) IsCaller(opts *bind.CallOpts, addr common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _GatewayVote.contract.Call(opts, out, "isCaller", addr)
	return *ret0, err
}

// IsCaller is a free data retrieval call binding the contract method 0x7ac07dcc.
//
// Solidity: function isCaller(addr address) constant returns(bool)
func (_GatewayVote *GatewayVoteSession) IsCaller(addr common.Address) (bool, error) {
	return _GatewayVote.Contract.IsCaller(&_GatewayVote.CallOpts, addr)
}

// IsCaller is a free data retrieval call binding the contract method 0x7ac07dcc.
//
// Solidity: function isCaller(addr address) constant returns(bool)
func (_GatewayVote *GatewayVoteCallerSession) IsCaller(addr common.Address) (bool, error) {
	return _GatewayVote.Contract.IsCaller(&_GatewayVote.CallOpts, addr)
}

// IsChain is a free data retrieval call binding the contract method 0xe0f6cc07.
//
// Solidity: function isChain(chain string) constant returns(bool)
func (_GatewayVote *GatewayVoteCaller) IsChain(opts *bind.CallOpts, chain string) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _GatewayVote.contract.Call(opts, out, "isChain", chain)
	return *ret0, err
}

// IsChain is a free data retrieval call binding the contract method 0xe0f6cc07.
//
// Solidity: function isChain(chain string) constant returns(bool)
func (_GatewayVote *GatewayVoteSession) IsChain(chain string) (bool, error) {
	return _GatewayVote.Contract.IsChain(&_GatewayVote.CallOpts, chain)
}

// IsChain is a free data retrieval call binding the contract method 0xe0f6cc07.
//
// Solidity: function isChain(chain string) constant returns(bool)
func (_GatewayVote *GatewayVoteCallerSession) IsChain(chain string) (bool, error) {
	return _GatewayVote.Contract.IsChain(&_GatewayVote.CallOpts, chain)
}

// IsChainCode is a free data retrieval call binding the contract method 0x0995efb5.
//
// Solidity: function isChainCode(code uint32) constant returns(bool)
func (_GatewayVote *GatewayVoteCaller) IsChainCode(opts *bind.CallOpts, code uint32) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _GatewayVote.contract.Call(opts, out, "isChainCode", code)
	return *ret0, err
}

// IsChainCode is a free data retrieval call binding the contract method 0x0995efb5.
//
// Solidity: function isChainCode(code uint32) constant returns(bool)
func (_GatewayVote *GatewayVoteSession) IsChainCode(code uint32) (bool, error) {
	return _GatewayVote.Contract.IsChainCode(&_GatewayVote.CallOpts, code)
}

// IsChainCode is a free data retrieval call binding the contract method 0x0995efb5.
//
// Solidity: function isChainCode(code uint32) constant returns(bool)
func (_GatewayVote *GatewayVoteCallerSession) IsChainCode(code uint32) (bool, error) {
	return _GatewayVote.Contract.IsChainCode(&_GatewayVote.CallOpts, code)
}

// IsVoter is a free data retrieval call binding the contract method 0xa7771ee3.
//
// Solidity: function isVoter(voter address) constant returns(bool)
func (_GatewayVote *GatewayVoteCaller) IsVoter(opts *bind.CallOpts, voter common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _GatewayVote.contract.Call(opts, out, "isVoter", voter)
	return *ret0, err
}

// IsVoter is a free data retrieval call binding the contract method 0xa7771ee3.
//
// Solidity: function isVoter(voter address) constant returns(bool)
func (_GatewayVote *GatewayVoteSession) IsVoter(voter common.Address) (bool, error) {
	return _GatewayVote.Contract.IsVoter(&_GatewayVote.CallOpts, voter)
}

// IsVoter is a free data retrieval call binding the contract method 0xa7771ee3.
//
// Solidity: function isVoter(voter address) constant returns(bool)
func (_GatewayVote *GatewayVoteCallerSession) IsVoter(voter common.Address) (bool, error) {
	return _GatewayVote.Contract.IsVoter(&_GatewayVote.CallOpts, voter)
}

// MMaxAppCode is a free data retrieval call binding the contract method 0x41fb55c6.
//
// Solidity: function mMaxAppCode() constant returns(uint32)
func (_GatewayVote *GatewayVoteCaller) MMaxAppCode(opts *bind.CallOpts) (uint32, error) {
	var (
		ret0 = new(uint32)
	)
	out := ret0
	err := _GatewayVote.contract.Call(opts, out, "mMaxAppCode")
	return *ret0, err
}

// MMaxAppCode is a free data retrieval call binding the contract method 0x41fb55c6.
//
// Solidity: function mMaxAppCode() constant returns(uint32)
func (_GatewayVote *GatewayVoteSession) MMaxAppCode() (uint32, error) {
	return _GatewayVote.Contract.MMaxAppCode(&_GatewayVote.CallOpts)
}

// MMaxAppCode is a free data retrieval call binding the contract method 0x41fb55c6.
//
// Solidity: function mMaxAppCode() constant returns(uint32)
func (_GatewayVote *GatewayVoteCallerSession) MMaxAppCode() (uint32, error) {
	return _GatewayVote.Contract.MMaxAppCode(&_GatewayVote.CallOpts)
}

// MMaxChainCode is a free data retrieval call binding the contract method 0x5c5e274e.
//
// Solidity: function mMaxChainCode() constant returns(uint32)
func (_GatewayVote *GatewayVoteCaller) MMaxChainCode(opts *bind.CallOpts) (uint32, error) {
	var (
		ret0 = new(uint32)
	)
	out := ret0
	err := _GatewayVote.contract.Call(opts, out, "mMaxChainCode")
	return *ret0, err
}

// MMaxChainCode is a free data retrieval call binding the contract method 0x5c5e274e.
//
// Solidity: function mMaxChainCode() constant returns(uint32)
func (_GatewayVote *GatewayVoteSession) MMaxChainCode() (uint32, error) {
	return _GatewayVote.Contract.MMaxChainCode(&_GatewayVote.CallOpts)
}

// MMaxChainCode is a free data retrieval call binding the contract method 0x5c5e274e.
//
// Solidity: function mMaxChainCode() constant returns(uint32)
func (_GatewayVote *GatewayVoteCallerSession) MMaxChainCode() (uint32, error) {
	return _GatewayVote.Contract.MMaxChainCode(&_GatewayVote.CallOpts)
}

// MNumVoters is a free data retrieval call binding the contract method 0xd454f92e.
//
// Solidity: function mNumVoters() constant returns(uint256)
func (_GatewayVote *GatewayVoteCaller) MNumVoters(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _GatewayVote.contract.Call(opts, out, "mNumVoters")
	return *ret0, err
}

// MNumVoters is a free data retrieval call binding the contract method 0xd454f92e.
//
// Solidity: function mNumVoters() constant returns(uint256)
func (_GatewayVote *GatewayVoteSession) MNumVoters() (*big.Int, error) {
	return _GatewayVote.Contract.MNumVoters(&_GatewayVote.CallOpts)
}

// MNumVoters is a free data retrieval call binding the contract method 0xd454f92e.
//
// Solidity: function mNumVoters() constant returns(uint256)
func (_GatewayVote *GatewayVoteCallerSession) MNumVoters() (*big.Int, error) {
	return _GatewayVote.Contract.MNumVoters(&_GatewayVote.CallOpts)
}

// MStopped is a free data retrieval call binding the contract method 0xadce80d5.
//
// Solidity: function mStopped() constant returns(bool)
func (_GatewayVote *GatewayVoteCaller) MStopped(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _GatewayVote.contract.Call(opts, out, "mStopped")
	return *ret0, err
}

// MStopped is a free data retrieval call binding the contract method 0xadce80d5.
//
// Solidity: function mStopped() constant returns(bool)
func (_GatewayVote *GatewayVoteSession) MStopped() (bool, error) {
	return _GatewayVote.Contract.MStopped(&_GatewayVote.CallOpts)
}

// MStopped is a free data retrieval call binding the contract method 0xadce80d5.
//
// Solidity: function mStopped() constant returns(bool)
func (_GatewayVote *GatewayVoteCallerSession) MStopped() (bool, error) {
	return _GatewayVote.Contract.MStopped(&_GatewayVote.CallOpts)
}

// AddApp is a paid mutator transaction binding the contract method 0x2fee26ce.
//
// Solidity: function addApp(app address, chain uint32, token uint32, proposal string) returns()
func (_GatewayVote *GatewayVoteTransactor) AddApp(opts *bind.TransactOpts, app common.Address, chain uint32, token uint32, proposal string) (*types.Transaction, error) {
	return _GatewayVote.contract.Transact(opts, "addApp", app, chain, token, proposal)
}

// AddApp is a paid mutator transaction binding the contract method 0x2fee26ce.
//
// Solidity: function addApp(app address, chain uint32, token uint32, proposal string) returns()
func (_GatewayVote *GatewayVoteSession) AddApp(app common.Address, chain uint32, token uint32, proposal string) (*types.Transaction, error) {
	return _GatewayVote.Contract.AddApp(&_GatewayVote.TransactOpts, app, chain, token, proposal)
}

// AddApp is a paid mutator transaction binding the contract method 0x2fee26ce.
//
// Solidity: function addApp(app address, chain uint32, token uint32, proposal string) returns()
func (_GatewayVote *GatewayVoteTransactorSession) AddApp(app common.Address, chain uint32, token uint32, proposal string) (*types.Transaction, error) {
	return _GatewayVote.Contract.AddApp(&_GatewayVote.TransactOpts, app, chain, token, proposal)
}

// AddChain is a paid mutator transaction binding the contract method 0xa9f940c3.
//
// Solidity: function addChain(chain string, proposal string) returns()
func (_GatewayVote *GatewayVoteTransactor) AddChain(opts *bind.TransactOpts, chain string, proposal string) (*types.Transaction, error) {
	return _GatewayVote.contract.Transact(opts, "addChain", chain, proposal)
}

// AddChain is a paid mutator transaction binding the contract method 0xa9f940c3.
//
// Solidity: function addChain(chain string, proposal string) returns()
func (_GatewayVote *GatewayVoteSession) AddChain(chain string, proposal string) (*types.Transaction, error) {
	return _GatewayVote.Contract.AddChain(&_GatewayVote.TransactOpts, chain, proposal)
}

// AddChain is a paid mutator transaction binding the contract method 0xa9f940c3.
//
// Solidity: function addChain(chain string, proposal string) returns()
func (_GatewayVote *GatewayVoteTransactorSession) AddChain(chain string, proposal string) (*types.Transaction, error) {
	return _GatewayVote.Contract.AddChain(&_GatewayVote.TransactOpts, chain, proposal)
}

// AddVoter is a paid mutator transaction binding the contract method 0xd9e95a98.
//
// Solidity: function addVoter(newer address, proposal string) returns()
func (_GatewayVote *GatewayVoteTransactor) AddVoter(opts *bind.TransactOpts, newer common.Address, proposal string) (*types.Transaction, error) {
	return _GatewayVote.contract.Transact(opts, "addVoter", newer, proposal)
}

// AddVoter is a paid mutator transaction binding the contract method 0xd9e95a98.
//
// Solidity: function addVoter(newer address, proposal string) returns()
func (_GatewayVote *GatewayVoteSession) AddVoter(newer common.Address, proposal string) (*types.Transaction, error) {
	return _GatewayVote.Contract.AddVoter(&_GatewayVote.TransactOpts, newer, proposal)
}

// AddVoter is a paid mutator transaction binding the contract method 0xd9e95a98.
//
// Solidity: function addVoter(newer address, proposal string) returns()
func (_GatewayVote *GatewayVoteTransactorSession) AddVoter(newer common.Address, proposal string) (*types.Transaction, error) {
	return _GatewayVote.Contract.AddVoter(&_GatewayVote.TransactOpts, newer, proposal)
}

// Burn is a paid mutator transaction binding the contract method 0xba9d3297.
//
// Solidity: function burn(wad uint64, dstDescribe string, fee uint64) returns()
func (_GatewayVote *GatewayVoteTransactor) Burn(opts *bind.TransactOpts, wad uint64, dstDescribe string, fee uint64) (*types.Transaction, error) {
	return _GatewayVote.contract.Transact(opts, "burn", wad, dstDescribe, fee)
}

// Burn is a paid mutator transaction binding the contract method 0xba9d3297.
//
// Solidity: function burn(wad uint64, dstDescribe string, fee uint64) returns()
func (_GatewayVote *GatewayVoteSession) Burn(wad uint64, dstDescribe string, fee uint64) (*types.Transaction, error) {
	return _GatewayVote.Contract.Burn(&_GatewayVote.TransactOpts, wad, dstDescribe, fee)
}

// Burn is a paid mutator transaction binding the contract method 0xba9d3297.
//
// Solidity: function burn(wad uint64, dstDescribe string, fee uint64) returns()
func (_GatewayVote *GatewayVoteTransactorSession) Burn(wad uint64, dstDescribe string, fee uint64) (*types.Transaction, error) {
	return _GatewayVote.Contract.Burn(&_GatewayVote.TransactOpts, wad, dstDescribe, fee)
}

// BurnForGateway is a paid mutator transaction binding the contract method 0xdf958553.
//
// Solidity: function burnForGateway(from address, dstDescribe string, wad uint64, fee uint64) returns()
func (_GatewayVote *GatewayVoteTransactor) BurnForGateway(opts *bind.TransactOpts, from common.Address, dstDescribe string, wad uint64, fee uint64) (*types.Transaction, error) {
	return _GatewayVote.contract.Transact(opts, "burnForGateway", from, dstDescribe, wad, fee)
}

// BurnForGateway is a paid mutator transaction binding the contract method 0xdf958553.
//
// Solidity: function burnForGateway(from address, dstDescribe string, wad uint64, fee uint64) returns()
func (_GatewayVote *GatewayVoteSession) BurnForGateway(from common.Address, dstDescribe string, wad uint64, fee uint64) (*types.Transaction, error) {
	return _GatewayVote.Contract.BurnForGateway(&_GatewayVote.TransactOpts, from, dstDescribe, wad, fee)
}

// BurnForGateway is a paid mutator transaction binding the contract method 0xdf958553.
//
// Solidity: function burnForGateway(from address, dstDescribe string, wad uint64, fee uint64) returns()
func (_GatewayVote *GatewayVoteTransactorSession) BurnForGateway(from common.Address, dstDescribe string, wad uint64, fee uint64) (*types.Transaction, error) {
	return _GatewayVote.Contract.BurnForGateway(&_GatewayVote.TransactOpts, from, dstDescribe, wad, fee)
}

// ChangeGatewayAddr is a paid mutator transaction binding the contract method 0x6f62e755.
//
// Solidity: function changeGatewayAddr(appCode uint32, newer address, proposal string) returns()
func (_GatewayVote *GatewayVoteTransactor) ChangeGatewayAddr(opts *bind.TransactOpts, appCode uint32, newer common.Address, proposal string) (*types.Transaction, error) {
	return _GatewayVote.contract.Transact(opts, "changeGatewayAddr", appCode, newer, proposal)
}

// ChangeGatewayAddr is a paid mutator transaction binding the contract method 0x6f62e755.
//
// Solidity: function changeGatewayAddr(appCode uint32, newer address, proposal string) returns()
func (_GatewayVote *GatewayVoteSession) ChangeGatewayAddr(appCode uint32, newer common.Address, proposal string) (*types.Transaction, error) {
	return _GatewayVote.Contract.ChangeGatewayAddr(&_GatewayVote.TransactOpts, appCode, newer, proposal)
}

// ChangeGatewayAddr is a paid mutator transaction binding the contract method 0x6f62e755.
//
// Solidity: function changeGatewayAddr(appCode uint32, newer address, proposal string) returns()
func (_GatewayVote *GatewayVoteTransactorSession) ChangeGatewayAddr(appCode uint32, newer common.Address, proposal string) (*types.Transaction, error) {
	return _GatewayVote.Contract.ChangeGatewayAddr(&_GatewayVote.TransactOpts, appCode, newer, proposal)
}

// ChangeVoter is a paid mutator transaction binding the contract method 0xd4975d71.
//
// Solidity: function changeVoter(older address, newer address, proposal string) returns()
func (_GatewayVote *GatewayVoteTransactor) ChangeVoter(opts *bind.TransactOpts, older common.Address, newer common.Address, proposal string) (*types.Transaction, error) {
	return _GatewayVote.contract.Transact(opts, "changeVoter", older, newer, proposal)
}

// ChangeVoter is a paid mutator transaction binding the contract method 0xd4975d71.
//
// Solidity: function changeVoter(older address, newer address, proposal string) returns()
func (_GatewayVote *GatewayVoteSession) ChangeVoter(older common.Address, newer common.Address, proposal string) (*types.Transaction, error) {
	return _GatewayVote.Contract.ChangeVoter(&_GatewayVote.TransactOpts, older, newer, proposal)
}

// ChangeVoter is a paid mutator transaction binding the contract method 0xd4975d71.
//
// Solidity: function changeVoter(older address, newer address, proposal string) returns()
func (_GatewayVote *GatewayVoteTransactorSession) ChangeVoter(older common.Address, newer common.Address, proposal string) (*types.Transaction, error) {
	return _GatewayVote.Contract.ChangeVoter(&_GatewayVote.TransactOpts, older, newer, proposal)
}

// MintByGateway is a paid mutator transaction binding the contract method 0xb0036d32.
//
// Solidity: function mintByGateway(appCode uint32, wad uint64, receiver address, proposal string) returns()
func (_GatewayVote *GatewayVoteTransactor) MintByGateway(opts *bind.TransactOpts, appCode uint32, wad uint64, receiver common.Address, proposal string) (*types.Transaction, error) {
	return _GatewayVote.contract.Transact(opts, "mintByGateway", appCode, wad, receiver, proposal)
}

// MintByGateway is a paid mutator transaction binding the contract method 0xb0036d32.
//
// Solidity: function mintByGateway(appCode uint32, wad uint64, receiver address, proposal string) returns()
func (_GatewayVote *GatewayVoteSession) MintByGateway(appCode uint32, wad uint64, receiver common.Address, proposal string) (*types.Transaction, error) {
	return _GatewayVote.Contract.MintByGateway(&_GatewayVote.TransactOpts, appCode, wad, receiver, proposal)
}

// MintByGateway is a paid mutator transaction binding the contract method 0xb0036d32.
//
// Solidity: function mintByGateway(appCode uint32, wad uint64, receiver address, proposal string) returns()
func (_GatewayVote *GatewayVoteTransactorSession) MintByGateway(appCode uint32, wad uint64, receiver common.Address, proposal string) (*types.Transaction, error) {
	return _GatewayVote.Contract.MintByGateway(&_GatewayVote.TransactOpts, appCode, wad, receiver, proposal)
}

// RemoveApp is a paid mutator transaction binding the contract method 0x917ec8e2.
//
// Solidity: function removeApp(code uint32, proposal string) returns()
func (_GatewayVote *GatewayVoteTransactor) RemoveApp(opts *bind.TransactOpts, code uint32, proposal string) (*types.Transaction, error) {
	return _GatewayVote.contract.Transact(opts, "removeApp", code, proposal)
}

// RemoveApp is a paid mutator transaction binding the contract method 0x917ec8e2.
//
// Solidity: function removeApp(code uint32, proposal string) returns()
func (_GatewayVote *GatewayVoteSession) RemoveApp(code uint32, proposal string) (*types.Transaction, error) {
	return _GatewayVote.Contract.RemoveApp(&_GatewayVote.TransactOpts, code, proposal)
}

// RemoveApp is a paid mutator transaction binding the contract method 0x917ec8e2.
//
// Solidity: function removeApp(code uint32, proposal string) returns()
func (_GatewayVote *GatewayVoteTransactorSession) RemoveApp(code uint32, proposal string) (*types.Transaction, error) {
	return _GatewayVote.Contract.RemoveApp(&_GatewayVote.TransactOpts, code, proposal)
}

// RemoveVoter is a paid mutator transaction binding the contract method 0xd9037522.
//
// Solidity: function removeVoter(older address, proposal string) returns()
func (_GatewayVote *GatewayVoteTransactor) RemoveVoter(opts *bind.TransactOpts, older common.Address, proposal string) (*types.Transaction, error) {
	return _GatewayVote.contract.Transact(opts, "removeVoter", older, proposal)
}

// RemoveVoter is a paid mutator transaction binding the contract method 0xd9037522.
//
// Solidity: function removeVoter(older address, proposal string) returns()
func (_GatewayVote *GatewayVoteSession) RemoveVoter(older common.Address, proposal string) (*types.Transaction, error) {
	return _GatewayVote.Contract.RemoveVoter(&_GatewayVote.TransactOpts, older, proposal)
}

// RemoveVoter is a paid mutator transaction binding the contract method 0xd9037522.
//
// Solidity: function removeVoter(older address, proposal string) returns()
func (_GatewayVote *GatewayVoteTransactorSession) RemoveVoter(older common.Address, proposal string) (*types.Transaction, error) {
	return _GatewayVote.Contract.RemoveVoter(&_GatewayVote.TransactOpts, older, proposal)
}

// Revoke is a paid mutator transaction binding the contract method 0x20c5429b.
//
// Solidity: function revoke(operation uint256) returns()
func (_GatewayVote *GatewayVoteTransactor) Revoke(opts *bind.TransactOpts, operation *big.Int) (*types.Transaction, error) {
	return _GatewayVote.contract.Transact(opts, "revoke", operation)
}

// Revoke is a paid mutator transaction binding the contract method 0x20c5429b.
//
// Solidity: function revoke(operation uint256) returns()
func (_GatewayVote *GatewayVoteSession) Revoke(operation *big.Int) (*types.Transaction, error) {
	return _GatewayVote.Contract.Revoke(&_GatewayVote.TransactOpts, operation)
}

// Revoke is a paid mutator transaction binding the contract method 0x20c5429b.
//
// Solidity: function revoke(operation uint256) returns()
func (_GatewayVote *GatewayVoteTransactorSession) Revoke(operation *big.Int) (*types.Transaction, error) {
	return _GatewayVote.Contract.Revoke(&_GatewayVote.TransactOpts, operation)
}

// Start is a paid mutator transaction binding the contract method 0x258e60b6.
//
// Solidity: function start(proposal string) returns()
func (_GatewayVote *GatewayVoteTransactor) Start(opts *bind.TransactOpts, proposal string) (*types.Transaction, error) {
	return _GatewayVote.contract.Transact(opts, "start", proposal)
}

// Start is a paid mutator transaction binding the contract method 0x258e60b6.
//
// Solidity: function start(proposal string) returns()
func (_GatewayVote *GatewayVoteSession) Start(proposal string) (*types.Transaction, error) {
	return _GatewayVote.Contract.Start(&_GatewayVote.TransactOpts, proposal)
}

// Start is a paid mutator transaction binding the contract method 0x258e60b6.
//
// Solidity: function start(proposal string) returns()
func (_GatewayVote *GatewayVoteTransactorSession) Start(proposal string) (*types.Transaction, error) {
	return _GatewayVote.Contract.Start(&_GatewayVote.TransactOpts, proposal)
}

// Stop is a paid mutator transaction binding the contract method 0x745dcd4d.
//
// Solidity: function stop(proposal string) returns()
func (_GatewayVote *GatewayVoteTransactor) Stop(opts *bind.TransactOpts, proposal string) (*types.Transaction, error) {
	return _GatewayVote.contract.Transact(opts, "stop", proposal)
}

// Stop is a paid mutator transaction binding the contract method 0x745dcd4d.
//
// Solidity: function stop(proposal string) returns()
func (_GatewayVote *GatewayVoteSession) Stop(proposal string) (*types.Transaction, error) {
	return _GatewayVote.Contract.Stop(&_GatewayVote.TransactOpts, proposal)
}

// Stop is a paid mutator transaction binding the contract method 0x745dcd4d.
//
// Solidity: function stop(proposal string) returns()
func (_GatewayVote *GatewayVoteTransactorSession) Stop(proposal string) (*types.Transaction, error) {
	return _GatewayVote.Contract.Stop(&_GatewayVote.TransactOpts, proposal)
}

// GatewayVoteAppAddedIterator is returned from FilterAppAdded and is used to iterate over the raw logs and unpacked data for AppAdded events raised by the GatewayVote contract.
type GatewayVoteAppAddedIterator struct {
	Event *GatewayVoteAppAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayVoteAppAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayVoteAppAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayVoteAppAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayVoteAppAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayVoteAppAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayVoteAppAdded represents a AppAdded event raised by the GatewayVote contract.
type GatewayVoteAppAdded struct {
	App       common.Address
	Chain     uint32
	Token     uint32
	Operation *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAppAdded is a free log retrieval operation binding the contract event 0x114bbbbe49914625d327d46098e0b09cbbbedbeb3d55151b2c278f81e0714824.
//
// Solidity: event AppAdded(app address, chain uint32, token uint32, operation indexed uint256)
func (_GatewayVote *GatewayVoteFilterer) FilterAppAdded(opts *bind.FilterOpts, operation []*big.Int) (*GatewayVoteAppAddedIterator, error) {

	var operationRule []interface{}
	for _, operationItem := range operation {
		operationRule = append(operationRule, operationItem)
	}

	logs, sub, err := _GatewayVote.contract.FilterLogs(opts, "AppAdded", operationRule)
	if err != nil {
		return nil, err
	}
	return &GatewayVoteAppAddedIterator{contract: _GatewayVote.contract, event: "AppAdded", logs: logs, sub: sub}, nil
}

// WatchAppAdded is a free log subscription operation binding the contract event 0x114bbbbe49914625d327d46098e0b09cbbbedbeb3d55151b2c278f81e0714824.
//
// Solidity: event AppAdded(app address, chain uint32, token uint32, operation indexed uint256)
func (_GatewayVote *GatewayVoteFilterer) WatchAppAdded(opts *bind.WatchOpts, sink chan<- *GatewayVoteAppAdded, operation []*big.Int) (event.Subscription, error) {

	var operationRule []interface{}
	for _, operationItem := range operation {
		operationRule = append(operationRule, operationItem)
	}

	logs, sub, err := _GatewayVote.contract.WatchLogs(opts, "AppAdded", operationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayVoteAppAdded)
				if err := _GatewayVote.contract.UnpackLog(event, "AppAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// GatewayVoteAppRemovedIterator is returned from FilterAppRemoved and is used to iterate over the raw logs and unpacked data for AppRemoved events raised by the GatewayVote contract.
type GatewayVoteAppRemovedIterator struct {
	Event *GatewayVoteAppRemoved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayVoteAppRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayVoteAppRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayVoteAppRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayVoteAppRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayVoteAppRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayVoteAppRemoved represents a AppRemoved event raised by the GatewayVote contract.
type GatewayVoteAppRemoved struct {
	Code      uint32
	Operation *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAppRemoved is a free log retrieval operation binding the contract event 0x423cd729ce1dc9d7c6999dfe1e3dfc33f730375427ec73ab1a1c3ca05e2c41ce.
//
// Solidity: event AppRemoved(code uint32, operation indexed uint256)
func (_GatewayVote *GatewayVoteFilterer) FilterAppRemoved(opts *bind.FilterOpts, operation []*big.Int) (*GatewayVoteAppRemovedIterator, error) {

	var operationRule []interface{}
	for _, operationItem := range operation {
		operationRule = append(operationRule, operationItem)
	}

	logs, sub, err := _GatewayVote.contract.FilterLogs(opts, "AppRemoved", operationRule)
	if err != nil {
		return nil, err
	}
	return &GatewayVoteAppRemovedIterator{contract: _GatewayVote.contract, event: "AppRemoved", logs: logs, sub: sub}, nil
}

// WatchAppRemoved is a free log subscription operation binding the contract event 0x423cd729ce1dc9d7c6999dfe1e3dfc33f730375427ec73ab1a1c3ca05e2c41ce.
//
// Solidity: event AppRemoved(code uint32, operation indexed uint256)
func (_GatewayVote *GatewayVoteFilterer) WatchAppRemoved(opts *bind.WatchOpts, sink chan<- *GatewayVoteAppRemoved, operation []*big.Int) (event.Subscription, error) {

	var operationRule []interface{}
	for _, operationItem := range operation {
		operationRule = append(operationRule, operationItem)
	}

	logs, sub, err := _GatewayVote.contract.WatchLogs(opts, "AppRemoved", operationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayVoteAppRemoved)
				if err := _GatewayVote.contract.UnpackLog(event, "AppRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// GatewayVoteBurnForGatewayIterator is returned from FilterBurnForGateway and is used to iterate over the raw logs and unpacked data for BurnForGateway events raised by the GatewayVote contract.
type GatewayVoteBurnForGatewayIterator struct {
	Event *GatewayVoteBurnForGateway // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayVoteBurnForGatewayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayVoteBurnForGateway)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayVoteBurnForGateway)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayVoteBurnForGatewayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayVoteBurnForGatewayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayVoteBurnForGateway represents a BurnForGateway event raised by the GatewayVote contract.
type GatewayVoteBurnForGateway struct {
	AppCode     uint32
	From        common.Address
	DstDescribe string
	Wad         uint64
	Fee         uint64
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBurnForGateway is a free log retrieval operation binding the contract event 0x170dc623c0294307cd20a299e1b1285180614d84477e13c10e2ab290ff50223a.
//
// Solidity: event BurnForGateway(appCode uint32, from address, dstDescribe string, wad uint64, fee uint64)
func (_GatewayVote *GatewayVoteFilterer) FilterBurnForGateway(opts *bind.FilterOpts) (*GatewayVoteBurnForGatewayIterator, error) {

	logs, sub, err := _GatewayVote.contract.FilterLogs(opts, "BurnForGateway")
	if err != nil {
		return nil, err
	}
	return &GatewayVoteBurnForGatewayIterator{contract: _GatewayVote.contract, event: "BurnForGateway", logs: logs, sub: sub}, nil
}

// WatchBurnForGateway is a free log subscription operation binding the contract event 0x170dc623c0294307cd20a299e1b1285180614d84477e13c10e2ab290ff50223a.
//
// Solidity: event BurnForGateway(appCode uint32, from address, dstDescribe string, wad uint64, fee uint64)
func (_GatewayVote *GatewayVoteFilterer) WatchBurnForGateway(opts *bind.WatchOpts, sink chan<- *GatewayVoteBurnForGateway) (event.Subscription, error) {

	logs, sub, err := _GatewayVote.contract.WatchLogs(opts, "BurnForGateway")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayVoteBurnForGateway)
				if err := _GatewayVote.contract.UnpackLog(event, "BurnForGateway", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// GatewayVoteChainAddedIterator is returned from FilterChainAdded and is used to iterate over the raw logs and unpacked data for ChainAdded events raised by the GatewayVote contract.
type GatewayVoteChainAddedIterator struct {
	Event *GatewayVoteChainAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayVoteChainAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayVoteChainAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayVoteChainAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayVoteChainAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayVoteChainAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayVoteChainAdded represents a ChainAdded event raised by the GatewayVote contract.
type GatewayVoteChainAdded struct {
	Chain     string
	Operation *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterChainAdded is a free log retrieval operation binding the contract event 0x8398b0e4d6d15e6038c42e15bf0f0f5466dc36dfd49979c0599bc2eb0fb58302.
//
// Solidity: event ChainAdded(chain string, operation indexed uint256)
func (_GatewayVote *GatewayVoteFilterer) FilterChainAdded(opts *bind.FilterOpts, operation []*big.Int) (*GatewayVoteChainAddedIterator, error) {

	var operationRule []interface{}
	for _, operationItem := range operation {
		operationRule = append(operationRule, operationItem)
	}

	logs, sub, err := _GatewayVote.contract.FilterLogs(opts, "ChainAdded", operationRule)
	if err != nil {
		return nil, err
	}
	return &GatewayVoteChainAddedIterator{contract: _GatewayVote.contract, event: "ChainAdded", logs: logs, sub: sub}, nil
}

// WatchChainAdded is a free log subscription operation binding the contract event 0x8398b0e4d6d15e6038c42e15bf0f0f5466dc36dfd49979c0599bc2eb0fb58302.
//
// Solidity: event ChainAdded(chain string, operation indexed uint256)
func (_GatewayVote *GatewayVoteFilterer) WatchChainAdded(opts *bind.WatchOpts, sink chan<- *GatewayVoteChainAdded, operation []*big.Int) (event.Subscription, error) {

	var operationRule []interface{}
	for _, operationItem := range operation {
		operationRule = append(operationRule, operationItem)
	}

	logs, sub, err := _GatewayVote.contract.WatchLogs(opts, "ChainAdded", operationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayVoteChainAdded)
				if err := _GatewayVote.contract.UnpackLog(event, "ChainAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// GatewayVoteConfirmationIterator is returned from FilterConfirmation and is used to iterate over the raw logs and unpacked data for Confirmation events raised by the GatewayVote contract.
type GatewayVoteConfirmationIterator struct {
	Event *GatewayVoteConfirmation // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayVoteConfirmationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayVoteConfirmation)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayVoteConfirmation)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayVoteConfirmationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayVoteConfirmationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayVoteConfirmation represents a Confirmation event raised by the GatewayVote contract.
type GatewayVoteConfirmation struct {
	Voter     common.Address
	Operation *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterConfirmation is a free log retrieval operation binding the contract event 0x4a504a94899432a9846e1aa406dceb1bcfd538bb839071d49d1e5e23f5be30ef.
//
// Solidity: event Confirmation(voter address, operation indexed uint256)
func (_GatewayVote *GatewayVoteFilterer) FilterConfirmation(opts *bind.FilterOpts, operation []*big.Int) (*GatewayVoteConfirmationIterator, error) {

	var operationRule []interface{}
	for _, operationItem := range operation {
		operationRule = append(operationRule, operationItem)
	}

	logs, sub, err := _GatewayVote.contract.FilterLogs(opts, "Confirmation", operationRule)
	if err != nil {
		return nil, err
	}
	return &GatewayVoteConfirmationIterator{contract: _GatewayVote.contract, event: "Confirmation", logs: logs, sub: sub}, nil
}

// WatchConfirmation is a free log subscription operation binding the contract event 0x4a504a94899432a9846e1aa406dceb1bcfd538bb839071d49d1e5e23f5be30ef.
//
// Solidity: event Confirmation(voter address, operation indexed uint256)
func (_GatewayVote *GatewayVoteFilterer) WatchConfirmation(opts *bind.WatchOpts, sink chan<- *GatewayVoteConfirmation, operation []*big.Int) (event.Subscription, error) {

	var operationRule []interface{}
	for _, operationItem := range operation {
		operationRule = append(operationRule, operationItem)
	}

	logs, sub, err := _GatewayVote.contract.WatchLogs(opts, "Confirmation", operationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayVoteConfirmation)
				if err := _GatewayVote.contract.UnpackLog(event, "Confirmation", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// GatewayVoteGatewayAddrChangedIterator is returned from FilterGatewayAddrChanged and is used to iterate over the raw logs and unpacked data for GatewayAddrChanged events raised by the GatewayVote contract.
type GatewayVoteGatewayAddrChangedIterator struct {
	Event *GatewayVoteGatewayAddrChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayVoteGatewayAddrChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayVoteGatewayAddrChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayVoteGatewayAddrChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayVoteGatewayAddrChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayVoteGatewayAddrChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayVoteGatewayAddrChanged represents a GatewayAddrChanged event raised by the GatewayVote contract.
type GatewayVoteGatewayAddrChanged struct {
	AppCode   uint32
	Newer     common.Address
	Operation *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterGatewayAddrChanged is a free log retrieval operation binding the contract event 0xbdfdf0ba24577a2e7e41b362393cc097b4bfcccd5b96a7f9e2a1a68d1871a9b3.
//
// Solidity: event GatewayAddrChanged(appCode uint32, newer address, operation indexed uint256)
func (_GatewayVote *GatewayVoteFilterer) FilterGatewayAddrChanged(opts *bind.FilterOpts, operation []*big.Int) (*GatewayVoteGatewayAddrChangedIterator, error) {

	var operationRule []interface{}
	for _, operationItem := range operation {
		operationRule = append(operationRule, operationItem)
	}

	logs, sub, err := _GatewayVote.contract.FilterLogs(opts, "GatewayAddrChanged", operationRule)
	if err != nil {
		return nil, err
	}
	return &GatewayVoteGatewayAddrChangedIterator{contract: _GatewayVote.contract, event: "GatewayAddrChanged", logs: logs, sub: sub}, nil
}

// WatchGatewayAddrChanged is a free log subscription operation binding the contract event 0xbdfdf0ba24577a2e7e41b362393cc097b4bfcccd5b96a7f9e2a1a68d1871a9b3.
//
// Solidity: event GatewayAddrChanged(appCode uint32, newer address, operation indexed uint256)
func (_GatewayVote *GatewayVoteFilterer) WatchGatewayAddrChanged(opts *bind.WatchOpts, sink chan<- *GatewayVoteGatewayAddrChanged, operation []*big.Int) (event.Subscription, error) {

	var operationRule []interface{}
	for _, operationItem := range operation {
		operationRule = append(operationRule, operationItem)
	}

	logs, sub, err := _GatewayVote.contract.WatchLogs(opts, "GatewayAddrChanged", operationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayVoteGatewayAddrChanged)
				if err := _GatewayVote.contract.UnpackLog(event, "GatewayAddrChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// GatewayVoteMintByGatewayIterator is returned from FilterMintByGateway and is used to iterate over the raw logs and unpacked data for MintByGateway events raised by the GatewayVote contract.
type GatewayVoteMintByGatewayIterator struct {
	Event *GatewayVoteMintByGateway // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayVoteMintByGatewayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayVoteMintByGateway)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayVoteMintByGateway)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayVoteMintByGatewayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayVoteMintByGatewayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayVoteMintByGateway represents a MintByGateway event raised by the GatewayVote contract.
type GatewayVoteMintByGateway struct {
	AppCode   uint32
	Receiver  common.Address
	Wad       uint64
	Operation *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterMintByGateway is a free log retrieval operation binding the contract event 0x7638d347efa67409b7eeb07e8ff0aa4e6cea7d75b3dffa02e2baf972b484b628.
//
// Solidity: event MintByGateway(appCode uint32, receiver address, wad uint64, operation indexed uint256)
func (_GatewayVote *GatewayVoteFilterer) FilterMintByGateway(opts *bind.FilterOpts, operation []*big.Int) (*GatewayVoteMintByGatewayIterator, error) {

	var operationRule []interface{}
	for _, operationItem := range operation {
		operationRule = append(operationRule, operationItem)
	}

	logs, sub, err := _GatewayVote.contract.FilterLogs(opts, "MintByGateway", operationRule)
	if err != nil {
		return nil, err
	}
	return &GatewayVoteMintByGatewayIterator{contract: _GatewayVote.contract, event: "MintByGateway", logs: logs, sub: sub}, nil
}

// WatchMintByGateway is a free log subscription operation binding the contract event 0x7638d347efa67409b7eeb07e8ff0aa4e6cea7d75b3dffa02e2baf972b484b628.
//
// Solidity: event MintByGateway(appCode uint32, receiver address, wad uint64, operation indexed uint256)
func (_GatewayVote *GatewayVoteFilterer) WatchMintByGateway(opts *bind.WatchOpts, sink chan<- *GatewayVoteMintByGateway, operation []*big.Int) (event.Subscription, error) {

	var operationRule []interface{}
	for _, operationItem := range operation {
		operationRule = append(operationRule, operationItem)
	}

	logs, sub, err := _GatewayVote.contract.WatchLogs(opts, "MintByGateway", operationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayVoteMintByGateway)
				if err := _GatewayVote.contract.UnpackLog(event, "MintByGateway", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// GatewayVoteOperationDoneIterator is returned from FilterOperationDone and is used to iterate over the raw logs and unpacked data for OperationDone events raised by the GatewayVote contract.
type GatewayVoteOperationDoneIterator struct {
	Event *GatewayVoteOperationDone // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayVoteOperationDoneIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayVoteOperationDone)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayVoteOperationDone)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayVoteOperationDoneIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayVoteOperationDoneIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayVoteOperationDone represents a OperationDone event raised by the GatewayVote contract.
type GatewayVoteOperationDone struct {
	Voter     common.Address
	Operation *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterOperationDone is a free log retrieval operation binding the contract event 0x82e3a8c792e0e06c7ec9b50e8245ce36ad45e4ea305aebcf6de24e7820e1d59c.
//
// Solidity: event OperationDone(voter address, operation indexed uint256)
func (_GatewayVote *GatewayVoteFilterer) FilterOperationDone(opts *bind.FilterOpts, operation []*big.Int) (*GatewayVoteOperationDoneIterator, error) {

	var operationRule []interface{}
	for _, operationItem := range operation {
		operationRule = append(operationRule, operationItem)
	}

	logs, sub, err := _GatewayVote.contract.FilterLogs(opts, "OperationDone", operationRule)
	if err != nil {
		return nil, err
	}
	return &GatewayVoteOperationDoneIterator{contract: _GatewayVote.contract, event: "OperationDone", logs: logs, sub: sub}, nil
}

// WatchOperationDone is a free log subscription operation binding the contract event 0x82e3a8c792e0e06c7ec9b50e8245ce36ad45e4ea305aebcf6de24e7820e1d59c.
//
// Solidity: event OperationDone(voter address, operation indexed uint256)
func (_GatewayVote *GatewayVoteFilterer) WatchOperationDone(opts *bind.WatchOpts, sink chan<- *GatewayVoteOperationDone, operation []*big.Int) (event.Subscription, error) {

	var operationRule []interface{}
	for _, operationItem := range operation {
		operationRule = append(operationRule, operationItem)
	}

	logs, sub, err := _GatewayVote.contract.WatchLogs(opts, "OperationDone", operationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayVoteOperationDone)
				if err := _GatewayVote.contract.UnpackLog(event, "OperationDone", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// GatewayVoteRevokeIterator is returned from FilterRevoke and is used to iterate over the raw logs and unpacked data for Revoke events raised by the GatewayVote contract.
type GatewayVoteRevokeIterator struct {
	Event *GatewayVoteRevoke // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayVoteRevokeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayVoteRevoke)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayVoteRevoke)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayVoteRevokeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayVoteRevokeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayVoteRevoke represents a Revoke event raised by the GatewayVote contract.
type GatewayVoteRevoke struct {
	Revoker   common.Address
	Operation *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRevoke is a free log retrieval operation binding the contract event 0xec9ab91322523c899ede7830ec9bfc992b5981cdcc27b91162fb23de5791117b.
//
// Solidity: event Revoke(revoker address, operation indexed uint256)
func (_GatewayVote *GatewayVoteFilterer) FilterRevoke(opts *bind.FilterOpts, operation []*big.Int) (*GatewayVoteRevokeIterator, error) {

	var operationRule []interface{}
	for _, operationItem := range operation {
		operationRule = append(operationRule, operationItem)
	}

	logs, sub, err := _GatewayVote.contract.FilterLogs(opts, "Revoke", operationRule)
	if err != nil {
		return nil, err
	}
	return &GatewayVoteRevokeIterator{contract: _GatewayVote.contract, event: "Revoke", logs: logs, sub: sub}, nil
}

// WatchRevoke is a free log subscription operation binding the contract event 0xec9ab91322523c899ede7830ec9bfc992b5981cdcc27b91162fb23de5791117b.
//
// Solidity: event Revoke(revoker address, operation indexed uint256)
func (_GatewayVote *GatewayVoteFilterer) WatchRevoke(opts *bind.WatchOpts, sink chan<- *GatewayVoteRevoke, operation []*big.Int) (event.Subscription, error) {

	var operationRule []interface{}
	for _, operationItem := range operation {
		operationRule = append(operationRule, operationItem)
	}

	logs, sub, err := _GatewayVote.contract.WatchLogs(opts, "Revoke", operationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayVoteRevoke)
				if err := _GatewayVote.contract.UnpackLog(event, "Revoke", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// GatewayVoteStartedIterator is returned from FilterStarted and is used to iterate over the raw logs and unpacked data for Started events raised by the GatewayVote contract.
type GatewayVoteStartedIterator struct {
	Event *GatewayVoteStarted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayVoteStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayVoteStarted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayVoteStarted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayVoteStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayVoteStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayVoteStarted represents a Started event raised by the GatewayVote contract.
type GatewayVoteStarted struct {
	Operation *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterStarted is a free log retrieval operation binding the contract event 0x006e0c97de781a7389d44ba8fd35d1467cabb17ed04d038d166d34ab819213f3.
//
// Solidity: event Started(operation indexed uint256)
func (_GatewayVote *GatewayVoteFilterer) FilterStarted(opts *bind.FilterOpts, operation []*big.Int) (*GatewayVoteStartedIterator, error) {

	var operationRule []interface{}
	for _, operationItem := range operation {
		operationRule = append(operationRule, operationItem)
	}

	logs, sub, err := _GatewayVote.contract.FilterLogs(opts, "Started", operationRule)
	if err != nil {
		return nil, err
	}
	return &GatewayVoteStartedIterator{contract: _GatewayVote.contract, event: "Started", logs: logs, sub: sub}, nil
}

// WatchStarted is a free log subscription operation binding the contract event 0x006e0c97de781a7389d44ba8fd35d1467cabb17ed04d038d166d34ab819213f3.
//
// Solidity: event Started(operation indexed uint256)
func (_GatewayVote *GatewayVoteFilterer) WatchStarted(opts *bind.WatchOpts, sink chan<- *GatewayVoteStarted, operation []*big.Int) (event.Subscription, error) {

	var operationRule []interface{}
	for _, operationItem := range operation {
		operationRule = append(operationRule, operationItem)
	}

	logs, sub, err := _GatewayVote.contract.WatchLogs(opts, "Started", operationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayVoteStarted)
				if err := _GatewayVote.contract.UnpackLog(event, "Started", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// GatewayVoteStoppedIterator is returned from FilterStopped and is used to iterate over the raw logs and unpacked data for Stopped events raised by the GatewayVote contract.
type GatewayVoteStoppedIterator struct {
	Event *GatewayVoteStopped // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayVoteStoppedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayVoteStopped)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayVoteStopped)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayVoteStoppedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayVoteStoppedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayVoteStopped represents a Stopped event raised by the GatewayVote contract.
type GatewayVoteStopped struct {
	Operation *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterStopped is a free log retrieval operation binding the contract event 0x5378e99f1c6cc344f21e74745742dd30c73797900a542307d5d109383b01bf28.
//
// Solidity: event Stopped(operation indexed uint256)
func (_GatewayVote *GatewayVoteFilterer) FilterStopped(opts *bind.FilterOpts, operation []*big.Int) (*GatewayVoteStoppedIterator, error) {

	var operationRule []interface{}
	for _, operationItem := range operation {
		operationRule = append(operationRule, operationItem)
	}

	logs, sub, err := _GatewayVote.contract.FilterLogs(opts, "Stopped", operationRule)
	if err != nil {
		return nil, err
	}
	return &GatewayVoteStoppedIterator{contract: _GatewayVote.contract, event: "Stopped", logs: logs, sub: sub}, nil
}

// WatchStopped is a free log subscription operation binding the contract event 0x5378e99f1c6cc344f21e74745742dd30c73797900a542307d5d109383b01bf28.
//
// Solidity: event Stopped(operation indexed uint256)
func (_GatewayVote *GatewayVoteFilterer) WatchStopped(opts *bind.WatchOpts, sink chan<- *GatewayVoteStopped, operation []*big.Int) (event.Subscription, error) {

	var operationRule []interface{}
	for _, operationItem := range operation {
		operationRule = append(operationRule, operationItem)
	}

	logs, sub, err := _GatewayVote.contract.WatchLogs(opts, "Stopped", operationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayVoteStopped)
				if err := _GatewayVote.contract.UnpackLog(event, "Stopped", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// GatewayVoteVoterAddedIterator is returned from FilterVoterAdded and is used to iterate over the raw logs and unpacked data for VoterAdded events raised by the GatewayVote contract.
type GatewayVoteVoterAddedIterator struct {
	Event *GatewayVoteVoterAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayVoteVoterAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayVoteVoterAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayVoteVoterAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayVoteVoterAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayVoteVoterAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayVoteVoterAdded represents a VoterAdded event raised by the GatewayVote contract.
type GatewayVoteVoterAdded struct {
	NewVoter  common.Address
	Operation *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterVoterAdded is a free log retrieval operation binding the contract event 0x56aa1768b52ddcdda65c5561181844717c2eed4bb3c848376d4420860da04dae.
//
// Solidity: event VoterAdded(newVoter address, operation indexed uint256)
func (_GatewayVote *GatewayVoteFilterer) FilterVoterAdded(opts *bind.FilterOpts, operation []*big.Int) (*GatewayVoteVoterAddedIterator, error) {

	var operationRule []interface{}
	for _, operationItem := range operation {
		operationRule = append(operationRule, operationItem)
	}

	logs, sub, err := _GatewayVote.contract.FilterLogs(opts, "VoterAdded", operationRule)
	if err != nil {
		return nil, err
	}
	return &GatewayVoteVoterAddedIterator{contract: _GatewayVote.contract, event: "VoterAdded", logs: logs, sub: sub}, nil
}

// WatchVoterAdded is a free log subscription operation binding the contract event 0x56aa1768b52ddcdda65c5561181844717c2eed4bb3c848376d4420860da04dae.
//
// Solidity: event VoterAdded(newVoter address, operation indexed uint256)
func (_GatewayVote *GatewayVoteFilterer) WatchVoterAdded(opts *bind.WatchOpts, sink chan<- *GatewayVoteVoterAdded, operation []*big.Int) (event.Subscription, error) {

	var operationRule []interface{}
	for _, operationItem := range operation {
		operationRule = append(operationRule, operationItem)
	}

	logs, sub, err := _GatewayVote.contract.WatchLogs(opts, "VoterAdded", operationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayVoteVoterAdded)
				if err := _GatewayVote.contract.UnpackLog(event, "VoterAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// GatewayVoteVoterChangedIterator is returned from FilterVoterChanged and is used to iterate over the raw logs and unpacked data for VoterChanged events raised by the GatewayVote contract.
type GatewayVoteVoterChangedIterator struct {
	Event *GatewayVoteVoterChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayVoteVoterChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayVoteVoterChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayVoteVoterChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayVoteVoterChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayVoteVoterChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayVoteVoterChanged represents a VoterChanged event raised by the GatewayVote contract.
type GatewayVoteVoterChanged struct {
	OldVoter  common.Address
	NewVoter  common.Address
	Operation *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterVoterChanged is a free log retrieval operation binding the contract event 0x67c6491ab3a0c25039987fa6abbc4dfdba14ea7d61df894465777da1bb1605b9.
//
// Solidity: event VoterChanged(oldVoter address, newVoter address, operation indexed uint256)
func (_GatewayVote *GatewayVoteFilterer) FilterVoterChanged(opts *bind.FilterOpts, operation []*big.Int) (*GatewayVoteVoterChangedIterator, error) {

	var operationRule []interface{}
	for _, operationItem := range operation {
		operationRule = append(operationRule, operationItem)
	}

	logs, sub, err := _GatewayVote.contract.FilterLogs(opts, "VoterChanged", operationRule)
	if err != nil {
		return nil, err
	}
	return &GatewayVoteVoterChangedIterator{contract: _GatewayVote.contract, event: "VoterChanged", logs: logs, sub: sub}, nil
}

// WatchVoterChanged is a free log subscription operation binding the contract event 0x67c6491ab3a0c25039987fa6abbc4dfdba14ea7d61df894465777da1bb1605b9.
//
// Solidity: event VoterChanged(oldVoter address, newVoter address, operation indexed uint256)
func (_GatewayVote *GatewayVoteFilterer) WatchVoterChanged(opts *bind.WatchOpts, sink chan<- *GatewayVoteVoterChanged, operation []*big.Int) (event.Subscription, error) {

	var operationRule []interface{}
	for _, operationItem := range operation {
		operationRule = append(operationRule, operationItem)
	}

	logs, sub, err := _GatewayVote.contract.WatchLogs(opts, "VoterChanged", operationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayVoteVoterChanged)
				if err := _GatewayVote.contract.UnpackLog(event, "VoterChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// GatewayVoteVoterRemovedIterator is returned from FilterVoterRemoved and is used to iterate over the raw logs and unpacked data for VoterRemoved events raised by the GatewayVote contract.
type GatewayVoteVoterRemovedIterator struct {
	Event *GatewayVoteVoterRemoved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GatewayVoteVoterRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayVoteVoterRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GatewayVoteVoterRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GatewayVoteVoterRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayVoteVoterRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayVoteVoterRemoved represents a VoterRemoved event raised by the GatewayVote contract.
type GatewayVoteVoterRemoved struct {
	OldVoter  common.Address
	Operation *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterVoterRemoved is a free log retrieval operation binding the contract event 0x98a7f87f8e2aa2f23f43769eff67782bb12946384b142d1ce1e8e38e05d9a3e6.
//
// Solidity: event VoterRemoved(oldVoter address, operation indexed uint256)
func (_GatewayVote *GatewayVoteFilterer) FilterVoterRemoved(opts *bind.FilterOpts, operation []*big.Int) (*GatewayVoteVoterRemovedIterator, error) {

	var operationRule []interface{}
	for _, operationItem := range operation {
		operationRule = append(operationRule, operationItem)
	}

	logs, sub, err := _GatewayVote.contract.FilterLogs(opts, "VoterRemoved", operationRule)
	if err != nil {
		return nil, err
	}
	return &GatewayVoteVoterRemovedIterator{contract: _GatewayVote.contract, event: "VoterRemoved", logs: logs, sub: sub}, nil
}

// WatchVoterRemoved is a free log subscription operation binding the contract event 0x98a7f87f8e2aa2f23f43769eff67782bb12946384b142d1ce1e8e38e05d9a3e6.
//
// Solidity: event VoterRemoved(oldVoter address, operation indexed uint256)
func (_GatewayVote *GatewayVoteFilterer) WatchVoterRemoved(opts *bind.WatchOpts, sink chan<- *GatewayVoteVoterRemoved, operation []*big.Int) (event.Subscription, error) {

	var operationRule []interface{}
	for _, operationItem := range operation {
		operationRule = append(operationRule, operationItem)
	}

	logs, sub, err := _GatewayVote.contract.WatchLogs(opts, "VoterRemoved", operationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayVoteVoterRemoved)
				if err := _GatewayVote.contract.UnpackLog(event, "VoterRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}
