package hopper

import (
	"crypto/aes"
	"crypto/cipher"
	"io"
	"net"

	"github.com/gavrylenkoIvan/hopper/public/cfb8"
	"github.com/gavrylenkoIvan/hopper/public/packet"
	"github.com/gavrylenkoIvan/hopper/public/types"
)

type Conn struct {
	// connection encrypter and decrypter
	encrypter cipher.Stream
	decrypter cipher.Stream

	// raw connection
	net.Conn
}

func NewConn(raw net.Conn) *Conn {
	return &Conn{nil, nil, raw}
}

// Sets encryption and decryption for conn
func (c *Conn) SetEncryption(sharedSecret []byte) error {
	block, err := aes.NewCipher(sharedSecret)
	if err != nil {
		return err
	}

	c.encrypter = cfb8.NewEncrypter(block, sharedSecret)
	c.decrypter = cfb8.NewDecrypter(block, sharedSecret)

	return nil
}

// Reads packet from conn and returns
// it's size, id and error if occurred
func (c *Conn) ReadPacket(fields ...io.ReaderFrom) (size, packetID types.VarInt, err error) {
	size, packetID, err = c.ReadPacketInfo()
	if err != nil {
		return
	}

	return size, packetID, packet.Unmarshal(c, fields...)
}

// Reads packet size and id from conn
func (c *Conn) ReadPacketInfo() (size, packetID types.VarInt, err error) {
	err = packet.Unmarshal(c, &size, &packetID)
	return
}

func (c *Conn) Read(b []byte) (n int, err error) {
	if c.decrypter != nil {
		c.decrypter.XORKeyStream(b, b)
	}

	return c.Conn.Read(b)
}

func (c *Conn) Write(b []byte) (n int, err error) {
	if c.encrypter != nil {
		c.encrypter.XORKeyStream(b, b)
	}

	return c.Conn.Write(b)
}
