package level

import (
	"main/utils"
)

type FloodTile struct {
	Pos           utils.Vec2
	Speed         int
	SendSignal    Signal
	ReceiveSignal Signal
}

func (flood_tile *FloodTile) Update(level *Level) {
	if !flood_tile.ReceiveSignal.Active {
		for _, signal := range level.Send_Signals {
			if signal.Id == flood_tile.ReceiveSignal.Id {
				if signal.Active {
					if level.Signal_Controlling_Water != nil {
						for i := range level.FloodTiles {
							opp_flood_tile := *&level.FloodTiles[i]
							if opp_flood_tile.ReceiveSignal.Id == level.Signal_Controlling_Water.Id {
								for j := range level.Send_Signals {
									send_signal := level.Send_Signals[j]
									if send_signal.Id == opp_flood_tile.SendSignal.Id {
										send_signal.Active = false
									}
								}
								opp_flood_tile.ReceiveSignal.Active = false
							}
						}
					}
					flood_tile.ReceiveSignal.Active = true
					level.Signal_Controlling_Water = &flood_tile.ReceiveSignal
					level.WaterLevelTarget = flood_tile.Pos.Y
				}
			}
		}
	}
}
