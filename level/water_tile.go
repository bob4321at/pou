package level

import (
	"fmt"
	"main/shaders"
	"main/utils"

	"github.com/bob4321at/textures"
)

var Water_Tile_Images map[bool]*textures.Texture

func init() {
	img_top := textures.NewTexture("./art/special_tiles/water_tile.png", shaders.Chunk_Shader)
	img_bottom := textures.NewTexture("./art/special_tiles/water_tile_bottom.png", shaders.Chunk_Shader)

	Water_Tile_Images = map[bool]*textures.Texture{
		false: img_top,
		true:  img_bottom,
	}

}

type WaterTile struct {
	Pos                               utils.Vec2
	Top_Bottom                        bool
	ReceiveSignal                     Signal
	Dissapear_or_Appear               bool
	Dissapear_Or_Appear_Top_Or_Bottom bool
	There_or_Not                      bool
	Image                             textures.Texture
}

func (water_tile *WaterTile) Update(level *Level) {
	if water_tile.ReceiveSignal.Id != 0 {
		if water_tile.Dissapear_or_Appear {
			if water_tile.ReceiveSignal.Active {
				water_tile.There_or_Not = false
			} else {
				for _, signal := range level.Send_Signals {
					if signal.Id == water_tile.ReceiveSignal.Id {
						if signal.Active == true {
							water_tile.ReceiveSignal.Active = true
						}
					}
				}
			}
		} else {
			if water_tile.ReceiveSignal.Active {
				water_tile.There_or_Not = true
			}
			if !water_tile.ReceiveSignal.Active {
				for _, signal := range level.Send_Signals {
					if signal.Id == water_tile.ReceiveSignal.Id {
						if signal.Active == true {
							water_tile.ReceiveSignal.Active = true
						}
					}
				}
			}
		}
	}
	if water_tile.Dissapear_Or_Appear_Top_Or_Bottom {
		for i := range level.WaterTiles {
			other_water_tile := &level.WaterTiles[i]
			if other_water_tile != water_tile {
				if other_water_tile.Pos.Y+32 == water_tile.Pos.Y {
					if other_water_tile.Pos.X == water_tile.Pos.X {
						if !other_water_tile.There_or_Not {
							water_tile.Image.Img = Water_Tile_Images[!water_tile.There_or_Not].Img
							fmt.Println(water_tile)
						}
					}
				}
			}
		}
	}
}
