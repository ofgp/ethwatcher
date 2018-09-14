// Copyright 2016 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

// Package ethclient provides a client for the Ethereum RPC API.
package ethwatcher

import (
	"context"
	// "encoding/hex"
	"encoding/json"
	// "errors"
	"fmt"
	"math/big"
	"reflect"
	"strings"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	// "github.com/ethereum/go-ethereum/crypto/sha3"
	"github.com/ethereum/go-ethereum/ethdb"
	// "github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/spf13/viper"
	ldberrors "github.com/syndtr/goleveldb/leveldb/errors"
)

func init() {
	viper.SetDefault("LEVELDB.ew_nonce_db_path", "./nonceStore")
	viper.SetDefault("ETHWATCHER.vote_contract", "")
	viper.SetDefault("ETHWATCHER.min_gasprice", 1000000000)
	viper.SetDefault("ETHWATCHER.max_gasprice", 80000000000)
}

var (
	voteAbi, _  = abi.JSON(strings.NewReader(GatewayVoteABI))
	ewLogger    = NewLogger("debug", "ethwatcher")
	minGasPrice = new(big.Int).SetInt64(1000000000)
	maxGasPrice = new(big.Int).SetInt64(80000000000)
)

// Client defines typed wrappers for the Ethereum RPC API.
type Client struct {
	nonceLock        *sync.Mutex
	c                *rpc.Client
	confirmHeight    *big.Int
	currentHeight    *big.Int
	ldb              *ethdb.LDBDatabase
	mapAppCodeToInfo *sync.Map
	mapAppAddrToCode *sync.Map
	mapCodeToChain   *sync.Map
	mapChainToCode   *sync.Map
	mapContracts     *sync.Map
	voteContract     common.Address
	voteEvents       map[string]string
}

func NewEthWatcher(rawurl string, confirmHeight int64, sPub string) (*Client, error) {
	return Dial(rawurl, confirmHeight, sPub)
}

// Dial connects a client to the given URL.
func Dial(rawurl string, confirmHeight int64, sPub string) (*Client, error) {
	return DialContext(context.Background(), rawurl, confirmHeight, sPub)
}

func DialContext(ctx context.Context, rawurl string, confirmHeight int64, sPub string) (*Client, error) {
	c, err := rpc.DialContext(ctx, rawurl)
	if err != nil {
		return nil, err
	}
	if confirmHeight < int64(3) {
		confirmHeight = int64(3)
	}
	return NewClient(c, confirmHeight, sPub)
}

// NewClient creates a client that uses the given RPC client.
func NewClient(c *rpc.Client, confirmHeight int64, sPub string) (ec *Client, err error) {
	ldb, err := ethdb.NewLDBDatabase(viper.GetString("LEVELDB.ew_nonce_db_path"), 16, 16)
	if err != nil {
		return nil, err
	}

	ewLogger = NewLogger(viper.GetString("loglevel"), "ethwatcher")
	minGasPrice.SetInt64(viper.GetInt64("ETHWATCHER.min_gasprice"))
	maxGasPrice.SetInt64(viper.GetInt64("ETHWATCHER.max_gasprice"))

	ec = &Client{
		nonceLock:        new(sync.Mutex),
		c:                c,
		confirmHeight:    new(big.Int).SetInt64(confirmHeight),
		currentHeight:    new(big.Int).SetInt64(0),
		ldb:              ldb,
		mapAppAddrToCode: new(sync.Map),
		mapAppCodeToInfo: new(sync.Map),
		mapChainToCode:   new(sync.Map),
		mapCodeToChain:   new(sync.Map),
		mapContracts:     new(sync.Map),
		voteContract:     common.HexToAddress(viper.GetString("ETHWATCHER.vote_contract")),
	}

	idMapName := make(map[string]string)

	for _, event := range voteAbi.Events {
		idMapName[event.Id().Hex()] = event.Name
	}
	ec.voteEvents = idMapName

	ec.ReviewNonce(sPub)

	return ec, nil
}

func (ec *Client) Close() {
	ec.c.Close()
	ec.ldb.Close()
}

func (ec *Client) ReviewBlock(start big.Int, ch chan<- *big.Int) {
	cur, err := ec.BlockNumber()
	if err != nil {
		panic("ethwatcher ReviewBlock, get current height error!")
	}
	ec.currentHeight = cur

	ewLogger.Debug("ethwatcher review block range", "from", start.Uint64(), "to", cur.Uint64())

	curConfirm := new(big.Int).Sub(cur, ec.confirmHeight)
	increaser := big.NewInt(1)
	go func() {
		for i := &start; i.Cmp(curConfirm) <= 0; i = new(big.Int).Add(i, increaser) {
			select {
			case ch <- i:
			}
		}
		close(ch)
	}()
}

func (ec *Client) StartWatchBlock(start big.Int, heightCh chan<- *big.Int) {
	wCh := make(chan *rpcHeader, 1000)
	rCh := make(chan *big.Int, 1000)
	sub, err := ec.SubscribeNewHead(ensureContext(nil), wCh)
	if err != nil {
		panic(fmt.Sprintf("ethwatcher StartWatch, subscribe new block error! e is %v!", err.Error()))
	}
	ec.ReviewBlock(start, rCh)
	bigConfirmH := ec.confirmHeight
	increaser := big.NewInt(1)
	go func() {
		defer sub.Unsubscribe()
		for blockHeight := range rCh {
			select {
			case heightCh <- blockHeight:
			}
		}
		for {
			select {
			case blockHeader := <-wCh:
				bigIntNumber := (*big.Int)(blockHeader.Number)
				if bigIntNumber.Cmp(ec.currentHeight) > 0 {
					startH := new(big.Int).Add(ec.currentHeight, increaser)
					for i := startH; i.Cmp(bigIntNumber) <= 0; i = new(big.Int).Add(i, increaser) {
						pushH := new(big.Int).Sub(i, bigConfirmH)
						select {
						case heightCh <- pushH:
							ewLogger.Debug("ethwatcher watched block", "height", pushH.Uint64())
						}
					}
					ec.currentHeight = bigIntNumber
				} else {
					ewLogger.Warn("ethwatcher receive under current block", "height", bigIntNumber.Uint64(), "current", ec.currentHeight.Uint64())
				}
			case err := <-sub.Err():
				ewLogger.Error("ethwatcher watch block error", "error", err.Error())
			}
		}
	}()
}

func (ec *Client) GetAppCode(sAppAddr string) uint32 {
	appCode, ok := ec.mapAppAddrToCode.Load(common.HexToAddress(sAppAddr).Hex())
	if !ok {
		ewLogger.Warn("ethwatcher GetAppCode, bad app addr", "hash", "appAddr", sAppAddr)
		return uint32(0)
	}

	return appCode.(uint32)
}

func (ec *Client) GetChainCode(sChainName string) uint32 {
	chainCode, ok := ec.mapChainToCode.Load(sChainName)
	if !ok {
		ewLogger.Warn("ethwatcher GetChainCode, bad chain name", "chain", sChainName)
		return uint32(0)
	}
	return chainCode.(uint32)
}

func (ec *Client) VerifyAppInfo(sChain string, tokenCode uint32, appCode uint32) bool {
	iAppInfo, ok := ec.mapAppCodeToInfo.Load(appCode)
	if !ok {
		ewLogger.Warn("ethwatcher VarifyAppInfo, no appinfo for appCode", "appCode", appCode)
		return false
	}
	oAppInfo := iAppInfo.(*AppInfo)
	if oAppInfo.Chain == sChain && oAppInfo.TokenCode == tokenCode {
		return true
	}
	return false

}

