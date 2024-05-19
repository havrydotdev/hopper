package hopper

import (
	"bytes"
	"io"

	"github.com/gavrylenkoIvan/hopper/public/types"
)

func marshalPacket(id int, p io.WriterTo) ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	// write packet id into buf
	_, err := types.VarInt(id).WriteTo(buf)
	if err != nil {
		return nil, err
	}

	// write packet's content into buf
	_, err = p.WriteTo(buf)

	return buf.Bytes(), err
}
