package main

import (
	"github.com/dpatrie/go-bots/server"
)

func main() {
	config := server.ServerConfig{
		Host:             "0.0.0.0:6666",
		MaxGame:          10,
		MaxClientPerGame: 10,
	}

	s := server.New(config)
	s.Listen()
}
