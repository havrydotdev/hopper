package types

import (
	"encoding/binary"
	"io"
)

const (
	LongBytes = 8
)

type Long int64

func (l *Long) ReadFrom(r io.Reader) (int64, error) {
	return LongBytes, binary.Read(r, binary.BigEndian, l)
}

func (l Long) WriteTo(w io.Writer) (int64, error) {
	return LongBytes, binary.Write(w, binary.BigEndian, l)
}
