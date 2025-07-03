package level

import (
	"encoding/json"
	"image"
	"main/enemies"
	"main/utils"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type TileJson struct {
	ID  int
	Pos utils.Vec2
}

type EnemySpawnerJson struct {
	Pos     utils.Vec2
	Enemies []int
	Signal  int
}

type BreakableTileJson struct {
	Pos    utils.Vec2
	Signal int
}

type LevelJson struct {
	Player_Spawn  utils.Vec2
	Tiles         []TileJson
	Enemies       []EnemySpawnerJson
	BreakableTile []BreakableTileJson
}

func LoadLevel(level_name string) Level {
	level := Level{}

	temp_img, _, err := ebitenutil.NewImageFromFile("./levels/" + level_name + "/tileset.png")
	if err != nil {
		panic(err)
	}

	level.TileSet_Img = temp_img

	level.TileSet = append(level.TileSet, ebiten.NewImageFromImage(level.TileSet_Img.SubImage(image.Rect(0, 0, 32, 32))))
	level.TileSet = append(level.TileSet, ebiten.NewImageFromImage(level.TileSet_Img.SubImage(image.Rect(32, 0, 64, 32))))
	level.TileSet = append(level.TileSet, ebiten.NewImageFromImage(level.TileSet_Img.SubImage(image.Rect(64, 0, 96, 32))))

	level.TileSet = append(level.TileSet, ebiten.NewImageFromImage(level.TileSet_Img.SubImage(image.Rect(0, 32, 32, 64))))
	level.TileSet = append(level.TileSet, ebiten.NewImageFromImage(level.TileSet_Img.SubImage(image.Rect(32, 32, 64, 64))))
	level.TileSet = append(level.TileSet, ebiten.NewImageFromImage(level.TileSet_Img.SubImage(image.Rect(64, 32, 96, 64))))

	level.TileSet = append(level.TileSet, ebiten.NewImageFromImage(level.TileSet_Img.SubImage(image.Rect(0, 64, 32, 96))))
	level.TileSet = append(level.TileSet, ebiten.NewImageFromImage(level.TileSet_Img.SubImage(image.Rect(32, 64, 64, 96))))
	level.TileSet = append(level.TileSet, ebiten.NewImageFromImage(level.TileSet_Img.SubImage(image.Rect(64, 64, 96, 96))))

	level.TileSet = append(level.TileSet, ebiten.NewImageFromImage(level.TileSet_Img.SubImage(image.Rect(96, 0, 128, 32))))
	level.TileSet = append(level.TileSet, ebiten.NewImageFromImage(level.TileSet_Img.SubImage(image.Rect(96, 32, 128, 64))))
	level.TileSet = append(level.TileSet, ebiten.NewImageFromImage(level.TileSet_Img.SubImage(image.Rect(96, 64, 128, 96))))

	level.TileSet = append(level.TileSet, ebiten.NewImageFromImage(level.TileSet_Img.SubImage(image.Rect(0, 96, 32, 128))))
	level.TileSet = append(level.TileSet, ebiten.NewImageFromImage(level.TileSet_Img.SubImage(image.Rect(32, 96, 64, 128))))
	level.TileSet = append(level.TileSet, ebiten.NewImageFromImage(level.TileSet_Img.SubImage(image.Rect(64, 96, 96, 128))))

	level.TileSet = append(level.TileSet, ebiten.NewImageFromImage(level.TileSet_Img.SubImage(image.Rect(96, 96, 128, 128))))

	level_file, err := os.ReadFile("./levels/" + level_name + "/level")
	if err != nil {
		panic(err)
	}

	temp_level_json := LevelJson{}

	if err := json.Unmarshal(level_file, &temp_level_json); err != nil {
		panic(err)
	}

	temp_level_tiles := []Tile{}

	for _, tile := range temp_level_json.Tiles {
		temp_level_tiles = append(temp_level_tiles, Tile{level.TileSet[tile.ID-1], tile.Pos})
		enemies.Level_Hitbox = append(enemies.Level_Hitbox, tile.Pos)
	}

	level.Tiles = temp_level_tiles
	level.Player_Spawn = utils.Vec2{X: temp_level_json.Player_Spawn.X - 320, Y: temp_level_json.Player_Spawn.Y - 240}
	level.Player_Loaded = false

	level.Enemy_Spawners = nil
	for _, spawner := range temp_level_json.Enemies {
		level.Enemy_Spawners = append(level.Enemy_Spawners, EnemySpawner{spawner.Pos, spawner.Enemies, nil, 10, 0, Signal{spawner.Signal, false}})
	}

	level.BreakableTile = nil
	breakable_tile_img, _, err := ebitenutil.NewImageFromFile("./levels/" + level_name + "/breakable_tile.png")
	if err != nil {
		panic(err)
	}
	for _, breakable_tile := range temp_level_json.BreakableTile {
		level.BreakableTile = append(level.BreakableTile, BreakableTile{breakable_tile.Pos, breakable_tile.Signal, false, breakable_tile_img})
	}

	return level
}
