package level

import (
	"main/utils"

	"github.com/bob4321at/textures"
)

type BreakableTile struct {
	Pos           utils.Vec2
	ReceiveSignal Signal
	Img           *textures.Texture
}

func (tile *BreakableTile) Update(level *Level) {
	for _, signal := range level.Send_Signals {
		if signal.Id == tile.ReceiveSignal.Id {
			if signal.Active {
				tile.ReceiveSignal.Active = true
			}
		}
	}
}
