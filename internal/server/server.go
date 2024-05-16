package server

import (
	"fmt"
	"log/slog"
	"net"

	"havry.dev/havry/hopper/internal/config"
)

// packet handler function
type HandleConn func(conn net.Conn) error

type Hopper struct {
	// port to start server on
	Config *config.Config

	// base64 encoded favicon
	favicon *string
}

// create new hopper server
func New(
	cfg *config.Config,
	faviconContent *string,
) *Hopper {
	return &Hopper{
		Config:  cfg,
		favicon: faviconContent,
	}
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
			// TODO replace with custom logger
			slog.Error("Error occurred while handling tcp conn: " + err.Error())
			continue
		}

		slog.Debug("New client connected", slog.String("addr", conn.RemoteAddr().String()))

		// handle connection in separate goroutine
		go func() {
			err := h.handshake(conn)
			if err != nil {
				slog.Error(err.Error())
			}
		}()
	}
}
