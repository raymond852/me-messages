package messages

//go:generate msgp
type CodeError struct {
	CommonHeader
	Code    int    `msg:"code"`
	Message string `msg:"message"`
}
