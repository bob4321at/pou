package level

import (
	"main/camera"
	"main/enemies"
	"main/gun"
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
	Tiles           []Tile
	TileSet_Img     *ebiten.Image
	TileSet         []*ebiten.Image
	Player_Loaded   bool
	Player_Spawn    utils.Vec2
	End_Pos         utils.Vec2
	Sock_Img        *ebiten.Image
	Enemy_Spawners  []EnemySpawner
	BreakableTile   []BreakableTile
	TriggerTile     []TriggerTile
	GunTiles        []GunTile
	SpikeTiles      []SpikeTile
	Send_Signals    []*Signal
	Receive_Signals []*Signal
	Enemies         []enemies.Enemy
	Dropped_Guns    []gun.DroppedGunStruct
}

func (level *Level) Update() {
	enemies.Breakable_Tile_Hitboxes = nil

	for _, enemy := range enemies.Enemies_To_Add {
		level.Enemy_Spawners[0].Responsible_For = append(level.Enemy_Spawners[0].Responsible_For, enemy)
	}

	enemies.Enemies_To_Add = nil

	for i := range level.Enemy_Spawners {
		spawner := &level.Enemy_Spawners[i]
		spawner.Update(level)
	}

	for i := range level.BreakableTile {
		breakable_tile := &level.BreakableTile[i]
		breakable_tile.Update(level)
		if !breakable_tile.ReceiveSignal.Active {
			enemies.Breakable_Tile_Hitboxes = append(enemies.Breakable_Tile_Hitboxes, breakable_tile.Pos)
		}
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

	for i := range level.Dropped_Guns {
		dropped_gun := &level.Dropped_Guns[i]
		dropped_gun.Update()
	}

	for i := range level.GunTiles {
		gun_spawner := &level.GunTiles[i]
		gun_spawner.Update(level)
	}

	enemies.AllEnemies = level.Enemies
	gun.Enemies_In_Level = level.Enemies
}

func (level *Level) Draw(screen *ebiten.Image) {
	op := ebiten.DrawImageOptions{}
	op.GeoM.Translate(level.End_Pos.X+camera.Camera.Pos.X, level.End_Pos.Y+camera.Camera.Pos.Y)
	screen.DrawImage(level.Sock_Img, &op)

	for _, tile := range level.Tiles {
		op := ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(tile.Pos.X)+(camera.Camera.Pos.X), float64(tile.Pos.Y)+(camera.Camera.Pos.Y))

		screen.DrawImage(tile.Img, &op)
	}

	for _, breakable_tile := range level.BreakableTile {
		op := ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(breakable_tile.Pos.X)+(camera.Camera.Pos.X), float64(breakable_tile.Pos.Y)+(camera.Camera.Pos.Y))
		if !breakable_tile.ReceiveSignal.Active {
			screen.DrawImage(breakable_tile.Img, &op)
		}
	}

	for _, spike_tile := range level.SpikeTiles {
		op := ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(spike_tile.Pos.X)+(camera.Camera.Pos.X), float64(spike_tile.Pos.Y)+(camera.Camera.Pos.Y))
		screen.DrawImage(spike_tile.Img, &op)
	}

	for i := range level.Enemy_Spawners {
		spawner := &level.Enemy_Spawners[i]
		for _, enemy := range spawner.Responsible_For {
			enemy.Draw(screen)
		}
	}

	for _, dropped_gun := range level.Dropped_Guns {
		dropped_gun.Draw(screen)
	}
}

func (level *Level) Reset() {
	level.Enemies = nil

	for i := range level.Enemy_Spawners {
		level.Enemy_Spawners[i].Index = 0
		level.Enemy_Spawners[i].Responsible_For = nil
		level.Enemy_Spawners[i].Timer = 10
		level.Enemy_Spawners[i].SendSignal.Active = false
	}

	for i := range level.BreakableTile {
		level.BreakableTile[i].ReceiveSignal.Active = false
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
		if !breakable_tile.ReceiveSignal.Active {
			check := utils.Collide(pos, size, utils.Vec2{X: float64(breakable_tile.Pos.X), Y: float64(breakable_tile.Pos.Y)}, utils.Vec2{X: 32, Y: 32})

			if check {
				hit = true
				tile_pos = utils.Vec2{X: float64(breakable_tile.Pos.X), Y: float64(breakable_tile.Pos.Y)}
			}
		}
	}

	for _, spike_tile := range level.SpikeTiles {
		check := utils.Collide(pos, size, utils.Vec2{X: float64(spike_tile.Pos.X + 8), Y: float64(spike_tile.Pos.Y + 8)}, utils.Vec2{X: 16, Y: 16})

		if check {
			hit = true
			tile_pos = utils.Vec2{X: float64(spike_tile.Pos.X), Y: float64(spike_tile.Pos.Y)}
		}
	}

	return hit, tile_pos
}