func (ec *Client) InitMapInfo() {
	voteCaller, err := NewGatewayVoteCaller(ec.voteContract, ec)
	if err != nil {
		panic(fmt.Sprintf("ethwatcher InitAppInfo, get contract caller error! e is %v \n", err.Error()))
	}
	ec.mapContracts.Store(ec.voteContract.Hex(), true)

	maxChainCode, err := voteCaller.MMaxChainCode(nil)
	if err != nil {
		panic(fmt.Sprintf("ethwatcher InitAppInfo, get max chain code error! e is %v \n", err.Error()))
	}

	var i uint32

	for i = 1; i <= maxChainCode; i++ {
		chain, err := voteCaller.GetChainName(nil, i)
		if err != nil {
			panic(fmt.Sprintf("ethwatcher InitAppInfo, get chain name error! chain code is %v, e is %v \n", i, err.Error()))
		}
		ec.mapChainToCode.Store(chain, i)
		ec.mapCodeToChain.Store(i, chain)
	}

	maxAppCode, err := voteCaller.MMaxAppCode(nil)
	if err != nil {
		panic(fmt.Sprintf("ethwatcher InitAppInfo, get max app code error! e is %v \n", err.Error()))
	}

	for i = 1; i <= maxAppCode; i++ {
		if isAppCode, err := voteCaller.IsAppCode(nil, i); err != nil {
			panic(fmt.Sprintf("ethwatcher InitAppInfo, call IsAppCode error! e is %v \n", err.Error()))

		} else if !isAppCode {
			continue
		}

		addr, chainCode, tokenCode, err := voteCaller.GetAppInfo(nil, i)
		if err != nil {
			panic(fmt.Sprintf("ethwatcher InitAppInfo, get app info  error! e is %v \n", err.Error()))
		}
		ewLogger.Info("find app", "addr", addr.Hex(), "chainCode", chainCode, "tokenCode", tokenCode)

		ichain, ok := ec.mapCodeToChain.Load(chainCode)
		if !ok {
			panic(fmt.Sprintf("ethwatcher InitAppInfo, find chain name from map error! chain code is %v, e is %v \n", chainCode, err.Error()))
		}

		appInfo := AppInfo{AppCode: i, ChainCode: chainCode, TokenCode: tokenCode, Addr: addr, Chain: reflect.ValueOf(ichain).Interface().(string)}
		ec.mapAppCodeToInfo.Store(i, &appInfo)
		ec.mapAppAddrToCode.Store(addr.Hex(), i)
		ec.mapContracts.Store(addr.Hex(), true)
	}

}

func (ec *Client) parseTxEventStatus(tx *rpcTx, r *types.Receipt) uint64 {
	var eventStatus uint64 = VOTE_STATUS_OTHER

	if r == nil {
		eventStatus |= TX_STATUS_PENDING // pending tx

	} else if r.Status == uint(0) {
		eventStatus |= TX_STATUS_FAILED // tx excuted failed

	} else {
		for _, txLog := range r.Logs {
			if txLog.Address.Hex() != ec.voteContract.Hex() {
				continue
			}
			logEventName := ec.voteEvents[txLog.Topics[0].Hex()]
			switch logEventName {

			// VOTE_EVENT_MINT          = "MintByGateway"
			case VOTE_EVENT_MINT:
				eventStatus |= VOTE_TX_MINT

			// VOTE_EVENT_BURN          = "BurnForGateway"
			case VOTE_EVENT_CONFIRM:
				eventStatus |= VOTE_TX_CONFIRM

			// VOTE_EVENT_CONFIRM       = "Confirmation"
			case VOTE_EVENT_BURN:
				eventStatus |= TOKEN_TX_BURN

			// VOTE_EVENT_STARTED       = "Started"
			case VOTE_EVENT_STARTED:
				eventStatus |= VOTE_TX_STARTED

			// VOTE_EVENT_STOPPED       = "Stopped"
			case VOTE_EVENT_STOPPED:
				eventStatus |= VOTE_TX_STOPPED

			// VOTE_EVENT_OPTDONE       = "OperationDone"
			case VOTE_EVENT_OPTDONE:
				eventStatus |= VOTE_TX_OPTDONE

			// VOTE_EVENT_REVOKE        = "Revoke"
			case VOTE_EVENT_REVOKE:
				eventStatus |= VOTER_TX_REVOKE

			// VOTE_EVENT_VOTERCHANGED  = "VoterChanged"
			case VOTE_EVENT_VOTERCHANGED:
				eventStatus |= VOTE_TX_VOTERCHANGED

			// VOTE_EVNET_VOTERADDED    = "VoterAdded"
			case VOTE_EVENT_VOTERADDED:
				eventStatus |= VOTE_TX_VOTERADDED

			// VOTE_EVNET_VOTERREMOVED  = "VoterRemoved"
			case VOTE_EVENT_VOTERREMOVED:
				eventStatus |= VOTE_TX_VOTERREMOVED

			// VOTE_EVENT_CHAINADDED    = "ChainAdded"
			case VOTE_EVENT_CHAINADDED:
				eventStatus |= VOTE_TX_CHAINADDED

			// VOTE_EVENT_APPADDED      = "AppAdded"
			case VOTE_EVENT_APPADDED:
				eventStatus |= VOTE_TX_APPADDED

			// VOTE_EVENT_APPREMOVED    = "AppRemoved"
			case VOTE_EVENT_APPREMOVED:
				eventStatus |= VOTE_TX_APPREMOVED

			// VOTE_EVENT_CHANGEGATEWAY = "GatewayAddrChanged"
			case VOTE_EVENT_CHANGEGATEWAY:
				eventStatus |= VOTE_TX_GATEWAYCHANGED
			}
		}
	}
	return eventStatus
}

func (ec *Client) parseBurnTx(tx *rpcTx, method *abi.Method) (*PushEvent, error) {
	var args burnCallArguments

	err := method.Inputs.Unpack(&args, []byte(*tx.Payload)[4:])
	if err != nil {
		ewLogger.Error("ethwatcher parseBurnTx, Unpack input error!", "hash", tx.TxHash.Hex(), "error", err.Error())
		return nil, err
	}

	var r *types.Receipt
	var confirmations int64 = int64(0)
	if tx.BlockNumber != nil {
		ctx := ensureContext(nil)
		r, err = ec.TransactionReceipt(ctx, *tx.TxHash)
		if err != nil {
			ewLogger.Error("ethwatcher parseChangeGatewayTx, get tx receipt error", "hash", tx.TxHash.Hex(), "error", err.Error())
			return nil, err
		}
		confirmations = new(big.Int).Sub(ec.currentHeight, (*big.Int)(tx.BlockNumber)).Int64()
	}

	var extraData ExtraBurnData
	extraData.ScTxid = tx.TxHash.Hex()

	appCode, ok := ec.mapAppAddrToCode.Load(tx.CallTo.Hex())
	if !ok {
		ewLogger.Error("ethwatcher parseBurnTx, bad app addr", "hash", tx.TxHash.Hex(), "app", tx.CallTo.Hex())
		return nil, fmt.Errorf("ethwatcher parseBurnTx, bad app addr!")
	}

	appInfo, ok := ec.mapAppCodeToInfo.Load(appCode.(uint32))
	if !ok {
		ewLogger.Error("ethwatcher parseBurnTx, bad app code", "hash", tx.TxHash.Hex(), "app", tx.CallTo.Hex(), "appCode", appCode)
		return nil, fmt.Errorf("ethwatcher parseBurnTx, bad app code!")
	}
	nApp := appInfo.(*AppInfo)

	extraData.From = CHAIN_SELF_ETH
	extraData.To = nApp.Chain
	extraData.TokenTo = nApp.TokenCode
	extraData.TokenFrom = nApp.AppCode
	extraData.Amount = args.Wad

	recharge := AssetInfo{Amount: args.Wad, Address: args.Receiver}
	extraData.RechargeList = []*AssetInfo{&recharge}

	return &PushEvent{
		Operation:     new(big.Int).SetBytes(crypto.Keccak256(*tx.Payload)),
		Tx:            createTxInfo(tx, r),
		Confirmations: confirmations,
		Method:        method.Name,
		Events:        ec.parseTxEventStatus(tx, r),
		ExtraData:     &extraData,
	}, nil

}

func (ec *Client) parseMintTx(tx *rpcTx, method *abi.Method) (*PushEvent, error) {

	var extraData ExtraMintData

	err := method.Inputs.Unpack(&extraData, []byte(*tx.Payload)[4:])
	if err != nil {
		ewLogger.Error("ethwatcher parseMintTx, Unpack input error!", "hash", tx.TxHash.Hex(), "error", err.Error())
		return nil, err
	}

	var r *types.Receipt
	var confirmations int64 = int64(0)
	if tx.BlockNumber != nil {
		ctx := ensureContext(nil)
		r, err = ec.TransactionReceipt(ctx, *tx.TxHash)
		if err != nil {
			ewLogger.Error("ethwatcher parseChangeGatewayTx, get tx receipt error", "hash", tx.TxHash.Hex(), "error", err.Error())
			return nil, err
		}
		confirmations = new(big.Int).Sub(ec.currentHeight, (*big.Int)(tx.BlockNumber)).Int64()
	}

	return &PushEvent{
		Operation:     new(big.Int).SetBytes(crypto.Keccak256(*tx.Payload)),
		Tx:            createTxInfo(tx, r),
		Confirmations: confirmations,
		Method:        method.Name,
		Events:        ec.parseTxEventStatus(tx, r),
		ExtraData:     &extraData,
	}, nil

}

