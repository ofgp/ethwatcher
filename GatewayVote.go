// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ethwatcher

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

// GatewayVoteABI is the input ABI used to generate the binding from.
const GatewayVoteABI = "[{\"stateMutability\":\"view\",\"constant\":true,\"type\":\"function\",\"inputs\":[{\"name\":\"operation\",\"type\":\"uint256\"},{\"name\":\"voter\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"name\":\"hasConfirmed\",\"payable\":false},{\"stateMutability\":\"nonpayable\",\"constant\":false,\"type\":\"function\",\"inputs\":[{\"name\":\"wad\",\"type\":\"uint64\"},{\"name\":\"receiver\",\"type\":\"string\"}],\"outputs\":[],\"name\":\"burn\",\"payable\":false},{\"stateMutability\":\"view\",\"constant\":true,\"type\":\"function\",\"inputs\":[{\"name\":\"code\",\"type\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"name\":\"isChainCode\",\"payable\":false},{\"stateMutability\":\"nonpayable\",\"constant\":false,\"type\":\"function\",\"inputs\":[{\"name\":\"operation\",\"type\":\"uint256\"}],\"outputs\":[],\"name\":\"revoke\",\"payable\":false},{\"stateMutability\":\"nonpayable\",\"constant\":false,\"type\":\"function\",\"inputs\":[{\"name\":\"proposal\",\"type\":\"string\"}],\"outputs\":[],\"name\":\"start\",\"payable\":false},{\"stateMutability\":\"nonpayable\",\"constant\":false,\"type\":\"function\",\"inputs\":[{\"name\":\"app\",\"type\":\"address\"},{\"name\":\"chain\",\"type\":\"uint32\"},{\"name\":\"token\",\"type\":\"uint32\"},{\"name\":\"proposal\",\"type\":\"string\"}],\"outputs\":[],\"name\":\"addApp\",\"payable\":false},{\"stateMutability\":\"nonpayable\",\"constant\":false,\"type\":\"function\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"receiver\",\"type\":\"string\"},{\"name\":\"wad\",\"type\":\"uint64\"}],\"outputs\":[],\"name\":\"burnForGateway\",\"payable\":false},{\"stateMutability\":\"view\",\"constant\":true,\"type\":\"function\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\"}],\"name\":\"mMaxAppCode\",\"payable\":false},{\"stateMutability\":\"view\",\"constant\":true,\"type\":\"function\",\"inputs\":[{\"name\":\"code\",\"type\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"name\":\"getChainName\",\"payable\":false},{\"stateMutability\":\"view\",\"constant\":true,\"type\":\"function\",\"inputs\":[{\"name\":\"app\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\"}],\"name\":\"getAppCode\",\"payable\":false},{\"stateMutability\":\"view\",\"constant\":true,\"type\":\"function\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\"}],\"name\":\"mMaxChainCode\",\"payable\":false},{\"stateMutability\":\"nonpayable\",\"constant\":false,\"type\":\"function\",\"inputs\":[{\"name\":\"appCode\",\"type\":\"uint32\"},{\"name\":\"newer\",\"type\":\"address\"},{\"name\":\"proposal\",\"type\":\"string\"}],\"outputs\":[],\"name\":\"changeGatewayAddr\",\"payable\":false},{\"stateMutability\":\"nonpayable\",\"constant\":false,\"type\":\"function\",\"inputs\":[{\"name\":\"proposal\",\"type\":\"string\"}],\"outputs\":[],\"name\":\"stop\",\"payable\":false},{\"stateMutability\":\"view\",\"constant\":true,\"type\":\"function\",\"inputs\":[{\"name\":\"code\",\"type\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"getAppAddress\",\"payable\":false},{\"stateMutability\":\"view\",\"constant\":true,\"type\":\"function\",\"inputs\":[{\"name\":\"code\",\"type\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"name\":\"isAppCode\",\"payable\":false},{\"stateMutability\":\"view\",\"constant\":true,\"type\":\"function\",\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"name\":\"isCaller\",\"payable\":false},{\"stateMutability\":\"view\",\"constant\":true,\"type\":\"function\",\"inputs\":[{\"name\":\"code\",\"type\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\"}],\"name\":\"getAppTokenCode\",\"payable\":false},{\"stateMutability\":\"nonpayable\",\"constant\":false,\"type\":\"function\",\"inputs\":[{\"name\":\"code\",\"type\":\"uint32\"},{\"name\":\"proposal\",\"type\":\"string\"}],\"outputs\":[],\"name\":\"removeApp\",\"payable\":false},{\"stateMutability\":\"view\",\"constant\":true,\"type\":\"function\",\"inputs\":[{\"name\":\"code\",\"type\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\"}],\"name\":\"getAppChainCode\",\"payable\":false},{\"stateMutability\":\"view\",\"constant\":true,\"type\":\"function\",\"inputs\":[{\"name\":\"code\",\"type\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint32\"},{\"name\":\"\",\"type\":\"uint32\"}],\"name\":\"getAppInfo\",\"payable\":false},{\"stateMutability\":\"view\",\"constant\":true,\"type\":\"function\",\"inputs\":[{\"name\":\"voter\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"name\":\"isVoter\",\"payable\":false},{\"stateMutability\":\"nonpayable\",\"constant\":false,\"type\":\"function\",\"inputs\":[{\"name\":\"chain\",\"type\":\"string\"},{\"name\":\"proposal\",\"type\":\"string\"}],\"outputs\":[],\"name\":\"addChain\",\"payable\":false},{\"stateMutability\":\"view\",\"constant\":true,\"type\":\"function\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"name\":\"mStopped\",\"payable\":false},{\"stateMutability\":\"nonpayable\",\"constant\":false,\"type\":\"function\",\"inputs\":[{\"name\":\"appCode\",\"type\":\"uint32\"},{\"name\":\"wad\",\"type\":\"uint64\"},{\"name\":\"receiver\",\"type\":\"address\"},{\"name\":\"proposal\",\"type\":\"string\"}],\"outputs\":[],\"name\":\"mintByGateway\",\"payable\":false},{\"stateMutability\":\"view\",\"constant\":true,\"type\":\"function\",\"inputs\":[{\"name\":\"app\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"name\":\"isApper\",\"payable\":false},{\"stateMutability\":\"view\",\"constant\":true,\"type\":\"function\",\"inputs\":[{\"name\":\"chain\",\"type\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\"}],\"name\":\"getChainCode\",\"payable\":false},{\"stateMutability\":\"view\",\"constant\":true,\"type\":\"function\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"mNumVoters\",\"payable\":false},{\"stateMutability\":\"nonpayable\",\"constant\":false,\"type\":\"function\",\"inputs\":[{\"name\":\"older\",\"type\":\"address\"},{\"name\":\"newer\",\"type\":\"address\"},{\"name\":\"proposal\",\"type\":\"string\"}],\"outputs\":[],\"name\":\"changeVoter\",\"payable\":false},{\"stateMutability\":\"nonpayable\",\"constant\":false,\"type\":\"function\",\"inputs\":[{\"name\":\"older\",\"type\":\"address\"},{\"name\":\"proposal\",\"type\":\"string\"}],\"outputs\":[],\"name\":\"removeVoter\",\"payable\":false},{\"stateMutability\":\"nonpayable\",\"constant\":false,\"type\":\"function\",\"inputs\":[{\"name\":\"newer\",\"type\":\"address\"},{\"name\":\"proposal\",\"type\":\"string\"}],\"outputs\":[],\"name\":\"addVoter\",\"payable\":false},{\"stateMutability\":\"view\",\"constant\":true,\"type\":\"function\",\"inputs\":[{\"name\":\"chain\",\"type\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"name\":\"isChain\",\"payable\":false},{\"type\":\"constructor\",\"inputs\":[{\"name\":\"voters\",\"type\":\"address[]\"}],\"stateMutability\":\"nonpayable\",\"payable\":false},{\"anonymous\":false,\"type\":\"event\",\"inputs\":[{\"indexed\":true,\"type\":\"uint256\",\"name\":\"operation\"}],\"name\":\"Stopped\"},{\"anonymous\":false,\"type\":\"event\",\"inputs\":[{\"indexed\":true,\"type\":\"uint256\",\"name\":\"operation\"}],\"name\":\"Started\"},{\"anonymous\":false,\"type\":\"event\",\"inputs\":[{\"indexed\":false,\"type\":\"address\",\"name\":\"voter\"},{\"indexed\":true,\"type\":\"uint256\",\"name\":\"operation\"}],\"name\":\"Confirmation\"},{\"anonymous\":false,\"type\":\"event\",\"inputs\":[{\"indexed\":false,\"type\":\"address\",\"name\":\"voter\"},{\"indexed\":true,\"type\":\"uint256\",\"name\":\"operation\"}],\"name\":\"OperationDone\"},{\"anonymous\":false,\"type\":\"event\",\"inputs\":[{\"indexed\":false,\"type\":\"address\",\"name\":\"revoker\"},{\"indexed\":true,\"type\":\"uint256\",\"name\":\"operation\"}],\"name\":\"Revoke\"},{\"anonymous\":false,\"type\":\"event\",\"inputs\":[{\"indexed\":false,\"type\":\"address\",\"name\":\"oldVoter\"},{\"indexed\":false,\"type\":\"address\",\"name\":\"newVoter\"},{\"indexed\":true,\"type\":\"uint256\",\"name\":\"operation\"}],\"name\":\"VoterChanged\"},{\"anonymous\":false,\"type\":\"event\",\"inputs\":[{\"indexed\":false,\"type\":\"address\",\"name\":\"newVoter\"},{\"indexed\":true,\"type\":\"uint256\",\"name\":\"operation\"}],\"name\":\"VoterAdded\"},{\"anonymous\":false,\"type\":\"event\",\"inputs\":[{\"indexed\":false,\"type\":\"address\",\"name\":\"oldVoter\"},{\"indexed\":true,\"type\":\"uint256\",\"name\":\"operation\"}],\"name\":\"VoterRemoved\"},{\"anonymous\":false,\"type\":\"event\",\"inputs\":[{\"indexed\":false,\"type\":\"string\",\"name\":\"chain\"},{\"indexed\":true,\"type\":\"uint256\",\"name\":\"operation\"}],\"name\":\"ChainAdded\"},{\"anonymous\":false,\"type\":\"event\",\"inputs\":[{\"indexed\":false,\"type\":\"address\",\"name\":\"app\"},{\"indexed\":false,\"type\":\"uint32\",\"name\":\"chain\"},{\"indexed\":false,\"type\":\"uint32\",\"name\":\"token\"},{\"indexed\":true,\"type\":\"uint256\",\"name\":\"operation\"}],\"name\":\"AppAdded\"},{\"anonymous\":false,\"type\":\"event\",\"inputs\":[{\"indexed\":false,\"type\":\"uint32\",\"name\":\"code\"},{\"indexed\":true,\"type\":\"uint256\",\"name\":\"operation\"}],\"name\":\"AppRemoved\"},{\"anonymous\":false,\"type\":\"event\",\"inputs\":[{\"indexed\":false,\"type\":\"uint32\",\"name\":\"appCode\"},{\"indexed\":false,\"type\":\"address\",\"name\":\"receiver\"},{\"indexed\":false,\"type\":\"uint64\",\"name\":\"wad\"},{\"indexed\":true,\"type\":\"uint256\",\"name\":\"operation\"}],\"name\":\"MintByGateway\"},{\"anonymous\":false,\"type\":\"event\",\"inputs\":[{\"indexed\":false,\"type\":\"uint32\",\"name\":\"appCode\"},{\"indexed\":false,\"type\":\"address\",\"name\":\"from\"},{\"indexed\":false,\"type\":\"string\",\"name\":\"receiver\"},{\"indexed\":false,\"type\":\"uint64\",\"name\":\"wad\"}],\"name\":\"BurnForGateway\"},{\"anonymous\":false,\"type\":\"event\",\"inputs\":[{\"indexed\":false,\"type\":\"uint32\",\"name\":\"appCode\"},{\"indexed\":false,\"type\":\"address\",\"name\":\"newer\"},{\"indexed\":true,\"type\":\"uint256\",\"name\":\"operation\"}],\"name\":\"GatewayAddrChanged\"}]"

// GatewayVoteCaller is an auto generated read-only Go binding around an Ethereum contract.
type GatewayVoteCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NewGatewayVoteCaller creates a new read-only instance of GatewayVote, bound to a specific deployed contract.
func NewGatewayVoteCaller(address common.Address, caller bind.ContractCaller) (*GatewayVoteCaller, error) {
	contract, err := bindGatewayVote(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GatewayVoteCaller{contract: contract}, nil
}

// bindGatewayVote binds a generic wrapper to an already deployed contract.
func bindGatewayVote(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GatewayVoteABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
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
