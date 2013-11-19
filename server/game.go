package server

import (
	"net"
	"sync"
)

type GameList []*Game

type Game struct {
	sync.Mutex
	Name    string
	Players PlayerList
	Board   Board
}

func NewGame(gameName string, botName string, conn net.Conn) (*Game, error) {
	g := &Game{
		Name:  gameName,
		Board: NewBoard(20, 20),
	}

	g.Lock()
	defer g.Unlock()

	x, y := g.Board.GetRandomSpawnXY()
	p, err := NewPlayer(botName, x, y, conn)
	if err != nil {
		return nil, err
	}

	g.Players = append(g.Players, p)

	return g, nil
}
