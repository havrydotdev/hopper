package types

import (
	"io"
)

type ByteArr []byte

func (b *ByteArr) ReadFrom(r io.Reader) (int64, error) {
	var size VarInt

	// read byte array size
	sizeN, err := size.ReadFrom(r)
	if err != nil {
		return 0, err
	}

	// allocate an array
	*b = ByteArr(make([]byte, size))

	// read all bytes from reader
	n, err := io.ReadFull(r, *b)

	return int64(n) + sizeN, err
}

func (b ByteArr) WriteTo(w io.Writer) (int64, error) {
	// write array size to writer
	sizeN, err := VarInt(len(b)).WriteTo(w)
	if err != nil {
		return 0, err
	}

	// write array content
	n, err := w.Write(b)

	return sizeN + int64(n), err
}
