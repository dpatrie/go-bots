package server

const (
	BOT_DEFAULT_HP            = 100
	BOT_DEFAULT_MISSILE_COUNT = 2
)

type Bot struct {
	X             int
	Y             int
	Name          string
	HitPoint      int
	MissileCount  int
	UnderAttack   bool
	ActivePowerUp []PowerUp
}

type BotAction string

const (
	ACTION_LOOK_AROUND BotAction = "look"
	ACTION_DEFEND      BotAction = "defend"
	ACTION_MOVE        BotAction = "move"
	ACTION_ATTACK      BotAction = "attack"
	ACTION_MISSILE     BotAction = "missile"
)
