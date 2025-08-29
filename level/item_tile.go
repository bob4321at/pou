package level

import (
	"main/item"
	"main/utils"
)

type ItemTile struct {
	Pos           utils.Vec2
	CatagoryId    int
	ItemId        int
	SendSignal    Signal
	ReceiveSignal Signal
	Spawned       bool
	Item          item.Item
}

func (tile *ItemTile) Update(level *Level) {
	if tile.ReceiveSignal.Id == 0 {
		if tile.Spawned == false {
			itm := item.ItemCreateFuncs[tile.CatagoryId][tile.ItemId](tile.Pos)
			if itm != nil {
				level.Items = append(level.Items, itm)
				tile.Item = level.Items[len(level.Items)-1]
			}
			tile.Spawned = true
		} else {
			not_there := true

			for i := range level.Items {
				if tile.Item == level.Items[i] {
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
