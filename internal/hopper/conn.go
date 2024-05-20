package hopper

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"net"

	"github.com/gavrylenkoIvan/hopper/internal/cfb8"
	cbound "github.com/gavrylenkoIvan/hopper/public/clientbound"
	sbound "github.com/gavrylenkoIvan/hopper/public/serverbound"
	"github.com/gavrylenkoIvan/hopper/public/types"
)

var (
	nilVarInt = types.VarInt(0)
)

type Conn struct {
	encrypted    bool
	sharedSecret []byte

	net.Conn
}

func NewConn(raw net.Conn) *Conn {
	return &Conn{false, nil, raw}
}

func (c *Conn) SetSharedSecret(sharedSecret []byte) {
	c.encrypted = true
	c.sharedSecret = sharedSecret
}

// Reads packet from conn
// Returns packet's size, id and error if occurred
func (c *Conn) ReadPacket(p sbound.Packet) (size, packetID types.VarInt, err error) {
	size, packetID, err = c.ReadPacketInfo()
	if err != nil {
		return nilVarInt, nilVarInt, err
	}

	_, err = p.ReadFrom(c)
	if err != nil {
		return nilVarInt, nilVarInt, err
	}

	return
}

// Reads packet size and id from conn
func (c *Conn) ReadPacketInfo() (size, packetID types.VarInt, err error) {
	_, err = size.ReadFrom(c)
	if err != nil {
		return nilVarInt, nilVarInt, err
	}

	_, err = packetID.ReadFrom(c)
	if err != nil {
		return nilVarInt, nilVarInt, err
	}

	return
}

func (c *Conn) WritePacket(p cbound.Packet) (size types.VarInt, err error) {
	// marshal packet
	buf, err := marshalPacket(p.ID(), p)
	if err != nil {
		return nilVarInt, err
	}

	return c.WriteRaw(buf)
}

// Writes buf into conn, appending it's length with types.VarInt
func (c *Conn) WriteRaw(buf []byte) (types.VarInt, error) {
	res := bytes.NewBuffer(nil)
	size := types.VarInt(len(buf))
	// write response size into buffer
	_, err := size.WriteTo(c)
	if err != nil {
		return nilVarInt, err
	}

	// write response body into buffer
	_, err = res.Write(buf)
	if err != nil {
		return nilVarInt, err
	}

	// encrypt response
	dst := res.Bytes()
	if c.encrypted {
		var block cipher.Block
		block, err = aes.NewCipher(c.sharedSecret)
		if err != nil {
			return nilVarInt, err
		}

		stream := cfb8.NewEncrypter(block, c.sharedSecret)

		stream.XORKeyStream(dst, dst)
	}

	// write all buffer content into w
	_, err = c.Write(dst)

	return size, err
}
