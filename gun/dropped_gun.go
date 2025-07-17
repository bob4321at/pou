package gun

import (
	"main/camera"
	"main/enemies"
	"main/utils"

	"github.com/hajimehoshi/ebiten/v2"
)

type DroppedGunStruct struct {
	Img       *ebiten.Image
	Pos       utils.Vec2
	Vel       utils.Vec2
	GiveFunc  func() Gun
	Picked_Up bool
}

func (dropped_gun *DroppedGunStruct) Draw(screen *ebiten.Image) {
	op := ebiten.DrawImageOptions{}
	op.GeoM.Translate(dropped_gun.Pos.X+camera.Camera.Pos.X, dropped_gun.Pos.Y+camera.Camera.Pos.Y)

	screen.DrawImage(dropped_gun.Img, &op)
}

func (dropped_gun *DroppedGunStruct) Update() {
	dropped_gun.Vel.Y += 0.1

	collide_y, _ := enemies.CheckLevelCollision(utils.Vec2{X: dropped_gun.Pos.X, Y: dropped_gun.Pos.Y + dropped_gun.Vel.Y}, utils.Vec2{X: float64(dropped_gun.Img.Bounds().Dx()), Y: float64(dropped_gun.Img.Bounds().Dy())})

	if collide_y {
		dropped_gun.Vel.Y = 0
	}

	if utils.Collide(dropped_gun.Pos, utils.Vec2{X: float64(dropped_gun.Img.Bounds().Dx()), Y: float64(dropped_gun.Img.Bounds().Dy())}, utils.Vec2{X: Player_Pos.X + 320, Y: Player_Pos.Y + 180}, utils.Vec2{X: 32, Y: 48}) {
		if ebiten.IsKeyPressed(ebiten.KeyE) {
			dropped_gun.Picked_Up = true
		}
	}

	dropped_gun.Pos.Y += dropped_gun.Vel.Y
}

func CreateDroppedGun(gun func() Gun, pos utils.Vec2) DroppedGunStruct {
	dropped_gun := DroppedGunStruct{}

	dropped_gun.Img = gun().GetDroppedImg().GetTexture()
	dropped_gun.Pos = pos
	dropped_gun.GiveFunc = gun

	return dropped_gun
}
