package enemies

import (
	"main/camera"
	"main/utils"

	"github.com/bob4321at/textures"
	"github.com/hajimehoshi/ebiten/v2"
)

var TURRET_BULLET_IMAGE = textures.NewTexture("./art/enemies/turret_bullet.png", "")

type TurretBullet struct {
	Pos             utils.Vec2
	Vel             utils.Vec2
	ShouldRemoveVar bool
}

func (bullet *TurretBullet) Draw(screen *ebiten.Image) {
	op := ebiten.DrawImageOptions{}
	op.GeoM.Translate(bullet.Pos.X+camera.Camera.Pos.X, bullet.Pos.Y+camera.Camera.Pos.Y)
	TURRET_BULLET_IMAGE.Draw(screen, &op)
}

func (bullet *TurretBullet) Update() {
	bullet.Pos.X += bullet.Vel.X
	bullet.Pos.Y += bullet.Vel.Y
}

func (bullet *TurretBullet) Hit() {
	*Player_Health -= 1
	bullet.ShouldRemoveVar = true
}

func (bullet *TurretBullet) GetSize() utils.Vec2 {
	return utils.Vec2{X: float64(TURRET_BULLET_IMAGE.Img.Bounds().Dx()), Y: float64(TURRET_BULLET_IMAGE.Img.Bounds().Dy())}
}

func (bullet *TurretBullet) GetPos() utils.Vec2 {
	return bullet.Pos
}

func (bullet *TurretBullet) ShouldRemove() bool {
	return bullet.ShouldRemoveVar
}

func NewTurretBullet(pos, vel utils.Vec2) *TurretBullet {
	turret_bullet := TurretBullet{
		Pos:             pos,
		Vel:             vel,
		ShouldRemoveVar: false,
	}

	return &turret_bullet
}
