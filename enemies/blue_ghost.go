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

	if Player_Pos.X+310 > blue_ghost.Pos.X {
		blue_ghost.Vel.X += 0.1
	} else if Player_Pos.X+310 < blue_ghost.Pos.X {
		blue_ghost.Vel.X -= 0.1
	}

	if Player_Pos.Y+160 > blue_ghost.Pos.Y {
		blue_ghost.Vel.Y += 0.1
	} else if Player_Pos.Y+160 < blue_ghost.Pos.Y {
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

	blue_ghost_enemy := Enemy(blue_ghost)
	for i := range AllEnemies {
		enemy := AllEnemies[i]
		_, ok := enemy.(*BlueGhost)
		if ok {
			if enemy != blue_ghost_enemy {
				if utils.Collide(blue_ghost.Pos, blue_ghost.GetSize(), enemy.GetPosition(), enemy.GetSize()) {
					if blue_ghost.Pos.X+(blue_ghost.GetSize().X/2) > enemy.GetPosition().X+(enemy.GetSize().X/2) {
						blue_ghost.Vel.X = 1
					} else {
						blue_ghost.Vel.X = -1
					}
					if blue_ghost.Pos.Y+(blue_ghost.GetSize().Y/2) > enemy.GetPosition().Y+(enemy.GetSize().Y/2) {
						blue_ghost.Vel.Y = 1
					} else {
						blue_ghost.Vel.Y = -1
					}
				}
			}
		}
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

func (blue_ghost *BlueGhost) GetHealth() int {
	return blue_ghost.Health
}

func (blue_ghost *BlueGhost) HitPlayer() {
	*Player_Health -= 1
}

func NewBlueGhost(pos utils.Vec2) Enemy {
	blue_ghost := BlueGhost{}
	blue_ghost.Pos = pos
	blue_ghost.Health = 10

	blue_ghost.Img = textures.NewTexture("./art/enemies/blue_ghost.png", shaders.Flash_Shader)

	return &blue_ghost
}
