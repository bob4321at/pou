package enemies

import (
	"main/camera"
	"main/shaders"
	"main/utils"

	"github.com/bob4321at/textures"
	"github.com/hajimehoshi/ebiten/v2"
)

type Target struct {
	Pos utils.Vec2

	Img      textures.RenderableTexture
	Health   int
	I_Frames float64
}

func (target *Target) Draw(screen *ebiten.Image) {
	op := ebiten.DrawImageOptions{}
	op.GeoM.Translate(target.Pos.X+camera.Camera.Pos.X, target.Pos.Y+camera.Camera.Pos.Y)

	target.Img.SetUniforms(map[string]any{
		"I": 1000,
	})

	target.Img.Draw(screen, &op)
}

func (target *Target) Update() {
	if target.I_Frames > 0 {
		target.I_Frames -= 0.1
	} else {
		target.I_Frames = 0
	}

}

func (target *Target) Hit(damage int) {
	if target.I_Frames <= 0 {
		target.Health -= damage
		target.I_Frames = 2
	}
}

func (target *Target) GetPosition() utils.Vec2 {
	return target.Pos
}

func (target *Target) GetSize() utils.Vec2 {
	return utils.Vec2{X: float64(target.Img.GetTexture().Bounds().Dx()), Y: float64(target.Img.GetTexture().Bounds().Dy())}
}

func (target *Target) GetHealth() int {
	return target.Health
}

func NewTarget(pos utils.Vec2) *Target {
	target := Target{}
	target.Pos = pos
	target.Health = 10

	target.Img = textures.NewTexture("./art/target.png", shaders.Enemy_Shader)

	return &target
}
