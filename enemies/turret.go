package enemies

import (
	"main/camera"
	"main/shaders"
	"main/utils"
	"math"

	"github.com/bob4321at/textures"
	"github.com/hajimehoshi/ebiten/v2"
)

type Turret struct {
	Pos   utils.Vec2
	Angle float64

	Img       textures.RenderableTexture
	Health    int
	I_Frames  float64
	Timer     float64
	Max_Timer float64
}

func (turret *Turret) Draw(screen *ebiten.Image) {
	op := ebiten.DrawImageOptions{}
	op.GeoM.Translate(-float64(turret.Img.GetTexture().Bounds().Dx())/2, -float64(turret.Img.GetTexture().Bounds().Dy())/2)
	op.GeoM.Rotate(utils.Deg2Rad(turret.Angle))
	op.GeoM.Translate(float64(turret.Img.GetTexture().Bounds().Dx())/2, float64(turret.Img.GetTexture().Bounds().Dy())/2)
	op.GeoM.Translate(turret.Pos.X+camera.Camera.Pos.X, turret.Pos.Y+camera.Camera.Pos.Y)

	turret.Img.SetUniforms(map[string]any{
		"I": turret.I_Frames,
	})

	turret.Img.Draw(screen, &op)
}

func (turret *Turret) Update() {
	if turret.I_Frames > 0 {
		turret.I_Frames -= 0.1
	} else {
		turret.I_Frames = 0
	}

	if turret.Timer > 0 {
		turret.Timer -= 0.1
	} else {
		angle := utils.GetAngle(utils.Vec2{X: Player_Pos.X + 320, Y: Player_Pos.Y + 180}, utils.Vec2{X: turret.Pos.X + 12, Y: turret.Pos.Y + 12})
		SpawnProjectile(NewTurretBullet(utils.Vec2{X: turret.Pos.X + 12, Y: turret.Pos.Y + 12}, utils.Vec2{X: math.Cos(angle-utils.Deg2Rad(90)) * 10, Y: -math.Sin(angle-utils.Deg2Rad(90)) * 10}))
		turret.Timer = turret.Max_Timer
	}
}

func (turret *Turret) Hit(damage int) {
	if turret.I_Frames <= 0 {
		turret.Health -= damage
		turret.I_Frames = 2
	}
}

func (turret *Turret) GetPosition() utils.Vec2 {
	return turret.Pos
}

func (turret *Turret) GetSize() utils.Vec2 {
	return utils.Vec2{X: float64(turret.Img.GetTexture().Bounds().Dx()), Y: float64(turret.Img.GetTexture().Bounds().Dy())}
}

func (turret *Turret) GetHealth() int {
	return turret.Health
}

func (turret *Turret) HitPlayer() {}

func (turret *Turret) GetWallDist(dir utils.Vec2) (float64, utils.Vec2) {
	closeset_hit := 100000000000.0
	closeset_hit_pos := utils.Vec2{X: 0, Y: 0}

	for _, hitbox := range Level_Hitbox {
		where_hit, hit := utils.Raycast(turret.Pos, dir, 10, hitbox, utils.Vec2{X: 32, Y: 32})
		if hit {
			if utils.GetDist(where_hit, turret.Pos) < closeset_hit {
				closeset_hit = utils.GetDist(where_hit, turret.Pos)
				closeset_hit_pos = where_hit
			}
		}
	}

	return closeset_hit, closeset_hit_pos
}

func NewTurret(pos utils.Vec2) Enemy {
	turret := Turret{}
	turret.Pos = pos
	turret.Health = 30

	go func() {
		wall_right_dist, wall_right_pos := turret.GetWallDist(utils.Vec2{X: 16, Y: 0})
		wall_right_pos = utils.Vec2{X: wall_right_pos.X - 32, Y: wall_right_pos.Y}
		wall_left_dist, wall_left_pos := turret.GetWallDist(utils.Vec2{X: -16, Y: 0})
		wall_left_pos = utils.Vec2{X: wall_left_pos.X + 16, Y: wall_left_pos.Y}
		wall_down_dist, wall_down_pos := turret.GetWallDist(utils.Vec2{X: 0, Y: 16})
		wall_down_pos = utils.Vec2{X: wall_down_pos.X, Y: wall_down_pos.Y - 32}
		wall_up_dist, wall_up_pos := turret.GetWallDist(utils.Vec2{X: 0, Y: -16})
		wall_up_pos = utils.Vec2{X: wall_up_pos.X, Y: wall_up_pos.Y + 16}

		sorted := utils.BubbleSort([]*float64{&wall_right_dist, &wall_left_dist, &wall_up_dist, &wall_down_dist})

		if sorted[0] == &wall_right_dist {
			turret.Pos = wall_right_pos

			turret.Angle = 270
		} else if sorted[0] == &wall_left_dist {
			turret.Pos = wall_left_pos

			turret.Angle = 90
		} else if sorted[0] == &wall_up_dist {
			turret.Pos = wall_up_pos

			turret.Angle = 180
		} else if sorted[0] == &wall_down_dist {
			turret.Pos = wall_down_pos

			turret.Angle = 0
		}
	}()
	turret.Img = textures.NewTexture("./art/enemies/turret.png", shaders.Flash_Shader)

	turret.Timer = 5
	turret.Max_Timer = turret.Timer

	return &turret
}
