package worldmap

import (
	"main/utils"

	"github.com/hajimehoshi/ebiten/v2"
)

var Current_Scene_Id *int

type PointOfIntrest interface {
	GoTo()
	GetPos() utils.Vec2
	Draw(screen *ebiten.Image)
}
