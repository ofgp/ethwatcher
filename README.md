# ETH-WATCHER

## **introduction**
1. 网关合约和应用TOKEN合约分离。
2. 网关合约处理网关节点共识，达成后执行具体的操作。
3. 应用TOKEN合约的增发控制权仅能由网关合约控制。
4. 整体流程：
    - 使用预制的网关节点地址列表作为constructor参数部署网关合约，获得网关合约地址《A》。
    - 使用《A》作为应用合约的constructor参数部署应用合约，获得应用合约地址《B》。
    - 网关节点使用 链《C》+proposal 在《A》中声明链，获得链的code 《CC》。
    - 网关节点使用 《B》+《CC》+《C》中代币的TokenCode + proposal 在《A》中声明应用，确定“应用合约、链、链下Token” 三者的绑定关系，获得应用的Code 《BC》。 
    - 网关节点使用《BC》+铸币信息向《A》发起铸币，共识达成后，促成《B》中TOKEN的增发。
    - 用户向《B》发起熔币交易，交易成功后促成网关发起《C》上的提币请求。
    - 《A》中可以存在多个应用合约，网关节点可以通过共识移除某个应用合约。
    - 网关节点可以通过共识在《A》中增删改有投票权的网关节点。
5. 所有推送消息为PushEvent，ExtraData依据Method不同而使用不同的struct。Method的值定义在常量中。Events是本交易触发的事件所代表的bit位集合。

##  **structs**
```
type TxInfo struct {
	BlockNumber *big.Int
	BlockHash   *common.Hash

	TxHash  *common.Hash
	TxIndex int // the index that this tx be sorted in it's block

	From *common.Address // tx sender
	To   *common.Address // tx receiver

	Status uint // result status for this tx. 1 > success; 0 > failed

	Fee *big.Int
}

type PushEvent struct {
	Operation *big.Int // keccak256(tx.input)

	Tx *TxInfo

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
```

## **constants**
```
const VOTE_STATUS_OTHER = 0 // uncatched tx
const (
	TX_STATUS_PENDING      = 1 << iota // 2**0 1
	TX_STATUS_FAILED                   // 2**1 2
	VOTE_TX_CONFIRM                    // 2**2 4  a vote tx
	VOTE_TX_OPTDONE                    // 2**3 8  vote done
	VOTE_TX_MINT                       // 2**4 16
	VOTE_TX_STOPPED                    // 2**5 32
	VOTE_TX_STARTED                    // 2**6 64
	VOTE_TX_VOTERCHANGED               // 2**7 128
	VOTE_TX_VOTERADDED                 // 2**8 256
	VOTE_TX_VOTERREMOVED               // 2**9 512
	VOTE_TX_CHAINADDED                 // 2**10 1024
	VOTE_TX_APPADDED                   // 2**11 2048
	VOTE_TX_APPREMOVED                 // 2**12 4096
	VOTE_TX_GATEWAYCHANGED             // 2**13 8192  change the gatewayAddr of an app
	VOTER_TX_REVOKE                    // 2**14 16384
	TOKEN_TX_BURN                      // 2**15 32768
)

const (
	VALID_TX_VALID_TX = iota
	VALID_TX_RPC_ERROR
	VALID_TX_NOT_FOUND
	VALID_TX_JSON_ERROR
)

const (
	CHAIN_SELF_ETH = "eth"

	VOTE_METHOD_MINT          = "mintByGateway"
	VOTE_METHOD_STOP          = "stop"
	VOTE_METHOD_START         = "start"
	VOTE_METHOD_REVOKE        = "revoke"
	VOTE_METHOD_CHANGEVOTER   = "changeVoter"
	VOTE_METHOD_ADDVOTER      = "addVoter"
	VOTE_METHOD_REMOVEVOTER   = "removeVoter"
	VOTE_METHOD_ADDCHAIN      = "addChain"
	VOTE_METHOD_ADDAPP        = "addApp"
	VOTE_METHOD_REMOVEAPP     = "removeApp"
	VOTE_METHOD_CHANGEGATEWAY = "changeGatewayAddr"

	TOKEN_METHOD_BURN = "burn"

	VOTE_EVENT_MINT          = "MintByGateway"
	VOTE_EVENT_BURN          = "BurnForGateway"
	VOTE_EVENT_CONFIRM       = "Confirmation"
	VOTE_EVENT_STARTED       = "Started"
	VOTE_EVENT_STOPPED       = "Stopped"
	VOTE_EVENT_OPTDONE       = "OperationDone"
	VOTE_EVENT_REVOKE        = "Revoke"
	VOTE_EVENT_VOTERCHANGED  = "VoterChanged"
	VOTE_EVENT_VOTERADDED    = "VoterAdded"
	VOTE_EVENT_VOTERREMOVED  = "VoterRemoved"
	VOTE_EVENT_CHAINADDED    = "ChainAdded"
	VOTE_EVENT_APPADDED      = "AppAdded"
	VOTE_EVENT_APPREMOVED    = "AppRemoved"
	VOTE_EVENT_CHANGEGATEWAY = "GatewayAddrChanged"
)
```

