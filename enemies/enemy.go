package enemies

import (
	"main/utils"

	"github.com/hajimehoshi/ebiten/v2"
)

type Enemy interface {
	Draw(screen *ebiten.Image)
	Update()
	Hit(damage int)
	HitPlayer()
	GetPosition() utils.Vec2
	GetSize() utils.Vec2
	GetHealth() int
}

var EnemySpawnFuncs = map[int]func(pos utils.Vec2) Enemy{
	1: NewTarget,
	2: NewBlueGhost,
	3: NewOrangeGuy,
	4: NewDummy,
}

var Level_Hitbox []utils.Vec2
var Breakable_Tile_Hitboxes []utils.Vec2
var AllEnemies []Enemy

var Player_Pos *utils.Vec2
var Player_Health *int

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
