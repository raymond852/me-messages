package messages

//go:generate msgp
type Symbol struct {
	CommonHeader
	BaseAssetId     int32  `msg:"b"`
	QuoteAssetId    int32  `msg:"q"`
	Id              int32  `msg:"i"`
	Name            string `msg:"n"`
	QuantityDecimal int8   `msg:"Q"`
	PriceDecimal    int8   `msg:"P"`
	MinNotional     int64  `msg:"m"`
	StepSize        int64  `msg:"s"`
	TickSize        int64  `msg:"t"`
}
