package types

import "io"

// Read exactly one byte from io.Reader
func readByte(r io.Reader) (byte, error) {
	// if io.Reader is io.ByteReader, just call ReadByte function
	if r, ok := r.(io.ByteReader); ok {
		return r.ReadByte()
	}

	var b [1]byte
	if _, err := io.ReadFull(r, b[:]); err != nil {
		return 0, err
	}

	return b[0], nil
}

// write one byte to io.Writer
func writeByte(w io.Writer, c byte) error {
	if w, ok := w.(io.ByteWriter); ok {
		return w.WriteByte(c)
	}

	if _, err := w.Write([]byte{c}); err != nil {
		return err
	}

	return nil
}
