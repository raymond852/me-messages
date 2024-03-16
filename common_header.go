package messages

//go:generate msgp

type CommonHeader struct {
	SourceKafkaTopic  int8  `msg:"ct,omitempty"`
	SourceKafkaOffset int64 `msg:"cO,omitempty"`
	SourceSeqNum      int64 `msg:"cN,omitempty"`
	SeqNum            int64 `msg:"cS,omitempty"`
	KafkaOffset       int64 `msg:"cK,omitempty"`
	TransactionId     int64 `msg:"cT,omitempty"`
	InputTime         int64 `msg:"cA,omitempty"`
	SendTime          int64 `msg:"ca,omitempty"`
	IsLastTransaction bool  `msg:"cL"`
}

func (z *CommonHeader) SetSourceKafkaTopic(topicId int8) {
	z.SourceKafkaTopic = topicId
}

func (z *CommonHeader) GetSourceKafkaTopic() int8 {
	return z.SourceKafkaTopic
}

func (z *CommonHeader) SetSourceKafkaOffset(offset int64) {
	z.SourceKafkaOffset = offset
}

func (z *CommonHeader) GetSourceKafkaOffset() int64 {
	return z.SourceKafkaOffset
}

func (z *CommonHeader) SetSourceSeqNum(seqNum int64) {
	z.SourceSeqNum = seqNum
}

func (z *CommonHeader) GetSourceSeqNum() int64 {
	return z.SourceSeqNum
}

func (z *CommonHeader) SetSeqNum(seqNum int64) {
	z.SeqNum = seqNum
}

func (z *CommonHeader) GetSeqNum() int64 {
	return z.SeqNum
}

func (z *CommonHeader) SetKafkaOffset(offset int64) {
	z.KafkaOffset = offset
}

func (z *CommonHeader) GetKafkaOffset() int64 {
	return z.KafkaOffset
}

func (z *CommonHeader) SetTransactionId(transId int64) {
	z.TransactionId = transId
}

func (z *CommonHeader) GetTransactionId() int64 {
	return z.TransactionId
}

func (z *CommonHeader) SetIsLastTransaction(isLastTransaction bool) {
	z.IsLastTransaction = isLastTransaction
}

func (z *CommonHeader) GetIsLastTransaction() bool {
	return z.IsLastTransaction
}

func (z *CommonHeader) SetInputTime(inputTime int64) {
	z.InputTime = inputTime
}

func (z *CommonHeader) GetInputTime() int64 {
	return z.InputTime
}

func (z *CommonHeader) SetSendTime(sendTime int64) {
	z.SendTime = sendTime
}
func (z *CommonHeader) GetSendTime() int64 {
	return z.SendTime
}
