package level

import (
	"main/utils"

	"github.com/hajimehoshi/ebiten/v2"
)

type SpringTile struct {
	Pos       utils.Vec2
	Power     float64
	Img       *ebiten.Image
	Direction int
}
