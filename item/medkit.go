package item

import (
	"main/camera"
	"main/utils"

	"github.com/bob4321at/textures"
	"github.com/hajimehoshi/ebiten/v2"
)

type MedKitItem struct {
	Pos         utils.Vec2
	Vel         utils.Vec2
	Img         textures.Texture
	Picked_Up   bool
	PickUpTimer float64
}

func (item *MedKitItem) PickUp() {
	if ebiten.IsKeyPressed(ebiten.KeyE) {
		item.PickUpTimer += 2
	}
	if item.PickUpTimer > 50 {
		*PlayerHealth += 50
		if *PlayerHealth > PlayerMaxHealth {
			*PlayerHealth = PlayerMaxHealth
			item.SetPickedUp()
		}
	}
}

func (item *MedKitItem) Update() {
	if item.PickUpTimer >= 1 {
		item.PickUpTimer -= 1
	}
	item.Vel.Y += 0.1

	if hit, _ := CheckLevelCollision(utils.Vec2{X: item.Pos.X, Y: item.Pos.Y + item.Vel.Y}, item.GetSize()); hit == true {
		item.Vel.Y = 0
	}

	item.Pos.Y += item.Vel.Y
}

func (item *MedKitItem) Draw(screen *ebiten.Image) {
	op := ebiten.DrawImageOptions{}
	op.GeoM.Translate(item.Pos.X+camera.Camera.Pos.X, item.Pos.Y+camera.Camera.Pos.Y)

	item.Img.Draw(screen, &op)

	Interacting_Icon.SetUniforms(map[string]any{
		"Percent": float64(item.PickUpTimer / 50),
	})

	if utils.Collide(utils.Vec2{X: PlayerPos.X + 320, Y: PlayerPos.Y + 180}, utils.Vec2{X: 32, Y: 48}, item.GetPos(), item.GetSize()) {
		op.GeoM.Translate(4, 2)
		Interacting_Icon.Draw(screen, &op)
	}
}

func (item *MedKitItem) GetPos() utils.Vec2 {
	return item.Pos
}

func (item *MedKitItem) GetSize() utils.Vec2 {
	return utils.Vec2{X: float64(item.Img.Img.Bounds().Dx()), Y: float64(item.Img.Img.Bounds().Dy())}
}

func (item *MedKitItem) PickedUp() bool {
	return item.Picked_Up
}

func (item *MedKitItem) SetPickedUp() {
	item.Picked_Up = true
}

func NewMedKit(Pos utils.Vec2) Item {
	medkit := &MedKitItem{}

	medkit.Pos = Pos
	medkit.Img = *textures.NewTexture("./art/items/heal/medkit.png", "")

	return medkit
}
