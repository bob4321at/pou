package level

import (
	"main/utils"

	"github.com/hajimehoshi/ebiten/v2"
)

type SpikeTile struct {
	Pos    utils.Vec2
	Damage int
	Img    *ebiten.Image
}
