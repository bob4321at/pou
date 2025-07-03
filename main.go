package main

import (
	"main/scenes"
	"main/utils"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct{}

func (g *Game) Update() error {
	rmx, rmy := ebiten.CursorPosition()
	utils.Mouse_X = float64(rmx)
	utils.Mouse_Y = float64(rmy)

	if !scenes.Scenes[scenes.Current_Scene_Id].GetSetup() {
		scenes.Scenes[scenes.Current_Scene_Id].Setup()
	}

	scenes.Scenes[scenes.Current_Scene_Id].Update()

	return nil
}

func (g *Game) Draw(display *ebiten.Image) {
	if !scenes.Scenes[scenes.Current_Scene_Id].GetSetup() {
		scenes.Scenes[scenes.Current_Scene_Id].Setup()
	}

	scenes.Scenes[scenes.Current_Scene_Id].Draw(display)
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
