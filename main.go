package main

import (
	"main/level"
	"main/player"
	"main/utils"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct{}

func (g *Game) Update() error {
	rmx, rmy := ebiten.CursorPosition()
	utils.Mouse_X = float64(rmx)
	utils.Mouse_Y = float64(rmy)

	player.Player.Update()

	level.Temp_Level.Update()

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, strconv.Itoa(int(ebiten.ActualFPS())))

	level.Temp_Level.Draw(screen)

	player.Player.Draw(screen)
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
