package gun

import (
	"main/camera"
	"main/level"
	"main/shaders"
	"main/utils"
	"math"
	"time"

	"github.com/bob4321at/textures"
	"github.com/hajimehoshi/ebiten/v2"
)

type NerfBullet struct {
	Img      textures.RenderableTexture
	Position utils.Vec2
	Vel      utils.Vec2
	Rotation float64
	Hit      bool
	Damage   int
	Remove   bool
}

func (bullet *NerfBullet) Update(current_level *level.Level) {
	bullet.Vel.X = math.Cos(utils.Deg2Rad(bullet.Rotation))
	bullet.Vel.Y = math.Sin(utils.Deg2Rad(bullet.Rotation))

	bullet.Position.X += bullet.Vel.X * 5
	bullet.Position.Y += bullet.Vel.Y * 5

	if (bullet.Rotation+180)+1710 > 0 {
		if bullet.Rotation+180+1710 < 150 {
			bullet.Rotation += 1
		}
	} else if (bullet.Rotation+180)+1710 < 0 {
		if bullet.Rotation+180+1710 > -150 {
			bullet.Rotation -= 1
		}
	}
}

func (bullet *NerfBullet) Draw(screen *ebiten.Image) {
	op := ebiten.DrawImageOptions{}
	op.GeoM.Translate(-4, -4)
	op.GeoM.Rotate(utils.Deg2Rad(bullet.Rotation))
	op.GeoM.Translate(bullet.Position.X+camera.Camera.Pos.X, bullet.Position.Y+camera.Camera.Pos.Y)

	bullet.Img.Draw(screen, &op)
}

func (bullet *NerfBullet) Collide(position utils.Vec2, size utils.Vec2) bool {
	if utils.Collide(bullet.Position, utils.Vec2{X: float64(bullet.Img.GetTexture().Bounds().Dx()), Y: float64(bullet.Img.GetTexture().Bounds().Dy())}, position, size) {
		bullet.Remove = true

		return true
	}
	return false
}

func (bullet *NerfBullet) GetDamage() int {
	return bullet.Damage
}

func (bullet *NerfBullet) CheckRemoval() bool {
	return bullet.Remove
}

func CreateNerfBullet(pos, vel utils.Vec2, rotation float64) *NerfBullet {
	bullet := NerfBullet{}

	bullet.Damage = 3

	bullet.Position = pos
	bullet.Vel = vel
	bullet.Rotation = rotation

	bullet.Img = textures.NewTexture("./art/guns/nerfgun/bullet.png", "")

	return &bullet
}

type NerfGun struct {
	Bullets []Bullet
	Rot     float64

	Cooldown float64
	Charge   int

	Rendering_Rot float64
	Img           textures.RenderableTexture
}

func (gun *NerfGun) Shoot(Charged int) {
	if gun.Cooldown < 0 {
		if Charged > gun.Charge {
			go func() {
				gun.Bullets = append(gun.Bullets, CreateNerfBullet(utils.Vec2{X: Player_Pos.X + (640 / 2), Y: Player_Pos.Y + (360 / 2)}, utils.Vec2{X: math.Cos(utils.Deg2Rad(gun.Rot)), Y: math.Sin(utils.Deg2Rad(gun.Rot))}, gun.Rot))
				time.Sleep(time.Second / 6)
				gun.Bullets = append(gun.Bullets, CreateNerfBullet(utils.Vec2{X: Player_Pos.X + (640 / 2), Y: Player_Pos.Y + (360 / 2)}, utils.Vec2{X: math.Cos(utils.Deg2Rad(gun.Rot)), Y: math.Sin(utils.Deg2Rad(gun.Rot))}, gun.Rot))
				time.Sleep(time.Second / 6)
				gun.Bullets = append(gun.Bullets, CreateNerfBullet(utils.Vec2{X: Player_Pos.X + (640 / 2), Y: Player_Pos.Y + (360 / 2)}, utils.Vec2{X: math.Cos(utils.Deg2Rad(gun.Rot)), Y: math.Sin(utils.Deg2Rad(gun.Rot))}, gun.Rot))
			}()
		} else {
			gun.Bullets = append(gun.Bullets, CreateNerfBullet(utils.Vec2{X: Player_Pos.X + (640 / 2), Y: Player_Pos.Y + (360 / 2)}, utils.Vec2{X: math.Cos(utils.Deg2Rad(gun.Rot)), Y: math.Sin(utils.Deg2Rad(gun.Rot))}, gun.Rot))
		}
		gun.Cooldown = 4
	}
}

func (gun *NerfGun) Update(current_level *level.Level) {
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
		bullet.Update(current_level)
		for i := range current_level.Enemies {
			enemy := current_level.Enemies[i]
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

func (gun *NerfGun) Draw(screen *ebiten.Image) {
	op := ebiten.DrawImageOptions{}
	op.GeoM.Translate(-24, -24)
	if 640/2 >= utils.Mouse_X {
		op.GeoM.Scale(-1, 1)
		op.GeoM.Rotate(utils.Deg2Rad(-gun.Rendering_Rot))
	} else {
		op.GeoM.Rotate(utils.Deg2Rad(gun.Rendering_Rot))
	}

	op.GeoM.Translate(640/2, 360/2)

	gun.Img.SetUniforms(map[string]any{
		"I": Player_I_Frames,
	})

	gun.Img.Draw(screen, &op)

	for i := range gun.Bullets {
		gun.Bullets[i].Draw(screen)
	}
}

func (gun *NerfGun) GetImg() textures.RenderableTexture {
	return gun.Img
}

func (gun *NerfGun) GetCharge() int {
	return gun.Charge
}

func (gun *NerfGun) CanShoot() bool {
	return gun.Cooldown <= 0
}

func CreateNerfGun() *NerfGun {
	gun := NerfGun{}

	gun.Img = textures.NewTexture("./art/guns/nerfgun/gun.png", shaders.Flash_Shader)
	gun.Charge = 20

	return &gun
}
