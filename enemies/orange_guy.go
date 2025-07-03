package enemies

import (
	"main/camera"
	"main/shaders"
	"main/utils"

	"github.com/bob4321at/textures"
	"github.com/hajimehoshi/ebiten/v2"
)

type OrangeGuy struct {
	Pos utils.Vec2
	Vel utils.Vec2

	Img      textures.RenderableTexture
	Health   int
	I_Frames float64
}

func (orange_guy *OrangeGuy) Draw(screen *ebiten.Image) {
	op := ebiten.DrawImageOptions{}
	op.GeoM.Translate(orange_guy.Pos.X+camera.Camera.Pos.X, orange_guy.Pos.Y+camera.Camera.Pos.Y)

	orange_guy.Img.SetUniforms(map[string]any{
		"I": orange_guy.I_Frames,
	})

	orange_guy.Img.Draw(screen, &op)
}

func (orange_guy *OrangeGuy) Update() {
	if orange_guy.I_Frames > 0 {
		orange_guy.I_Frames -= 0.1
	} else {
		orange_guy.I_Frames = 0
	}

	orange_guy.Vel.Y += 0.1

	if Player_Pos.X+320 > orange_guy.Pos.X {
		orange_guy.Vel.X += 0.1
	} else if Player_Pos.X+320 < orange_guy.Pos.X {
		orange_guy.Vel.X -= 0.1
	}

	y_collision, _ := CheckLevelCollision(utils.Vec2{X: orange_guy.Pos.X, Y: orange_guy.Pos.Y + orange_guy.Vel.Y}, utils.Vec2{X: 48, Y: 64})
	if y_collision {
		orange_guy.Vel.Y = 0
	}

	x_collision, _ := CheckLevelCollision(utils.Vec2{X: orange_guy.Pos.X + orange_guy.Vel.X, Y: orange_guy.Pos.Y}, utils.Vec2{X: 48, Y: 64})
	if x_collision {
		orange_guy.Vel.X = 0
	}

	orange_guy.Pos.Y += orange_guy.Vel.Y
	orange_guy.Pos.X += orange_guy.Vel.X
}

func (orange_guy *OrangeGuy) Hit(damage int) {
	if orange_guy.I_Frames <= 0 {
		orange_guy.Health -= damage
		orange_guy.I_Frames = 2
	}
}

func (orange_guy *OrangeGuy) GetPosition() utils.Vec2 {
	return orange_guy.Pos
}

func (orange_guy *OrangeGuy) GetSize() utils.Vec2 {
	return utils.Vec2{X: float64(orange_guy.Img.GetTexture().Bounds().Dx()), Y: float64(orange_guy.Img.GetTexture().Bounds().Dy())}
}

func (orange_guy *OrangeGuy) GetHealth() int {
	return orange_guy.Health
}

func (orange_guy *OrangeGuy) HitPlayer() {}

func NewOrangeGuy(pos utils.Vec2) Enemy {
	orange_guy := OrangeGuy{}
	orange_guy.Pos = pos
	orange_guy.Health = 10

	orange_guy.Img = textures.NewTexture("./art/enemies/orange_guy.png", shaders.Flash_Shader)

	return &orange_guy
}
