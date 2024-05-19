package server

import (
	"bytes"
	"io"

	"github.com/gavrylenkoIvan/hopper/public/types"
)

var (
	nilVarInt = types.VarInt(0)
)

// Reads packet from io.Reader
// Second parameter is point to packet
// Returns packet's size, id and error if occurred
func ReadPacket(r io.Reader, p io.ReaderFrom) (size, packetID types.VarInt, err error) {
	size, packetID, err = ReadPacketInfo(r)
	if err != nil {
		return nilVarInt, nilVarInt, err
	}

	_, err = p.ReadFrom(r)
	if err != nil {
		return nilVarInt, nilVarInt, err
	}

	return
}

// Reads packet size and id from io.Reader
func ReadPacketInfo(r io.Reader) (size, packetID types.VarInt, err error) {
	_, err = size.ReadFrom(r)
	if err != nil {
		return nilVarInt, nilVarInt, err
	}

	_, err = packetID.ReadFrom(r)
	if err != nil {
		return nilVarInt, nilVarInt, err
	}

	return
}

func WritePacket(w io.Writer, packetID int, p io.WriterTo) (size types.VarInt, err error) {
	// marshal packet
	buf, err := MarshalPacket(packetID, p)
	if err != nil {
		return nilVarInt, err
	}

	return WriteRaw(w, buf)
}

// Writes buf to io.Writer, appending it's length with types.VarInt
func WriteRaw(w io.Writer, buf []byte) (size types.VarInt, err error) {
	res := bytes.NewBuffer(nil)
	size = types.VarInt(len(buf))
	// write response size to buffer
	_, err = size.WriteTo(w)
	if err != nil {
		return nilVarInt, err
	}

	// write response body to buffer
	_, err = res.Write(buf)
	if err != nil {
		return nilVarInt, err
	}

	// write all buffer content to w
	_, err = w.Write(res.Bytes())

	return
}

func MarshalPacket(id int, p io.WriterTo) ([]byte, error) {
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

func PrependID(id int, p []byte) ([]byte, error) {
	res := bytes.NewBuffer(nil)
	_, err := types.VarInt(id).WriteTo(res)
	if err != nil {
		return nil, err
	}

	_, err = res.Write(p)

	return res.Bytes(), err
}
