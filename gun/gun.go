package gun

import (
	"main/enemies"
	"main/utils"

	"github.com/bob4321at/textures"
	"github.com/hajimehoshi/ebiten/v2"
)

var Player_Pos *utils.Vec2
var Player_Vel *utils.Vec2
var Player_I_Frames float64
var Player_Damage_Multiplier *float64

var Enemies_In_Level []enemies.Enemy

type Bullet interface {
	Update()
	Draw(screen *ebiten.Image)
	Collide(position utils.Vec2, size utils.Vec2) bool
	GetDamage() int
	CheckRemoval() bool
}

type Gun interface {
	Shoot(Charged int)
	Update()
	Draw(screen *ebiten.Image)
	GetImg() textures.RenderableTexture
	GetDroppedImg() textures.RenderableTexture
	GetCharge() int
	CanShoot() bool
}

var Guns = map[int]func() Gun{
	1: CreateNerfGun,
	2: CreateBeeGun,
	3: CreateShotgun,
	4: CreateTwinMagGun,
	5: CreateMechaGun,
}
