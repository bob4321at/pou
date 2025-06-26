package level

import (
	"main/camera"
	"main/enemies"
	"main/utils"

	"github.com/hajimehoshi/ebiten/v2"
)

type Tile struct {
	Img *ebiten.Image
	Pos utils.Vec2
}

type Level struct {
	Tiles         []Tile
	TileSet_Img   *ebiten.Image
	TileSet       []*ebiten.Image
	Enemies       []enemies.Enemy
	Player_Loaded bool
	Player_Spawn  utils.Vec2
}

func (level *Level) Update() {
	for i, enemy := range level.Enemies {
		enemy.Update()
		if enemy.GetHealth() <= 0 {
			utils.RemoveArrayElement(i, &level.Enemies)
		}
	}
}

func (level *Level) Draw(screen *ebiten.Image) {
	for _, tile := range level.Tiles {
		op := ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(tile.Pos.X)+(camera.Camera.Pos.X), float64(tile.Pos.Y)+(camera.Camera.Pos.Y))

		screen.DrawImage(tile.Img, &op)
	}

	for _, enemy := range level.Enemies {
		enemy.Draw(screen)
	}
}

func (level *Level) AddEnemy(enemy enemies.Enemy) {
	level.Enemies = append(level.Enemies, enemy)
}

func (level *Level) CheckCollision(pos utils.Vec2, size utils.Vec2) (bool, utils.Vec2) {
	hit := false
	tile_pos := utils.Vec2{}

	for _, tile := range level.Tiles {
		check := utils.Collide(pos, size, utils.Vec2{X: float64(tile.Pos.X), Y: float64(tile.Pos.Y)}, utils.Vec2{X: 32, Y: 32})

		if check {
			hit = true
			tile_pos = utils.Vec2{X: float64(tile.Pos.X), Y: float64(tile.Pos.Y)}
		}
	}

	return hit, tile_pos
}

var Temp_Level = LoadLevel("test")
