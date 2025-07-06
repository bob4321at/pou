package scenes

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Death_Menu_scene struct {
	Img     *ebiten.Image
	SetedUp bool
}

func (scene *Death_Menu_scene) Setup() {
	scene.Img, _, _ = ebitenutil.NewImageFromFile("./art/menus/dead.png")
	scene.SetedUp = true
}

func (scene *Death_Menu_scene) Draw(display *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(10, 10)
	display.DrawImage(scene.Img, op)
}

func (scene *Death_Menu_scene) Update() {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButton0) {
		// if utils.Collide(utils.Vec2{X: 2, Y: 10}, utils.Vec2{X: 19, Y: 8}, utils.Vec2{X: utils.Mouse_X / 10, Y: utils.Mouse_Y / 10}, utils.Vec2{X: 1, Y: 1}) {
		Current_Scene_Id = 1
		// }
	}
}

func (scene *Death_Menu_scene) GetSetup() bool {
	return scene.SetedUp
}
