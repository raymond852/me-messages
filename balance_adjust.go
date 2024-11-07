package messages

//go:generate msgp

type BalanceAdjust struct {
	CommonHeader
	UserId  int32 `msg:"u"`
	AssetId int32 `msg:"a"`
	Delta   int64 `msg:"d"`
	Type    int8  `msg:"p"`
}
