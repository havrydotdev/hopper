package types

import (
	"encoding/binary"
	"io"
)

const (
	// Short is always 2 bytes long (from docs)
	UShortBytes = 2
)

type UShort uint16

func (u *UShort) ReadFrom(r io.Reader) (n int64, err error) {
	return UShortBytes, binary.Read(r, binary.BigEndian, u)
}
