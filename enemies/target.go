package enemies

import (
	"main/camera"
	"main/utils"

	"github.com/bob4321at/textures"
	"github.com/hajimehoshi/ebiten/v2"
)

type Target struct {
	Pos utils.Vec2

	Img    textures.RenderableTexture
	Health int
}

func (target *Target) Draw(screen *ebiten.Image) {
	op := ebiten.DrawImageOptions{}
	op.GeoM.Translate(target.Pos.X+camera.Camera.Pos.X, target.Pos.Y+camera.Camera.Pos.Y)

	target.Img.Draw(screen, &op)
}

func (target *Target) Update() {}

func (target *Target) Hit(damage int) {
	target.Health -= damage
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
	target.Health = 1

	target.Img = textures.NewTexture("./art/target.png", "")

	return &target
}
