package server

import (
	"net"
)

func NewPlayer(botName string, botSpawnX int, botSpawnY int, conn net.Conn) (Player, error) {
	b := Bot{
		X:            botSpawnX,
		Y:            botSpawnY,
		Name:         botName,
		HitPoint:     BOT_DEFAULT_HP,
		MissileCount: BOT_DEFAULT_MISSILE_COUNT,
		UnderAttack:  false,
	}

	p := Player{
		Bot:  b,
		Conn: conn,
	}

	return p, nil
}

type Player struct {
	Bot  Bot
	Conn net.Conn
}

type PlayerList []Player

func (pl PlayerList) BroadCast(msg string) {

}
