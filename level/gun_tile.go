package level

import (
	"main/gun"
	"main/utils"
)

type GunTile struct {
	Pos           utils.Vec2
	GunId         int
	SendSignal    Signal
	ReceiveSignal Signal
	Spawned       bool
}

func (tile *GunTile) Update(level *Level) {
	if tile.ReceiveSignal.Id == 0 {
		if tile.Spawned == false {
			level.Dropped_Guns = append(level.Dropped_Guns, gun.CreateDroppedGun(gun.Guns[tile.GunId], tile.Pos))
			tile.Spawned = true
		}
	}
}
