package types

import (
	"io"
)

type Array[T io.WriterTo] []T

func (a Array[T]) WriteTo(w io.Writer) (n int64, err error) {
	var lenN int64
	lenN, err = VarInt(len(a)).WriteTo(w)
	if err != nil {
		return
	}
	n += lenN

	for _, elem := range []T(a) {
		var elemN int64
		elemN, err = elem.WriteTo(w)
		if err != nil {
			return
		}

		n += elemN
	}

	return
}
