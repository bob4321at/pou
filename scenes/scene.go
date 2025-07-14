package scenes

import "github.com/hajimehoshi/ebiten/v2"

type Scene interface {
	Setup()
	Update()
	Draw(display *ebiten.Image)
	GetSetup() bool
}

var Scenes = []Scene{&Main_Menu_Scene{}, &GameScene{}, &Death_Menu_scene{}, &Win_Menu_Scene{}}
var Current_Scene_Id = 0
