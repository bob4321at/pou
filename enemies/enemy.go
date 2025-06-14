package enemies

import (
	"main/utils"

	"github.com/hajimehoshi/ebiten/v2"
)

type Enemy interface {
	Draw(screen *ebiten.Image)
	Update()
	Hit(damage int)
	GetPosition() utils.Vec2
	GetSize() utils.Vec2
	GetHealth() int
}
