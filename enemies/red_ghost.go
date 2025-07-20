package enemies

import (
	"main/camera"
	"main/shaders"
	"main/utils"
	"math"

	"github.com/bob4321at/textures"
	"github.com/hajimehoshi/ebiten/v2"
)

type RedGhost struct {
	Pos utils.Vec2
	Vel utils.Vec2

	Img            textures.RenderableTexture
	Health         int
	I_Frames       float64
	Dash_Timer     float64
	Max_Dash_Timer float64
}

func (red_ghost *RedGhost) Draw(screen *ebiten.Image) {
	op := ebiten.DrawImageOptions{}
	op.GeoM.Translate(red_ghost.Pos.X+camera.Camera.Pos.X, red_ghost.Pos.Y+camera.Camera.Pos.Y)

	red_ghost.Img.SetUniforms(map[string]any{
		"I": red_ghost.I_Frames,
	})

	red_ghost.Img.Draw(screen, &op)
}

func (red_ghost *RedGhost) Update() {
	if red_ghost.I_Frames > 0 {
		red_ghost.I_Frames -= 0.1
	} else {
		red_ghost.I_Frames = 0
	}

	if Player_Pos.X+310 > red_ghost.Pos.X {
		red_ghost.Vel.X += 0.1
	} else if Player_Pos.X+310 < red_ghost.Pos.X {
		red_ghost.Vel.X -= 0.1
	}

	if Player_Pos.Y+160 > red_ghost.Pos.Y {
		red_ghost.Vel.Y += 0.1
	} else if Player_Pos.Y+160 < red_ghost.Pos.Y {
		red_ghost.Vel.Y -= 0.1
	}

	if red_ghost.Dash_Timer > 0 {
		red_ghost.Dash_Timer -= 0.1

		if red_ghost.Vel.X > 3 {
			red_ghost.Vel.X -= 0.1
		} else if red_ghost.Vel.X < -3 {
			red_ghost.Vel.X -= -0.1
		}

		if red_ghost.Vel.Y > 3 {
			red_ghost.Vel.Y -= 0.1
		} else if red_ghost.Vel.Y < -3 {
			red_ghost.Vel.Y -= -0.1
		}
	} else {
		red_ghost.Dash_Timer = red_ghost.Max_Dash_Timer

		angle := utils.GetAngle(utils.Vec2{X: Player_Pos.X + 320, Y: Player_Pos.Y + 180}, red_ghost.Pos)

		red_ghost.Vel.X = -math.Cos(angle+90) * 10
		red_ghost.Vel.Y = math.Sin(angle+90) * 10
	}

	red_ghost_enemy := Enemy(red_ghost)
	for i := range AllEnemies {
		enemy := AllEnemies[i]
		_, ok := enemy.(*RedGhost)
		if ok {
			if enemy != red_ghost_enemy {
				if utils.Collide(red_ghost.Pos, red_ghost.GetSize(), enemy.GetPosition(), enemy.GetSize()) {
					if red_ghost.Pos.X+(red_ghost.GetSize().X/2) > enemy.GetPosition().X+(enemy.GetSize().X/2) {
						red_ghost.Vel.X = 1
					} else {
						red_ghost.Vel.X = -1
					}
					if red_ghost.Pos.Y+(red_ghost.GetSize().Y/2) > enemy.GetPosition().Y+(enemy.GetSize().Y/2) {
						red_ghost.Vel.Y = 1
					} else {
						red_ghost.Vel.Y = -1
					}
				}
			}
		}
	}

	red_ghost.Pos.Y += red_ghost.Vel.Y
	red_ghost.Pos.X += red_ghost.Vel.X
}

func (red_ghost *RedGhost) Hit(damage int) {
	if red_ghost.I_Frames <= 0 {
		red_ghost.Health -= damage
		red_ghost.I_Frames = 2
	}
}

func (red_ghost *RedGhost) GetPosition() utils.Vec2 {
	return red_ghost.Pos
}

func (red_ghost *RedGhost) GetSize() utils.Vec2 {
	return utils.Vec2{X: float64(red_ghost.Img.GetTexture().Bounds().Dx()), Y: float64(red_ghost.Img.GetTexture().Bounds().Dy())}
}

func (red_ghost *RedGhost) GetHealth() int {
	return red_ghost.Health
}

func (red_ghost *RedGhost) HitPlayer() {
	*Player_Health -= 1
}

func NewRedGhost(pos utils.Vec2) Enemy {
	red_ghost := RedGhost{}
	red_ghost.Pos = pos
	red_ghost.Health = 20
	red_ghost.Dash_Timer = 20
	red_ghost.Max_Dash_Timer = 20

	red_ghost.Img = textures.NewTexture("./art/enemies/red_ghost.png", shaders.Flash_Shader)

	return &red_ghost
}
