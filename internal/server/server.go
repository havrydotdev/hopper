package server

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"fmt"
	"log/slog"
	"net"

	"github.com/gavrylenkoIvan/hopper/internal/config"
	"github.com/gavrylenkoIvan/hopper/internal/hopper"
)

const PubKeyBits = 1024

type Hopper struct {
	// Port to start server on
	Config *config.Config

	// Base64 encoded favicon
	favicon *string

	privKey *rsa.PrivateKey
	// x509-encoded RSA public key
	pubKey []byte
}

// create new hopper server
func New(
	cfg *config.Config,
	faviconContent *string,
) (*Hopper, error) {
	privKey, err := rsa.GenerateKey(rand.Reader, PubKeyBits)
	if err != nil {
		return nil, err
	}

	pubKey, err := x509.MarshalPKIXPublicKey(&privKey.PublicKey)
	if err != nil {
		return nil, err
	}

	slog.Debug("Generated public key successfully")

	return &Hopper{
		Config:  cfg,
		favicon: faviconContent,

		privKey: privKey,
		pubKey:  pubKey,
	}, nil
}

// returns address to start on
func (h *Hopper) Addr() string {
	return fmt.Sprintf("0.0.0.0:%d", h.Config.Server.Port)
}

func (h *Hopper) Listen() error {
	serverAddr := h.Addr()

	// open tcp server on specified port
	listener, err := net.Listen("tcp", serverAddr)
	if err != nil {
		return err
	}
	defer listener.Close()

	slog.Info(fmt.Sprintf("Hopper Server Is Listening On %s", serverAddr))

	// start listening for tcp connections
	for {
		conn, err := listener.Accept()
		if err != nil {
			slog.Error("Error occurred while handling tcp conn: " + err.Error())
			continue
		}

		slog.Debug("New client connected", slog.String("addr", conn.RemoteAddr().String()))

		// handle connection in separate goroutine
		go func() {
			err := h.handshake(hopper.NewConn(conn))
			if err != nil {
				slog.Error(err.Error())
			}
		}()
	}
}
