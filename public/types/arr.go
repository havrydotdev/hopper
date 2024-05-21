package types

import (
	"io"
	"log/slog"
)

type Array[T io.WriterTo] []T

func (a Array[T]) WriteTo(w io.Writer) (n int64, err error) {
	var lenN int64
	lenN, err = VarInt(len(a)).WriteTo(w)
	if err != nil {
		return
	}
	slog.Debug("", slog.Int("len", len(a)))
	n += lenN

	for _, elem := range []T(a) {
		slog.Debug("", slog.Any("elem", elem))
		var elemN int64
		elemN, err = elem.WriteTo(w)
		if err != nil {
			return
		}

		n += elemN
	}

	return
}
