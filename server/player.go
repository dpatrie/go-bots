package server

import (
	"net"
)

type PlayerList []*Player

type Player struct {
	Bot  *Bot
	Conn net.Conn
}

func NewPlayer(botName string, conn net.Conn) *Player {
	p := &Player{
		Bot: &Bot{
			Name:         botName,
			HitPoint:     BOT_DEFAULT_HP,
			MissileCount: BOT_DEFAULT_MISSILE_COUNT,
			UnderAttack:  false,
		},
		Conn: conn,
	}

	return p
}
