package enemies

import (
	"main/camera"
	"main/shaders"
	"main/utils"
	"math"

	"github.com/bob4321at/textures"
	"github.com/hajimehoshi/ebiten/v2"
)

type TurretBullet struct {
	Pos utils.Vec2
	Vel utils.Vec2

	Img      textures.RenderableTexture
	Health   int
	I_Frames float64
}

func (bullet *TurretBullet) Draw(screen *ebiten.Image) {
	op := ebiten.DrawImageOptions{}
	op.GeoM.Translate(bullet.Pos.X+camera.Camera.Pos.X, bullet.Pos.Y+camera.Camera.Pos.Y)

	bullet.Img.SetUniforms(map[string]any{
		"I": bullet.I_Frames,
	})

	bullet.Img.Draw(screen, &op)
}

func (bullet *TurretBullet) Update() {
	if bullet.I_Frames > 0 {
		bullet.I_Frames -= 0.1
	} else {
		bullet.I_Frames = 0
	}

	bullet.Pos.X += bullet.Vel.X
	bullet.Pos.Y += bullet.Vel.Y
}

func (bullet *TurretBullet) Hit(damage int) {
	if bullet.I_Frames <= 0 {
		bullet.Health -= damage
		bullet.I_Frames = 2
	}
}

func (bullet *TurretBullet) GetPosition() utils.Vec2 {
	return bullet.Pos
}

func (bullet *TurretBullet) GetSize() utils.Vec2 {
	return utils.Vec2{X: float64(bullet.Img.GetTexture().Bounds().Dx()), Y: float64(bullet.Img.GetTexture().Bounds().Dy())}
}

func (bullet *TurretBullet) GetHealth() int {
	return bullet.Health
}

func (bullet *TurretBullet) HitPlayer() {
	*Player_Health -= 1
}

func NewTurretBullet(pos utils.Vec2) Enemy {
	bullet := TurretBullet{}
	bullet.Pos = pos
	bullet.Health = 10

	bullet.Img = textures.NewTexture("./art/enemies/turret_bullet.png", shaders.Flash_Shader)

	angle := utils.GetAngle(utils.Vec2{X: Player_Pos.X + 320, Y: Player_Pos.Y + 180}, bullet.Pos)

	bullet.Vel.X = math.Cos(angle-utils.Deg2Rad(90)) * 10
	bullet.Vel.Y = -math.Sin(angle-utils.Deg2Rad(90)) * 10

	return &bullet
}
