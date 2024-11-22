package messages

//go:generate msgp

type PlaceOrder struct {
	CommonHeader
	ClOrderId               string `msg:"c"`
	UserId                  int32  `msg:"u"`
	SymbolId                int32  `msg:"s"`
	Side                    int8   `msg:"S"`
	Type                    int8   `msg:"t"`
	TimeInForce             int8   `msg:"T"`
	Quantity                int64  `msg:"q"`
	QuoteOrderQty           int64  `msg:"Q"`
	Price                   int64  `msg:"p"`
	NewClientOrderId        string `msg:"C"`
	StopPrice               int64  `msg:"P"`
	SelfTradePreventionMode int8   `msg:"V"`
}
