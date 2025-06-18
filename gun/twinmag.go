package gun

import (
	"main/camera"
	"main/level"
	"main/music"
	"main/utils"
	"math"

	"github.com/bob4321at/textures"
	"github.com/hajimehoshi/ebiten/v2"
)

type TwinMagBullet struct {
	Img      textures.RenderableTexture
	Position utils.Vec2
	Vel      utils.Vec2
	Rotation float64
	Hit      bool
	Damage   int
	Remove   bool
}

func (bullet *TwinMagBullet) Update() {
	bullet.Position.X += bullet.Vel.X * 5
	bullet.Position.Y += bullet.Vel.Y * 5
}

func (bullet *TwinMagBullet) Draw(screen *ebiten.Image) {
	op := ebiten.DrawImageOptions{}
	op.GeoM.Translate(-4, -4)
	op.GeoM.Rotate(utils.Deg2Rad(bullet.Rotation))
	op.GeoM.Translate(bullet.Position.X+camera.Camera.Pos.X, bullet.Position.Y+camera.Camera.Pos.Y)

	bullet.Img.Draw(screen, &op)
}

func (bullet *TwinMagBullet) Collide(position utils.Vec2, size utils.Vec2) bool {
	if utils.Collide(bullet.Position, utils.Vec2{X: float64(bullet.Img.GetTexture().Bounds().Dx()), Y: float64(bullet.Img.GetTexture().Bounds().Dy())}, position, size) {
		bullet.Remove = true

		return true
	}
	return false
}

func (bullet *TwinMagBullet) GetDamage() int {
	return bullet.Damage
}

func (bullet *TwinMagBullet) CheckRemoval() bool {
	return bullet.Remove
}

func CreateTwinMagBullet(pos, vel utils.Vec2, rotation float64) *TwinMagBullet {
	bullet := TwinMagBullet{}

	bullet.Damage = 1

	bullet.Position = pos
	bullet.Vel = vel
	bullet.Rotation = rotation

	bullet.Img = textures.NewTexture("./art/guns/twinmag/bullet.png", "")

	return &bullet
}

type BigTwinMagBullet struct {
	Img      textures.RenderableTexture
	Position utils.Vec2
	Vel      utils.Vec2
	Rotation float64
	Hit      bool
	Damage   int
	Remove   bool
}

func (bullet *BigTwinMagBullet) Update() {
	bullet.Position.X += bullet.Vel.X * 5
	bullet.Position.Y += bullet.Vel.Y * 5
}

func (bullet *BigTwinMagBullet) Draw(screen *ebiten.Image) {
	op := ebiten.DrawImageOptions{}
	op.GeoM.Translate(-4, -4)
	op.GeoM.Rotate(utils.Deg2Rad(bullet.Rotation))
	op.GeoM.Translate(bullet.Position.X+camera.Camera.Pos.X, bullet.Position.Y+camera.Camera.Pos.Y)

	bullet.Img.Draw(screen, &op)
}

func (bullet *BigTwinMagBullet) Collide(position utils.Vec2, size utils.Vec2) bool {
	if utils.Collide(bullet.Position, utils.Vec2{X: float64(bullet.Img.GetTexture().Bounds().Dx()), Y: float64(bullet.Img.GetTexture().Bounds().Dy())}, position, size) {
		bullet.Remove = true

		return true
	}
	return false
}

func (bullet *BigTwinMagBullet) GetDamage() int {
	return bullet.Damage
}

func (bullet *BigTwinMagBullet) CheckRemoval() bool {
	return bullet.Remove
}

func CreateBigTwinMagBullet(pos, vel utils.Vec2, rotation float64) *BigTwinMagBullet {
	bullet := BigTwinMagBullet{}

	bullet.Damage = 5

	bullet.Position = pos
	bullet.Vel = vel
	bullet.Rotation = rotation

	bullet.Img = textures.NewTexture("./art/guns/twinmag/big_bullet.png", "")

	return &bullet
}

type TwinMagGun struct {
	Bullets []Bullet
	Rot     float64

	Cooldown float64

	Rendering_Rot float64
	Img           textures.RenderableTexture
}

func (gun *TwinMagGun) Shoot() {
	if gun.Cooldown < 0 {
		if music.AtPeak {
			gun.Bullets = append(gun.Bullets, CreateBigTwinMagBullet(utils.Vec2{X: Player_Pos.X + (640 / 2), Y: Player_Pos.Y + (360 / 2)}, utils.Vec2{X: math.Cos(utils.Deg2Rad(gun.Rot)), Y: math.Sin(utils.Deg2Rad(gun.Rot))}, gun.Rot))
			gun.Cooldown = 3
		} else {
			gun.Bullets = append(gun.Bullets, CreateTwinMagBullet(utils.Vec2{X: Player_Pos.X + (640 / 2), Y: Player_Pos.Y + (360 / 2)}, utils.Vec2{X: math.Cos(utils.Deg2Rad(gun.Rot + 5)), Y: math.Sin(utils.Deg2Rad(gun.Rot + 5))}, gun.Rot+5))
			gun.Bullets = append(gun.Bullets, CreateTwinMagBullet(utils.Vec2{X: Player_Pos.X + (640 / 2), Y: Player_Pos.Y + (360 / 2)}, utils.Vec2{X: math.Cos(utils.Deg2Rad(gun.Rot - 5)), Y: math.Sin(utils.Deg2Rad(gun.Rot - 5))}, gun.Rot-5))
			gun.Cooldown = 3
		}
	}
}

func (gun *TwinMagGun) Update() {
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
		}
	}
}

func (gun *TwinMagGun) Draw(screen *ebiten.Image) {
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

func (gun *TwinMagGun) GetImg() textures.RenderableTexture {
	return gun.Img
}

func CreateTwinMagGun() *TwinMagGun {
	gun := TwinMagGun{}

	gun.Img = textures.NewTexture("./art/guns/twinmag/gun.png", "")

	return &gun
}
