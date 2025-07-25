package level

import (
	"main/utils"

	"github.com/hajimehoshi/ebiten/v2"
)

type TriggerTile struct {
	Pos        utils.Vec2
	SendSignal Signal
	Visible    bool
	Img        *ebiten.Image
}

func (tile *TriggerTile) Update(level *Level) {
	for _, spawner := range level.Enemy_Spawners {
		if spawner.ReceiveSignal.Id == tile.SendSignal.Id {
			if tile.SendSignal.Active {
				spawner.ReceiveSignal.Active = true
			}
		}
	}
}
