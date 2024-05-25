package types

import "io"

const (
	trueVal  byte = 0x01
	falseVal byte = 0x00
)

type Boolean bool

func (b Boolean) WriteTo(w io.Writer) (int64, error) {
	n, err := w.Write([]byte{b.GetValue()})

	return int64(n), err
}

func (b Boolean) GetValue() byte {
	if b {
		return trueVal
	}

	return falseVal
}
