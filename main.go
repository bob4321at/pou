package main

import (
	"main/gun"
	"main/level"
	"main/music"
	"main/player"
	"main/utils"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var Shoot_Now_Ui *ebiten.Image

func init() {
	var err error
	Shoot_Now_Ui, _, err = ebitenutil.NewImageFromFile("./art/shoot_now_ui.png")
	if err != nil {
		panic(err)
	}
}

type Game struct{}

func (g *Game) Update() error {
	rmx, rmy := ebiten.CursorPosition()
	utils.Mouse_X = float64(rmx)
	utils.Mouse_Y = float64(rmy)

	if ebiten.IsKeyPressed(ebiten.Key1) {
		player.Player.Gun = gun.CreateNerfGun()
	} else if ebiten.IsKeyPressed(ebiten.Key2) {
		player.Player.Gun = gun.CreateBeeGun()
	} else if ebiten.IsKeyPressed(ebiten.Key3) {
		player.Player.Gun = gun.CreateTwinMagGun()
	} else if ebiten.IsKeyPressed(ebiten.Key4) {
		player.Player.Gun = gun.CreateShotgun()
	}

	player.Player.Update()

	level.Temp_Level.Update()

	utils.GameTime += 1

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, strconv.Itoa(int(ebiten.ActualFPS())))

	level.Temp_Level.Draw(screen)

	player.Player.Draw(screen)

	if music.AtPeak {
		screen.DrawImage(Shoot_Now_Ui, &ebiten.DrawImageOptions{})
	}
}

func (g *Game) Layout(ow, oh int) (sw, sh int) {
	return 640, 360
}

func main() {
	ebiten.SetWindowSize(1920, 1080)

	if err := ebiten.RunGame(&Game{}); err != nil {
		panic(err)
	}
}
