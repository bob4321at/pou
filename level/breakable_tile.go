package level

import (
	"main/utils"

	"github.com/hajimehoshi/ebiten/v2"
)

type BreakableTile struct {
	Pos    utils.Vec2
	Signal int
	Active bool
	Img    *ebiten.Image
}

func (tile *BreakableTile) Update(level *Level) {
	for _, spawner := range level.Enemy_Spawners {
		if spawner.SendSignal.Id == tile.Signal {
			if spawner.SendSignal.Active {
				tile.Active = true
			}
		}
	}
	for _, trigger := range level.TriggerTile {
		if trigger.Signal == tile.Signal {
			if trigger.Active {
				tile.Active = true
			}
		}
	}
}
