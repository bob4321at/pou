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
	Selected_Gun  *gun.DroppedGunStruct
}

func (tile *GunTile) Update(level *Level) {
	if tile.ReceiveSignal.Id == 0 {
		if tile.Spawned == false {
			level.Dropped_Guns = append(level.Dropped_Guns, gun.CreateDroppedGun(gun.Guns[tile.GunId], tile.Pos))
			tile.Selected_Gun = &level.Dropped_Guns[len(level.Dropped_Guns)-1]
			tile.Spawned = true
		} else {
			not_there := true

			for i := range level.Dropped_Guns {
				if tile.Selected_Gun == &level.Dropped_Guns[i] {
					not_there = false
				}
			}

			if not_there {
				tile.SendSignal.Active = true
			}
		}
	} else {
		for _, signal := range level.Send_Signals {
			if signal.Id == tile.ReceiveSignal.Id {
				if signal.Active {
					tile.ReceiveSignal.Id = 0
				}
			}
		}
	}
}
