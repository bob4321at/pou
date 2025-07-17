package worldmap

import (
	"main/utils"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var Level_To_Load *string

type LevelPointOfIntrest struct {
	LevelName string
	Pos       utils.Vec2
	Img       *ebiten.Image
}

func (point *LevelPointOfIntrest) GoTo() {
	*Current_Scene_Id = 2
	Level_To_Load = &point.LevelName
}

func (point *LevelPointOfIntrest) Draw(screen *ebiten.Image) {
	op := ebiten.DrawImageOptions{}
	op.GeoM.Translate(point.Pos.X, point.Pos.Y)

	screen.DrawImage(point.Img, &op)
}

func (point *LevelPointOfIntrest) GetPos() utils.Vec2 {
	return point.Pos
}

func NewLevelPoint(level_name string, pos utils.Vec2) *LevelPointOfIntrest {
	point := LevelPointOfIntrest{}
	point.LevelName = level_name
	point.Pos = utils.Vec2{X: pos.X - 16, Y: pos.Y - 16}

	var err error
	point.Img, _, err = ebitenutil.NewImageFromFile("./art/ui/level_point.png")
	if err != nil {
		panic(err)
	}

	return &point
}
