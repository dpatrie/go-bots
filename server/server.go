package server

import (
	"encoding/json"
	"errors"
	"log"
	"net"
	"sync"
	"time"
)

const (
	RCV_BUFFER_SIZE = 4096
	SND_BUFFER_SIZE = 4096
)

type ServerConfig struct {
	Host             string
	MaxGame          int
	MaxPlayerPerGame int
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
	case RequestListGame:
		reply = getListGameResponse(s.listGame())
		closeConn = true
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
	if len(s.Games) >= s.Config.MaxGame {
		return 0, errors.New("Server is full")
	}

	gn, found := req.Param["name"]
	if !found {
		return 0, errors.New("Missing parameter: name")
	}

	gameName, ok := gn.(string)
	if !ok {
		return 0, errors.New("Invalid parameter: name must be a string")
	}

	bn, found := req.Param["botName"]
	if !found {
		return 0, errors.New("Missing parameter: botName")
	}

	botName, ok := bn.(string)
	if !ok {
		return 0, errors.New("Invalid parameter: botName must be a string")
	}

	width := BOARD_DEFAULT_WIDTH
	height := BOARD_DEFAULT_HEIGHT

	if wi, found := req.Param["width"]; found {
		w, ok := wi.(float64)
		if !ok {
			return 0, errors.New("Invalid parameter: width must be numeric")
		}
		width = int(w)
	}

	if hi, found := req.Param["height"]; found {
		h, ok := hi.(float64)
		if !ok {
			return 0, errors.New("Invalid parameter: height must be numeric")
		}
		height = int(h)
	}

	s.Lock()
	defer s.Unlock()
	g := NewGame(gameName, width, height)
	p := NewPlayer(botName, conn)
	g.Players = append(g.Players, p)
	g.Board.Spawn(p.Bot)

	s.Games = append(s.Games, g)

	go func(g *Game) {
		log.Printf("Game %s will start in %s", g.Name, GAME_DEFAULT_START_TIME)
		time.Sleep(GAME_DEFAULT_START_TIME)
		log.Printf("Game %s is starting", g.Name)
		g.Start()
	}(g)

	return len(s.Games), nil
}

func (s *Server) listGame() []GameDescription {
	var games []GameDescription

	s.Lock()
	defer s.Unlock()

	for id, game := range s.Games {
		if len(game.Players) >= s.Config.MaxPlayerPerGame {
			continue
		}

		d := GameDescription{
			Id:   id + 1,
			Name: game.Name,
		}

		for _, p := range game.Players {
			d.Bots = append(d.Bots, p.Bot.Name)
		}

		games = append(games, d)
	}

	return games
}

func (s *Server) joinGame(req StandardRequest, conn net.Conn) error {

	id, found := req.Param["gameId"]
	if !found {
		return errors.New("Missing parameter: gameId")
	}

	gameId, ok := id.(float64)
	if !ok {
		return errors.New("Invalid parameter: gameId must be numeric")
	}

	bn, found := req.Param["botName"]
	if !found {
		return errors.New("Missing parameter: botName")
	}

	botName, ok := bn.(string)
	if !ok {
		errors.New("Invalid parameter: botName must be a string")
	}

	if gameId <= 0 {
		return errors.New("Invalid parameter: gameId must be > 0")
	}

	game := s.Games[int(gameId)-1]
	if game == nil {
		return errors.New("Invalid parameter: game does not exist")
	}

	game.Lock()
	defer game.Unlock()

	if game.InProgress {
		return errors.New("Game is already started")
	}

	if len(game.Players) >= s.Config.MaxPlayerPerGame {
		return errors.New("Game is full")
	}

	p := NewPlayer(botName, conn)
	game.Players = append(game.Players, p)
	game.Board.Spawn(p.Bot)

	return nil
}
