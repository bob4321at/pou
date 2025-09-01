package item

import (
	"main/camera"
	"main/utils"

	"github.com/bob4321at/textures"
	"github.com/hajimehoshi/ebiten/v2"
)

type Coffee struct {
	Pos         utils.Vec2
	Vel         utils.Vec2
	Img         textures.Texture
	Picked_Up   bool
	PickUpTimer float64
}

func (item *Coffee) PickUp() {
	if ebiten.IsKeyPressed(ebiten.KeyE) {
		item.PickUpTimer += 2
	}
	if item.PickUpTimer > 50 {
		*PlayerHealth -= 30
		*PlayerDamageMultiplier = 2
		item.SetPickedUp()
	}
}

func (item *Coffee) Update() {
	if item.PickUpTimer >= 1 {
		item.PickUpTimer -= 1
	}
	item.Vel.Y += 0.1

	if hit, _ := CheckLevelCollision(utils.Vec2{X: item.Pos.X, Y: item.Pos.Y + item.Vel.Y}, item.GetSize()); hit == true {
		item.Vel.Y = 0
	}

	item.Pos.Y += item.Vel.Y
}

func (item *Coffee) Draw(screen *ebiten.Image) {
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

func (item *Coffee) GetPos() utils.Vec2 {
	return item.Pos
}

func (item *Coffee) GetSize() utils.Vec2 {
	return utils.Vec2{X: float64(item.Img.Img.Bounds().Dx()), Y: float64(item.Img.Img.Bounds().Dy())}
}

func (item *Coffee) PickedUp() bool {
	return item.Picked_Up
}

func (item *Coffee) SetPickedUp() {
	item.Picked_Up = true
}

func NewCoffeeBoot(Pos utils.Vec2) Item {
	medkit := &Coffee{}

	medkit.Pos = Pos
	medkit.Img = *textures.NewTexture("./art/items/boosts/coffee.png", "")

	return medkit
}
