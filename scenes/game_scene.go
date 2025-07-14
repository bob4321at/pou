package scenes

import (
	"image/color"
	"main/gun"
	"main/level"
	"main/music"
	"main/player"
	"main/shaders"
	"main/utils"
	"strconv"

	"github.com/bob4321at/textures"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type GameScene struct {
	Shoot_Now_Ui  *ebiten.Image
	SetedUp       bool
	Music         music.MusicStruct
	Player        player.PlayerStruct
	Screen        *textures.Texture
	Current_Level level.Level
}

func (scene *GameScene) Setup() {
	var err error
	scene.Shoot_Now_Ui, _, err = ebitenutil.NewImageFromFile("./art/shoot_now_ui.png")
	if err != nil {
		panic(err)
	}
	scene.Music = music.NewMusic("./music/song.mp3")

	scene.Screen = textures.NewTexture("./art/display.png", shaders.Test_Refraction_Shader)

	scene.Current_Level = level.LoadLevel("test_real_level")

	scene.Player = player.NewPlayer(utils.Vec2{X: 100, Y: 100})

	scene.SetedUp = true
}

func (scene *GameScene) Update() {
	if ebiten.IsKeyPressed(ebiten.Key1) {
		scene.Player.Gun = gun.CreateNerfGun()
	} else if ebiten.IsKeyPressed(ebiten.Key2) {
		scene.Player.Gun = gun.CreateBeeGun()
	} else if ebiten.IsKeyPressed(ebiten.Key3) {
		scene.Player.Gun = gun.CreateTwinMagGun()
	} else if ebiten.IsKeyPressed(ebiten.Key4) {
		scene.Player.Gun = gun.CreateShotgun()
	} else if ebiten.IsKeyPressed(ebiten.Key5) {
		scene.Player.Gun = gun.CreateMechaGun()
	}

	if scene.Player.Health <= 0 {
		scene.Reset()
		Current_Scene_Id = 2
	}

	if scene.Player.Won {
		scene.Reset()
		Current_Scene_Id = 3
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
	display.Fill(color.RGBA{255, 217, 217, 255})

	scene.Screen.Img.Fill(color.RGBA{0, 0, 0, 0})

	scene.Current_Level.Draw(scene.Screen.Img)

	scene.Player.Draw(scene.Screen.Img)

	scene.Screen.Draw(display, &ebiten.DrawImageOptions{})

	op := ebiten.DrawImageOptions{}
	op.GeoM.Translate(10, 10)
	op.GeoM.Scale(float64(scene.Player.Health)/100, 1)

	display.DrawImage(scene.Player.Health_Bar_Img, &op)

	ebitenutil.DebugPrint(display, strconv.Itoa(int(ebiten.ActualFPS())))
}

func (scene *GameScene) GetSetup() bool {
	return scene.SetedUp
}
