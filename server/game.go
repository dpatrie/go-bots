package server

type GameList []Game

type Game struct {
	Name    string
	Clients ClientList
	Board   Board
}

func NewGame(gameName string) *Game {
	return nil
}
