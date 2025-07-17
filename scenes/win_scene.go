package scenes

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Win_Menu_Scene struct {
	Img     *ebiten.Image
	SetedUp bool
}

func (scene *Win_Menu_Scene) Setup() {
	scene.Img, _, _ = ebitenutil.NewImageFromFile("./art/menus/win.png")
	scene.SetedUp = true
}

func (scene *Win_Menu_Scene) Draw(display *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	display.DrawImage(scene.Img, op)
}

func (scene *Win_Menu_Scene) Update() {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButton0) {
		Current_Scene_Id = 1
	}
}

func (scene *Win_Menu_Scene) GetSetup() bool {
	return scene.SetedUp
}
