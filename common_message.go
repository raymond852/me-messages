package messages

type CommonMessage interface {
	GetCommonHeader() CommonHeader

	SetSourceKafkaTopic(topicId int8)
	GetSourceKafkaTopic() int8

	SetSourceKafkaOffset(offset int64)
	GetSourceKafkaOffset() int64

	SetSourceSeqNum(seqNum int64)
	GetSourceSeqNum() int64

	SetSeqNum(seqNum int64)
	GetSeqNum() int64

	SetKafkaOffset(offset int64)
	GetKafkaOffset() int64

	SetTransactionId(transId int64)
	GetTransactionId() int64

	SetIsLastTransaction(isLastTransaction bool)
	GetIsLastTransaction() bool

	SetInputTime(inputTime int64)
	GetInputTime() int64

	SetSendTime(sendTime int64)
	GetSendTime() int64
}
