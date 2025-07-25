package level

import (
	"main/utils"

	"github.com/hajimehoshi/ebiten/v2"
)

type SpringTile struct {
	Pos       utils.Vec2
	Power     int
	Img       *ebiten.Image
	Direction int
}
