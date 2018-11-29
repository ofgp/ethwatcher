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
const GatewayVoteABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"targetChain\",\"type\":\"string\"},{\"name\":\"targetAddr\",\"type\":\"string\"}],\"name\":\"recvEther\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"operation\",\"type\":\"uint256\"},{\"name\":\"voter\",\"type\":\"address\"}],\"name\":\"hasConfirmed\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"code\",\"type\":\"uint32\"}],\"name\":\"isChainCode\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"operation\",\"type\":\"uint256\"}],\"name\":\"revoke\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"proposal\",\"type\":\"string\"}],\"name\":\"start\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"app\",\"type\":\"address\"},{\"name\":\"chain\",\"type\":\"uint32\"},{\"name\":\"token\",\"type\":\"uint32\"},{\"name\":\"proposal\",\"type\":\"string\"}],\"name\":\"addApp\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"receiver\",\"type\":\"string\"},{\"name\":\"wad\",\"type\":\"uint64\"}],\"name\":\"burnForGateway\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"mMaxAppCode\",\"outputs\":[{\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"code\",\"type\":\"uint32\"}],\"name\":\"getChainName\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"app\",\"type\":\"address\"}],\"name\":\"getAppCode\",\"outputs\":[{\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"mMaxChainCode\",\"outputs\":[{\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"proposal\",\"type\":\"string\"}],\"name\":\"sendEther\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"appCode\",\"type\":\"uint32\"},{\"name\":\"newer\",\"type\":\"address\"},{\"name\":\"proposal\",\"type\":\"string\"}],\"name\":\"changeGatewayAddr\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"proposal\",\"type\":\"string\"}],\"name\":\"stop\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"code\",\"type\":\"uint32\"}],\"name\":\"getAppAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"code\",\"type\":\"uint32\"}],\"name\":\"isAppCode\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isCaller\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"code\",\"type\":\"uint32\"}],\"name\":\"getAppTokenCode\",\"outputs\":[{\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"code\",\"type\":\"uint32\"},{\"name\":\"proposal\",\"type\":\"string\"}],\"name\":\"removeApp\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"code\",\"type\":\"uint32\"}],\"name\":\"getAppChainCode\",\"outputs\":[{\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"code\",\"type\":\"uint32\"}],\"name\":\"getAppInfo\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint32\"},{\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"voter\",\"type\":\"address\"}],\"name\":\"isVoter\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"chain\",\"type\":\"string\"},{\"name\":\"proposal\",\"type\":\"string\"}],\"name\":\"addChain\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"mStopped\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"appCode\",\"type\":\"uint32\"},{\"name\":\"wad\",\"type\":\"uint64\"},{\"name\":\"receiver\",\"type\":\"address\"},{\"name\":\"proposal\",\"type\":\"string\"}],\"name\":\"mintByGateway\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"app\",\"type\":\"address\"}],\"name\":\"isApper\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"chain\",\"type\":\"string\"}],\"name\":\"getChainCode\",\"outputs\":[{\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"mNumVoters\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"older\",\"type\":\"address\"},{\"name\":\"newer\",\"type\":\"address\"},{\"name\":\"proposal\",\"type\":\"string\"}],\"name\":\"changeVoter\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"older\",\"type\":\"address\"},{\"name\":\"proposal\",\"type\":\"string\"}],\"name\":\"removeVoter\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newer\",\"type\":\"address\"},{\"name\":\"proposal\",\"type\":\"string\"}],\"name\":\"addVoter\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"chain\",\"type\":\"string\"}],\"name\":\"isChain\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"voters\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"operation\",\"type\":\"uint256\"}],\"name\":\"Stopped\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"operation\",\"type\":\"uint256\"}],\"name\":\"Started\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"operation\",\"type\":\"uint256\"}],\"name\":\"Confirmation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"operation\",\"type\":\"uint256\"}],\"name\":\"OperationDone\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"revoker\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"operation\",\"type\":\"uint256\"}],\"name\":\"Revoke\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"oldVoter\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"newVoter\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"operation\",\"type\":\"uint256\"}],\"name\":\"VoterChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"newVoter\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"operation\",\"type\":\"uint256\"}],\"name\":\"VoterAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"oldVoter\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"operation\",\"type\":\"uint256\"}],\"name\":\"VoterRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"chain\",\"type\":\"string\"},{\"indexed\":true,\"name\":\"operation\",\"type\":\"uint256\"}],\"name\":\"ChainAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"app\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"chain\",\"type\":\"uint32\"},{\"indexed\":false,\"name\":\"token\",\"type\":\"uint32\"},{\"indexed\":true,\"name\":\"operation\",\"type\":\"uint256\"}],\"name\":\"AppAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"code\",\"type\":\"uint32\"},{\"indexed\":true,\"name\":\"operation\",\"type\":\"uint256\"}],\"name\":\"AppRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"appCode\",\"type\":\"uint32\"},{\"indexed\":false,\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"wad\",\"type\":\"uint64\"},{\"indexed\":true,\"name\":\"operation\",\"type\":\"uint256\"}],\"name\":\"MintByGateway\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"appCode\",\"type\":\"uint32\"},{\"indexed\":false,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"receiver\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"wad\",\"type\":\"uint64\"}],\"name\":\"BurnForGateway\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"appCode\",\"type\":\"uint32\"},{\"indexed\":false,\"name\":\"newer\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"operation\",\"type\":\"uint256\"}],\"name\":\"GatewayAddrChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"targetChain\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"targetAddr\",\"type\":\"string\"}],\"name\":\"RecvEther\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"proposal\",\"type\":\"string\"}],\"name\":\"SendEther\",\"type\":\"event\"}]"

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