func (ec *Client) parseStartTx(tx *rpcTx, method *abi.Method) (*PushEvent, error) {

	var extraData ExtraStartedData

	err := method.Inputs.Unpack(&extraData, []byte(*tx.Payload)[4:])
	if err != nil {
		ewLogger.Error("ethwatcher parseStartTx, Unpack input error!", "hash", tx.TxHash.Hex(), "error", err.Error())
		return nil, err
	}

	var r *types.Receipt
	var confirmations int64 = int64(0)
	if tx.BlockNumber != nil {
		ctx := ensureContext(nil)
		r, err = ec.TransactionReceipt(ctx, *tx.TxHash)
		if err != nil {
			ewLogger.Error("ethwatcher parseChangeGatewayTx, get tx receipt error", "hash", tx.TxHash.Hex(), "error", err.Error())
			return nil, err
		}
		confirmations = new(big.Int).Sub(ec.currentHeight, (*big.Int)(tx.BlockNumber)).Int64()
	}

	return &PushEvent{
		Operation:     new(big.Int).SetBytes(crypto.Keccak256(*tx.Payload)),
		Tx:            createTxInfo(tx, r),
		Confirmations: confirmations,
		Method:        method.Name,
		Events:        ec.parseTxEventStatus(tx, r),
		ExtraData:     &extraData,
	}, nil

}

func (ec *Client) parseStopTx(tx *rpcTx, method *abi.Method) (*PushEvent, error) {

	var extraData ExtraStoppedData

	err := method.Inputs.Unpack(&extraData, []byte(*tx.Payload)[4:])
	if err != nil {
		ewLogger.Error("ethwatcher parseStopTx, Unpack input error!", "hash", tx.TxHash.Hex(), "error", err.Error())
		return nil, err
	}

	var r *types.Receipt
	var confirmations int64 = int64(0)
	if tx.BlockNumber != nil {
		ctx := ensureContext(nil)
		r, err = ec.TransactionReceipt(ctx, *tx.TxHash)
		if err != nil {
			ewLogger.Error("ethwatcher parseChangeGatewayTx, get tx receipt error", "hash", tx.TxHash.Hex(), "error", err.Error())
			return nil, err
		}
		confirmations = new(big.Int).Sub(ec.currentHeight, (*big.Int)(tx.BlockNumber)).Int64()
	}

	return &PushEvent{
		Operation:     new(big.Int).SetBytes(crypto.Keccak256(*tx.Payload)),
		Tx:            createTxInfo(tx, r),
		Confirmations: confirmations,
		Method:        method.Name,
		Events:        ec.parseTxEventStatus(tx, r),
		ExtraData:     &extraData,
	}, nil

}

func (ec *Client) parseRevokeTx(tx *rpcTx, method *abi.Method) (*PushEvent, error) {

	var extraData ExtraRevokeData

	err := method.Inputs.Unpack(&extraData, []byte(*tx.Payload)[4:])
	if err != nil {
		ewLogger.Error("ethwatcher parseRevokeTx, Unpack input error!", "hash", tx.TxHash.Hex(), "error", err.Error())
		return nil, err
	}

	var r *types.Receipt
	var confirmations int64 = int64(0)
	if tx.BlockNumber != nil {
		ctx := ensureContext(nil)
		r, err = ec.TransactionReceipt(ctx, *tx.TxHash)
		if err != nil {
			ewLogger.Error("ethwatcher parseChangeGatewayTx, get tx receipt error", "hash", tx.TxHash.Hex(), "error", err.Error())
			return nil, err
		}
		confirmations = new(big.Int).Sub(ec.currentHeight, (*big.Int)(tx.BlockNumber)).Int64()
	}

	return &PushEvent{
		Operation:     new(big.Int).SetBytes(crypto.Keccak256(*tx.Payload)),
		Tx:            createTxInfo(tx, r),
		Confirmations: confirmations,
		Method:        method.Name,
		Events:        ec.parseTxEventStatus(tx, r),
		ExtraData:     &extraData,
	}, nil

}

func (ec *Client) parseAddAppTx(tx *rpcTx, method *abi.Method) (*PushEvent, error) {

	var extraData ExtraAppAddedData

	err := method.Inputs.Unpack(&extraData, []byte(*tx.Payload)[4:])
	if err != nil {
		ewLogger.Error("ethwatcher parseAddAppTx, Unpack input error!", "hash", tx.TxHash.Hex(), "error", err.Error())
		return nil, err
	}

	var r *types.Receipt
	var confirmations int64 = int64(0)
	if tx.BlockNumber != nil {
		ctx := ensureContext(nil)
		r, err = ec.TransactionReceipt(ctx, *tx.TxHash)
		if err != nil {
			ewLogger.Error("ethwatcher parseChangeGatewayTx, get tx receipt error", "hash", tx.TxHash.Hex(), "error", err.Error())
			return nil, err
		}
		confirmations = new(big.Int).Sub(ec.currentHeight, (*big.Int)(tx.BlockNumber)).Int64()
	}

	eventS := ec.parseTxEventStatus(tx, r)
	if eventS&VOTE_TX_APPADDED == VOTE_TX_APPADDED {

		ewLogger.Info("ethwatcher parseAddAppTx, received add app tx", "hash", tx.TxHash.Hex(), "app", extraData.App.Hex())
		iAppCode, ok := ec.mapAppAddrToCode.Load(extraData.App.Hex())
		if !ok {
			voteCaller, err := NewGatewayVoteCaller(ec.voteContract, ec)
			if err != nil {
				panic(fmt.Sprintf("ethwatcher parseAddAppTx, get contract caller error! e is %v \n", err.Error()))
			}

			if isApper, err := voteCaller.IsApper(nil, extraData.App); err != nil {
				panic(fmt.Sprintf("ethwatcher parseAddAppTx, call IsAppCode error! e is %v \n", err.Error()))

			} else if !isApper {
				return nil, fmt.Errorf("bad app address for real addApp tx. tx is %v", tx.TxHash.Hex())
			}

			appCode, err := voteCaller.GetAppCode(nil, extraData.App)
			if err != nil {
				panic(fmt.Sprintf("ethwatcher parseAddAppTx, get app code error! e is %v \n", err.Error()))
			}

			ichain, ok := ec.mapCodeToChain.Load(extraData.Chain)
			if !ok {
				panic(fmt.Sprintf("ethwatcher parseAddAppTx, find chain name from map error! chain code is %v, e is %v \n", extraData.Chain, err.Error()))
			}

			appInfo := AppInfo{
				AppCode:   appCode,
				ChainCode: extraData.Chain,
				TokenCode: extraData.Token,
				Addr:      extraData.App,
				Chain:     ichain.(string),
			}
			ec.mapAppCodeToInfo.Store(appCode, &appInfo)
			ec.mapAppAddrToCode.Store(extraData.App.Hex(), appCode)
			ec.mapContracts.Store(extraData.App.Hex(), true)

			ewLogger.Info("ethwatcher parseAddAppTx, add app to memory success", "hash", tx.TxHash.Hex(), "app", extraData.App.Hex())
		} else {
			ewLogger.Info("ethwatcher parseAddAppTx, app existed in  memory", "hash", tx.TxHash.Hex(), "app", extraData.App.Hex(), "appCode", iAppCode.(uint32))
		}
	}

	return &PushEvent{
		Operation:     new(big.Int).SetBytes(crypto.Keccak256(*tx.Payload)),
		Tx:            createTxInfo(tx, r),
		Confirmations: confirmations,
		Method:        method.Name,
		Events:        eventS,
		ExtraData:     &extraData,
	}, nil

}

