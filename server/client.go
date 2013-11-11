package server

import (
	"net"
)

type Client struct {
	Id   int
	Bot  *Bot
	Conn net.Conn
}

type ClientList map[int]*Client

func (cl ClientList) BroadCast(msg string) {

}
