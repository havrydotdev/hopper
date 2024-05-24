package packet

import (
	"bytes"
	"io"
)

func Marshal(fields ...io.WriterTo) ([]byte, error) {
	buf := bytes.NewBuffer(nil)

	for _, field := range fields {
		_, err := field.WriteTo(buf)
		if err != nil {
			return nil, err
		}
	}

	return buf.Bytes(), nil
}

func Unmarshal(r io.Reader, fields ...io.ReaderFrom) error {
	for _, field := range fields {
		_, err := field.ReadFrom(r)
		if err != nil {
			return err
		}
	}

	return nil
}