func (ec *Client) parseRemoveAppTx(tx *rpcTx, method *abi.Method) (*PushEvent, error) {

	var extraData ExtraAppRemovedData

	err := method.Inputs.Unpack(&extraData, []byte(*tx.Payload)[4:])
	if err != nil {
		ewLogger.Error("ethwatcher parseRemoveAppTx, Unpack input error!", "hash", tx.TxHash.Hex(), "error", err.Error())
		return nil, err
	}

	var r *types.Receipt
	var confirmations int64 = int64(0)
	if tx.BlockNumber != nil {
		ctx := ensureContext(nil)
		r, err = ec.TransactionReceipt(ctx, *tx.TxHash)
		if err != nil {
			ewLogger.Error("ethwatcher parseChangeGatewayTx, get tx receipt error", "hash", tx.TxHash.Hex(), "error", err.Error())
			return nil, err
		}
		confirmations = new(big.Int).Sub(ec.currentHeight, (*big.Int)(tx.BlockNumber)).Int64()
	}

	eventS := ec.parseTxEventStatus(tx, r)
	if eventS&VOTE_TX_APPREMOVED == VOTE_TX_APPREMOVED {
		ewLogger.Info("ethwatcher parseRemoveAppTx, received removed app tx", "hash", tx.TxHash.Hex(), "appCode", extraData.Code)
		iAppInfo, ok := ec.mapAppCodeToInfo.Load(extraData.Code)
		if ok {
			oAppInfo := iAppInfo.(*AppInfo)
			ec.mapAppAddrToCode.Delete(oAppInfo.Addr.Hex())
			ec.mapAppCodeToInfo.Delete(extraData.Code)
			ec.mapContracts.Delete(oAppInfo.Addr.Hex())
			ewLogger.Info("ethwatcher parseRemoveAppTx, remove app from memory suceess", "hash", tx.TxHash.Hex(), "appCode", extraData.Code)
		} else {
			ewLogger.Info("ethwatcher parseRemoveAppTx, app code not in memory", "hash", tx.TxHash.Hex(), "appCode", extraData.Code)
		}
	}

	return &PushEvent{
		Operation:     new(big.Int).SetBytes(crypto.Keccak256(*tx.Payload)),
		Tx:            createTxInfo(tx, r),
		Confirmations: confirmations,
		Method:        method.Name,
		Events:        eventS,
		ExtraData:     &extraData,
	}, nil

}

func (ec *Client) parseAddChainTx(tx *rpcTx, method *abi.Method) (*PushEvent, error) {

	var extraData ExtraChainAddedData

	err := method.Inputs.Unpack(&extraData, []byte(*tx.Payload)[4:])
	if err != nil {
		ewLogger.Error("ethwatcher parseAddChainTx, Unpack input error!", "hash", tx.TxHash.Hex(), "error", err.Error())
		return nil, err
	}

	var r *types.Receipt
	var confirmations int64 = int64(0)
	if tx.BlockNumber != nil {
		ctx := ensureContext(nil)
		r, err = ec.TransactionReceipt(ctx, *tx.TxHash)
		if err != nil {
			ewLogger.Error("ethwatcher parseChangeGatewayTx, get tx receipt error", "hash", tx.TxHash.Hex(), "error", err.Error())
			return nil, err
		}
		confirmations = new(big.Int).Sub(ec.currentHeight, (*big.Int)(tx.BlockNumber)).Int64()
	}

	eventS := ec.parseTxEventStatus(tx, r)
	if eventS&VOTE_TX_CHAINADDED == VOTE_TX_CHAINADDED {
		ewLogger.Info("ethwatcher parseAddChainTx, received chain added tx", "hash", tx.TxHash.Hex(), "chain", extraData.Chain)
		iChainCode, ok := ec.mapChainToCode.Load(extraData.Chain)
		if !ok {
			voteCaller, err := NewGatewayVoteCaller(ec.voteContract, ec)
			if err != nil {
				panic(fmt.Sprintf("ethwatcher error! e is %v \n", err.Error()))
			}

			if isChain, err := voteCaller.IsChain(nil, extraData.Chain); err != nil {
				panic(fmt.Sprintf("ethwatcher error, e is %v \n", err.Error()))

			} else if !isChain {
				ewLogger.Error("ethwatcher find chain added tx, but not a chain in contract", "hash", tx.TxHash.Hex(), "chain", extraData.Chain)
				return nil, fmt.Errorf("bad chain for real add chain tx. tx is %v", tx.TxHash.Hex())
			}

			chainCode, err := voteCaller.GetChainCode(nil, extraData.Chain)
			if err != nil {
				panic(fmt.Sprintf("ethwatcher error! e is %v \n", err.Error()))
			}
			ec.mapChainToCode.Store(extraData.Chain, chainCode)
			ec.mapCodeToChain.Store(chainCode, extraData.Chain)

			ewLogger.Info("ethwatcher parseAddChainTx, add chain to memory success", "hash", tx.TxHash.Hex(), "chain", extraData.Chain)
		} else {
			ewLogger.Info("ethwatcher parseAddChainTx, chain existed in memory", "hash", tx.TxHash.Hex(), "chain", extraData.Chain, "chainCode", iChainCode.(uint32))
		}
	}

	return &PushEvent{
		Operation:     new(big.Int).SetBytes(crypto.Keccak256(*tx.Payload)),
		Tx:            createTxInfo(tx, r),
		Confirmations: confirmations,
		Method:        method.Name,
		Events:        eventS,
		ExtraData:     &extraData,
	}, nil

}

func (ec *Client) parseChangeVoterTx(tx *rpcTx, method *abi.Method) (*PushEvent, error) {

	var extraData ExtraVoterChangedData

	err := method.Inputs.Unpack(&extraData, []byte(*tx.Payload)[4:])
	if err != nil {
		ewLogger.Error("ethwatcher parseChangeVoterTx, Unpack input error!", "hash", tx.TxHash.Hex(), "error", err.Error())
		return nil, err
	}

	var r *types.Receipt
	var confirmations int64 = int64(0)
	if tx.BlockNumber != nil {
		ctx := ensureContext(nil)
		r, err = ec.TransactionReceipt(ctx, *tx.TxHash)
		if err != nil {
			ewLogger.Error("ethwatcher parseChangeGatewayTx, get tx receipt error", "hash", tx.TxHash.Hex(), "error", err.Error())
			return nil, err
		}
		confirmations = new(big.Int).Sub(ec.currentHeight, (*big.Int)(tx.BlockNumber)).Int64()
	}

	return &PushEvent{
		Operation:     new(big.Int).SetBytes(crypto.Keccak256(*tx.Payload)),
		Tx:            createTxInfo(tx, r),
		Confirmations: confirmations,
		Method:        method.Name,
		Events:        ec.parseTxEventStatus(tx, r),
		ExtraData:     &extraData,
	}, nil

}

func (ec *Client) parseAddVoterTx(tx *rpcTx, method *abi.Method) (*PushEvent, error) {

	var extraData ExtraVoterAddedData

	err := method.Inputs.Unpack(&extraData, []byte(*tx.Payload)[4:])
	if err != nil {
		ewLogger.Error("ethwatcher parseAddVoterTx, Unpack input error!", "hash", tx.TxHash.Hex(), "error", err.Error())
		return nil, err
	}

	var r *types.Receipt
	var confirmations int64 = int64(0)
	if tx.BlockNumber != nil {
		ctx := ensureContext(nil)
		r, err = ec.TransactionReceipt(ctx, *tx.TxHash)
		if err != nil {
			ewLogger.Error("ethwatcher parseChangeGatewayTx, get tx receipt error", "hash", tx.TxHash.Hex(), "error", err.Error())
			return nil, err
		}
		confirmations = new(big.Int).Sub(ec.currentHeight, (*big.Int)(tx.BlockNumber)).Int64()
	}

	return &PushEvent{
		Operation:     new(big.Int).SetBytes(crypto.Keccak256(*tx.Payload)),
		Tx:            createTxInfo(tx, r),
		Confirmations: confirmations,
		Method:        method.Name,
		Events:        ec.parseTxEventStatus(tx, r),
		ExtraData:     &extraData,
	}, nil

}

