package level

import (
	"encoding/json"
	"image/color"
	"main/enemies"
	"main/shaders"
	"main/utils"
	"os"

	"github.com/bob4321at/textures"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type TileJson struct {
	ID  int
	Pos utils.Vec2
}

type EnemySpawnerJson struct {
	Pos           utils.Vec2
	Enemies       []int
	SendSignal    int
	ReceiveSignal int
}

type BreakableTileJson struct {
	Pos    utils.Vec2
	Signal int
}

type TriggerTileJson struct {
	Pos       utils.Vec2
	Signal    int
	Visible   bool
	Direction int
}

type GunTileJson struct {
	Pos           utils.Vec2
	GunId         int
	SendSignal    int
	ReceiveSignal int
}

type SpikeTileJson struct {
	Pos       utils.Vec2
	Damage    int
	Direction int
}

type SpringTileJson struct {
	Pos       utils.Vec2
	Power     int
	Direction int
}

type LevelJson struct {
	Player_Spawn  utils.Vec2
	End           utils.Vec2
	Tiles         []TileJson
	Enemies       []EnemySpawnerJson
	BreakableTile []BreakableTileJson
	TriggerTile   []TriggerTileJson
	GunTiles      []GunTileJson
	SpikeTiles    []SpikeTileJson
	SpringTiles   []SpringTileJson

	TileBorderColor color.RGBA
	TileColor       color.RGBA
	BackgroundColor color.RGBA
}

var TileSet []textures.Texture

func init() {
	TileSet = append(TileSet, *textures.NewTexture("./art/tileset/top_left.png", shaders.Chunk_Shader))
	TileSet = append(TileSet, *textures.NewTexture("./art/tileset/top_center.png", shaders.Chunk_Shader))
	TileSet = append(TileSet, *textures.NewTexture("./art/tileset/top_right.png", shaders.Chunk_Shader))

	TileSet = append(TileSet, *textures.NewTexture("./art/tileset/middle_left.png", shaders.Chunk_Shader))
	TileSet = append(TileSet, *textures.NewTexture("./art/tileset/middle_center.png", shaders.Chunk_Shader))
	TileSet = append(TileSet, *textures.NewTexture("./art/tileset/middle_right.png", shaders.Chunk_Shader))

	TileSet = append(TileSet, *textures.NewTexture("./art/tileset/bottom_left.png", shaders.Chunk_Shader))
	TileSet = append(TileSet, *textures.NewTexture("./art/tileset/bottom_center.png", shaders.Chunk_Shader))
	TileSet = append(TileSet, *textures.NewTexture("./art/tileset/bottom_right.png", shaders.Chunk_Shader))

	TileSet = append(TileSet, *textures.NewTexture("./art/tileset/vertical_top.png", shaders.Chunk_Shader))
	TileSet = append(TileSet, *textures.NewTexture("./art/tileset/vertical_middle.png", shaders.Chunk_Shader))
	TileSet = append(TileSet, *textures.NewTexture("./art/tileset/vertical_bottom.png", shaders.Chunk_Shader))

	TileSet = append(TileSet, *textures.NewTexture("./art/tileset/horizontal_left.png", shaders.Chunk_Shader))
	TileSet = append(TileSet, *textures.NewTexture("./art/tileset/horizontal_center.png", shaders.Chunk_Shader))
	TileSet = append(TileSet, *textures.NewTexture("./art/tileset/horizontal_right.png", shaders.Chunk_Shader))

	TileSet = append(TileSet, *textures.NewTexture("./art/tileset/center.png", shaders.Chunk_Shader))
}

func LoadLevel(level_name string) Level {
	level := Level{}

	level_file, err := os.ReadFile("./levels/" + level_name + "/level")
	if err != nil {
		panic(err)
	}

	temp_level_json := LevelJson{}

	if err := json.Unmarshal(level_file, &temp_level_json); err != nil {
		panic(err)
	}

	temp_level_tiles := []Tile{}

	enemies.Level_Hitbox = nil
	for _, tile := range temp_level_json.Tiles {
		temp_level_tiles = append(temp_level_tiles, Tile{TileSet[tile.ID-1], tile.Pos})
		enemies.Level_Hitbox = append(enemies.Level_Hitbox, tile.Pos)
	}

	level.Tiles = temp_level_tiles
	level.Player_Spawn = utils.Vec2{X: temp_level_json.Player_Spawn.X - 320, Y: temp_level_json.Player_Spawn.Y - 240}
	level.Player_Loaded = false

	level.End_Pos = temp_level_json.End
	level.Sock_Img, _, err = ebitenutil.NewImageFromFile("./levels/" + level_name + "/sock.png")
	if err != nil {
		panic(err)
	}

	level.Enemy_Spawners = nil

	level.Enemy_Spawners = append(level.Enemy_Spawners, EnemySpawner{utils.Vec2{X: 0, Y: 0}, []int{}, nil, 10, 0, Signal{0, false}, Signal{0, false}, true})

	for _, spawner := range temp_level_json.Enemies {
		if spawner.ReceiveSignal == 0 {
			level.Enemy_Spawners = append(level.Enemy_Spawners, EnemySpawner{spawner.Pos, spawner.Enemies, nil, 10, 0, Signal{spawner.SendSignal, false}, Signal{spawner.ReceiveSignal, false}, true})
		} else {
			level.Enemy_Spawners = append(level.Enemy_Spawners, EnemySpawner{spawner.Pos, spawner.Enemies, nil, 10, 0, Signal{spawner.SendSignal, false}, Signal{spawner.ReceiveSignal, false}, false})
		}
	}

	breakable_tile_img := textures.NewTexture("./art/special_tiles/breakable_tile.png", shaders.Chunk_Shader)
	for _, breakable_tile := range temp_level_json.BreakableTile {
		level.BreakableTile = append(level.BreakableTile, BreakableTile{breakable_tile.Pos, Signal{breakable_tile.Signal, false}, breakable_tile_img})
	}

	imga, _, err := ebitenutil.NewImageFromFile("./art/special_tiles/buttonup.png")
	if err != nil {
		panic(err)
	}
	imgb, _, err := ebitenutil.NewImageFromFile("./art/special_tiles/buttonright.png")
	if err != nil {
		panic(err)
	}
	imgc, _, err := ebitenutil.NewImageFromFile("./art/special_tiles/buttondown.png")
	if err != nil {
		panic(err)
	}
	imgd, _, err := ebitenutil.NewImageFromFile("./art/special_tiles/buttonleft.png")
	if err != nil {
		panic(err)
	}

	trigger_images := []*ebiten.Image{
		imga,
		imgb,
		imgc,
		imgd,
	}

	level.TriggerTile = nil
	for _, trigger_tile := range temp_level_json.TriggerTile {
		level.TriggerTile = append(level.TriggerTile, TriggerTile{trigger_tile.Pos, Signal{trigger_tile.Signal, false}, trigger_tile.Visible, trigger_images[trigger_tile.Direction]})
	}

	level.GunTiles = nil
	for _, gun_tile := range temp_level_json.GunTiles {
		level.GunTiles = append(level.GunTiles, GunTile{gun_tile.Pos, gun_tile.GunId, Signal{gun_tile.SendSignal, false}, Signal{gun_tile.ReceiveSignal, false}, false, nil})
	}

	level.SpikeTiles = nil

	imga, _, err = ebitenutil.NewImageFromFile("./art/special_tiles/spikeup.png")
	if err != nil {
		panic(err)
	}
	imgb, _, err = ebitenutil.NewImageFromFile("./art/special_tiles/spikeright.png")
	if err != nil {
		panic(err)
	}
	imgc, _, err = ebitenutil.NewImageFromFile("./art/special_tiles/spikedown.png")
	if err != nil {
		panic(err)
	}
	imgd, _, err = ebitenutil.NewImageFromFile("./art/special_tiles/spikeleft.png")
	if err != nil {
		panic(err)
	}

	spike_images := []*ebiten.Image{
		imga,
		imgb,
		imgc,
		imgd,
	}

	for _, spike_tile := range temp_level_json.SpikeTiles {
		level.SpikeTiles = append(level.SpikeTiles, SpikeTile{spike_tile.Pos, spike_tile.Damage, spike_images[spike_tile.Direction]})
	}

	imga, _, err = ebitenutil.NewImageFromFile("./art/special_tiles/springup.png")
	if err != nil {
		panic(err)
	}
	imgb, _, err = ebitenutil.NewImageFromFile("./art/special_tiles/springright.png")
	if err != nil {
		panic(err)
	}
	imgc, _, err = ebitenutil.NewImageFromFile("./art/special_tiles/springdown.png")
	if err != nil {
		panic(err)
	}
	imgd, _, err = ebitenutil.NewImageFromFile("./art/special_tiles/springleft.png")
	if err != nil {
		panic(err)
	}

	spring_images := []*ebiten.Image{
		imga,
		imgb,
		imgc,
		imgd,
	}

	for _, spring_tiles := range temp_level_json.SpringTiles {
		level.SpringTiles = append(level.SpringTiles, SpringTile{spring_tiles.Pos, spring_tiles.Power, spring_images[spring_tiles.Direction], spring_tiles.Direction})
	}

	for i, spawner := range level.Enemy_Spawners {
		if spawner.ReceiveSignal.Id == 0 {
			level.Send_Signals = append(level.Send_Signals, &level.Enemy_Spawners[i].SendSignal)
		} else {
			level.Receive_Signals = append(level.Receive_Signals, &level.Enemy_Spawners[i].ReceiveSignal)
			level.Send_Signals = append(level.Send_Signals, &level.Enemy_Spawners[i].SendSignal)
		}
	}

	for i := range level.BreakableTile {
		level.Receive_Signals = append(level.Receive_Signals, &level.BreakableTile[i].ReceiveSignal)
	}

	for i := range level.TriggerTile {
		level.Send_Signals = append(level.Send_Signals, &level.TriggerTile[i].SendSignal)
	}

	for i := range level.GunTiles {
		level.Send_Signals = append(level.Send_Signals, &level.GunTiles[i].SendSignal)
		level.Receive_Signals = append(level.Receive_Signals, &level.GunTiles[i].ReceiveSignal)
	}

	level.TileBorderColor = temp_level_json.TileBorderColor
	level.TileColor = temp_level_json.TileColor
	level.BackgroundColor = temp_level_json.BackgroundColor

	return level
}
