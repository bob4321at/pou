package player

import (
	"main/camera"
	"main/level"
	"main/music"
	"main/utils"
	"math"

	"github.com/bob4321at/textures"

	"github.com/hajimehoshi/ebiten/v2"
)

type Bullet struct {
	Pos    utils.Vec2
	Vel    utils.Vec2
	Img    textures.RenderableTexture
	Damage int
	Hit    bool
}

func NewBullet(Pos, Vel utils.Vec2, damage int) (bullet Bullet) {
	bullet.Pos = Pos
	bullet.Vel = Vel
	bullet.Damage = damage

	bullet.Img = textures.NewTexture("./art/bullet.png", "")

	return bullet
}

func (bullet *Bullet) Update() {
	bullet.Pos.X += bullet.Vel.X * 10
	bullet.Pos.Y += bullet.Vel.Y * 10

	for i := range level.Temp_Level.Enemies {
		enemy := level.Temp_Level.Enemies[i]
		if utils.Collide(bullet.Pos, utils.Vec2{X: 4, Y: 4}, enemy.GetPosition(), enemy.GetSize()) {
			enemy.Hit(bullet.Damage)
			bullet.Hit = true
		}
	}
}

func (bullet *Bullet) Draw(screen *ebiten.Image) {
	op := ebiten.DrawImageOptions{}
	op.GeoM.Translate(bullet.Pos.X+camera.Camera.Pos.X, bullet.Pos.Y+camera.Camera.Pos.Y)

	bullet.Img.Draw(screen, &op)
}

type PlayerStruct struct {
	Pos utils.Vec2
	Vel utils.Vec2
	Dir bool
	img textures.RenderableTexture

	Gun_Rot      float64
	Real_Gun_Rot float64
	Gun_Img      textures.RenderableTexture

	Bullets  []Bullet
	Cooldown float64
}

func (player *PlayerStruct) Update() {
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
		player.Vel.Y = 0
		if ebiten.IsKeyPressed(ebiten.KeyW) || ebiten.IsKeyPressed(ebiten.KeySpace) {
			player.Vel.Y = -4
		}
		if !ebiten.IsKeyPressed(ebiten.KeyA) && !ebiten.IsKeyPressed(ebiten.KeyD) {
			player.Vel.X -= player.Vel.X / 5
		}
	}

	camera.Camera.Pos.X = -player.Pos.X
	camera.Camera.Pos.Y = -player.Pos.Y

	player.Pos.X += player.Vel.X
	player.Pos.Y += player.Vel.Y

	if 640/2 < utils.Mouse_X {
		player.Real_Gun_Rot = utils.Rad2Deg(-utils.GetAngle(utils.Vec2{X: 640 / 2, Y: 360 / 2}, utils.Vec2{X: utils.Mouse_X, Y: utils.Mouse_Y}) - 33)
		player.Gun_Rot = utils.Rad2Deg(-utils.GetAngle(utils.Vec2{X: 640 / 2, Y: 360 / 2}, utils.Vec2{X: utils.Mouse_X, Y: utils.Mouse_Y}) - 33)
	} else {
		player.Real_Gun_Rot = utils.Rad2Deg(-utils.GetAngle(utils.Vec2{X: 640 / 2, Y: 360 / 2}, utils.Vec2{X: utils.Mouse_X, Y: utils.Mouse_Y}) - 33)
		player.Gun_Rot = utils.Rad2Deg(utils.GetAngle(utils.Vec2{X: 640 / 2, Y: 360 / 2}, utils.Vec2{X: utils.Mouse_X, Y: utils.Mouse_Y}) - 33)
	}

	if ebiten.IsMouseButtonPressed(ebiten.MouseButton0) && player.Cooldown < 1 {
		if music.AtPeak {
			player.Bullets = append(player.Bullets, NewBullet(utils.Vec2{X: player.Pos.X + (640 / 2), Y: player.Pos.Y + (360 / 2)}, utils.Vec2{X: math.Cos(utils.Deg2Rad(player.Real_Gun_Rot)) * 2, Y: math.Sin(utils.Deg2Rad(player.Real_Gun_Rot)) * 2}, 5))
			player.Cooldown = 2

			player.Vel.X -= math.Cos(utils.Deg2Rad(player.Real_Gun_Rot)) * 10
			player.Vel.Y -= math.Sin(utils.Deg2Rad(player.Real_Gun_Rot)) * 5
		} else {
			player.Bullets = append(player.Bullets, NewBullet(utils.Vec2{X: player.Pos.X + (640 / 2), Y: player.Pos.Y + (360 / 2)}, utils.Vec2{X: math.Cos(utils.Deg2Rad(player.Real_Gun_Rot)), Y: math.Sin(utils.Deg2Rad(player.Real_Gun_Rot))}, 1))
			player.Cooldown = 7
		}
	}
	player.Cooldown -= 0.1

	for b := range player.Bullets {
		if b < len(player.Bullets) {
			player.Bullets[b].Update()
			if player.Bullets[b].Hit {
				utils.RemoveArrayElement(b, &player.Bullets)
			}
		}
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

	op.GeoM.Reset()

	op.GeoM.Translate(-24, -24)
	if 640/2 >= utils.Mouse_X {
		op.GeoM.Scale(-1, 1)
		op.GeoM.Rotate(utils.Deg2Rad(-player.Gun_Rot))
	} else {
		op.GeoM.Rotate(utils.Deg2Rad(player.Gun_Rot))
	}

	op.GeoM.Translate(640/2, 360/2)

	player.Gun_Img.Draw(screen, &op)

	for b := range player.Bullets {
		player.Bullets[b].Draw(screen)
	}
}

func NewPlayer(pos utils.Vec2) (player PlayerStruct) {
	player.Pos = pos

	player.img = textures.NewTexture("./art/player.png", "")
	player.Gun_Img = textures.NewTexture("./art/gun.png", "")

	return player
}

var Player = NewPlayer(utils.Vec2{X: 100, Y: 100})
