package messages

//go:generate msgp

type UserTierCommission struct {
	CommonHeader
	Tier       int32
	Commission int32
}
