package enemies

import (
	"main/camera"
	"main/shaders"
	"main/utils"

	"github.com/bob4321at/textures"
	"github.com/hajimehoshi/ebiten/v2"
)

type GreenGuy struct {
	Pos utils.Vec2
	Vel utils.Vec2

	Img        textures.RenderableTexture
	Health     int
	Max_Health int
	I_Frames   float64

	Heal_Timer     float64
	Max_Heal_timer float64
}

func (green_guy *GreenGuy) Draw(screen *ebiten.Image) {
	op := ebiten.DrawImageOptions{}
	op.GeoM.Translate(green_guy.Pos.X+camera.Camera.Pos.X, green_guy.Pos.Y+camera.Camera.Pos.Y)

	green_guy.Img.SetUniforms(map[string]any{
		"I": green_guy.I_Frames,
	})

	green_guy.Img.Draw(screen, &op)
}

func (green_guy *GreenGuy) Update() {
	if green_guy.I_Frames > 0 {
		green_guy.I_Frames -= 0.1
	} else {
		green_guy.I_Frames = 0
	}

	green_guy.Vel.Y += 0.1

	if Player_Pos.X+320 > green_guy.Pos.X {
		green_guy.Vel.X += 0.1
	} else if Player_Pos.X+320 < green_guy.Pos.X {
		green_guy.Vel.X -= 0.1
	}

	green_guy_enemy := Enemy(green_guy)
	for i := range AllEnemies {
		enemy := AllEnemies[i]
		if enemy != green_guy_enemy {
			if utils.Collide(green_guy.Pos, green_guy.GetSize(), enemy.GetPosition(), enemy.GetSize()) {
				if green_guy.Pos.X+(green_guy.GetSize().X/2) > enemy.GetPosition().X+(enemy.GetSize().X/2) {
					green_guy.Vel.X = 1
				} else {
					green_guy.Vel.X = -1
				}
			}
		}
	}

	if green_guy.Heal_Timer > 0 {
		green_guy.Heal_Timer -= 0.1
	} else {
		green_guy.Heal_Timer = green_guy.Max_Heal_timer
		green_guy.Health += 1
		if green_guy.Health > green_guy.Max_Health {
			green_guy.Health = green_guy.Max_Health
		}
	}

	y_collision, _ := CheckLevelCollision(utils.Vec2{X: green_guy.Pos.X, Y: green_guy.Pos.Y + 2 + green_guy.Vel.Y}, utils.Vec2{X: 32, Y: 40})
	if y_collision {
		green_guy.Vel.Y = 0
	}

	x_collision, _ := CheckLevelCollision(utils.Vec2{X: green_guy.Pos.X + green_guy.Vel.X, Y: green_guy.Pos.Y + 2}, utils.Vec2{X: 32, Y: 40})
	if x_collision {
		green_guy.Vel.X = 0
	}

	green_guy.Pos.Y += green_guy.Vel.Y
	green_guy.Pos.X += green_guy.Vel.X
}

func (green_guy *GreenGuy) Hit(damage int) {
	if green_guy.I_Frames <= 0 {
		green_guy.Health -= damage
		green_guy.I_Frames = 2
	}
}

func (green_guy *GreenGuy) GetPosition() utils.Vec2 {
	return green_guy.Pos
}

func (green_guy *GreenGuy) GetSize() utils.Vec2 {
	return utils.Vec2{X: float64(green_guy.Img.GetTexture().Bounds().Dx()), Y: float64(green_guy.Img.GetTexture().Bounds().Dy())}
}

func (green_guy *GreenGuy) GetHealth() int {
	return green_guy.Health
}

func (green_guy *GreenGuy) HitPlayer() {
	*Player_Health -= 3
}

func NewGreenGuy(pos utils.Vec2) Enemy {
	blue_guy := GreenGuy{}
	blue_guy.Pos = pos
	blue_guy.Health = 16
	blue_guy.Max_Health = blue_guy.Health

	blue_guy.Heal_Timer = 10
	blue_guy.Max_Heal_timer = blue_guy.Heal_Timer

	blue_guy.Img = textures.NewTexture("./art/enemies/green_guy.png", shaders.Flash_Shader)

	return &blue_guy
}
