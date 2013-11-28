package server

import (
	"math/rand"
	"time"
)

type PowerUp string

const (
	POWER_UP_NONE          PowerUp = ""
	POWER_UP_NITRO_BOOST   PowerUp = "nitroBoost"
	POWER_UP_MISSILE_CRATE PowerUp = "missileCrate"
	POWER_UP_SUPER_VISION  PowerUp = "superVision"
	POWER_UP_REPAIR_KIT    PowerUp = "repairKit"
	POWER_UP_SUPER_SHIELD  PowerUp = "superShield"
)

func getRandomPowerUp() PowerUp {
	rand.Seed(time.Now().UnixNano())

	powerups := []PowerUp{
		POWER_UP_NITRO_BOOST,
		POWER_UP_MISSILE_CRATE,
		POWER_UP_SUPER_VISION,
		POWER_UP_REPAIR_KIT,
		POWER_UP_SUPER_SHIELD,
	}
	return powerups[rand.Intn(len(powerups))]
}
