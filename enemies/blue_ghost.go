package enemies

import (
	"main/camera"
	"main/shaders"
	"main/utils"

	"github.com/bob4321at/textures"
	"github.com/hajimehoshi/ebiten/v2"
)

type BlueGhost struct {
	Pos utils.Vec2
	Vel utils.Vec2

	Img      textures.RenderableTexture
	Health   int
	I_Frames float64
}

func (blue_ghost *BlueGhost) Draw(screen *ebiten.Image) {
	op := ebiten.DrawImageOptions{}
	op.GeoM.Translate(blue_ghost.Pos.X+camera.Camera.Pos.X, blue_ghost.Pos.Y+camera.Camera.Pos.Y)

	blue_ghost.Img.SetUniforms(map[string]any{
		"I": blue_ghost.I_Frames,
	})

	blue_ghost.Img.Draw(screen, &op)
}

func (blue_ghost *BlueGhost) Update() {
	if blue_ghost.I_Frames > 0 {
		blue_ghost.I_Frames -= 0.1
	} else {
		blue_ghost.I_Frames = 0
	}

	if Player_Pos.X+320 > blue_ghost.Pos.X {
		blue_ghost.Vel.X += 0.1
	} else if Player_Pos.X+320 < blue_ghost.Pos.X {
		blue_ghost.Vel.X -= 0.1
	}

	if Player_Pos.Y+240 > blue_ghost.Pos.Y {
		blue_ghost.Vel.Y += 0.1
	} else if Player_Pos.Y+240 < blue_ghost.Pos.Y {
		blue_ghost.Vel.Y -= 0.1
	}

	if blue_ghost.Vel.X > 3 {
		blue_ghost.Vel.X = 3
	} else if blue_ghost.Vel.X < -3 {
		blue_ghost.Vel.X = -3
	}

	if blue_ghost.Vel.Y > 3 {
		blue_ghost.Vel.Y = 3
	} else if blue_ghost.Vel.Y < -3 {
		blue_ghost.Vel.Y = -3
	}

	blue_ghost.Pos.Y += blue_ghost.Vel.Y
	blue_ghost.Pos.X += blue_ghost.Vel.X
}

func (blue_ghost *BlueGhost) Hit(damage int) {
	if blue_ghost.I_Frames <= 0 {
		blue_ghost.Health -= damage
		blue_ghost.I_Frames = 2
	}
}

func (blue_ghost *BlueGhost) GetPosition() utils.Vec2 {
	return blue_ghost.Pos
}

func (blue_ghost *BlueGhost) GetSize() utils.Vec2 {
	return utils.Vec2{X: float64(blue_ghost.Img.GetTexture().Bounds().Dx()), Y: float64(blue_ghost.Img.GetTexture().Bounds().Dy())}
}

func (orange_guy *BlueGhost) GetHealth() int {
	return orange_guy.Health
}

func NewBlueGhost(pos utils.Vec2) Enemy {
	blue_ghost := BlueGhost{}
	blue_ghost.Pos = pos
	blue_ghost.Health = 10

	blue_ghost.Img = textures.NewTexture("./art/enemies/blue_ghost.png", shaders.Enemy_Shader)

	return &blue_ghost
}
