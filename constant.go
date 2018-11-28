package ethwatcher

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
	VOTE_TX_RECVETHER                  //2**16 65536
	VOTE_TX_SENDETHER                  //2**17 131072
)

const (
	VALID_TX_VALID_TX = iota
	VALID_TX_RPC_ERROR
	VALID_TX_NOT_FOUND
	VALID_TX_JSON_ERROR
)

const (
	CHAIN_SELF_ETH = "eth"

	BLOCK_DONE_METHOD = "blockPushDone"

	VOTE_METHOD_MINT = "mintByGateway"
	// VOTE_METHOD_BURN          = "burnForGateway"
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
	VOTE_METHOD_RECVETHER     = "recvEther"
	VOTE_METHOD_SENDETHER     = "sendEther"

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
	VOTE_EVENT_RECVETHER     = "RecvEther"
	VOTE_EVENT_SENDETHER     = "SendEther"
)
