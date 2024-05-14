package server

import (
	"encoding"
	"io"

	"havry.dev/havry/hopper/internal/protocol/types"
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

func WriteResp(w io.Writer, packetID int, resp encoding.BinaryMarshaler) (size types.VarInt, err error) {
	respEncoded, err := resp.MarshalBinary()
	if err != nil {
		return nilVarInt, err
	}

	packetIdEncoded, err := types.VarInt(packetID).MarshalBinary()
	if err != nil {
		return nilVarInt, err
	}

	size = types.VarInt(len(respEncoded) + len(packetIdEncoded))
	sizeEncoded, err := size.MarshalBinary()
	if err != nil {
		return nilVarInt, err
	}

	res := append(sizeEncoded, packetIdEncoded...)
	res = append(res, respEncoded...)
	_, err = w.Write(res)

	return size, err
}
