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

func NewRandomSquare(x int, y int) *Square {
	s := &Square{X: x, Y: y}

	s.Type = getRandomSquareType()
	if s.Type == SQUARE_POWERUP {
		s.PowerUp = getRandomPowerUp()
	}

	return s
}

//TODO: Make it easy to change probabilities
func getRandomSquareType() SquareType {
	rand.Seed(time.Now().Unix())
	dice := rand.Intn(100)

	if dice < 75 {
		return SQUARE_EMPTY
	} else if dice < 95 {
		return SQUARE_WALL
	} else {
		return SQUARE_POWERUP
	}
}

func getRandomPowerUp() PowerUp {
	rand.Seed(time.Now().Unix())

	powerups := []PowerUp{
		POWER_UP_NITRO_BOOST,
		POWER_UP_MISSILE_CRATE,
		POWER_UP_SUPER_VISION,
		POWER_UP_REPAIR_KIT,
		POWER_UP_SUPER_SHIELD,
	}
	return powerups[rand.Intn(len(powerups))]
}

type Square struct {
	X       int
	Y       int
	Type    SquareType
	PowerUp PowerUp
	Bot     *Bot
}

type BoardRow []*Square
type Board []BoardRow

func (b Board) GetRandomSpawnXY() (x int, y int) {
	// for {

	// }
	x = 0
	y = 0
	return
}

func NewBoard(width int, height int) Board {
	board := make(Board, height)
	for x, _ := range board {
		board[x] = make(BoardRow, width)
		for y, _ := range board[x] {
			board[x][y] = NewRandomSquare(x, y)
		}
	}

	return board
}
