package gun

import (
	"main/level"
	"main/utils"

	"github.com/bob4321at/textures"
	"github.com/hajimehoshi/ebiten/v2"
)

var Player_Pos *utils.Vec2
var Player_Vel *utils.Vec2
var Player_I_Frames float64

type Bullet interface {
	Update(level *level.Level)
	Draw(screen *ebiten.Image)
	Collide(position utils.Vec2, size utils.Vec2) bool
	GetDamage() int
	CheckRemoval() bool
}

type Gun interface {
	Shoot(Charged int)
	Update(current_level *level.Level)
	Draw(screen *ebiten.Image)
	GetImg() textures.RenderableTexture
	GetCharge() int
	CanShoot() bool
}
