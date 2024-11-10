package messages

//go:generate msgp

type DeleteSymbolTierCommission struct {
	CommonHeader
	SymbolId int32
	Tier       int32
}
