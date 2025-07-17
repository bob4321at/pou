package scenes

import (
	worldmap "main/world_map"

	"github.com/hajimehoshi/ebiten/v2"
)

type WorldMapScene struct {
	World_Map worldmap.WorldMapStruct
	SetedUp   bool
}

func (scene *WorldMapScene) Setup() {
	scene.SetedUp = true
}

func (scene *WorldMapScene) Draw(display *ebiten.Image) {
	scene.World_Map.Draw(display)
}

func (scene *WorldMapScene) Update() {
	scene.World_Map.Update()
}

func (scene *WorldMapScene) GetSetup() bool {
	return scene.SetedUp
}

func NewWorldMapScene() Scene {
	scene := &WorldMapScene{}

	scene.World_Map = worldmap.NewWorldMap()

	return scene
}