func (ec *Client) parseRemoveVoterTx(tx *rpcTx, method *abi.Method) (*PushEvent, error) {

	var extraData ExtraVoterRemovedData

	err := method.Inputs.Unpack(&extraData, []byte(*tx.Payload)[4:])
	if err != nil {
		ewLogger.Error("ethwatcher parseRemoveVoterTx, Unpack input error!", "hash", tx.TxHash.Hex(), "error", err.Error())
		return nil, err
	}

	var r *types.Receipt
	var confirmations int64 = int64(0)
	if tx.BlockNumber != nil {
		ctx := ensureContext(nil)
		r, err = ec.TransactionReceipt(ctx, *tx.TxHash)
		if err != nil {
			ewLogger.Error("ethwatcher parseChangeGatewayTx, get tx receipt error", "hash", tx.TxHash.Hex(), "error", err.Error())
			return nil, err
		}
		confirmations = new(big.Int).Sub(ec.currentHeight, (*big.Int)(tx.BlockNumber)).Int64()
	}

	return &PushEvent{
		Operation:     new(big.Int).SetBytes(crypto.Keccak256(*tx.Payload)),
		Tx:            createTxInfo(tx, r),
		Confirmations: confirmations,
		Method:        method.Name,
		Events:        ec.parseTxEventStatus(tx, r),
		ExtraData:     &extraData,
	}, nil

}

func (ec *Client) parseChangeGatewayTx(tx *rpcTx, method *abi.Method) (*PushEvent, error) {

	var extraData ExtraGatewayAddrChangedData

	err := method.Inputs.Unpack(&extraData, []byte(*tx.Payload)[4:])
	if err != nil {
		ewLogger.Error("ethwatcher parseChangeGatewayTx, Unpack input error!", "hash", tx.TxHash.Hex(), "error", err.Error())
		return nil, err
	}

	var r *types.Receipt
	var confirmations int64 = int64(0)
	if tx.BlockNumber != nil {
		ctx := ensureContext(nil)
		r, err = ec.TransactionReceipt(ctx, *tx.TxHash)
		if err != nil {
			ewLogger.Error("ethwatcher parseChangeGatewayTx, get tx receipt error", "hash", tx.TxHash.Hex(), "error", err.Error())
			return nil, err
		}
		confirmations = new(big.Int).Sub(ec.currentHeight, (*big.Int)(tx.BlockNumber)).Int64()
	}

	return &PushEvent{
		Operation:     new(big.Int).SetBytes(crypto.Keccak256(*tx.Payload)),
		Tx:            createTxInfo(tx, r),
		Confirmations: confirmations,
		Method:        method.Name,
		Events:        ec.parseTxEventStatus(tx, r),
		ExtraData:     &extraData,
	}, nil

}

func (ec *Client) parseRpcTx(tx *rpcTx) (*PushEvent, error) {
	method, err := voteAbi.MethodById([]byte(*tx.Payload)[:4])
	if err != nil {
		ewLogger.Warn("ethwatcher PushTranxEvent, unknown contract method", "txhash", tx.TxHash.Hex())
		return nil, err
	}

	switch method.Name {
	case TOKEN_METHOD_BURN:
		return ec.parseBurnTx(tx, method)

	case VOTE_METHOD_MINT:
		return ec.parseMintTx(tx, method)

	case VOTE_METHOD_STOP:
		return ec.parseStopTx(tx, method)

	case VOTE_METHOD_START:
		return ec.parseStartTx(tx, method)

	case VOTE_METHOD_REVOKE:
		return ec.parseRevokeTx(tx, method)

	case VOTE_METHOD_CHANGEVOTER:
		return ec.parseChangeVoterTx(tx, method)

	case VOTE_METHOD_ADDVOTER:
		return ec.parseAddVoterTx(tx, method)

	case VOTE_METHOD_REMOVEVOTER:
		return ec.parseRemoveVoterTx(tx, method)

	case VOTE_METHOD_ADDCHAIN:
		return ec.parseAddChainTx(tx, method)

	case VOTE_METHOD_ADDAPP:
		return ec.parseAddAppTx(tx, method)

	case VOTE_METHOD_REMOVEAPP:
		return ec.parseRemoveAppTx(tx, method)

	case VOTE_METHOD_CHANGEGATEWAY:
		return ec.parseChangeGatewayTx(tx, method)

	default:
		ewLogger.Warn("ethwatcher uncatched input method", "method", method.Name)
		return nil, fmt.Errorf("ethwatcher uncatched input method")
	}
}

func (ec *Client) GetEventByHash(sHash string) (*PushEvent, error) {
	ctx := ensureContext(nil)
	tx, err := ec.TransactionByHash(ctx, common.HexToHash(sHash))
	if err != nil {
		return nil, err
	}
	return ec.parseRpcTx(tx)
}

func (ec *Client) pushTranxEvent(tranxes *[]*rpcTx, startIx int, eventCh chan<- *PushEvent) {
	for _, tx := range *tranxes {
		if int(*tx.TxIndex) < startIx {
			continue
		}

		if tx.CallTo == nil {
			continue
		}

		if _, ok := ec.mapContracts.Load(tx.CallTo.Hex()); !ok {
			continue
		}

		if len(*tx.Payload) == 0 {
			continue
		}

		event, err := ec.parseRpcTx(tx)
		if err == nil {
			select {
			case eventCh <- event:
				ewLogger.Debug("ethwatcher event pushed", "method", event.Method, "events", event.Events, "txHash", event.Tx.TxHash.Hex())
			}
		} else {
			ewLogger.Warn("ethwatcher PushEvent error", "txHash", tx.TxHash.Hex(), "error", err.Error())
		}
	}
}

func (ec *Client) StartWatch(start big.Int, tranxIx int, eventCh chan<- *PushEvent) {
	blkCh := make(chan *big.Int, 1000)
	ec.StartWatchBlock(start, blkCh)

	ec.InitMapInfo()

	ctx := ensureContext(nil)
	go func() {
		for {
			select {
			case blkHeight := <-blkCh:
				blockInfo, err := ec.BlockByNumber(ctx, blkHeight)
				if err != nil {
					ewLogger.Error("ethwatcher StartWatch, get block info error", "height", blkHeight.Uint64(), "error", err.Error())
					break
				}
				if blkHeight.Cmp(&start) > 0 {
					tranxIx = 0
				}
				ec.pushTranxEvent(&blockInfo.Transactions, tranxIx, eventCh)

				mockTx := &TxInfo{
					BlockNumber: blkHeight,
					TxIndex:     int(len(blockInfo.Transactions)),
				}

				blockDoneEvent := &PushEvent{
					Operation:     new(big.Int).SetInt64(0),
					Tx:            mockTx,
					Confirmations: int64(0),
					Method:        BLOCK_DONE_METHOD,
					Events:        uint64(0),
				}
				select {
				case eventCh <- blockDoneEvent:
					ewLogger.Debug("ethwatcher block push done", "height", blkHeight.Uint64())
				}
			}
		}
	}()

}

func (ec *Client) ReviewNonce(sPub string) bool {
	account, _, err := GetAddressFromPub(sPub)
	if err != nil {
		panic("calculate the address of a pubkey error when reviewing ldb store nonce!")
	}

	ec.nonceLock.Lock()

	defer ec.nonceLock.Unlock()

	var result hexutil.Uint64
	ctx := ensureContext(nil)

	err = ec.c.CallContext(ctx, &result, "eth_getTransactionCount", account, "pending")

	uintR := uint64(result)
	if err != nil {
		panic("get address nonce from chain error when reviewing ldb store nonce!")
	}

	addrTable := ethdb.NewTable(ec.ldb, account.Hex())
	var storeNonce uint64
	storeBytes, err := ec.ldb.Get(account.Bytes())
	if err == nil {
		storeNonce = hexutil.MustDecodeUint64(string(storeBytes))
		if storeNonce < uintR {
			return true
		}

		var realNonce uint64 = 0
		for i := storeNonce; i >= uintR; i-- {
			storeHashBytes, err := addrTable.Get([]byte(hexutil.EncodeUint64(i)))
			if err == nil {
				timestamp, err := hexutil.DecodeUint64(string(storeHashBytes))
				if err != nil {
					txHash := common.BytesToHash(storeHashBytes)
					validTxFlag := ec.IsValidTx(ctx, txHash)
					if validTxFlag == VALID_TX_VALID_TX {
						ewLogger.Debug("ethwatcher review nonce, find hash for nonce", "address", account.Hex(), "nonce", i, "hash", txHash.Hex())
					} else if validTxFlag == VALID_TX_NOT_FOUND {
						addrTable.Delete([]byte(hexutil.EncodeUint64(i)))
						realNonce = i - uint64(1)
						ewLogger.Debug("ethwatcher review nonce, find invalid hash nonce", "address", account.Hex(), "nonce", i, "hash", txHash.Hex())
					} else {
						panic(fmt.Sprintf("ethwatcher review nonce, get tranx by hash error, address: %v, nonce: %v, hash: %v!", account.Hex(), i, txHash.Hex()))
					}
				} else {
					addrTable.Delete([]byte(hexutil.EncodeUint64(i)))
					realNonce = i - uint64(1)
					ewLogger.Debug("ethwatcher review nonce, find timestamp nonce", "address", account.Hex(), "nonce", i, "timestamp", timestamp)
				}

			} else if err.Error() == ldberrors.ErrNotFound.Error() {
				ewLogger.Debug("ethwatcher review nonce, find nothing for this nonce", "address", account.Hex(), "nonce", i)
				realNonce = i - uint64(1)
			} else {
				panic(fmt.Sprintf("ethwatcher review nonce, read the hash of a nonce error, address: %v, nonce: %v!", account.Hex(), i))
			}
		}
		if realNonce > uint64(0) {
			err = ec.ldb.Put(account.Bytes(), []byte(hexutil.EncodeUint64(realNonce)))
			if err != nil {
				panic(fmt.Sprintf("ethwatcher review nonce set realNonce error, address: %v, realNonce: %v, storeNonce: %v!", account.Hex(), realNonce, storeNonce))
			}
		}

		return true

	} else if err.Error() == ldberrors.ErrNotFound.Error() {
		return true

	} else {
		panic(fmt.Sprintf(" ethwatcher review nonce read storeNonce from ldb error, address: %v, error: %v!", account.Hex(), err.Error()))
	}

}

