package server

import (
	"net"
)

func NewPlayer(botName string, conn net.Conn) (*Player, error) {
	b := &Bot{
		Name:         botName,
		HitPoint:     BOT_DEFAULT_HP,
		MissileCount: BOT_DEFAULT_MISSILE_COUNT,
		UnderAttack:  false,
	}

	p := &Player{
		Bot:  b,
		Conn: conn,
	}

	return p, nil
}

type Player struct {
	Bot  *Bot
	Conn net.Conn
}

type PlayerList []*Player
