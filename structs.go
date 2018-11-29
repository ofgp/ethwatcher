package ethwatcher

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type TxInfo struct {
	BlockNumber *big.Int
	BlockHash   *common.Hash

	TxHash  *common.Hash
	TxIndex int // the index that this tx be sorted in it's block

	From *common.Address // tx sender
	To   *common.Address // tx receiver

	Status uint64 // result status for this tx. 1 > success; 0 > failed

	Fee *big.Int
}

type PushEvent struct {
	Operation *big.Int // keccak256(tx.input)

	Tx *TxInfo

	Confirmations int64

	Method string // specify which ExtraData struct type should this event to use.
	Events uint64 // tx log flags. see constant.go -> VOTE_TX_*

	ExtraData interface{}
}

type AssetInfo struct {
	Address string
	Amount  uint64
}

type ExtraBurnData struct { // TOKEN_METHOD_BURN
	ScTxid       string // tx hash be coded to hex string
	Amount       uint64
	RechargeList []*AssetInfo
	From         string // the chain that this event from
	To           string // the chain that this event will go to
	TokenFrom    uint32 // the code of the app which created this event
	TokenTo      uint32 // the code of the token that this event will go to
}

type ExtraMintData struct { // VOTE_METHOD_MINT
	AppCode  uint32
	Receiver common.Address
	Wad      uint64
	Proposal string
}

type ExtraAppAddedData struct { // VOTE_METHOD_ADDAPP
	App      common.Address
	Chain    uint32
	Token    uint32
	Proposal string
}

type ExtraAppRemovedData struct { // VOTE_METHOD_REMOVEAPP
	Code     uint32
	Proposal string
}

type ExtraChainAddedData struct { // VOTE_METHOD_ADDCHAIN
	Chain    string
	Proposal string
}

type ExtraGatewayAddrChangedData struct { // VOTE_METHOD_CHANGEGATEWAY
	AppCode  uint32
	Newer    common.Address
	Proposal string
}

type ExtraRevokeData struct { // VOTE_METHOD_REVOKE
	Operation *big.Int
}

type ExtraStartedData struct { // VOTE_METHOD_START
	Proposal string
}

type ExtraStoppedData struct { // VOTE_METHOD_STOP
	Proposal string
}

type ExtraVoterAddedData struct { // VOTE_METHOD_ADDVOTER
	NewVoter common.Address
	Proposal string
}

type ExtraVoterChangedData struct { // VOTE_METHOD_CHANGEVOTER
	OldVoter common.Address
	NewVoter common.Address
	Proposal string
}

type ExtraVoterRemovedData struct { // VOTE_METHOD_REMOVEVOTER
	OldVoter common.Address
	Proposal string
}

type AppInfo struct {
	AppCode   uint32
	ChainCode uint32
	TokenCode uint32
	Addr      common.Address
	Chain     string
}

//ExtraEther 作为侧链接收eth
type ExtraEther struct {
	Chain  string   //目标链
	Addr   string   //就收地址
	Amount *big.Int //接收金额
	ScTxid string   //txid
	From   string
	To     string
}

// ExtraSendEther eth 转账
type ExtraSendEther struct {
	Receiver common.Address // 接收eth地址
	Amount   *big.Int       //接收金额
	Proposal string         //原始兑换交易txhash
}
