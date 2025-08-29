package item

import (
	"main/shaders"
	"main/utils"

	"github.com/bob4321at/textures"
	"github.com/hajimehoshi/ebiten/v2"
)

var PlayerHealth *int
var PlayerMaxHealth int
var PlayerPos *utils.Vec2

var Level_Hitbox []utils.Vec2
var Breakable_Tile_Hitboxes []utils.Vec2

type Item interface {
	Update()
	Draw(screen *ebiten.Image)
	PickUp()
	GetPos() utils.Vec2
	GetSize() utils.Vec2
	PickedUp() bool
	SetPickedUp()
}

var Interacting_Icon = textures.NewTexture("./art/ui/interact_icon.png", shaders.Fill_Shader)

func CheckLevelCollision(pos utils.Vec2, size utils.Vec2) (bool, utils.Vec2) {
	hit := false
	tile_pos := utils.Vec2{}

	for _, tile := range Level_Hitbox {
		check := utils.Collide(pos, size, utils.Vec2{X: float64(tile.X), Y: float64(tile.Y)}, utils.Vec2{X: 32, Y: 32})

		if check {
			hit = true
			tile_pos = utils.Vec2{X: float64(tile.X), Y: float64(tile.Y)}
		}
	}

	for _, breakable_tile := range Breakable_Tile_Hitboxes {
		check := utils.Collide(pos, size, utils.Vec2{X: float64(breakable_tile.X), Y: float64(breakable_tile.Y)}, utils.Vec2{X: 32, Y: 32})

		if check {
			hit = true
			tile_pos = utils.Vec2{X: float64(breakable_tile.X), Y: float64(breakable_tile.Y)}
		}
	}

	return hit, tile_pos
}

var ItemCreateFuncs = [][]func(Pos utils.Vec2) Item{
	nil,
	{
		NewMedKit,
	},
}
