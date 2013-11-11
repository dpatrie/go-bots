package main

import (
	"github.com/dpatrie/go-bots/server"
	"log"
	"net"
	"time"
)

func main() {
	config := server.ServerConfig{
		Host:             "0.0.0.0:6666",
		MaxGame:          10,
		MaxClientPerGame: 10,
	}

	s := server.New(config)
	go s.Listen()

	time.Sleep(2 * time.Second)
	conn, err := net.Dial("tcp", "0.0.0.0:6666")
	if err != nil {
		log.Println(err.Error())
	} else {
		buf := make([]byte, 1024)
		// conn.Write([]byte(`{"request": "create_game"}`))
		conn.Write([]byte(`Craps out`))
		conn.Read(buf)
		log.Printf("[Client]: %s", buf)
		time.Sleep(10 * time.Second)
	}
}
