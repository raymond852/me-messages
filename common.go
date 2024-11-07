package messages

var KafkaTopic = struct {
	MeIn  int8
	MeOut int8
}{
	MeIn:  1,
	MeOut: 2,
}

var OrderType = struct {
	Limit           int8
	Market          int8
	StopLoss        int8
	StopLossLimit   int8
	TakeProfit      int8
	TakeProfitLimit int8
	LimitMaker      int8
}{
	Limit:           1,
	Market:          2,
	StopLoss:        3,
	StopLossLimit:   4,
	TakeProfit:      5,
	TakeProfitLimit: 6,
	LimitMaker:      7,
}

var Side = struct {
	Buy  int8
	Sell int8
}{
	Buy:  1,
	Sell: 2,
}

var TimeInForce = struct {
	GTC int8
	IOC int8
	FOK int8
}{
	GTC: 1,
	IOC: 2,
	FOK: 3,
}

var OrderStatus = struct {
	New             int8
	PartiallyFilled int8
	Filled          int8
	Expired         int8
	Cancelled       int8
	ExpiredInMatch  int8
}{
	New:             1,
	PartiallyFilled: 2,
	Filled:          3,
	Expired:         4,
	Cancelled:       5,
	ExpiredInMatch:  6,
}

var OrderExecType = struct {
	New             int8
	Trade           int8
	Cancelled       int8
	Expired         int8
	TradePrevention int8
}{
	New:             1,
	Trade:           2,
	Cancelled:       3,
	Expired:         4,
	TradePrevention: 5,
}

var BalanceAdjustType = struct {
	Deposit  int8
	Withdraw int8
}{
	Deposit:  1,
	Withdraw: 2,
}

var BalanceUpdateType = struct {
	Order    int8
	Deposit  int8
	Withdraw int8
}{
	Order:    1,
	Deposit:  2,
	Withdraw: 3,
}
