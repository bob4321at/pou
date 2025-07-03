package scenes

import (
	"main/utils"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Main_Menu_Scene struct {
	Img     *ebiten.Image
	SetedUp bool
}

func (scene *Main_Menu_Scene) Setup() {
	scene.Img, _, _ = ebitenutil.NewImageFromFile("./art/menus/main_menu.png")
	scene.SetedUp = true
}

func (scene *Main_Menu_Scene) Draw(display *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(10, 10)
	display.DrawImage(scene.Img, op)
}

func (scene *Main_Menu_Scene) Update() {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButton0) {
		if utils.Collide(utils.Vec2{X: 2, Y: 10}, utils.Vec2{X: 19, Y: 8}, utils.Vec2{X: utils.Mouse_X / 10, Y: utils.Mouse_Y / 10}, utils.Vec2{X: 1, Y: 1}) {
			Current_Scene_Id = 1
		} else if utils.Collide(utils.Vec2{X: 2, Y: 18}, utils.Vec2{X: 20, Y: 7}, utils.Vec2{X: utils.Mouse_X / 10, Y: utils.Mouse_Y / 10}, utils.Vec2{X: 1, Y: 1}) {
			os.Exit(0)
		}
	}
}

func (scene *Main_Menu_Scene) GetSetup() bool {
	return scene.SetedUp
}
