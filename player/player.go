package player

import (
	"main/camera"
	"main/gun"
	"main/level"
	"main/utils"

	"github.com/bob4321at/textures"

	"github.com/hajimehoshi/ebiten/v2"
)

type PlayerStruct struct {
	Pos utils.Vec2
	Vel utils.Vec2
	Dir bool
	img textures.RenderableTexture

	Gun gun.Gun
}

func (player *PlayerStruct) Update() {
	gun.Player_Pos = &player.Pos
	gun.Player_Vel = &player.Vel

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
	} else {
		player.Dir = true
	}

	collision_x, _, _ := level.Temp_Level.CheckCollision(utils.Vec2{X: player.Pos.X + player.Vel.X + 640/2 - 16, Y: player.Pos.Y + 360/2 - 24}, utils.Vec2{X: 32, Y: 48})
	if collision_x {
		player.Vel.X = 0
	}

	collision_y, _, _ := level.Temp_Level.CheckCollision(utils.Vec2{X: player.Pos.X + 640/2 - 16, Y: player.Pos.Y + player.Vel.Y + 360/2 - 24}, utils.Vec2{X: 32, Y: 48})
	if collision_y {
		if player.Vel.Y > 0 {
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
	}

	camera.Camera.Pos.X = -player.Pos.X
	camera.Camera.Pos.Y = -player.Pos.Y

	player.Pos.X += player.Vel.X
	player.Pos.Y += player.Vel.Y

	player.Gun.Update()

	if ebiten.IsMouseButtonPressed(ebiten.MouseButton0) {
		player.Gun.Shoot()
	}
}

func (player *PlayerStruct) Draw(screen *ebiten.Image) {
	op := ebiten.DrawImageOptions{}

	if player.Dir {
		op.GeoM.Translate(-16, 0)
		op.GeoM.Scale(-1, 1)
		op.GeoM.Translate(16, 0)
	}

	op.GeoM.Translate(640/2-16, 360/2-24)

	player.img.Draw(screen, &op)

	player.Gun.Draw(screen)
}

func NewPlayer(pos utils.Vec2) (player PlayerStruct) {
	player.Pos = pos

	player.img = textures.NewTexture("./art/player.png", "")
	player.Gun = gun.CreateTwinMagGun()

	return player
}

var Player = NewPlayer(utils.Vec2{X: 100, Y: 100})
