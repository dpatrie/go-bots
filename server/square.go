package server

import (
	"math/rand"
	"time"
)

type SquareType string

const (
	SQUARE_EMPTY   SquareType = "empty"
	SQUARE_WALL    SquareType = "wall"
	SQUARE_BOT     SquareType = "bot"
	SQUARE_POWERUP SquareType = "powerup"
)

type Square struct {
	X       int
	Y       int
	Type    SquareType
	PowerUp PowerUp
	Bot     *Bot
}

func NewRandomSquare(x int, y int) *Square {
	s := &Square{X: x, Y: y}
	s.Type = getRandomSquareType()
	if s.Type == SQUARE_POWERUP {
		s.PowerUp = getRandomPowerUp()
	}

	return s
}

//TODO: Make it easy to change probabilities
//TODO: Make it more probable to get wall when there's other wall in
//surrounding
func getRandomSquareType() SquareType {
	rand.Seed(time.Now().UnixNano())
	dice := rand.Intn(100)
	if dice < 75 {
		return SQUARE_EMPTY
	} else if dice < 98 {
		return SQUARE_WALL
	} else {
		return SQUARE_POWERUP
	}
}
