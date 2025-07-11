package level

import (
	"main/utils"
)

type TriggerTile struct {
	Pos    utils.Vec2
	Signal int
	Active bool
}

func (tile *TriggerTile) Update(level *Level) {
	for _, spawner := range level.Enemy_Spawners {
		if spawner.ReceiveSignal.Id == tile.Signal {
			if tile.Active {
				spawner.ReceiveSignal.Active = true
			}
		}
	}
}
