package level

import (
	"image/color"
	"main/camera"
	"main/enemies"
	"main/gun"
	"main/item"
	"main/utils"

	"github.com/bob4321at/textures"
	"github.com/hajimehoshi/ebiten/v2"
)

type Tile struct {
	Img textures.Texture
	Pos utils.Vec2
}

type Signal struct {
	Id     int
	Active bool
}

type Level struct {
	Player_Loaded bool
	Player_Spawn  utils.Vec2

	End_Pos  utils.Vec2
	Sock_Img *ebiten.Image

	Tiles                    []Tile
	TileSet_Img              *ebiten.Image
	TileSet                  []*ebiten.Image
	Enemy_Spawners           []EnemySpawner
	BreakableTile            []BreakableTile
	TriggerTile              []TriggerTile
	GunTiles                 []GunTile
	ItemTiles                []ItemTile
	SpikeTiles               []SpikeTile
	SpringTiles              []SpringTile
	WaterTiles               []WaterTile
	FloodTiles               []FloodTile
	Signal_Controlling_Water *Signal

	MovingPlatforms         []MovingPlatform
	MovingPlatformPaths     map[int]map[int]utils.Vec2
	MovingPlatformLoopOrNot map[int]bool

	Send_Signals    []*Signal
	Receive_Signals []*Signal

	Enemies []enemies.Enemy

	Dropped_Guns []gun.DroppedGunStruct
	Items        []item.Item

	TileBorderColor color.RGBA
	TileColor       color.RGBA
	BackgroundColor color.RGBA

	WaterLevel       float64
	WaterLevelTarget float64
}

func (level *Level) Update() {
	if ebiten.IsKeyPressed(ebiten.KeyY) {
		level.WaterLevel -= 0.1
		if ebiten.IsKeyPressed(ebiten.KeyShift) {
			level.WaterLevel += 0.2
		}
	}
	enemies.Breakable_Tile_Hitboxes = nil
	item.Breakable_Tile_Hitboxes = nil

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
			item.Breakable_Tile_Hitboxes = append(item.Breakable_Tile_Hitboxes, breakable_tile.Pos)
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

	for i := range enemies.Projectiles {
		projectile := enemies.Projectiles[i]
		projectile.Update()
	}

	enemies.ManageProjectiles()

	for i := range level.Dropped_Guns {
		dropped_gun := &level.Dropped_Guns[i]
		dropped_gun.Update()
	}

	for i := range level.Items {
		item := level.Items[i]
		item.Update()
	}

	for i := range level.GunTiles {
		gun_spawner := &level.GunTiles[i]
		gun_spawner.Update(level)
	}

	for i := range level.ItemTiles {
		item_tile := &level.ItemTiles[i]
		item_tile.Update(level)
	}

	for i := range level.MovingPlatforms {
		platform := &level.MovingPlatforms[i]
		platform.Update(level)
	}

	for i := range level.WaterTiles {
		water_tile := &level.WaterTiles[i]
		water_tile.Update(level)
	}

	for i := range level.FloodTiles {
		flood_tile := &level.FloodTiles[i]
		flood_tile.Update(level)
	}

	enemies.AllEnemies = level.Enemies
	gun.Enemies_In_Level = level.Enemies

	if level.Signal_Controlling_Water != nil {
		if level.WaterLevel > level.WaterLevelTarget {
			level.WaterLevel -= 1
		} else if level.WaterLevel < level.WaterLevelTarget {
			level.WaterLevel += 1
		}

		if level.WaterLevel > level.WaterLevelTarget-5 && level.WaterLevel < level.WaterLevelTarget+5 {
			level.Signal_Controlling_Water.Active = false
			level.Signal_Controlling_Water = nil
		}
	}

	for i := len(level.Items) - 1; i >= 0; i-- {
		if level.Items[i].PickedUp() {
			utils.RemoveArrayElement(i, &level.Items)
		}
	}
}

func (level *Level) Draw(screen *ebiten.Image) {
	op := ebiten.DrawImageOptions{}
	op.GeoM.Translate(level.End_Pos.X+camera.Camera.Pos.X, level.End_Pos.Y+camera.Camera.Pos.Y)
	screen.DrawImage(level.Sock_Img, &op)

	for _, tile := range level.Tiles {
		op := ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(tile.Pos.X)+(camera.Camera.Pos.X), float64(tile.Pos.Y)+(camera.Camera.Pos.Y))

		tile.Img.SetUniforms(map[string]any{
			"R": float64(level.TileBorderColor.R) / 255,
			"G": float64(level.TileBorderColor.G) / 255,
			"B": float64(level.TileBorderColor.B) / 255,

			"RR": float64(level.TileColor.R) / 255,
			"GG": float64(level.TileColor.G) / 255,
			"BB": float64(level.TileColor.B) / 255,
		})

		tile.Img.Draw(screen, &op)
	}

	for _, breakable_tile := range level.BreakableTile {
		op := ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(breakable_tile.Pos.X)+(camera.Camera.Pos.X), float64(breakable_tile.Pos.Y)+(camera.Camera.Pos.Y))
		if !breakable_tile.ReceiveSignal.Active {
			breakable_tile.Img.SetUniforms(map[string]any{
				"R": float64(level.TileBorderColor.R) / 255,
				"G": float64(level.TileBorderColor.G) / 255,
				"B": float64(level.TileBorderColor.B) / 255,

				"RR": float64(level.TileColor.R) / 255,
				"GG": float64(level.TileColor.G) / 255,
				"BB": float64(level.TileColor.B) / 255,
			})

			breakable_tile.Img.Draw(screen, &op)
		}
	}

	for _, spike_tile := range level.SpikeTiles {
		op := ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(spike_tile.Pos.X)+(camera.Camera.Pos.X), float64(spike_tile.Pos.Y)+(camera.Camera.Pos.Y))
		screen.DrawImage(spike_tile.Img, &op)
	}

	for _, spring_tiles := range level.SpringTiles {
		op := ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(spring_tiles.Pos.X)+(camera.Camera.Pos.X), float64(spring_tiles.Pos.Y)+(camera.Camera.Pos.Y))
		screen.DrawImage(spring_tiles.Img, &op)
	}

	for _, trigger_tile := range level.TriggerTile {
		if trigger_tile.Visible {
			op := ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(trigger_tile.Pos.X)+(camera.Camera.Pos.X), float64(trigger_tile.Pos.Y)+(camera.Camera.Pos.Y))
			screen.DrawImage(trigger_tile.Img, &op)
		}
	}

	for _, moving_platform := range level.MovingPlatforms {
		moving_platform.Draw(screen)
	}

	for i := range level.Enemy_Spawners {
		spawner := &level.Enemy_Spawners[i]
		for _, enemy := range spawner.Responsible_For {
			enemy.Draw(screen)
		}
	}

	for i := range enemies.Projectiles {
		projectile := enemies.Projectiles[i]
		projectile.Draw(screen)
	}

	for _, dropped_gun := range level.Dropped_Guns {
		dropped_gun.Draw(screen)
	}

	for _, item := range level.Items {
		item.Draw(screen)
	}
}

func (level *Level) Reset() {
	level.Enemies = nil
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
