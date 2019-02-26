package rpc

import "github.com/zjykzk/rocketmq-client-go/remote"

// request code
const (
	SendMessage                      = remote.Code(10)
	PullMessage                      = remote.Code(11)
	QueryMessage                     = remote.Code(12)
	QueryBrokerOffset                = remote.Code(13)
	QueryConsumerOffset              = remote.Code(14)
	UpdateConsumerOffset             = remote.Code(15)
	UpdateAndCreateTopic             = remote.Code(17)
	GetAllTopicConfig                = remote.Code(21)
	GetTopicConfigList               = remote.Code(22)
	GetTopicNameList                 = remote.Code(23)
	UpdateBrokerConfig               = remote.Code(25)
	GetBrokerConfig                  = remote.Code(26)
	TriggerDeleteFiles               = remote.Code(27)
	GetBrokerRuntimeInfo             = remote.Code(28)
	SearchOffsetByTimestamp          = remote.Code(29)
	GetMaxOffset                     = remote.Code(30)
	GetMinOffset                     = remote.Code(31)
	GetEarliestMsgStoretime          = remote.Code(32)
	ViewMessageByID                  = remote.Code(33)
	HeartBeat                        = remote.Code(34)
	UnregisterClient                 = remote.Code(35)
	ConsumerSendMsgBack              = remote.Code(36)
	EndTransaction                   = remote.Code(37)
	GetConsumerListByGroup           = remote.Code(38)
	CheckTransactionState            = remote.Code(39)
	NotifyConsumerIdsChanged         = remote.Code(40)
	LockBatchMq                      = remote.Code(41)
	UnlockBatchMq                    = remote.Code(42)
	GetAllConsumerOffset             = remote.Code(43)
	GetAllDelayOffset                = remote.Code(45)
	CheckClientConfig                = remote.Code(46)
	PutKvConfig                      = remote.Code(100)
	GetKvConfig                      = remote.Code(101)
	DeleteKvConfig                   = remote.Code(102)
	RegisterBroker                   = remote.Code(103)
	UnregisterBroker                 = remote.Code(104)
	GetRouteintoByTopic              = remote.Code(105)
	GetBrokerClusterInfo             = remote.Code(106)
	UpdateAndCreateSubscriptiongroup = remote.Code(200)
	GetAllSubscriptiongroupConfig    = remote.Code(201)
	GetTopicStatsInfo                = remote.Code(202)
	GetConsumerConnectionList        = remote.Code(203)
	GetProducerConnectionList        = remote.Code(204)
	WipeWritePermOfBroker            = remote.Code(205)
	GetAllTopicListFromNameserver    = remote.Code(206)
	DeleteSubscriptiongroup          = remote.Code(207)
	GetConsumeStats                  = remote.Code(208)
	SuspendConsumer                  = remote.Code(209)
	ResumeConsumer                   = remote.Code(210)
	ResetConsumerOffsetInConsumer    = remote.Code(211)
	ResetConsumerOffsetInBroker      = remote.Code(212)
	AdjustConsumerThreadPool         = remote.Code(213)
	WhoConsumeTheMessage             = remote.Code(214)
	DeleteTopicInBroker              = remote.Code(215)
	DeleteTopicInNamesrv             = remote.Code(216)
	GetKvlistByNamespace             = remote.Code(219)
	ResetConsumerClientOffset        = remote.Code(220)
	GetConsumerStatusFromClient      = remote.Code(221)
	InvokeBrokerToResetOffset        = remote.Code(222)
	InvokeBrokerToGetConsumerStatus  = remote.Code(223)
	QueryTopicConsumeByWho           = remote.Code(300)
	GetTopicsByCluster               = remote.Code(224)
	RegisterFilterServer             = remote.Code(301)
	RegisterMessageFilterClass       = remote.Code(302)
	QueryConsumeTimeSpan             = remote.Code(303)
	GetSystemTopicListFromNs         = remote.Code(304)
	GetSystemTopicListFromBroker     = remote.Code(305)
	CleanExpiredConsumequeue         = remote.Code(306)
	GetConsumerRunningInfo           = remote.Code(307)
	QueryCorrectionOffset            = remote.Code(308)
	ConsumeMessageDirectly           = remote.Code(309)
	SendMessageV2                    = remote.Code(310)
	GetUnitTopicList                 = remote.Code(311)
	GetHasUnitSubTopicList           = remote.Code(312)
	GetHasUnitSubUnunitTopicList     = remote.Code(313)
	CloneGroupOffset                 = remote.Code(314)
	ViewBrokerStatsData              = remote.Code(315)
	CleanUnusedTopic                 = remote.Code(316)
	GetBrokerConsumeStats            = remote.Code(317)
	UpdateNamesrvConfig              = remote.Code(318)
	GetNamesrvConfig                 = remote.Code(319)
	SendBatchMessage                 = remote.Code(320)
	QueryConsumeQueue                = remote.Code(321)
)

// response code
const (
	UnknowError                = remote.Code(-1)
	Success                    = remote.Code(0)
	SystemError                = remote.Code(1)
	SystemBusy                 = remote.Code(2)
	RequestCodeNotSupported    = remote.Code(3)
	TransactionFailed          = remote.Code(4)
	FlushDiskTimeout           = remote.Code(10)
	SlaveNotAvailable          = remote.Code(11)
	FlushSlaveTimeout          = remote.Code(12)
	MessageIllegal             = remote.Code(13)
	ServiceNotAvailable        = remote.Code(14)
	VersionNotSupported        = remote.Code(15)
	NoPermission               = remote.Code(16)
	TopicNotExist              = remote.Code(17)
	TopicExistAlready          = remote.Code(18)
	PullNotFound               = remote.Code(19)
	PullRetryImmediately       = remote.Code(20)
	PullOffsetMoved            = remote.Code(21)
	QueryNotFound              = remote.Code(22)
	SubscriptionParseFailed    = remote.Code(23)
	SubscriptionNotExist       = remote.Code(24)
	SubscriptionNotLatest      = remote.Code(25)
	SubscriptionGroupNotExist  = remote.Code(26)
	FilterDataNotExist         = remote.Code(27)
	FilterDataNotLatest        = remote.Code(28)
	TransactionShouldCommit    = remote.Code(200)
	TransactionShouldRollback  = remote.Code(201)
	TransactionStateUnknow     = remote.Code(202)
	TransactionStateGroupWrong = remote.Code(203)
	NoBuyerID                  = remote.Code(204)
	NotInCurrentUnit           = remote.Code(205)
	ConsumerNotOnline          = remote.Code(206)
	ConsumeMsgTimeout          = remote.Code(207)
	NoMessage                  = remote.Code(208)
	ConnectBrokerException     = remote.Code(10001)
	AccessBrokerException      = remote.Code(10002)
	BrokerNotExistException    = remote.Code(10003)
	NoNameServerException      = remote.Code(10004)
	NotFoundTopicException     = remote.Code(10005)
)

// RPC contains the rpc
type RPC struct {
	client remote.Client
}

// NewRPC create the remoting rpc
func NewRPC(c Client) *RPC {
	return &RPC{client: c}
}