package level

import (
	"main/camera"
	"main/utils"
	"math"

	"github.com/bob4321at/textures"
	"github.com/hajimehoshi/ebiten/v2"
)

type Direction int

const (
	LEFT  Direction = 0
	RIGHT Direction = 1
)

type MovingPlatform struct {
	Pos                 utils.Vec2
	Vel                 utils.Vec2
	ReceiveSignal       Signal
	Img                 textures.Texture
	Track               int
	Current_Track_Index int
	Direction           Direction
}

func (tile *MovingPlatform) Draw(screen *ebiten.Image) {
	op := ebiten.DrawImageOptions{}
	op.GeoM.Translate(tile.Pos.X+camera.Camera.Pos.X-32, tile.Pos.Y+camera.Camera.Pos.Y-8)

	tile.Img.Draw(screen, &op)
}

func (platform *MovingPlatform) Update(level *Level) {
	if platform.Direction == RIGHT {
		if platform.Current_Track_Index+1 < len(level.MovingPlatformPaths[platform.Track]) {
			travel_angle := utils.Deg2Rad(utils.Rad2Deg(utils.GetAngle(platform.Pos, level.MovingPlatformPaths[platform.Track][platform.Current_Track_Index+1])) + 90)
			platform.Vel.X = math.Cos(travel_angle)
			platform.Vel.Y = -math.Sin(travel_angle)

			if utils.Collide(platform.Pos, utils.Vec2{X: 4, Y: 4}, level.MovingPlatformPaths[platform.Track][platform.Current_Track_Index+1], utils.Vec2{X: 4, Y: 4}) {
				platform.Current_Track_Index += 1
			}
		} else {
			if level.MovingPlatformLoopOrNot[platform.Track] {
				platform.Current_Track_Index = -1
			} else {
				platform.Direction = LEFT
			}
		}
	} else if platform.Direction == LEFT {
		if platform.Current_Track_Index-1 >= 0 {
			travel_angle := utils.Deg2Rad(utils.Rad2Deg(utils.GetAngle(platform.Pos, level.MovingPlatformPaths[platform.Track][platform.Current_Track_Index-1])) + 90)
			platform.Vel.X = math.Cos(travel_angle)
			platform.Vel.Y = -math.Sin(travel_angle)

			if utils.Collide(platform.Pos, utils.Vec2{X: 4, Y: 4}, level.MovingPlatformPaths[platform.Track][platform.Current_Track_Index-1], utils.Vec2{X: 4, Y: 4}) {
				platform.Current_Track_Index -= 1
			}
		} else {
			if level.MovingPlatformLoopOrNot[platform.Track] {
				platform.Current_Track_Index = -1
			} else {
				platform.Direction = RIGHT
			}
		}
	}

	platform.Pos.X += platform.Vel.X
	platform.Pos.Y += platform.Vel.Y
}

func NewMovingPlatform(Pos utils.Vec2, Track int) MovingPlatform {
	platform := MovingPlatform{}

	platform.Pos = Pos
	platform.Track = Track
	platform.Current_Track_Index = 0
	platform.Img = *textures.NewTexture("./art/moving_platform.png", "")
	platform.Direction = RIGHT

	return platform
}
