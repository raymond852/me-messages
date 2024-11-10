package messages

//go:generate msgp

type DeleteUserTierCommission struct {
	CommonHeader
	Tier int32
}
