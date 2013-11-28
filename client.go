package main

import (
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "0.0.0.0:6666")
	if err != nil {
		log.Println(err.Error())
	} else {
		buf := make([]byte, 1024)
		// conn.Write([]byte(`{"request": "createGame", "param":{"name":"Awesome game", "botName": "Awesome Bot", "width":10, "height":10}}`))
		// conn.Write([]byte(`{"request": "createGame", "param":{"name":"Awesome game2", "botName": "Awesome Bot2"}}`))
		conn.Write([]byte(`{"request": "joinGame", "param":{"gameId":1, "botName": "Awesome Bot3"}}`))
		// conn.Write([]byte(`{"request": "listGame"}`))
		// conn.Write([]byte(`Craps out`))
		conn.Read(buf)
		log.Printf("[Client]: %s", buf)
	}

}
