package scenes

import (
	"image/color"
	"main/level"
	"main/music"
	"main/player"
	"main/shaders"
	"main/utils"
	worldmap "main/world_map"
	"strconv"

	"github.com/bob4321at/textures"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type GameScene struct {
	SetedUp       bool
	Music         music.MusicStruct
	Player        player.PlayerStruct
	Screen        *textures.Texture
	Current_Level level.Level
}

func (scene *GameScene) Setup() {
	var err error
	if err != nil {
		panic(err)
	}
	scene.Music = music.NewMusic("./music/song.mp3")

	scene.Screen = textures.NewTexture("./art/display.png", shaders.Test_Refraction_Shader)

	scene.Player = player.NewPlayer(utils.Vec2{X: 100, Y: 100})

	scene.SetedUp = true
}

func (scene *GameScene) Update() {
	if worldmap.Level_To_Load != nil {
		scene.Current_Level = level.LoadLevel(*worldmap.Level_To_Load)
		worldmap.Level_To_Load = nil
	}

	for i := range scene.Current_Level.Dropped_Guns {
		dropped_gun := &scene.Current_Level.Dropped_Guns[i]
		if dropped_gun.Picked_Up {
			scene.Player.Gun = dropped_gun.GiveFunc()
			utils.RemoveArrayElement(i, &scene.Current_Level.Dropped_Guns)
			break
		}
	}

	if scene.Player.Health <= 0 {
		scene.Reset()
		Current_Scene_Id = 3
	}

	if scene.Player.Won {
		scene.Reset()
		Current_Scene_Id = 4
	}

	scene.Player.Update(&scene.Current_Level)

	scene.Current_Level.Update()

	utils.GameTime += 1
}

func (scene *GameScene) Reset() {
	scene.Current_Level.Reset()
	scene.Player.Reset(scene.Current_Level.Player_Spawn)
	scene.Music.Reset()
}

func (scene *GameScene) Draw(display *ebiten.Image) {
	display.Fill(scene.Current_Level.BackgroundColor)

	scene.Screen.Img.Fill(color.RGBA{0, 0, 0, 0})

	scene.Player.Draw(scene.Screen.Img)

	if scene.Current_Level.Tiles != nil {
		scene.Current_Level.Draw(scene.Screen.Img)
	}

	display_op := ebiten.DrawImageOptions{}
	display_op.GeoM.Translate(-2, -2)
	scene.Screen.Draw(display, &display_op)

	op := ebiten.DrawImageOptions{}
	op.GeoM.Translate(10, 10)
	op.GeoM.Scale(float64(scene.Player.Health)/100, 1)

	display.DrawImage(scene.Player.Health_Bar_Img, &op)

	ebitenutil.DebugPrint(display, strconv.Itoa(int(ebiten.ActualFPS())))
}

func (scene *GameScene) GetSetup() bool {
	return scene.SetedUp
}
