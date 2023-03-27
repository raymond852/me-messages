package messages

var KafkaTopic = struct {
	MeIn  int8
	MeOut int8
}{
	MeIn:  1,
	MeOut: 2,
}

var OrderType = struct {
	Limit  int8
	Market int8
}{
	Limit:  1,
	Market: 2,
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
