package server

type SquareType string

const (
	SQUARE_EMPTY   SquareType = "empty"
	SQUARE_WALL    SquareType = "wall"
	SQUARE_BOT     SquareType = "bot"
	SQUARE_POWERUP SquareType = "powerup"
)

type Square struct {
	Position Coord
	Type     SquareType
	PowerUp  PowerUp
	Bot      *Bot
}

type SquareList []*Square
type BoardRow SquareList
type Board []*BoardRow

type Coord struct {
	X int
	Y int
}

func NewBoard() *Board {
	return nil
}
