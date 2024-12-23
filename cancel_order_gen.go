package messages

// Code generated by github.com/tinylib/msgp DO NOT EDIT.

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *CancelOrder) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "CommonHeader":
			err = z.CommonHeader.DecodeMsg(dc)
			if err != nil {
				err = msgp.WrapError(err, "CommonHeader")
				return
			}
		case "u":
			z.UserId, err = dc.ReadInt32()
			if err != nil {
				err = msgp.WrapError(err, "UserId")
				return
			}
		case "s":
			z.SymbolId, err = dc.ReadInt32()
			if err != nil {
				err = msgp.WrapError(err, "SymbolId")
				return
			}
		case "o":
			z.OrderId, err = dc.ReadInt64()
			if err != nil {
				err = msgp.WrapError(err, "OrderId")
				return
			}
		case "c":
			z.ClOrderId, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "ClOrderId")
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *CancelOrder) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 5
	// write "CommonHeader"
	err = en.Append(0x85, 0xac, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72)
	if err != nil {
		return
	}
	err = z.CommonHeader.EncodeMsg(en)
	if err != nil {
		err = msgp.WrapError(err, "CommonHeader")
		return
	}
	// write "u"
	err = en.Append(0xa1, 0x75)
	if err != nil {
		return
	}
	err = en.WriteInt32(z.UserId)
	if err != nil {
		err = msgp.WrapError(err, "UserId")
		return
	}
	// write "s"
	err = en.Append(0xa1, 0x73)
	if err != nil {
		return
	}
	err = en.WriteInt32(z.SymbolId)
	if err != nil {
		err = msgp.WrapError(err, "SymbolId")
		return
	}
	// write "o"
	err = en.Append(0xa1, 0x6f)
	if err != nil {
		return
	}
	err = en.WriteInt64(z.OrderId)
	if err != nil {
		err = msgp.WrapError(err, "OrderId")
		return
	}
	// write "c"
	err = en.Append(0xa1, 0x63)
	if err != nil {
		return
	}
	err = en.WriteString(z.ClOrderId)
	if err != nil {
		err = msgp.WrapError(err, "ClOrderId")
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *CancelOrder) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 5
	// string "CommonHeader"
	o = append(o, 0x85, 0xac, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72)
	o, err = z.CommonHeader.MarshalMsg(o)
	if err != nil {
		err = msgp.WrapError(err, "CommonHeader")
		return
	}
	// string "u"
	o = append(o, 0xa1, 0x75)
	o = msgp.AppendInt32(o, z.UserId)
	// string "s"
	o = append(o, 0xa1, 0x73)
	o = msgp.AppendInt32(o, z.SymbolId)
	// string "o"
	o = append(o, 0xa1, 0x6f)
	o = msgp.AppendInt64(o, z.OrderId)
	// string "c"
	o = append(o, 0xa1, 0x63)
	o = msgp.AppendString(o, z.ClOrderId)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *CancelOrder) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "CommonHeader":
			bts, err = z.CommonHeader.UnmarshalMsg(bts)
			if err != nil {
				err = msgp.WrapError(err, "CommonHeader")
				return
			}
		case "u":
			z.UserId, bts, err = msgp.ReadInt32Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "UserId")
				return
			}
		case "s":
			z.SymbolId, bts, err = msgp.ReadInt32Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "SymbolId")
				return
			}
		case "o":
			z.OrderId, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "OrderId")
				return
			}
		case "c":
			z.ClOrderId, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "ClOrderId")
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *CancelOrder) Msgsize() (s int) {
	s = 1 + 13 + z.CommonHeader.Msgsize() + 2 + msgp.Int32Size + 2 + msgp.Int32Size + 2 + msgp.Int64Size + 2 + msgp.StringPrefixSize + len(z.ClOrderId)
	return
}
