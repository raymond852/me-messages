package messages

var MessageTypeEnum = struct {
	PlaceOrder      int8
	ExecutionReport int8
	Symbol          int8
	Asset           int8
	CodeError       int8
}{
	PlaceOrder:      1,
	ExecutionReport: 2,
	Symbol:          3,
	Asset:           4,
	CodeError:       100,
}
