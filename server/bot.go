package server

type Bot struct {
	Name          string
	HitPoint      int
	Position      Coord
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

type PowerUp string

const (
	POWER_UP_NITRO_BOOST   PowerUp = "nitroBoost"
	POWER_UP_MISSILE_CRATE PowerUp = "missileCrate"
	POWER_UP_SUPER_VISION  PowerUp = "superVision"
	POWER_UP_REPAIR_KIT    PowerUp = "repairKit"
	POWER_UP_SUPER_SHIELD  PowerUp = "superShield"
)
