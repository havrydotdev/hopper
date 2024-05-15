package types

import (
	"errors"
	"io"
)

const (
	MaxVarIntLen = 5

	segmentBits byte = 0x7F
	continueBit byte = 0x80
)

var ErrTooBig = errors.New("var int is too big")

// Protocol's VarInt data type
// VarInt on wiki.vg https://wiki.vg/Data_types#Type:VarInt
type VarInt int

// Implement io.ReaderFrom for VarInt
func (v *VarInt) ReadFrom(r io.Reader) (n int64, err error) {
	var val uint32

	for curr := continueBit; curr&continueBit != 0; n++ {
		if n > MaxVarIntLen {
			return 0, ErrTooBig
		}

		curr, err = readByte(r)
		if err != nil {
			return 0, err
		}

		val |= uint32(curr&segmentBits) << uint32(7*n)
	}

	*v = VarInt(val)

	return
}

// Implement io.Writer for VarInt
func (v VarInt) WriteTo(w io.Writer) (n int64, err error) {
	num := uint32(v)
	for {
		b := num & uint32(segmentBits)
		num >>= 7
		if num != 0 {
			b |= uint32(continueBit)
		}

		err = writeByte(w, byte(b))
		if err != nil {
			return 0, err
		}
		n++

		if num == 0 {
			break
		}
	}

	return n, nil
}
