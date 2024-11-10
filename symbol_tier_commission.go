package messages

//go:generate msgp

type SymbolTierCommission struct {
	CommonHeader
	SymbolId   int32
	Tier       int32
	Commission int32
}