// Blockchain Access

// BlockByNumber returns a block from the current canonical chain. If number is nil, the
// latest known block is returned.
//
// Note that loading full blocks requires two requests. Use HeaderByNumber
// if you don't need all transactions or uncle headers.
func (ec *Client) BlockByNumber(ctx context.Context, number *big.Int) (*rpcBlock, error) {
	return ec.getBlock(ctx, "eth_getBlockByNumber", toBlockNumArg(number), true)
}

func (ec *Client) getBlock(ctx context.Context, method string, args ...interface{}) (*rpcBlock, error) {
	var raw json.RawMessage

	err := ec.c.CallContext(ctx, &raw, method, args...)
	if err != nil {
		return nil, err

	} else if len(raw) == 0 {
		return nil, ethereum.NotFound
	}

	// fmt.Printf("raw prc rsp is %v \n", string(raw))

	// Decode header and transactions.
	var block *rpcBlock
	if err := json.Unmarshal(raw, &block); err != nil {
		return nil, err
	}

	// fmt.Printf("block hash is %v \n", block.BlockHash.Hex())

	return block, nil
}

// HeaderByHash returns the block header with the given hash.
func (ec *Client) HeaderByHash(ctx context.Context, hash common.Hash) (*rpcHeader, error) {
	var raw json.RawMessage

	err := ec.c.CallContext(ctx, &raw, "eth_getBlockByHash", hash, false)
	if err != nil {
		return nil, err

	} else if len(raw) == 0 {
		return nil, ethereum.NotFound
	}

	// fmt.Printf("raw prc rsp is %v \n", string(raw))

	// Decode header
	var head *rpcHeader
	if err := json.Unmarshal(raw, &head); err != nil {
		return nil, err
	}

	// fmt.Printf("block hash is %v \n", head.BlockHash.Hex())

	return head, nil

	// var head *types.Header
	// err := ec.c.CallContext(ctx, &head, "eth_getBlockByHash", hash, false)
	// if err == nil && head == nil {
	// 	err = ethereum.NotFound
	// }
	// return head, err
}

// HeaderByNumber returns a block header from the current canonical chain. If number is
// nil, the latest known header is returned.
func (ec *Client) HeaderByNumber(ctx context.Context, number *big.Int) (*rpcHeader, error) {
	var raw json.RawMessage

	err := ec.c.CallContext(ctx, &raw, "eth_getBlockByNumber", toBlockNumArg(number), false)
	if err != nil {
		return nil, err

	} else if len(raw) == 0 {
		return nil, ethereum.NotFound
	}

	// fmt.Printf("raw prc rsp is %v \n", string(raw))

	// Decode header
	var head *rpcHeader
	if err := json.Unmarshal(raw, &head); err != nil {
		return nil, err
	}

	// fmt.Printf("block hash is %v \n", head.BlockHash.Hex())

	return head, nil

	// var head *types.Header
	// err := ec.c.CallContext(ctx, &head, "eth_getBlockByNumber", toBlockNumArg(number), false)
	// if err == nil && head == nil {
	// 	err = ethereum.NotFound
	// }
	// return head, err
}

func (ec *Client) IsValidTx(ctx context.Context, hash common.Hash) int {
	var raw json.RawMessage
	var tx rpcTx

	err := ec.c.CallContext(ctx, &raw, "eth_getTransactionByHash", hash)
	if err != nil {
		return VALID_TX_RPC_ERROR

	} else if len(raw) == 0 {
		return VALID_TX_NOT_FOUND

	}

	if err = json.Unmarshal(raw, &tx); err != nil {
		return VALID_TX_JSON_ERROR
	}
	return VALID_TX_VALID_TX

}

// TransactionByHash returns the transaction with the given hash.
func (ec *Client) TransactionByHash(ctx context.Context, hash common.Hash) (tx *rpcTx, err error) {
	var raw json.RawMessage

	err = ec.c.CallContext(ctx, &raw, "eth_getTransactionByHash", hash)
	if err != nil {
		return nil, err

	} else if len(raw) == 0 {
		return nil, ethereum.NotFound

	}

	if err := json.Unmarshal(raw, &tx); err != nil {
		return nil, err
	}

	// fmt.Printf("raw tx is %v \n", string(raw))

	return tx, nil
}

// TransactionReceipt returns the receipt of a transaction by transaction hash.
// Note that the receipt is not available for pending transactions.
func (ec *Client) TransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error) {
	var r *types.Receipt
	err := ec.c.CallContext(ctx, &r, "eth_getTransactionReceipt", txHash)
	if err == nil {
		if r == nil {
			return nil, ethereum.NotFound
		}
	}
	return r, err
}

func (ec *Client) SendTranxByInput(sPub string, sPubHash string, input []byte) (string, error) {
	opts := CreateRemoteSignOpts(sPub, sPubHash, viper.GetString("KEYSTORE.url"), viper.GetString("KEYSTORE.service_id"), viper.GetString("KEYSTORE.keystore_private_key"))
	return ec.doSendTxByInput(input, opts)
}

func (ec *Client) GatewayTransaction(sPub string, sPubHash string, method string, args ...interface{}) (string, error) {
	input, err := ec.EncodeInput(method, args...)
	if err != nil {
		ewLogger.Error("encode input error!", "error", err.Error())
		return "", err
	}
	return ec.SendTranxByInput(sPub, sPubHash, input)
}

func (ec *Client) SendTxByInput(key []byte, input []byte) (string, error) {
	privKey := crypto.ToECDSAUnsafe(key)
	opts := bind.NewKeyedTransactor(privKey)
	return ec.doSendTxByInput(input, opts)
}

func (ec *Client) SendTxByArgs(key []byte, method string, args ...interface{}) (string, error) {
	input, err := ec.EncodeInput(method, args...)
	if err != nil {
		ewLogger.Error("encode input error!", "error", err.Error())
		return "", err
	}
	return ec.SendTxByInput(key, input)
}

