package level

import (
	"fmt"
	"main/camera"
	"main/enemies"
	"main/utils"

	"github.com/hajimehoshi/ebiten/v2"
)

type Tile struct {
	Img *ebiten.Image
	Pos utils.Vec2
}

type Signal struct {
	Id     int
	Active bool
}

type Level struct {
	Tiles          []Tile
	TileSet_Img    *ebiten.Image
	TileSet        []*ebiten.Image
	Player_Loaded  bool
	Player_Spawn   utils.Vec2
	Enemy_Spawners []EnemySpawner
	BreakableTile  []BreakableTile
	Enemies        []enemies.Enemy
}

func (level *Level) Update() {
	for i := range level.Enemy_Spawners {
		spawner := &level.Enemy_Spawners[i]
		spawner.Update()
	}

	for i := range level.BreakableTile {
		breakable_tile := &level.BreakableTile[i]
		breakable_tile.Update(level)
	}

	level.Enemies = nil

	for i := range level.Enemy_Spawners {
		spawner := &level.Enemy_Spawners[i]
		for i, enemy := range spawner.Responsible_For {
			if i < len(spawner.Responsible_For) {
				enemy.Update()

				if enemy.GetHealth() <= 0 {
					utils.RemoveArrayElement(i, &spawner.Responsible_For)
				}

				level.Enemies = append(level.Enemies, enemy)
			}
		}
	}
}

func (level *Level) Draw(screen *ebiten.Image) {
	for _, tile := range level.Tiles {
		op := ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(tile.Pos.X)+(camera.Camera.Pos.X), float64(tile.Pos.Y)+(camera.Camera.Pos.Y))

		screen.DrawImage(tile.Img, &op)
	}

	for i := range level.Enemy_Spawners {
		spawner := &level.Enemy_Spawners[i]
		for _, enemy := range spawner.Responsible_For {
			enemy.Draw(screen)
		}
	}

	for _, breakable_tile := range level.BreakableTile {
		op := ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(breakable_tile.Pos.X)+(camera.Camera.Pos.X), float64(breakable_tile.Pos.Y)+(camera.Camera.Pos.Y))
		screen.DrawImage(breakable_tile.Img, &op)
	}
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

	for _, breakable_tile := range level.BreakableTile {
		fmt.Println(breakable_tile.active)
		if !breakable_tile.active {
			check := utils.Collide(pos, size, utils.Vec2{X: float64(breakable_tile.Pos.X), Y: float64(breakable_tile.Pos.Y)}, utils.Vec2{X: 32, Y: 32})

			if check {
				hit = true
				tile_pos = utils.Vec2{X: float64(breakable_tile.Pos.X), Y: float64(breakable_tile.Pos.Y)}
			}
		}
	}

	return hit, tile_pos
}

var Temp_Level = LoadLevel("test_real_level")
