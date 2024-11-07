package messages

//go:generate msgp

type BalanceUpdate struct {
	CommonHeader
	UserId  int32 `msg:"u"`
	AssetId int32 `msg:"a"`
	Total   int64 `msg:"t"`
	Locked  int64 `msg:"l"`
	Delta   int64 `msg:"d"`
	Type    int8  `msg:"p"`
}
