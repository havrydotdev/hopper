package types

import (
	"io"
)

type String string

func (s *String) ReadFrom(r io.Reader) (n int64, err error) {
	var size VarInt
	sizeN, err := size.ReadFrom(r)
	if err != nil {
		return 0, err
	}

	b := make([]byte, size)
	if _, err := io.ReadFull(r, b); err != nil {
		return 0, err
	}

	*s = String(b)

	return int64(size) + sizeN, nil
}

func (s String) WriteTo(w io.Writer) (int64, error) {
	sizeN, err := VarInt(len([]byte(s))).WriteTo(w)
	if err != nil {
		return 0, err
	}

	n, err := w.Write([]byte(s))
	return int64(n) + sizeN, err
}
