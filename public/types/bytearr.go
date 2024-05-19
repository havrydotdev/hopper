package types

import "io"

type ByteArr []byte

func (b ByteArr) ReadFrom(r io.Reader) (int64, error) {
	n, err := io.ReadFull(r, b)
	return int64(n), err
}

func (b ByteArr) WriteTo(w io.Writer) (int64, error) {
	n, err := w.Write(b)
	return int64(n), err
}
