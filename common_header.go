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

func (z CommonHeader) GetCommonHeader() CommonHeader {
	return z
}
