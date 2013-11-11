package server

import (
	"github.com/bitly/go-simplejson"
	"log"
	"net"
)

const (
	RCV_BUFFER_SIZE = 1024
	SND_BUFFER_SIZE = 4096
)

type ServerConfig struct {
	Host             string
	MaxGame          int
	MaxClientPerGame int
}

func New(config ServerConfig) *Server {
	return &Server{
		Config: config,
		Games:  make(GameList, config.MaxGame),
	}
}

type Server struct {
	Config ServerConfig
	Games  GameList
}

func (s *Server) Listen() {

	listener, err := net.Listen("tcp", s.Config.Host)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println("Listening for connection on ", listener.Addr().String())

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err.Error())
		}

		go s.acceptConn(conn)
	}
}

func (s *Server) acceptConn(conn net.Conn) {
	//Look for the initial request. It should be a create or a join game.

	buf := make([]byte, RCV_BUFFER_SIZE)

	_, err := conn.Read(buf)
	if err != nil {
		log.Println(err.Error())
		//We won't accept this incomming connection
		return
	}

	_, err = simplejson.NewJson(buf)
	if err != nil {
		log.Println(err.Error())
		conn.Write(getErrorResponse(err.Error()))
		return
	}
}

// func (s *Server) readFromConn(conn net.Conn) {
// 	continueReading := true

// 	for continueReading {
// 		buf := make([]byte, RCV_BUFFER_SIZE)
// 		n, err := conn.Read(buf)
// 		if err != nil {
// 			continueReading = false
// 			log.Println(err.Error())
// 		}

// 		continueReading = s.DispatchMessage(buf)
// 	}
// }

// func (s *Server) NewGame(gameName string) (GameId string, ClientId string, err error) {
//Create new game struct
//Initialize board
//Initialize first client
// }

// func (s *Server) JoinGame(GameId string) (ClientId string, err error) {

// }
