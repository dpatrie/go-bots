package server

import (
	"log"
	"sync"
	"time"
)

const (
	GAME_DEFAULT_START_TIME = 60 * time.Second
	GAME_MAX_ROUND          = 1000
)

type GameList []*Game

type Game struct {
	sync.Mutex
	InProgress bool
	Name       string
	Players    PlayerList
	Board      Board
	Round      int
}

func (g *Game) Start() {
	g.Lock()
	g.InProgress = true
	log.Printf("\n%s\n", g.Board.ConsoleRender())
	g.Unlock()

	//Validate that we can indeed start...
	//We have at least 2 bots
	//TODO:Implement ping/pong between server and client to ensure that client is active...

	for g.Round < GAME_MAX_ROUND {
		g.PlayRound()
		g.Round++
	}
}

func (g *Game) PlayRound() {

}

func NewGame(gameName string, boardWidth int, boardHeight int) *Game {
	g := &Game{
		Name:  gameName,
		Board: NewBoard(boardWidth, boardHeight),
	}

	return g
}
