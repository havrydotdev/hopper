package server

import (
	"fmt"
	"log"
	"net"
)

// packet handler function
type HandleConn func(conn net.Conn) error

type Hopper struct {
	// port to start server on
	port int
}

// create new hopper server
func New(port int) *Hopper {
	return &Hopper{port}
}

func (h *Hopper) Listen() error {
	// open tcp server on specified port
	listener, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", h.port))
	if err != nil {
		return err
	}
	defer listener.Close()

	log.Printf("Hopper Server Is Listening On 0.0.0.0:%d", h.port)
	// start listening for tcp connections
	for {
		conn, err := listener.Accept()
		if err != nil {
			// TODO replace with custom logger
			log.Print("Error occurred when handling tcp conn: " + err.Error())
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
