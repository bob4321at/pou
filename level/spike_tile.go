package level

import (
	"main/utils"

	"github.com/hajimehoshi/ebiten/v2"
)

type SpikeTile struct {
	Pos utils.Vec2
	Img *ebiten.Image
}
