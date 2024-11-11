package messages

var MessageTypeEnum = struct {
	PlaceOrder                 int8
	ExecutionReport            int8
	Symbol                     int8
	Asset                      int8
	BalanceAdjust              int8
	BalanceUpdate              int8
	UserTierCommission         int8
	DeleteUserTierCommission   int8
	SymbolTierCommission       int8
	DeleteSymbolTierCommission int8
	CodeError                  int8
}{
	PlaceOrder:                 1,
	ExecutionReport:            2,
	Symbol:                     3,
	Asset:                      4,
	BalanceAdjust:              5,
	BalanceUpdate:              6,
	UserTierCommission:         7,
	DeleteUserTierCommission:   8,
	SymbolTierCommission:       9,
	DeleteSymbolTierCommission: 10,
	CodeError:                  100,
}
