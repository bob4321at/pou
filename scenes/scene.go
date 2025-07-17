package scenes

import (
	worldmap "main/world_map"

	"github.com/hajimehoshi/ebiten/v2"
)

type Scene interface {
	Setup()
	Update()
	Draw(display *ebiten.Image)
	GetSetup() bool
}

var Scenes = []Scene{&Main_Menu_Scene{}, NewWorldMapScene(), &GameScene{}, &Death_Menu_scene{}, &Win_Menu_Scene{}}
var Current_Scene_Id = 0

func init() {
	worldmap.Current_Scene_Id = &Current_Scene_Id
}
