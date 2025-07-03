package level

import (
	"main/utils"

	"github.com/hajimehoshi/ebiten/v2"
)

type BreakableTile struct {
	Pos    utils.Vec2
	Signal int
	active bool
	Img    *ebiten.Image
}

func (tile *BreakableTile) Update(level *Level) {
	for _, spawner := range level.Enemy_Spawners {
		if spawner.Signal.Id == tile.Signal {
			if spawner.Signal.Active {
				tile.Img = ebiten.NewImage(32, 32)
				tile.active = true
			}
		}
	}
}
