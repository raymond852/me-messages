package messages

//go:generate msgp

// ExecutionReport see https://binance-docs.github.io/apidocs/spot/en/#payload-order-update
type ExecutionReport struct {
	CommonHeader `msg:",flatten"`

	AccountId            int32  `msg:"a"`
	SymbolId             int32  `msg:"s"`
	Side                 int8   `msg:"S"`
	OrderType            int8   `msg:"o"`
	ClientOrderId        string `msg:"c"`
	TimeInForce          int8   `msg:"f,omitempty"`
	Price                int64  `msg:"p"`
	Quantity             int64  `msg:"q"`
	StopPrice            int64  `msg:"P"`
	OrderStatus          int8   `msg:"X"`
	ExecutionType        int8   `msg:"x"`
	OrderId              int64  `msg:"i"`
	LastQty              int64  `msg:"l"`
	LastPrice            int64  `msg:"L"`
	CumQty               int64  `msg:"z"`
	Commission           int64  `msg:"n"`
	CommissionAssetId    int32  `msg:"N"`
	TransactionTime      int64  `msg:"T"`
	TradeId              int64  `msg:"t"`
	IsMaker              bool   `msg:"m"`
	CreateTime           int64  `msg:"O"`
	CumQuoteQty          int64  `msg:"Z"`
	QuoteOrderQty        int64  `msg:"Q,omitempty"`
	StrategyId           int64  `msg:"j,omitempty"`
	StrategyType         int8   `msg:"J"`
	SelfTradePreventMode int8   `msg:"V"`
	TradeGroupId         int64  `msg:"u"`
	CounterOrderId       int64  `msg:"U,omitempty"`
	PreventedQty         int64  `msg:"A,omitempty"`
	LastPreventedQty     int64  `msg:"B,omitempty"`
	CounterPartyUserId   int32  `msg:"k,omitempty"`
	CounterPartyOrderId  int64  `msg:"K,omitempty"`
}
