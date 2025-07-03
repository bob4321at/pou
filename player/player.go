package player

import (
	"image/color"
	"main/camera"
	"main/enemies"
	"main/gun"
	"main/level"
	"main/shaders"
	"main/utils"
	"math"

	"github.com/bob4321at/textures"

	"github.com/hajimehoshi/ebiten/v2"
)

type PlayerStruct struct {
	Pos utils.Vec2
	Vel utils.Vec2
	Dir bool
	Img *textures.AnimatedTexture

	Gun gun.Gun

	Health          int
	Previous_Health int
	Health_Bar_Img  *ebiten.Image
	I_Frames        float64
}

func (player *PlayerStruct) Update() {
	gun.Player_Pos = &player.Pos
	gun.Player_Vel = &player.Vel
	enemies.Player_Pos = &player.Pos
	enemies.Player_Health = &player.Health

	if player.I_Frames > 0 {
		player.I_Frames -= 0.1
	} else {
		player.I_Frames = 0
		for ei := range level.Temp_Level.Enemies {
			enemy := level.Temp_Level.Enemies[ei]
			if utils.Collide(utils.Vec2{X: player.Pos.X + player.Vel.X + 640/2 - 14, Y: player.Pos.Y + 360/2 - 18}, utils.Vec2{X: 32, Y: 48}, enemy.GetPosition(), enemy.GetSize()) {
				enemy.HitPlayer()
			}
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeyT) {
		player.Health -= 1
	}

	if player.Health < player.Previous_Health {
		player.I_Frames = 1
	}

	player.Img.SetUniforms(map[string]any{
		"I": player.I_Frames,
	})

	player.Img.Current_Animation = 0

	if !level.Temp_Level.Player_Loaded {
		player.Vel.X = 0
		player.Vel.Y = 0
		player.Pos = level.Temp_Level.Player_Spawn
		level.Temp_Level.Player_Loaded = true
	}

	player.Vel.Y += 0.1

	if ebiten.IsKeyPressed(ebiten.KeyA) && player.Vel.X-0.1 > -4 {
		player.Vel.X -= 0.1
		if player.Vel.X > 0 {
			player.Vel.X -= 0.2
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) && player.Vel.X+0.1 < 4 {
		player.Vel.X += 0.1
		if player.Vel.X < 0 {
			player.Vel.X += 0.2
		}
	}

	if player.Vel.X > 0 {
		player.Dir = false
	} else if player.Vel.X < 0 {
		player.Dir = true
	}

	collision_x, _ := level.Temp_Level.CheckCollision(utils.Vec2{X: player.Pos.X + player.Vel.X + 640/2 - 14, Y: player.Pos.Y + 360/2 - 18}, utils.Vec2{X: 28, Y: 42})
	if collision_x {
		player.Vel.X = 0
	}

	collision_y, _ := level.Temp_Level.CheckCollision(utils.Vec2{X: player.Pos.X + 640/2 - 14, Y: player.Pos.Y + player.Vel.Y + 360/2 - 18}, utils.Vec2{X: 28, Y: 42})
	if collision_y {
		if player.Vel.Y > 0 {
			if math.Abs(player.Vel.X) >= 0.1 {
				player.Img.Current_Animation = 1
			} else {
				player.Img.Current_Animation = 0
			}

			player.Vel.Y = 0
			if ebiten.IsKeyPressed(ebiten.KeyW) || ebiten.IsKeyPressed(ebiten.KeySpace) {
				player.Vel.Y = -4
			}
			if !ebiten.IsKeyPressed(ebiten.KeyA) && !ebiten.IsKeyPressed(ebiten.KeyD) {
				player.Vel.X -= player.Vel.X / 5
			}
		} else {
			player.Vel.Y = 0
		}
	} else {
		if player.Vel.Y > 0 {
			player.Img.Current_Animation = 2
		} else {
			player.Img.Current_Animation = 3
		}
	}

	camera.Camera.Pos.X = -player.Pos.X
	camera.Camera.Pos.Y = -player.Pos.Y

	player.Pos.X += player.Vel.X
	player.Pos.Y += player.Vel.Y

	player.Gun.Update()

	if ebiten.IsMouseButtonPressed(ebiten.MouseButton0) {
		player.Gun.Shoot()
	}

	player.Previous_Health = player.Health
}

func (player *PlayerStruct) Draw(screen *ebiten.Image) {
	player.Img.Update()

	op := ebiten.DrawImageOptions{}

	if player.Dir {
		op.GeoM.Translate(-16, 0)
		op.GeoM.Scale(-1, 1)
		op.GeoM.Translate(16, 0)
	}

	op.GeoM.Translate(640/2-16, 360/2-24)

	player.Img.Draw(screen, &op)

	player.Gun.Draw(screen)
}

func NewPlayer(pos utils.Vec2) (player PlayerStruct) {
	player.Pos = pos

	player.Img = textures.NewAnimatedTexture("./art/player.png", shaders.Flash_Shader)
	player.Gun = gun.CreateMechaGun()

	player.Health = 100
	player.Health_Bar_Img = ebiten.NewImage(250, 24)
	player.Health_Bar_Img.Fill(color.RGBA{255, 50, 50, 255})

	return player
}
