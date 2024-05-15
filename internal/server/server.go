package server

import (
	"fmt"
	"log"
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

func (h *Hopper) Listen() error {
	// open tcp server on specified port
	listener, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", h.Config.Server.Port))
	if err != nil {
		return err
	}
	defer listener.Close()

	log.Printf("Hopper Server Is Listening On 0.0.0.0:%d", h.Config.Server.Port)
	// start listening for tcp connections
	for {
		conn, err := listener.Accept()
		if err != nil {
			// TODO replace with custom logger
			log.Print("Error occurred while handling tcp conn: " + err.Error())
			continue
		}

		// handle connection in separate goroutine
		go func() {
			err := h.handleConn(conn)
			if err != nil {
				log.Println(err)
			}
		}()
	}
}
