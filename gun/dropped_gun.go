package gun

import (
	"main/camera"
	"main/enemies"
	"main/shaders"
	"main/utils"

	"github.com/bob4321at/textures"
	"github.com/hajimehoshi/ebiten/v2"
)

type DroppedGunStruct struct {
	Img       textures.RenderableTexture
	Pos       utils.Vec2
	Vel       utils.Vec2
	GiveFunc  func() Gun
	Picked_Up bool
	Counter   float64
}

var Interacting_Icon = textures.NewTexture("./art/ui/interact_icon.png", shaders.Fill_Shader)

func (dropped_gun *DroppedGunStruct) Draw(screen *ebiten.Image) {
	op := ebiten.DrawImageOptions{}
	op.GeoM.Translate(dropped_gun.Pos.X+camera.Camera.Pos.X, dropped_gun.Pos.Y+camera.Camera.Pos.Y)

	dropped_gun.Img.Draw(screen, &op)

	if utils.Collide(dropped_gun.Pos, utils.Vec2{X: float64(dropped_gun.Img.GetTexture().Bounds().Dx()), Y: float64(dropped_gun.Img.GetTexture().Bounds().Dy())}, utils.Vec2{X: Player_Pos.X + 320, Y: Player_Pos.Y + 180}, utils.Vec2{X: 32, Y: 48}) {
		op.GeoM.Reset()
		op.GeoM.Translate(dropped_gun.Pos.X+camera.Camera.Pos.X+float64(dropped_gun.Img.GetTexture().Bounds().Dx()/2), dropped_gun.Pos.Y+camera.Camera.Pos.Y)

		Interacting_Icon.SetUniforms(map[string]any{
			"Percent": float64(dropped_gun.Counter / 5),
		})

		Interacting_Icon.Draw(screen, &op)
	}
}

func (dropped_gun *DroppedGunStruct) Update() {
	dropped_gun.Vel.Y += 0.1

	collide_y, _ := enemies.CheckLevelCollision(utils.Vec2{X: dropped_gun.Pos.X, Y: dropped_gun.Pos.Y + dropped_gun.Vel.Y}, utils.Vec2{X: float64(dropped_gun.Img.GetTexture().Bounds().Dx()), Y: float64(dropped_gun.Img.GetTexture().Bounds().Dy())})

	// collide_y = false

	if collide_y {
		dropped_gun.Vel.Y = 0
	}

	if utils.Collide(dropped_gun.Pos, utils.Vec2{X: float64(dropped_gun.Img.GetTexture().Bounds().Dx()), Y: float64(dropped_gun.Img.GetTexture().Bounds().Dy())}, utils.Vec2{X: Player_Pos.X + 320, Y: Player_Pos.Y + 180}, utils.Vec2{X: 32, Y: 48}) {
		if ebiten.IsKeyPressed(ebiten.KeyE) {
			dropped_gun.Counter += 0.1
		} else {
			dropped_gun.Counter = 0
		}
	} else {
		dropped_gun.Counter = 0
	}

	if dropped_gun.Counter > 5 {
		dropped_gun.Picked_Up = true
	}

	dropped_gun.Pos.Y += dropped_gun.Vel.Y
}

func CreateDroppedGun(gun func() Gun, pos utils.Vec2) DroppedGunStruct {
	dropped_gun := DroppedGunStruct{}

	dropped_gun.Img = gun().GetDroppedImg()
	dropped_gun.Pos = pos
	dropped_gun.GiveFunc = gun

	return dropped_gun
}