func (ec *Client) doSendTxByInput(input []byte, opts *bind.TransactOpts) (sTxHash string, sendErr error) {
	ctx := ensureContext(nil)

	tStart := time.Now()
	fromAddr := opts.From
	// Ensure a valid value field and resolve the account nonce
	value := new(big.Int)
	nonce, err := ec.PendingNonceAt(ctx, fromAddr)
	if err != nil {
		ewLogger.Error("ethwatcher send tx, retrieve account nonce error!", "error", err.Error())
		return "", err
	}
	// tNonce := time.Now()
	// ewLogger.Debug("ethwatcher send tx, nonce cost", "cost", tNonce.Sub(tStart))

	addrTable := ethdb.NewTable(ec.ldb, fromAddr.Hex())
	defer func() {
		if sTxHash == "" {
			storeHashBytes, err := addrTable.Get([]byte(hexutil.EncodeUint64(nonce)))
			if err == nil {
				_, err := hexutil.DecodeUint64(string(storeHashBytes))
				if err != nil {
					txHash := common.BytesToHash(storeHashBytes)
					validTxFlag := ec.IsValidTx(ctx, txHash)
					if validTxFlag == VALID_TX_NOT_FOUND {
						addrTable.Delete([]byte(hexutil.EncodeUint64(nonce)))
						ewLogger.Debug("ethwatcher send tx, delete nonce store for sended tx error", "address", fromAddr.Hex(), "nonce", nonce)
					}

				} else {
					addrTable.Delete([]byte(hexutil.EncodeUint64(nonce)))
					ewLogger.Debug("ethwatcher send tx, delete nonce store for sended tx error", "address", fromAddr.Hex(), "nonce", nonce)

				}

			} else {
				addrTable.Delete([]byte(hexutil.EncodeUint64(nonce)))
				ewLogger.Debug("ethwatcher send tx, delete nonce store for sended tx error", "address", fromAddr.Hex(), "nonce", nonce)
			}
		}
		ewLogger.Debug("ethwatcher send tx, sended tx", "address", fromAddr.Hex(), "nonce", nonce, "txHash", sTxHash)
		tEnd := time.Now()
		ewLogger.Debug("ethwatcher send tx, all time cost", "cost", tEnd.Sub(tStart))
	}()

	// Figure out the gas allowance and gas price values
	gasPrice, err := ec.SuggestGasPrice(ctx)
	if err != nil {
		ewLogger.Error("ethwatcher send tx, suggest gas price error!", "error", err.Error())
		return "", err
	}
	if gasPrice.Cmp(minGasPrice) < 0 {
		gasPrice = minGasPrice

	} else if gasPrice.Cmp(maxGasPrice) > 0 {
		gasPrice = maxGasPrice
	}

	// Gas estimation cannot succeed without code for method invocations
	contractAddr := ec.voteContract
	if code, err := ec.PendingCodeAt(ctx, contractAddr); err != nil {
		ewLogger.Error("ethwatcher send tx, retrieve contract code error!", "error", err.Error())
		return "", err

	} else if len(code) == 0 {
		ewLogger.Error("ethwatcher send tx, no code at vote contract address", "address", contractAddr.Hex())
		return "", bind.ErrNoCode
	}
	// If the contract surely has code (or code is not needed), estimate the transaction
	msg := ethereum.CallMsg{From: fromAddr, To: &contractAddr, Value: value, Data: input}
	gasLimit, err := ec.EstimateGas(ctx, msg)
	if err != nil {
		ewLogger.Error("ethwatcher send tx, estimate gas error!", "error", err.Error())
		return "", err
	}
	gasLimit += 200000
	// Create the transaction, sign it and schedule it for execution
	var rawTx *types.Transaction
	rawTx = types.NewTransaction(nonce, contractAddr, value, gasLimit, gasPrice, input)

	// tCreateTx := time.Now()
	// ewLogger.Debug("ethwatcher send tx, create tx cost", "cost", tCreateTx.Sub(tNonce))

	signedTx, err := opts.Signer(types.HomesteadSigner{}, opts.From, rawTx)
	if err != nil {
		return "", err
	}
	err = ec.SendTransaction(ctx, signedTx)
	if err != nil {
		ewLogger.Error("ethwatcher send tx, broadcast tx error!", "nonce", nonce, "error", err.Error())
		return "", err
	}

	// tSendTx := time.Now()
	// ewLogger.Debug("ethwatcher send tx, broadcast tx cost", "cost", tSendTx.Sub(tCreateTx))

	addrTable.Put([]byte(hexutil.EncodeUint64(nonce)), signedTx.Hash().Bytes())

	return signedTx.Hash().Hex(), nil
}

func (ec *Client) MethodById(sig []byte) (*abi.Method, error) {
	return voteAbi.MethodById(sig)
}

func (ec *Client) EncodeInput(method string, args ...interface{}) ([]byte, error) {
	return voteAbi.Pack(method, args...)
}

func (ec *Client) GetBlockNumber() int64 {
	return new(big.Int).Sub(ec.currentHeight, ec.confirmHeight).Int64()
}

// SubscribeNewHead subscribes to notifications about the current blockchain head
// on the given channel.
func (ec *Client) SubscribeNewHead(ctx context.Context, ch chan<- *rpcHeader) (ethereum.Subscription, error) {
	return ec.c.EthSubscribe(ctx, ch, "newHeads")
}

// State Access
// BalanceAt returns the wei balance of the given account.
// The block number can be nil, in which case the balance is taken from the latest known block.
func (ec *Client) BalanceAt(ctx context.Context, account common.Address, blockNumber *big.Int) (*big.Int, error) {
	var result hexutil.Big
	err := ec.c.CallContext(ctx, &result, "eth_getBalance", account, toBlockNumArg(blockNumber))
	return (*big.Int)(&result), err
}

// CodeAt returns the contract code of the given account.
// The block number can be nil, in which case the code is taken from the latest known block.
func (ec *Client) CodeAt(ctx context.Context, account common.Address, blockNumber *big.Int) ([]byte, error) {
	var result hexutil.Bytes
	err := ec.c.CallContext(ctx, &result, "eth_getCode", account, toBlockNumArg(blockNumber))
	return result, err
}

// NonceAt returns the account nonce of the given account.
// The block number can be nil, in which case the nonce is taken from the latest known block.
func (ec *Client) NonceAt(ctx context.Context, account common.Address, blockNumber *big.Int) (uint64, error) {
	var result hexutil.Uint64
	err := ec.c.CallContext(ctx, &result, "eth_getTransactionCount", account, toBlockNumArg(blockNumber))
	return uint64(result), err
}

// Filters

// FilterLogs executes a filter query.
func (ec *Client) FilterLogs(ctx context.Context, q ethereum.FilterQuery) ([]types.Log, error) {
	var result []types.Log
	err := ec.c.CallContext(ctx, &result, "eth_getLogs", toFilterArg(q))
	return result, err
}

// SubscribeFilterLogs subscribes to the results of a streaming filter query.
func (ec *Client) SubscribeFilterLogs(ctx context.Context, q ethereum.FilterQuery, ch chan<- types.Log) (ethereum.Subscription, error) {
	return ec.c.EthSubscribe(ctx, ch, "logs", toFilterArg(q))
}

// Pending State

// PendingCodeAt returns the contract code of the given account in the pending state.
func (ec *Client) PendingCodeAt(ctx context.Context, account common.Address) ([]byte, error) {
	var result hexutil.Bytes
	err := ec.c.CallContext(ctx, &result, "eth_getCode", account, "pending")
	return result, err
}

