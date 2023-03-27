package messages

//go:generate msgp
type Asset struct {
	CommonHeader
	AssetId int32  `msg:"a,omitempty"`
	Name    string `msg:"n"`
	Decimal int8   `msg:"d"`
}
