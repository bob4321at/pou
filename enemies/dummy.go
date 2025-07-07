package enemies

import (
	"main/camera"
	"main/shaders"
	"main/utils"

	"github.com/bob4321at/textures"
	"github.com/hajimehoshi/ebiten/v2"
)

type Dummy struct {
	Pos utils.Vec2
	Vel utils.Vec2

	Img      textures.RenderableTexture
	Health   int
	I_Frames float64
}

func (dummy *Dummy) Draw(screen *ebiten.Image) {
	op := ebiten.DrawImageOptions{}
	op.GeoM.Translate(dummy.Pos.X+camera.Camera.Pos.X, dummy.Pos.Y+camera.Camera.Pos.Y)

	dummy.Img.SetUniforms(map[string]any{
		"I": dummy.I_Frames,
	})

	dummy.Img.Draw(screen, &op)
}

func (dummy *Dummy) Update() {
	dummy.Vel.Y += 0.1

	if dummy.I_Frames > 0 {
		dummy.I_Frames -= 0.1
	} else {
		dummy.I_Frames = 0
	}

	y_collision, _ := CheckLevelCollision(utils.Vec2{X: dummy.Pos.X, Y: dummy.Pos.Y + dummy.Vel.Y}, utils.Vec2{X: 32, Y: 48})
	if y_collision {
		dummy.Vel.Y = 0
	}

	dummy.Pos.Y += dummy.Vel.Y
}

func (dummy *Dummy) Hit(damage int) {

	if dummy.I_Frames <= 0 {
		dummy.Health -= damage
		dummy.I_Frames = 2
	}
}

func (dummy *Dummy) GetPosition() utils.Vec2 {
	return dummy.Pos
}

func (dummy *Dummy) GetSize() utils.Vec2 {
	return utils.Vec2{X: float64(dummy.Img.GetTexture().Bounds().Dx()), Y: float64(dummy.Img.GetTexture().Bounds().Dy())}
}

func (target *Dummy) GetHealth() int {
	return target.Health
}

func (dummy *Dummy) HitPlayer() {}

func NewDummy(pos utils.Vec2) Enemy {
	target := Dummy{}
	target.Pos = pos
	target.Health = 100

	target.Img = textures.NewTexture("./art/enemies/dummy.png", shaders.Flash_Shader)

	return &target
}