// PendingNonceAt returns the account nonce of the given account in the pending state.
// This is the nonce that should be used for the next transaction.
func (ec *Client) PendingNonceAt(ctx context.Context, account common.Address) (uint64, error) {
	ec.nonceLock.Lock()
	defer ec.nonceLock.Unlock()

	var result hexutil.Uint64
	err := ec.c.CallContext(ctx, &result, "eth_getTransactionCount", account, "pending")

	uintR := uint64(result)
	if err != nil {
		return uintR, err
	}

	addrTable := ethdb.NewTable(ec.ldb, account.Hex())
	var storeNonce uint64
	storeBytes, err := ec.ldb.Get(account.Bytes())
	if err == nil {
		storeNonce = hexutil.MustDecodeUint64(string(storeBytes))
		storeNonce += uint64(1)
		if storeNonce <= uintR {
			storeNonce = uintR

		} else {
			storeHashBytes, err := addrTable.Get([]byte(hexutil.EncodeUint64(uintR)))
			if err == nil {
				timestamp, err := hexutil.DecodeUint64(string(storeHashBytes))
				if err != nil {
					txHash := common.BytesToHash(storeHashBytes)
					validTxFlag := ec.IsValidTx(ctx, txHash)
					if validTxFlag == VALID_TX_NOT_FOUND {
						addrTable.Put([]byte(hexutil.EncodeUint64(uintR)), []byte(hexutil.EncodeUint64(uint64(time.Now().Unix()))))
						ewLogger.Debug("ethwatcher get nonce, no spec tx for this nonce, using web3 nonce", "address", account.Hex(), "nonce", uintR)
						return uintR, nil
					}
				} else if time.Now().Unix()-int64(timestamp) > int64(10) {
					addrTable.Put([]byte(hexutil.EncodeUint64(uintR)), []byte(hexutil.EncodeUint64(uint64(time.Now().Unix()))))
					ewLogger.Debug("ethwatcher get nonce, wait nonce used timeout, using web3 nonce", "address", account.Hex(), "nonce", uintR)
					return uintR, nil

				}

			} else if err != nil && err.Error() == ldberrors.ErrNotFound.Error() {
				addrTable.Put([]byte(hexutil.EncodeUint64(uintR)), []byte(hexutil.EncodeUint64(uint64(time.Now().Unix()))))
				ewLogger.Debug("ethwatcher get nonce, no ldb store for this nonce, using web3 nonce", "address", account.Hex(), "nonce", uintR)
				return uintR, nil
			}
		}

	} else if err.Error() == ldberrors.ErrNotFound.Error() {
		storeNonce = uintR

	} else {
		ewLogger.Error("ethwatcher get nonce, read ldb error", "address", account.Hex())
		return uintR, err
	}

	for {
		storeHashBytes, err := addrTable.Get([]byte(hexutil.EncodeUint64(storeNonce)))
		if err == nil {
			timestamp, err := hexutil.DecodeUint64(string(storeHashBytes))
			if err != nil {
				txHash := common.BytesToHash(storeHashBytes)
				validTxFlag := ec.IsValidTx(ctx, txHash)
				if validTxFlag == VALID_TX_VALID_TX {
					ewLogger.Debug("ethwatcher get nonce, find spec tx for this nonce, use higher", "address", account.Hex(), "nonce", storeNonce, "hash", txHash.Hex())
					storeNonce += 1

				} else if validTxFlag == VALID_TX_NOT_FOUND {
					ewLogger.Debug("ethwatcher get nonce, no spec tx for this nonce, use it", "address", account.Hex(), "nonce", storeNonce, "hash", txHash.Hex())
					break
				} else {
					ewLogger.Warn("ethwatcher get nonce, get tranx error, wait 1s", "address", account.Hex(), "nonce", storeNonce, "hash", txHash.Hex())
					time.Sleep(1 * time.Second)
				}

			} else {
				ewLogger.Debug("ethwatcher get nonce, find timestamp nonce, use it", "address", account.Hex(), "nonce", storeNonce, "timestamp", timestamp)
				break

			}

		} else if err.Error() == ldberrors.ErrNotFound.Error() {
			break

		} else {
			ewLogger.Warn("ethwatcher get nonce, read ldb for storeNonce error, wait 1s", "address", account.Hex(), "nonce", storeNonce)
			time.Sleep(1 * time.Second)
		}

	}

	err = ec.ldb.Put(account.Bytes(), []byte(hexutil.EncodeUint64(storeNonce)))
	if err != nil {
		ewLogger.Error("ethwatcher get nonce, store nocne to ldb error", "address", account.Hex(), "nonce", storeNonce)
		return uintR, err
	}

	err = addrTable.Put([]byte(hexutil.EncodeUint64(storeNonce)), []byte(hexutil.EncodeUint64(uint64(time.Now().Unix()))))
	if err != nil {
		ewLogger.Error("ethwatcher get nonce, store nocne tx tag to ldb error", "address", account.Hex(), "nonce", storeNonce)
		return uintR, err
	}

	ewLogger.Debug("ethwatcher get nonce, returned", "address", account.Hex(), "nonce", storeNonce)
	return storeNonce, nil
}

// Contract Calling
// CallContract executes a message call transaction, which is directly executed in the VM
// of the node, but never mined into the blockchain.
// blockNumber selects the block height at which the call runs. It can be nil, in which
// case the code is taken from the latest known block. Note that state from very old
// blocks might not be available.
func (ec *Client) CallContract(ctx context.Context, msg ethereum.CallMsg, blockNumber *big.Int) ([]byte, error) {
	var hex hexutil.Bytes
	err := ec.c.CallContext(ctx, &hex, "eth_call", toCallArg(msg), toBlockNumArg(blockNumber))
	if err != nil {
		return nil, err
	}
	return hex, nil
}

func (ec *Client) BlockNumber() (*big.Int, error) {
	var hex hexutil.Big
	if err := ec.c.CallContext(ensureContext(nil), &hex, "eth_blockNumber"); err != nil {
		ewLogger.Error("Get blockNumber error!", "error", err)
		return nil, err
	}
	return (*big.Int)(&hex), nil
}

// SuggestGasPrice retrieves the currently suggested gas price to allow a timely
// execution of a transaction.
func (ec *Client) SuggestGasPrice(ctx context.Context) (*big.Int, error) {
	var hex hexutil.Big
	if err := ec.c.CallContext(ctx, &hex, "eth_gasPrice"); err != nil {
		return nil, err
	}
	return (*big.Int)(&hex), nil
}

// EstimateGas tries to estimate the gas needed to execute a specific transaction based on
// the current pending state of the backend blockchain. There is no guarantee that this is
// the true gas limit requirement as other transactions may be added or removed by miners,
// but it should provide a basis for setting a reasonable default.
func (ec *Client) EstimateGas(ctx context.Context, msg ethereum.CallMsg) (uint64, error) {
	var hex hexutil.Uint64
	err := ec.c.CallContext(ctx, &hex, "eth_estimateGas", toCallArg(msg))
	if err != nil {
		return 0, err
	}
	return uint64(hex), nil
}

// SendTransaction injects a signed transaction into the pending pool for execution.
// If the transaction was a contract creation use the TransactionReceipt method to get the
// contract address after the transaction has been mined.
func (ec *Client) SendTransaction(ctx context.Context, tx *types.Transaction) error {
	data, err := rlp.EncodeToBytes(tx)
	if err != nil {
		return err
	}
	return ec.c.CallContext(ctx, nil, "eth_sendRawTransaction", common.ToHex(data))
}

func toCallArg(msg ethereum.CallMsg) interface{} {
	arg := map[string]interface{}{
		"from": msg.From,
		"to":   msg.To,
	}
	if len(msg.Data) > 0 {
		arg["data"] = hexutil.Bytes(msg.Data)
	}
	if msg.Value != nil {
		arg["value"] = (*hexutil.Big)(msg.Value)
	}
	if msg.Gas != 0 {
		arg["gas"] = hexutil.Uint64(msg.Gas)
	}
	if msg.GasPrice != nil {
		arg["gasPrice"] = (*hexutil.Big)(msg.GasPrice)
	}
	return arg
}

func toBlockNumArg(number *big.Int) string {
	if number == nil {
		return "latest"
	}
	return hexutil.EncodeBig(number)
}

func toFilterArg(q ethereum.FilterQuery) interface{} {
	arg := map[string]interface{}{
		"fromBlock": toBlockNumArg(q.FromBlock),
		"toBlock":   toBlockNumArg(q.ToBlock),
		"address":   q.Addresses,
		"topics":    q.Topics,
	}
	if q.FromBlock == nil {
		arg["fromBlock"] = "0x0"
	}
	return arg
}

func createTxInfo(tx *rpcTx, r *types.Receipt) *TxInfo {
	txInfo := &TxInfo{
		BlockNumber: (*big.Int)(tx.BlockNumber),
		BlockHash:   tx.BlockHash,

		TxHash: tx.TxHash,

		From: tx.CallFrom,
		To:   tx.CallTo,
	}

	if r != nil {
		txInfo.TxIndex = int(*tx.TxIndex)
		txInfo.Status = r.Status
		txInfo.Fee = new(big.Int).Mul((*big.Int)(tx.GasPrice), new(big.Int).SetUint64(r.GasUsed))
	}

	return txInfo
}

type burnCallArguments struct {
	Wad      uint64
	Receiver string
}

type rpcTx struct {
	BlockNumber *hexutil.Big `json:"blockNumber" gencodec:"required"`
	BlockHash   *common.Hash `json:"blockHash" rlp:"-"`

	TxHash  *common.Hash  `json:"hash"     rlp:"-"`
	TxIndex *hexutil.Uint `json:"transactionIndex"`

	CallFrom *common.Address `json:"from"`
	CallTo   *common.Address `json:"to"       rlp:"nil"`

	GasPrice *hexutil.Big `json:"gasPrice" gencodec:"required"`

	EthAmount *hexutil.Big   `json:"value"    gencodec:"required"`
	Payload   *hexutil.Bytes `json:"input"    gencodec:"required"`
}

type rpcHeader struct {
	Number    *hexutil.Big `json:"number"`
	BlockHash *common.Hash `json:"hash"`
}

type rpcBlock struct {
	Transactions []*rpcTx     `json:"transactions"`
	BlockNumber  *hexutil.Big `json:"number"`
	BlockHash    *common.Hash `json:"hash"`
}
