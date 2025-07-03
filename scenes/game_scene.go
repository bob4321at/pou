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
	Shoot_Now_Ui *ebiten.Image
	SetedUp      bool
	Music        music.MusicStruct
	Player       player.PlayerStruct
	Screen       *textures.Texture
}

func (scene *GameScene) Setup() {
	var err error
	scene.Shoot_Now_Ui, _, err = ebitenutil.NewImageFromFile("./art/shoot_now_ui.png")
	if err != nil {
		panic(err)
	}
	scene.Music = music.NewMusic("./music/song.mp3")

	scene.Screen = textures.NewTexture("./art/display.png", shaders.Test_Refraction_Shader)

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
	}

	scene.Player.Update()

	level.Temp_Level.Update()

	utils.GameTime += 1
}

func (scene *GameScene) Draw(display *ebiten.Image) {
	display.Fill(color.RGBA{255, 217, 217, 255})

	scene.Screen.Img.Fill(color.RGBA{0, 0, 0, 0})

	level.Temp_Level.Draw(scene.Screen.Img)

	scene.Player.Draw(scene.Screen.Img)

	if music.AtPeak {
		scene.Screen.Img.DrawImage(scene.Shoot_Now_Ui, &ebiten.DrawImageOptions{})
	}

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
