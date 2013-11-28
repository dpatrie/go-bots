package server

import (
	"errors"
	"math/rand"
	"strings"
	"time"
)

const (
	BOARD_DEFAULT_WIDTH  = 20
	BOARD_DEFAULT_HEIGHT = 20
)

type Board []BoardRow
type BoardRow []*Square

func (b Board) ConsoleRender() string {
	var buf []string

	for _, row := range b {
		for _, s := range row {
			switch s.Type {
			case SQUARE_EMPTY:
				buf = append(buf, "\u2B1C ")
			case SQUARE_POWERUP:
				buf = append(buf, "P ")
			case SQUARE_WALL:
				buf = append(buf, "\u2B1B ")
			case SQUARE_BOT:
				buf = append(buf, "B ")
			}
		}
		buf = append(buf, "\n")
	}

	return strings.Join(buf, "")
}

func (b Board) getRandomEmptySquare() *Square {
	rand.Seed(time.Now().UnixNano())
	maxX := len(b) - 1
	maxY := len(b[0]) - 1

	i := 0
	for i < (maxX * maxY) {
		x := rand.Intn(maxX)
		y := rand.Intn(maxY)

		if s := b[x][y]; s.Type == SQUARE_EMPTY {
			return s
		}
		i++
	}

	return nil
}

func (b Board) Spawn(bot *Bot) error {
	s := b.getRandomEmptySquare()
	if s == nil {
		return errors.New("Unable to spawn")
	}
	s.Type = SQUARE_BOT
	s.Bot = bot
	return nil
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