## **interface**
```
▼+Client : struct     

    [fields]      
    
   -c : *rpc.Client  // ethereum client, 和以太坊进行交互           
   -confirmHeight : *big.Int // 块确认所需的高度，实例化时由外部传入
   -currentHeight : *big.Int // 当前块高度                      
   -ldb : *ethdb.LDBDatabase // 存储nonce相关信息                               
   -mapAppAddrToCode : *sync.Map // 缓存app合约地址和appcode的映射关系
   -mapAppCodeToInfo : *sync.Map // 缓存appcode和app信息的映射关系
   -mapChainToCode : *sync.Map   // 缓存链名和chaincode的映射关系
   -mapCodeToChain : *sync.Map   // 缓存chaincode和链名的映射关系
   -mapContracts : *sync.Map     // 缓存需要监听的合约地址，包括vote合约和app合约
   -nonceLock : *sync.Mutex      // 并发交易时nonce控制
   -voteContract : common.Address  // vote合约地址
   -voteEvents : map[string]string // 需要监听的合约事件
   
    [methods]
    
   // 传入合约方法名，及对应的方法参数，按abi规则序列化。
   // 参照GatewayVote.go及GatewayVote.json 
   +EncodeInput(method string, args ...interface{}) : []byte, error
   
   // 通过合约方法及参数直接发送交易
   +GatewayTransaction(sPub string, sPubHash string, method string, args ...interface{}) : string, error
   
   // 通过app合约地址获取appcode，合法的appcode是一个大于0的整数
   +GetAppCode(sAppAddr string) : uint32
   
   // 获取确认的块高度
   +GetBlockNumber() : int64
   
   // 通过链名获取chaincode，合法的chaincode是一个大于0的整数
   +GetChainCode(sChainName string) : uint32
   
   // 通过交易hash获得交易的event详情，用于验证交易。
   +GetEventByHash(sHash string) : *PushEvent, error
   
   // 使用EncodeInput的结果发送一个交易
   +SendTranxByInput(sPub string, sPubHash string, input []byte) : string, error、
   
   // 启动监听，传入起始监听高度，起始块交易索引值（跳过起始块的多少个交易）， channel
   +StartWatch(start big.Int, tranxIx int, eventCh chan *PushEvent)
   
   // 验证app合约的绑定关系
   // sChain 指代app合约所对应的源链
   // appCode是合约产生的，对应的是某个合约；tokenCode是bch链生成的，对应omni。
   // 铸币时，铸币信息里的TokenFrom对应的是tokenCode，TokenTo对应的是appCode；
   // 熔币时，熔币信息里的TokenFrom对应的是appCode，TokenTo对应的是tokenCode
   +VerifyAppInfo(sChain string, tokenCode uint32, appCode uint32) : bool
   
   [functions]
    
   // 实例化。 使用参数> 以太坊ws连接地址，确认块所需的高度，节点公钥
   +Dial(rawurl string, confirmHeight int64, sPub string) : *Client, error 
   +NewEthWatcher(rawurl string, confirmHeight int64, sPub string) : *Client, error  
```

