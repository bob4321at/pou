package gun

import (
	"main/camera"
	"main/enemies"
	"main/level"
	"main/music"
	"main/utils"
	"math"

	"github.com/bob4321at/textures"
	"github.com/hajimehoshi/ebiten/v2"
)

type BeeBullet struct {
	Img      textures.RenderableTexture
	Position utils.Vec2
	Vel      utils.Vec2
	Rotation float64
	Hit      bool
	Damage   int
	Remove   bool
	Lifttime float64
}

func (bullet *BeeBullet) Update() {
	bullet.Img.Update()

	bullet.Vel.X = math.Cos(utils.Deg2Rad(bullet.Rotation))
	bullet.Vel.Y = math.Sin(utils.Deg2Rad(bullet.Rotation))

	if len(level.Temp_Level.Enemies) != 0 {
		var closest_enemy enemies.Enemy
		closest_enemy_distance := -1.0

		closest_enemy = level.Temp_Level.Enemies[0]
		closest_enemy_distance = math.Abs(utils.GetDist(bullet.Position, closest_enemy.GetPosition()))

		for ei, enemy := range level.Temp_Level.Enemies {
			if math.Abs(utils.GetDist(bullet.Position, enemy.GetPosition())) < closest_enemy_distance {
				closest_enemy = level.Temp_Level.Enemies[ei]
				closest_enemy_distance = math.Abs(utils.GetDist(bullet.Position, enemy.GetPosition()))
			}
		}

		if math.Atan2(bullet.Vel.X, bullet.Vel.Y) > utils.GetAngle(closest_enemy.GetPosition(), bullet.Position) {
			bullet.Rotation += 2
		} else {
			bullet.Rotation -= 2
		}
	} else {
		if math.Atan2(bullet.Vel.X, bullet.Vel.Y) > utils.GetAngle(utils.Vec2{X: Player_Pos.X + 320, Y: Player_Pos.Y + 20}, bullet.Position) {
			bullet.Rotation += 2
		} else {
			bullet.Rotation -= 2
		}
	}

	bullet.Lifttime -= 0.1
	if bullet.Lifttime < 0 {
		bullet.Remove = true
	}

	bullet.Position.X += bullet.Vel.X * 1
	bullet.Position.Y += bullet.Vel.Y * 1
}

func (bullet *BeeBullet) Draw(screen *ebiten.Image) {
	op := ebiten.DrawImageOptions{}
	op.GeoM.Translate(-4, -4)
	op.GeoM.Rotate(utils.Deg2Rad(bullet.Rotation))
	op.GeoM.Translate(bullet.Position.X+camera.Camera.Pos.X, bullet.Position.Y+camera.Camera.Pos.Y)

	bullet.Img.Draw(screen, &op)
}

func (bullet *BeeBullet) Collide(position utils.Vec2, size utils.Vec2) bool {
	if utils.Collide(bullet.Position, utils.Vec2{X: float64(bullet.Img.GetTexture().Bounds().Dx()), Y: float64(bullet.Img.GetTexture().Bounds().Dy())}, position, size) {
		bullet.Remove = true

		return true
	}
	return false
}

func (bullet *BeeBullet) GetDamage() int {
	return bullet.Damage
}

func (bullet *BeeBullet) CheckRemoval() bool {
	return bullet.Remove
}

func CreateBeeBullet(pos, vel utils.Vec2, rotation float64) *BeeBullet {
	bullet := BeeBullet{}

	bullet.Damage = 3
	bullet.Lifttime = 100

	bullet.Position = pos
	bullet.Vel = vel
	bullet.Rotation = rotation

	bullet.Img = textures.NewAnimatedTexture("./art/guns/beegun/bullet.png", "")

	return &bullet
}

type BeeGun struct {
	Bullets []Bullet
	Rot     float64

	Cooldown float64

	Rendering_Rot float64
	Img           textures.RenderableTexture
}

func (gun *BeeGun) Shoot() {
	if gun.Cooldown < 0 {
		if music.AtPeak {
			gun.Bullets = append(gun.Bullets, CreateBeeBullet(utils.Vec2{X: Player_Pos.X + (640 / 2), Y: Player_Pos.Y + (360 / 2)}, utils.Vec2{X: math.Cos(utils.Deg2Rad(gun.Rot)), Y: math.Sin(utils.Deg2Rad(gun.Rot))}, gun.Rot))
			gun.Bullets = append(gun.Bullets, CreateBeeBullet(utils.Vec2{X: Player_Pos.X + (640 / 2), Y: Player_Pos.Y + (360 / 2)}, utils.Vec2{X: math.Cos(utils.Deg2Rad(gun.Rot + 10)), Y: math.Sin(utils.Deg2Rad(gun.Rot + 10))}, gun.Rot+10))
			gun.Bullets = append(gun.Bullets, CreateBeeBullet(utils.Vec2{X: Player_Pos.X + (640 / 2), Y: Player_Pos.Y + (360 / 2)}, utils.Vec2{X: math.Cos(utils.Deg2Rad(gun.Rot - 10)), Y: math.Sin(utils.Deg2Rad(gun.Rot - 10))}, gun.Rot-10))
			gun.Cooldown = 5
		} else {
			gun.Bullets = append(gun.Bullets, CreateBeeBullet(utils.Vec2{X: Player_Pos.X + (640 / 2), Y: Player_Pos.Y + (360 / 2)}, utils.Vec2{X: math.Cos(utils.Deg2Rad(gun.Rot)), Y: math.Sin(utils.Deg2Rad(gun.Rot))}, gun.Rot))
			gun.Cooldown = 5
		}
	}
}

func (gun *BeeGun) Update() {
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
			if bullet_index < len(gun.Bullets) {
				utils.RemoveArrayElement(bullet_index, &gun.Bullets)
			}
		}
	}
}

func (gun *BeeGun) Draw(screen *ebiten.Image) {
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

func (gun *BeeGun) GetImg() textures.RenderableTexture {
	return gun.Img
}

func CreateBeeGun() *BeeGun {
	gun := BeeGun{}

	gun.Img = textures.NewTexture("./art/guns/beegun/gun.png", "")

	return &gun
}
