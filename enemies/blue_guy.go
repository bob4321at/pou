package enemies

import (
	"main/camera"
	"main/shaders"
	"main/utils"

	"github.com/bob4321at/textures"
	"github.com/hajimehoshi/ebiten/v2"
)

type BlueGuy struct {
	Pos utils.Vec2
	Vel utils.Vec2

	Img      textures.RenderableTexture
	Health   int
	I_Frames float64
}

func (blue_guy *BlueGuy) Draw(screen *ebiten.Image) {
	op := ebiten.DrawImageOptions{}
	op.GeoM.Translate(blue_guy.Pos.X+camera.Camera.Pos.X, blue_guy.Pos.Y+camera.Camera.Pos.Y)

	blue_guy.Img.SetUniforms(map[string]any{
		"I": blue_guy.I_Frames,
	})

	blue_guy.Img.Draw(screen, &op)
}

func (blue_guy *BlueGuy) Update() {
	if blue_guy.I_Frames > 0 {
		blue_guy.I_Frames -= 0.1
	} else {
		blue_guy.I_Frames = 0
	}

	blue_guy.Vel.Y += 0.1

	if Player_Pos.X+320 > blue_guy.Pos.X {
		blue_guy.Vel.X += 0.2
	} else if Player_Pos.X+320 < blue_guy.Pos.X {
		blue_guy.Vel.X -= 0.2
	}

	blue_guy_enemy := Enemy(blue_guy)
	for i := range AllEnemies {
		enemy := AllEnemies[i]
		if enemy != blue_guy_enemy {
			if utils.Collide(blue_guy.Pos, blue_guy.GetSize(), enemy.GetPosition(), enemy.GetSize()) {
				if blue_guy.Pos.X+(blue_guy.GetSize().X/2) > enemy.GetPosition().X+(enemy.GetSize().X/2) {
					blue_guy.Vel.X = 1
				} else {
					blue_guy.Vel.X = -1
				}
			}
		}
	}

	y_collision, _ := CheckLevelCollision(utils.Vec2{X: blue_guy.Pos.X, Y: blue_guy.Pos.Y + 2 + blue_guy.Vel.Y}, utils.Vec2{X: 32, Y: 40})
	if y_collision {
		blue_guy.Vel.Y = 0
	}

	x_collision, _ := CheckLevelCollision(utils.Vec2{X: blue_guy.Pos.X + blue_guy.Vel.X, Y: blue_guy.Pos.Y + 2}, utils.Vec2{X: 32, Y: 40})
	if x_collision {
		blue_guy.Vel.X = 0
	}

	blue_guy.Pos.Y += blue_guy.Vel.Y
	blue_guy.Pos.X += blue_guy.Vel.X
}

func (blue_guy *BlueGuy) Hit(damage int) {
	if blue_guy.I_Frames <= 0 {
		blue_guy.Health -= damage
		blue_guy.I_Frames = 2
	}
}

func (blue_guy *BlueGuy) GetPosition() utils.Vec2 {
	return blue_guy.Pos
}

func (blue_guy *BlueGuy) GetSize() utils.Vec2 {
	return utils.Vec2{X: float64(blue_guy.Img.GetTexture().Bounds().Dx()), Y: float64(blue_guy.Img.GetTexture().Bounds().Dy())}
}

func (blue_guy *BlueGuy) GetHealth() int {
	return blue_guy.Health
}

func (blue_guy *BlueGuy) HitPlayer() {
	*Player_Health -= 3
}

func NewBlueGuy(pos utils.Vec2) Enemy {
	blue_guy := BlueGuy{}
	blue_guy.Pos = pos
	blue_guy.Health = 8

	blue_guy.Img = textures.NewTexture("./art/enemies/blue_guy.png", shaders.Flash_Shader)

	return &blue_guy
}
