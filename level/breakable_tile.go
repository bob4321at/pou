package level

import (
	"main/utils"

	"github.com/hajimehoshi/ebiten/v2"
)

type BreakableTile struct {
	Pos           utils.Vec2
	ReceiveSignal Signal
	Img           *ebiten.Image
}

func (tile *BreakableTile) Update(level *Level) {
	for _, signal := range level.Send_Signals {
		if signal.Id == tile.ReceiveSignal.Id {
			if signal.Active {
				tile.ReceiveSignal.Active = true
			}
		}
	}

	// for _, spawner := range level.Enemy_Spawners {
	// 	if spawner.SendSignal.Id == tile.ReceiveSignal.Id {
	// 		if spawner.SendSignal.Active {
	// 			tile.ReceiveSignal.Active = true
	// 		}
	// 	}
	// }
	// for _, trigger := range level.TriggerTile {
	// 	if trigger.SendSignal.Id == tile.ReceiveSignal.Id {
	// 		if trigger.SendSignal.Active {
	// 			tile.ReceiveSignal.Active = true
	// 		}
	// 	}
	// }
}
