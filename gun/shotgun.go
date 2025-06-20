package gun

import (
	"main/camera"
	"main/level"
	"main/music"
	"main/utils"
	"math"
	"math/rand"

	"github.com/bob4321at/textures"
	"github.com/hajimehoshi/ebiten/v2"
)

type ShotgunBullet struct {
	Img      textures.RenderableTexture
	Position utils.Vec2
	Vel      utils.Vec2
	Rotation float64
	Hit      bool
	Damage   int
	Remove   bool
}

func (bullet *ShotgunBullet) Update() {
	bullet.Position.X += bullet.Vel.X * 5
	bullet.Position.Y += bullet.Vel.Y * 5
}

func (bullet *ShotgunBullet) Draw(screen *ebiten.Image) {
	op := ebiten.DrawImageOptions{}
	op.GeoM.Translate(-4, -4)
	op.GeoM.Rotate(utils.Deg2Rad(bullet.Rotation))
	op.GeoM.Translate(bullet.Position.X+camera.Camera.Pos.X, bullet.Position.Y+camera.Camera.Pos.Y)

	bullet.Img.Draw(screen, &op)
}

func (bullet *ShotgunBullet) Collide(position utils.Vec2, size utils.Vec2) bool {
	if utils.Collide(bullet.Position, utils.Vec2{X: float64(bullet.Img.GetTexture().Bounds().Dx()), Y: float64(bullet.Img.GetTexture().Bounds().Dy())}, position, size) {
		bullet.Remove = true

		return true
	}
	return false
}

func (bullet *ShotgunBullet) GetDamage() int {
	return bullet.Damage
}

func (bullet *ShotgunBullet) CheckRemoval() bool {
	return bullet.Remove
}

func CreateShotgunBullet(pos, vel utils.Vec2, rotation float64) *ShotgunBullet {
	bullet := ShotgunBullet{}

	bullet.Damage = 3

	bullet.Position = pos
	bullet.Vel = vel
	bullet.Rotation = rotation

	bullet.Img = textures.NewTexture("./art/guns/shotgun/bullet.png", "")

	return &bullet
}

type Shotgun struct {
	Bullets []Bullet
	Rot     float64

	Cooldown float64

	Rendering_Rot float64
	Img           textures.RenderableTexture
}

func (gun *Shotgun) Shoot() {
	if gun.Cooldown < 0 {
		for _ = range 10 {
			gun.Bullets = append(gun.Bullets, CreateShotgunBullet(utils.Vec2{X: Player_Pos.X + (640 / 2), Y: Player_Pos.Y + (360 / 2)}, utils.Vec2{X: math.Cos(utils.Deg2Rad(gun.Rot + (rand.Float64() * 10))), Y: math.Sin(utils.Deg2Rad(gun.Rot + (rand.Float64() * 10)))}, gun.Rot+(rand.Float64()*10)))
		}
		gun.Cooldown = 4
		if music.AtPeak {
			Player_Vel.Y = -math.Sin(utils.Deg2Rad(gun.Rot)) * 4
			Player_Vel.X = -math.Cos(utils.Deg2Rad(gun.Rot)) * 8
		}
	}
}

func (gun *Shotgun) Update() {
	if gun.Cooldown >= 0 {
		gun.Cooldown -= 0.1
	}

	gun.Rot = utils.Rad2Deg(-utils.GetAngle(utils.Vec2{X: 640 / 2, Y: 360 / 2}, utils.Vec2{X: utils.Mouse_X, Y: utils.Mouse_Y}) - 33)

	if 640/2 < utils.Mouse_X {
		gun.Rendering_Rot = utils.Rad2Deg(-utils.GetAngle(utils.Vec2{X: 640 / 2, Y: 360 / 2}, utils.Vec2{X: utils.Mouse_X, Y: utils.Mouse_Y}) - 33)
	} else {
		gun.Rendering_Rot = utils.Rad2Deg(utils.GetAngle(utils.Vec2{X: 640 / 2, Y: 360 / 2}, utils.Vec2{X: utils.Mouse_X, Y: utils.Mouse_Y}) - 33)
	}

	for bullet_index, bullet := range gun.Bullets {
		bullet.Update()
		for i := range level.Temp_Level.Enemies {
			enemy := level.Temp_Level.Enemies[i]
			if bullet.Collide(enemy.GetPosition(), enemy.GetSize()) {
				enemy.Hit(bullet.GetDamage())
			}

		}

		if bullet.CheckRemoval() {
			utils.RemoveArrayElement(bullet_index, &gun.Bullets)
			break
		}
	}
}

func (gun *Shotgun) Draw(screen *ebiten.Image) {
	op := ebiten.DrawImageOptions{}
	op.GeoM.Translate(-24, -24)
	if 640/2 >= utils.Mouse_X {
		op.GeoM.Scale(-1, 1)
		op.GeoM.Rotate(utils.Deg2Rad(-gun.Rendering_Rot))
	} else {
		op.GeoM.Rotate(utils.Deg2Rad(gun.Rendering_Rot))
	}

	op.GeoM.Translate(640/2, 360/2)

	gun.Img.Draw(screen, &op)

	for i := range gun.Bullets {
		gun.Bullets[i].Draw(screen)
	}
}

func (gun *Shotgun) GetImg() textures.RenderableTexture {
	return gun.Img
}

func CreateShotgun() *Shotgun {
	gun := Shotgun{}

	gun.Img = textures.NewTexture("./art/guns/shotgun/gun.png", "")

	return &gun
}
