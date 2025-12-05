package scenes

import (
	"main/utils"
	worldmap "main/world_map"

	"github.com/hajimehoshi/ebiten/v2"
)

type WorldMapScene struct {
	World_Maps        []worldmap.WorldMapStruct
	Current_World_Map int
	SetedUp           bool
}

func (scene *WorldMapScene) Setup() {
	scene.SetedUp = true
}

func (scene *WorldMapScene) Draw(display *ebiten.Image) {
	scene.World_Maps[scene.Current_World_Map].Draw(display)
}

func (scene *WorldMapScene) Update() {
	scene.World_Maps[scene.Current_World_Map].Update()

	if scene.World_Maps[scene.Current_World_Map].Finished {
		scene.World_Maps[scene.Current_World_Map].Finished = false
		scene.Current_World_Map += 1
	}
}

func (scene *WorldMapScene) GetSetup() bool {
	return scene.SetedUp
}

func NewWorldMapScene() Scene {
	scene := &WorldMapScene{}

	scene.World_Maps = []worldmap.WorldMapStruct{
		worldmap.NewWorldMap(
			"./art/menus/world_maps/world_one_map.png",
			worldmap.NewLevelPoint("test", utils.Vec2{X: 0, Y: 0}),
			worldmap.NewLevelPoint("tutorial", utils.Vec2{X: 71, Y: 39}),
			worldmap.NewLevelPoint("test_real_level", utils.Vec2{X: 294, Y: 44}),
			worldmap.NewLevelPoint("spikes_come", utils.Vec2{X: 597, Y: 54}),
			worldmap.NewLevelPoint("spring_land", utils.Vec2{X: 615, Y: 187}),
			worldmap.NewLevelPoint("med_intro", utils.Vec2{X: 511, Y: 228}),
			worldmap.NewLevelPoint("moving_platform_intro", utils.Vec2{X: 483, Y: 141}),
			worldmap.NewLevelPoint("flying_boot_intro", utils.Vec2{X: 152, Y: 119}),
			worldmap.NewLevelPoint("coffee_intro", utils.Vec2{X: 56, Y: 204}),
			worldmap.NewLevelPoint("button_intro", utils.Vec2{X: 108, Y: 303}),
			worldmap.NewLevelPoint("turret_intro", utils.Vec2{X: 607, Y: 321}),
		),
		worldmap.NewWorldMap(
			"./art/menus/world_maps/world_two_map.png",
			worldmap.NewLevelPoint("flood_intro", utils.Vec2{X: 67, Y: 44}),
			worldmap.NewLevelPoint("water_intro", utils.Vec2{X: 45, Y: 189}),
			worldmap.NewLevelPoint("test_real_level", utils.Vec2{X: 25, Y: 317}),
			worldmap.NewLevelPoint("arson", utils.Vec2{X: 274, Y: 326}),
			worldmap.NewLevelPoint("red_jumping", utils.Vec2{X: 559, Y: 327}),
			worldmap.NewLevelPoint("hard", utils.Vec2{X: 398, Y: 253}),
			worldmap.NewLevelPoint("arena", utils.Vec2{X: 182, Y: 161}),
			worldmap.NewLevelPoint("plateform", utils.Vec2{X: 384, Y: 36}),
			worldmap.NewLevelPoint("spikes", utils.Vec2{X: 613, Y: 127}),
		),
	}

	return scene
}
