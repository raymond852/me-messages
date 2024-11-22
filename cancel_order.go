package messages

//go:generate msgp

type CancelOrder struct {
	CommonHeader
	UserId    int32  `msg:"u"`
	SymbolId  int32  `msg:"s"`
	OrderId   int64  `msg:"o"`
	ClOrderId string `msg:"c"`
}
