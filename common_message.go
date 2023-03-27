package messages

type CommonMessage interface {
	SetSourceKafkaTopic(topicId int8)
	SetSourceKafkaOffset(offset int64)
	SetSourceSeqNum(seqNum int64)
	SetSeqNum(seqNum int64)
	SetKafkaOffset(offset int64)
	SetTransactionId(transId int64)
	SetIsLastTransaction(isLastTransaction bool)
	SetInputTime(inputTime int64)
	SetSendTime(sendTime int64)
}
