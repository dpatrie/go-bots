package server

import (
	"encoding/json"
	"errors"
	"log"
	"net"
	"sync"
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

type Server struct {
	sync.Mutex
	Config ServerConfig
	Games  GameList
}

func New(config ServerConfig) *Server {
	return &Server{
		Config: config,
	}
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
	buf := make([]byte, RCV_BUFFER_SIZE)

	numBytes, err := conn.Read(buf)
	if err != nil {
		//We won't accept this incomming connection
		log.Println(err.Error())
		conn.Close()
		return
	}

	r := StandardRequest{}
	if err = json.Unmarshal(buf[:numBytes], &r); err != nil {
		conn.Write(getErrorResponse(err.Error()))
		conn.Close()
		return
	}

	var reply []byte
	closeConn := false

	switch r.Request {
	case RequestCreateGame:
		gameId, err := s.newGame(r, conn)
		if err != nil {
			reply = getErrorResponse(err.Error())
			closeConn = true
		} else {
			reply = getGameCreatedResponse(gameId)
		}
	case RequestJoinGame:
		err = s.joinGame(r, conn)
		if err != nil {
			reply = getErrorResponse(err.Error())
			closeConn = true
		} else {
			reply = getOkResponse()
		}
	default:
		reply = getErrorResponse("invalid request")
		closeConn = true
	}

	conn.Write(reply)
	if closeConn {
		conn.Close()
	}
}

func (s *Server) newGame(req StandardRequest, conn net.Conn) (int, error) {
	gameName, found := req.Param["name"]
	if !found {
		return 0, errors.New("Missing parameters: name")
	}

	botName, found := req.Param["botName"]
	if !found {
		return 0, errors.New("Missing parameters: botName")
	}

	s.Lock()
	defer s.Unlock()
	g, err := NewGame(gameName.(string), botName.(string), conn)
	if err != nil {
		return 0, err
	}

	s.Games = append(s.Games, g)

	log.Printf("\n%s\n", g.Board.ConsoleRender())

	return len(s.Games), nil
}

func (s *Server) joinGame(req StandardRequest, conn net.Conn) error {
	return nil
}
