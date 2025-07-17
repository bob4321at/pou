package worldmap

import (
	"main/utils"
	"math"

	"github.com/bob4321at/textures"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type WorldMapStruct struct {
	Points []PointOfIntrest
	Img    *ebiten.Image

	World_Map_Player            textures.AnimatedTexture
	World_Map_Player_Pos        utils.Vec2
	World_Map_Player_Index      int
	World_Map_Player_Index_Prev int
	World_Map_Player_Reached    bool
	World_Map_Player_Inputed    bool
	World_Map_Player_Dir        bool
}

func (world_map *WorldMapStruct) Update() {
	if world_map.World_Map_Player_Reached {
		if world_map.World_Map_Player_Index+1 < len(world_map.Points) {
			if ebiten.IsKeyPressed(ebiten.KeyD) && !world_map.World_Map_Player_Inputed {
				world_map.World_Map_Player_Index += 1
				world_map.World_Map_Player_Inputed = true
			}
		}

		if world_map.World_Map_Player_Index-1 >= 0 {
			if ebiten.IsKeyPressed(ebiten.KeyA) && !world_map.World_Map_Player_Inputed {
				world_map.World_Map_Player_Index -= 1
				world_map.World_Map_Player_Inputed = true
			}
		}

		world_map.World_Map_Player.Current_Animation = 0

		if ebiten.IsKeyPressed(ebiten.KeyEnter) {
			world_map.Points[world_map.World_Map_Player_Index].GoTo()
		}
	} else {
		world_map.World_Map_Player.Current_Animation = 1
	}

	if !ebiten.IsKeyPressed(ebiten.KeyD) && !ebiten.IsKeyPressed(ebiten.KeyA) {
		world_map.World_Map_Player_Inputed = false
	}

	if world_map.World_Map_Player_Index_Prev != world_map.World_Map_Player_Index {
		world_map.World_Map_Player_Reached = false
	}

	world_map.World_Map_Player.Update()
}

func (world_map *WorldMapStruct) Draw(screen *ebiten.Image) {
	op := ebiten.DrawImageOptions{}
	screen.DrawImage(world_map.Img, &op)

	for _, point := range world_map.Points {
		point.Draw(screen)
	}

	if !world_map.World_Map_Player_Reached {
		angle := utils.Deg2Rad(utils.Rad2Deg(utils.GetAngle(world_map.World_Map_Player_Pos, world_map.Points[world_map.World_Map_Player_Index].GetPos())) + 90)
		world_map.World_Map_Player_Pos.X += math.Cos(angle)
		world_map.World_Map_Player_Pos.Y -= math.Sin(angle)
	}

	if world_map.World_Map_Player_Pos.X > world_map.Points[world_map.World_Map_Player_Index].GetPos().X {
		world_map.World_Map_Player_Dir = true
	} else {
		world_map.World_Map_Player_Dir = false
	}

	if utils.GetDist(world_map.World_Map_Player_Pos, world_map.Points[world_map.World_Map_Player_Index].GetPos()) < 1 {
		world_map.World_Map_Player_Reached = true
		world_map.World_Map_Player_Index_Prev = world_map.World_Map_Player_Index
	}

	if !world_map.World_Map_Player_Dir {
		op.GeoM.Translate(world_map.World_Map_Player_Pos.X, world_map.World_Map_Player_Pos.Y)
	} else {
		op.GeoM.Scale(-1, 1)
		op.GeoM.Translate(world_map.World_Map_Player_Pos.X+32, world_map.World_Map_Player_Pos.Y)
	}
	world_map.World_Map_Player.Draw(screen, &op)
}

func NewWorldMap() WorldMapStruct {
	world_map := WorldMapStruct{}

	var err error
	world_map.Img, _, err = ebitenutil.NewImageFromFile("./art/menus/world_map.png")
	if err != nil {
		panic(err)
	}

	world_map.World_Map_Player = *textures.NewAnimatedTexture("./art/ui/world_map_player.png", "")

	world_map.Points = append(world_map.Points, NewLevelPoint("tutorial", utils.Vec2{X: 71, Y: 39}))
	world_map.Points = append(world_map.Points, NewLevelPoint("test_real_level", utils.Vec2{X: 294, Y: 44}))
	world_map.Points = append(world_map.Points, NewLevelPoint("signal_test", utils.Vec2{X: 597, Y: 54}))
	world_map.Points = append(world_map.Points, NewLevelPoint("test", utils.Vec2{X: 615, Y: 187}))
	world_map.Points = append(world_map.Points, NewLevelPoint("test", utils.Vec2{X: 511, Y: 228}))
	world_map.Points = append(world_map.Points, NewLevelPoint("test", utils.Vec2{X: 483, Y: 141}))
	world_map.Points = append(world_map.Points, NewLevelPoint("test", utils.Vec2{X: 152, Y: 119}))
	world_map.Points = append(world_map.Points, NewLevelPoint("test", utils.Vec2{X: 56, Y: 204}))
	world_map.Points = append(world_map.Points, NewLevelPoint("test", utils.Vec2{X: 108, Y: 303}))
	world_map.Points = append(world_map.Points, NewLevelPoint("test", utils.Vec2{X: 607, Y: 321}))

	return world_map
}
