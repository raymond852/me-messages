package messages

// Code generated by github.com/tinylib/msgp DO NOT EDIT.

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *DeleteUserTierCommission) DecodeMsg(dc *msgp.Reader) (err error) {
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
		case "Tier":
			z.Tier, err = dc.ReadInt32()
			if err != nil {
				err = msgp.WrapError(err, "Tier")
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
func (z *DeleteUserTierCommission) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 2
	// write "CommonHeader"
	err = en.Append(0x82, 0xac, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72)
	if err != nil {
		return
	}
	err = z.CommonHeader.EncodeMsg(en)
	if err != nil {
		err = msgp.WrapError(err, "CommonHeader")
		return
	}
	// write "Tier"
	err = en.Append(0xa4, 0x54, 0x69, 0x65, 0x72)
	if err != nil {
		return
	}
	err = en.WriteInt32(z.Tier)
	if err != nil {
		err = msgp.WrapError(err, "Tier")
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *DeleteUserTierCommission) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "CommonHeader"
	o = append(o, 0x82, 0xac, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72)
	o, err = z.CommonHeader.MarshalMsg(o)
	if err != nil {
		err = msgp.WrapError(err, "CommonHeader")
		return
	}
	// string "Tier"
	o = append(o, 0xa4, 0x54, 0x69, 0x65, 0x72)
	o = msgp.AppendInt32(o, z.Tier)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *DeleteUserTierCommission) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
		case "Tier":
			z.Tier, bts, err = msgp.ReadInt32Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Tier")
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
func (z *DeleteUserTierCommission) Msgsize() (s int) {
	s = 1 + 13 + z.CommonHeader.Msgsize() + 5 + msgp.Int32Size
	return
}
