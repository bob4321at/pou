package enemies

import (
	"main/camera"
	"main/shaders"
	"main/utils"

	"github.com/bob4321at/textures"
	"github.com/hajimehoshi/ebiten/v2"
)

type RedGuy struct {
	Pos utils.Vec2
	Vel utils.Vec2

	Img            textures.RenderableTexture
	Health         int
	I_Frames       float64
	Jump_Timer     float64
	Max_Jump_Timer float64
}

func (red_guy *RedGuy) Draw(screen *ebiten.Image) {
	op := ebiten.DrawImageOptions{}
	op.GeoM.Translate(red_guy.Pos.X+camera.Camera.Pos.X, red_guy.Pos.Y+camera.Camera.Pos.Y)

	red_guy.Img.SetUniforms(map[string]any{
		"I": red_guy.I_Frames,
	})

	red_guy.Img.Draw(screen, &op)
}

func (red_guy *RedGuy) Update() {
	if red_guy.I_Frames > 0 {
		red_guy.I_Frames -= 0.1
	} else {
		red_guy.I_Frames = 0
	}

	red_guy.Vel.Y += 0.1

	if Player_Pos.X+320 > red_guy.Pos.X {
		red_guy.Vel.X += 0.1
	} else if Player_Pos.X+320 < red_guy.Pos.X {
		red_guy.Vel.X -= 0.1
	}

	if red_guy.Jump_Timer > 0 {
		red_guy.Jump_Timer -= 0.1
	} else {
		red_guy.Jump_Timer = red_guy.Max_Jump_Timer
		red_guy.Vel.Y = -5
	}

	red_guy_enemy := Enemy(red_guy)
	for i := range AllEnemies {
		enemy := AllEnemies[i]
		if enemy != red_guy_enemy {
			if utils.Collide(red_guy.Pos, red_guy.GetSize(), enemy.GetPosition(), enemy.GetSize()) {
				if red_guy.Pos.X+(red_guy.GetSize().X/2) > enemy.GetPosition().X+(enemy.GetSize().X/2) {
					red_guy.Vel.X = 1
				} else {
					red_guy.Vel.X = -1
				}
			}
		}
	}

	y_collision, _ := CheckLevelCollision(utils.Vec2{X: red_guy.Pos.X, Y: red_guy.Pos.Y + 4 + red_guy.Vel.Y}, utils.Vec2{X: 48, Y: 60})
	if y_collision {
		red_guy.Vel.Y = 0
	}

	x_collision, _ := CheckLevelCollision(utils.Vec2{X: red_guy.Pos.X + red_guy.Vel.X, Y: red_guy.Pos.Y + 4}, utils.Vec2{X: 48, Y: 60})
	if x_collision {
		red_guy.Vel.X = 0
	}

	red_guy.Pos.Y += red_guy.Vel.Y
	red_guy.Pos.X += red_guy.Vel.X
}

func (red_guy *RedGuy) Hit(damage int) {
	if red_guy.I_Frames <= 0 {
		red_guy.Health -= damage
		red_guy.I_Frames = 2
	}
}

func (red_guy *RedGuy) GetPosition() utils.Vec2 {
	return red_guy.Pos
}

func (red_guy *RedGuy) GetSize() utils.Vec2 {
	return utils.Vec2{X: float64(red_guy.Img.GetTexture().Bounds().Dx()), Y: float64(red_guy.Img.GetTexture().Bounds().Dy())}
}

func (red_guy *RedGuy) GetHealth() int {
	return red_guy.Health
}

func (red_guy *RedGuy) HitPlayer() {
	*Player_Health -= 3
}

func NewRedGuy(pos utils.Vec2) Enemy {
	red_guy := RedGuy{}
	red_guy.Pos = pos
	red_guy.Health = 20
	red_guy.Jump_Timer = 20
	red_guy.Max_Jump_Timer = 20

	red_guy.Img = textures.NewTexture("./art/enemies/red_guy.png", shaders.Flash_Shader)

	return &red_guy
}
