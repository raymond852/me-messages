package messages

// Code generated by github.com/tinylib/msgp DO NOT EDIT.

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *Symbol) DecodeMsg(dc *msgp.Reader) (err error) {
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
		case "b":
			z.BaseAssetId, err = dc.ReadInt32()
			if err != nil {
				err = msgp.WrapError(err, "BaseAssetId")
				return
			}
		case "q":
			z.QuoteAssetId, err = dc.ReadInt32()
			if err != nil {
				err = msgp.WrapError(err, "QuoteAssetId")
				return
			}
		case "i":
			z.Id, err = dc.ReadInt32()
			if err != nil {
				err = msgp.WrapError(err, "Id")
				return
			}
		case "n":
			z.Name, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "Name")
				return
			}
		case "Q":
			z.QuantityDecimal, err = dc.ReadInt8()
			if err != nil {
				err = msgp.WrapError(err, "QuantityDecimal")
				return
			}
		case "P":
			z.PriceDecimal, err = dc.ReadInt8()
			if err != nil {
				err = msgp.WrapError(err, "PriceDecimal")
				return
			}
		case "m":
			z.MinNotional, err = dc.ReadInt64()
			if err != nil {
				err = msgp.WrapError(err, "MinNotional")
				return
			}
		case "s":
			z.StepSize, err = dc.ReadInt64()
			if err != nil {
				err = msgp.WrapError(err, "StepSize")
				return
			}
		case "t":
			z.TickSize, err = dc.ReadInt64()
			if err != nil {
				err = msgp.WrapError(err, "TickSize")
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
func (z *Symbol) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 10
	// write "CommonHeader"
	err = en.Append(0x8a, 0xac, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72)
	if err != nil {
		return
	}
	err = z.CommonHeader.EncodeMsg(en)
	if err != nil {
		err = msgp.WrapError(err, "CommonHeader")
		return
	}
	// write "b"
	err = en.Append(0xa1, 0x62)
	if err != nil {
		return
	}
	err = en.WriteInt32(z.BaseAssetId)
	if err != nil {
		err = msgp.WrapError(err, "BaseAssetId")
		return
	}
	// write "q"
	err = en.Append(0xa1, 0x71)
	if err != nil {
		return
	}
	err = en.WriteInt32(z.QuoteAssetId)
	if err != nil {
		err = msgp.WrapError(err, "QuoteAssetId")
		return
	}
	// write "i"
	err = en.Append(0xa1, 0x69)
	if err != nil {
		return
	}
	err = en.WriteInt32(z.Id)
	if err != nil {
		err = msgp.WrapError(err, "Id")
		return
	}
	// write "n"
	err = en.Append(0xa1, 0x6e)
	if err != nil {
		return
	}
	err = en.WriteString(z.Name)
	if err != nil {
		err = msgp.WrapError(err, "Name")
		return
	}
	// write "Q"
	err = en.Append(0xa1, 0x51)
	if err != nil {
		return
	}
	err = en.WriteInt8(z.QuantityDecimal)
	if err != nil {
		err = msgp.WrapError(err, "QuantityDecimal")
		return
	}
	// write "P"
	err = en.Append(0xa1, 0x50)
	if err != nil {
		return
	}
	err = en.WriteInt8(z.PriceDecimal)
	if err != nil {
		err = msgp.WrapError(err, "PriceDecimal")
		return
	}
	// write "m"
	err = en.Append(0xa1, 0x6d)
	if err != nil {
		return
	}
	err = en.WriteInt64(z.MinNotional)
	if err != nil {
		err = msgp.WrapError(err, "MinNotional")
		return
	}
	// write "s"
	err = en.Append(0xa1, 0x73)
	if err != nil {
		return
	}
	err = en.WriteInt64(z.StepSize)
	if err != nil {
		err = msgp.WrapError(err, "StepSize")
		return
	}
	// write "t"
	err = en.Append(0xa1, 0x74)
	if err != nil {
		return
	}
	err = en.WriteInt64(z.TickSize)
	if err != nil {
		err = msgp.WrapError(err, "TickSize")
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Symbol) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 10
	// string "CommonHeader"
	o = append(o, 0x8a, 0xac, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72)
	o, err = z.CommonHeader.MarshalMsg(o)
	if err != nil {
		err = msgp.WrapError(err, "CommonHeader")
		return
	}
	// string "b"
	o = append(o, 0xa1, 0x62)
	o = msgp.AppendInt32(o, z.BaseAssetId)
	// string "q"
	o = append(o, 0xa1, 0x71)
	o = msgp.AppendInt32(o, z.QuoteAssetId)
	// string "i"
	o = append(o, 0xa1, 0x69)
	o = msgp.AppendInt32(o, z.Id)
	// string "n"
	o = append(o, 0xa1, 0x6e)
	o = msgp.AppendString(o, z.Name)
	// string "Q"
	o = append(o, 0xa1, 0x51)
	o = msgp.AppendInt8(o, z.QuantityDecimal)
	// string "P"
	o = append(o, 0xa1, 0x50)
	o = msgp.AppendInt8(o, z.PriceDecimal)
	// string "m"
	o = append(o, 0xa1, 0x6d)
	o = msgp.AppendInt64(o, z.MinNotional)
	// string "s"
	o = append(o, 0xa1, 0x73)
	o = msgp.AppendInt64(o, z.StepSize)
	// string "t"
	o = append(o, 0xa1, 0x74)
	o = msgp.AppendInt64(o, z.TickSize)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Symbol) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
		case "b":
			z.BaseAssetId, bts, err = msgp.ReadInt32Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "BaseAssetId")
				return
			}
		case "q":
			z.QuoteAssetId, bts, err = msgp.ReadInt32Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "QuoteAssetId")
				return
			}
		case "i":
			z.Id, bts, err = msgp.ReadInt32Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Id")
				return
			}
		case "n":
			z.Name, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Name")
				return
			}
		case "Q":
			z.QuantityDecimal, bts, err = msgp.ReadInt8Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "QuantityDecimal")
				return
			}
		case "P":
			z.PriceDecimal, bts, err = msgp.ReadInt8Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "PriceDecimal")
				return
			}
		case "m":
			z.MinNotional, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "MinNotional")
				return
			}
		case "s":
			z.StepSize, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "StepSize")
				return
			}
		case "t":
			z.TickSize, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "TickSize")
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
func (z *Symbol) Msgsize() (s int) {
	s = 1 + 13 + z.CommonHeader.Msgsize() + 2 + msgp.Int32Size + 2 + msgp.Int32Size + 2 + msgp.Int32Size + 2 + msgp.StringPrefixSize + len(z.Name) + 2 + msgp.Int8Size + 2 + msgp.Int8Size + 2 + msgp.Int64Size + 2 + msgp.Int64Size + 2 + msgp.Int64Size
	return
}