package level

import (
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
	Pos           utils.Vec2
	Top_Bottom    bool
	ReceiveSignal Signal
	Image         *textures.Texture
}

func (water_tile WaterTile) Update(level *Level) {
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
