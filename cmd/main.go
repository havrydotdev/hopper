package main

import (
	"log"

	"havry.dev/havry/hopper/internal/server"
)

const (
	defaultPort = 25565
)

func main() {
	hopper := server.New(defaultPort)

	err := hopper.Listen()
	if err != nil {
		log.Fatal("Failed to start Hopper server: ", err)
	}
}
