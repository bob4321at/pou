package player

import (
	"fmt"
	"image/color"
	"main/camera"
	"main/enemies"
	"main/gun"
	"main/item"
	"main/level"
	"main/shaders"
	"main/utils"
	"math"

	"github.com/bob4321at/textures"

	"github.com/hajimehoshi/ebiten/v2"
)

const MAX_PLAYER_OXYGEN int = 500
const MAX_PLAYER_HEALTH int = 100

type PlayerStruct struct {
	Pos   utils.Vec2
	Vel   utils.Vec2
	Speed float64

	Moving_Platform_Vel utils.Vec2

	Dir                  bool
	Img                  *textures.AnimatedTexture
	WithFlyingBootsImage *textures.AnimatedTexture
	Charge_UI            textures.RenderableTexture

	Gun              gun.Gun
	DamageMultiplier float64

	Health          int
	Max_Health      int
	Previous_Health int
	Health_Bar_Img  *ebiten.Image
	I_Frames        float64

	UnderWater   bool
	Oxygen       int
	AirBubbleImg *textures.Texture

	Charged int

	Won bool
}

func (player *PlayerStruct) Update(current_level *level.Level) {
	gun.Player_Pos = &player.Pos
	gun.Player_Vel = &player.Vel
	gun.Player_I_Frames = player.I_Frames
	gun.Player_Damage_Multiplier = &player.DamageMultiplier
	enemies.Player_Pos = &player.Pos
	enemies.Player_Health = &player.Health

	item.PlayerHealth = &player.Health
	item.PlayerPos = &player.Pos
	item.PlayerSpeed = &player.Speed
	item.PlayerDamageMultiplier = &player.DamageMultiplier

	if player.Pos.Y+360/2 > current_level.WaterLevel {
		player.UnderWater = true
	} else {
		player.UnderWater = false
	}

	if player.UnderWater {
		player.Oxygen -= 1
	} else {
		if player.Oxygen < MAX_PLAYER_OXYGEN {
			player.Oxygen += 1
		}
	}

	if player.Oxygen <= 0 {
		player.Health = -1
	}

	if player.I_Frames > 0 {
		player.I_Frames -= 0.1
	} else {
		player.I_Frames = 0
		for ei := range current_level.Enemies {
			enemy := current_level.Enemies[ei]
			if utils.Collide(utils.Vec2{X: player.Pos.X + player.Vel.X*player.Speed + 640/2 - 14, Y: player.Pos.Y + 360/2 - 18}, utils.Vec2{X: 32, Y: 48}, enemy.GetPosition(), enemy.GetSize()) {
				enemy.HitPlayer()
			}
		}

		for i, spike := range current_level.SpikeTiles {
			tile := &current_level.SpikeTiles[i]

			if utils.Collide(utils.Vec2{X: player.Pos.X + 640/2 - 14, Y: player.Pos.Y + player.Vel.Y + 360/2 - 18}, utils.Vec2{X: 28, Y: 42}, tile.Pos, utils.Vec2{X: 32, Y: 32}) {
				player.Health -= spike.Damage
			}
		}

		for i, spring := range current_level.SpringTiles {
			tile := &current_level.SpringTiles[i]

			switch spring.Direction {
			case 0:
				if utils.Collide(utils.Vec2{X: player.Pos.X + 640/2 - 14 + player.Vel.X*player.Speed, Y: player.Pos.Y + player.Vel.Y + 360/2 - 18 + player.Vel.Y}, utils.Vec2{X: 28, Y: 42}, utils.Vec2{X: tile.Pos.X, Y: tile.Pos.Y + 28}, utils.Vec2{X: 32, Y: 4}) {
					player.Vel.Y = -float64(spring.Power) * 3.2
				}
			case 1:
				if utils.Collide(utils.Vec2{X: player.Pos.X + 640/2 - 14 + player.Vel.X*player.Speed, Y: player.Pos.Y + player.Vel.Y + 360/2 - 18 + player.Vel.Y}, utils.Vec2{X: 28, Y: 42}, utils.Vec2{X: tile.Pos.X, Y: tile.Pos.Y}, utils.Vec2{X: 4, Y: 32}) {
					player.Vel.X = float64(spring.Power) * 3.2
				}
			case 2:
				if utils.Collide(utils.Vec2{X: player.Pos.X + 640/2 - 14 + player.Vel.X*player.Speed, Y: player.Pos.Y + player.Vel.Y + 360/2 - 18 + player.Vel.Y}, utils.Vec2{X: 28, Y: 42}, utils.Vec2{X: tile.Pos.X, Y: tile.Pos.Y}, utils.Vec2{X: 32, Y: 4}) {
					player.Vel.Y = float64(spring.Power) * 3.2
				}
			case 3:
				if utils.Collide(utils.Vec2{X: player.Pos.X + 640/2 - 14 + player.Vel.X*player.Speed, Y: player.Pos.Y + player.Vel.Y + 360/2 - 18 + player.Vel.Y}, utils.Vec2{X: 28, Y: 42}, utils.Vec2{X: tile.Pos.X + 28, Y: tile.Pos.Y}, utils.Vec2{X: 32, Y: 4}) {
					player.Vel.X = -float64(spring.Power) * 3.2
				}
			}
		}
	}

	for _, item := range current_level.Items {
		if utils.Collide(utils.Vec2{X: player.Pos.X + 640/2 - 14, Y: player.Pos.Y + player.Vel.Y + 360/2 - 18}, utils.Vec2{X: 32, Y: 48}, item.GetPos(), item.GetSize()) {
			item.PickUp()
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeyT) {
		player.Health -= 1
	}

	if player.Health < player.Previous_Health {
		player.I_Frames = 1
	}

	player.Img.Update()
	player.Img.SetUniforms(map[string]any{
		"I": player.I_Frames,
	})

	player.WithFlyingBootsImage.Update()
	player.WithFlyingBootsImage.SetUniforms(map[string]any{
		"I": player.I_Frames,
	})

	player.Img.Current_Animation = 0
	player.WithFlyingBootsImage.Current_Animation = 0

	if !current_level.Player_Loaded {
		player.Vel.X = 0
		player.Vel.Y = 0
		player.Pos = current_level.Player_Spawn
		current_level.Player_Loaded = true
	}

	player.Vel.Y += 0.1

	if ebiten.IsKeyPressed(ebiten.KeyA) && player.Vel.X-0.1 > -4 {
		player.Vel.X -= 0.1
		if player.Vel.X > 0 {
			player.Vel.X -= 0.2
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) && player.Vel.X+0.1 < 4 {
		player.Vel.X += 0.1
		if player.Vel.X < 0 {
			player.Vel.X += 0.2
		}
	}

	if player.Vel.X > 0 {
		player.Dir = false
	} else if player.Vel.X < 0 {
		player.Dir = true
	}

	collision_y, _ := current_level.CheckCollision(utils.Vec2{X: player.Pos.X + 640/2 - 14, Y: player.Pos.Y + player.Vel.Y + 360/2 - 18}, utils.Vec2{X: 28, Y: 42})
	if collision_y {
		if player.Vel.Y > 0 {
			if math.Abs(player.Vel.X) >= 0.1 {
				player.Img.Current_Animation = 1
				player.WithFlyingBootsImage.Current_Animation = 1
			} else {
				player.Img.Current_Animation = 0
				player.WithFlyingBootsImage.Current_Animation = 0
			}

			player.Vel.Y = 0
			if ebiten.IsKeyPressed(ebiten.KeyW) || ebiten.IsKeyPressed(ebiten.KeySpace) {
				player.Vel.Y = -4
			}
			if !ebiten.IsKeyPressed(ebiten.KeyA) && !ebiten.IsKeyPressed(ebiten.KeyD) {
				player.Vel.X -= player.Vel.X / 5
			}
		} else {
			player.Vel.Y = 0
			player.Moving_Platform_Vel.Y = 0
		}
	} else {
		if player.Vel.Y > 0 {
			player.Img.Current_Animation = 2
			player.WithFlyingBootsImage.Current_Animation = 2
		} else {
			player.Img.Current_Animation = 3
			player.WithFlyingBootsImage.Current_Animation = 3
		}
	}

	player.Moving_Platform_Vel = utils.Vec2{}

	for _, movingplatform := range current_level.MovingPlatforms {
		if player.Vel.Y > 0 {
			if utils.Collide(utils.Vec2{X: player.Pos.X + 640/2 - 14, Y: player.Pos.Y + player.Vel.Y + 360/2 - 18 + player.Vel.Y + 38}, utils.Vec2{X: 28, Y: 4}, utils.Vec2{X: movingplatform.Pos.X - 32 + movingplatform.Vel.X, Y: movingplatform.Pos.Y - 8 + movingplatform.Vel.Y}, utils.Vec2{X: 64, Y: 4}) {
				if player.Vel.Y > 0 {
					if math.Abs(player.Vel.X) >= 0.1 {
						player.Img.Current_Animation = 1
						player.WithFlyingBootsImage.Current_Animation = 1
					} else {
						player.Img.Current_Animation = 0
						player.WithFlyingBootsImage.Current_Animation = 0
					}

					player.Vel.Y = 0
					if ebiten.IsKeyPressed(ebiten.KeyW) || ebiten.IsKeyPressed(ebiten.KeySpace) {
						player.Vel.Y = -4 + movingplatform.Vel.Y
					}
					if !ebiten.IsKeyPressed(ebiten.KeyA) && !ebiten.IsKeyPressed(ebiten.KeyD) {
						player.Vel.X -= player.Vel.X / 5
					}
				} else {
					player.Vel.Y = 0
				}
			}

			if movingplatform.Vel.Y > 0 {
				if utils.Collide(utils.Vec2{X: player.Pos.X + 640/2 - 14 + player.Vel.X*player.Speed, Y: player.Pos.Y + 360/2 - 18 + player.Vel.Y}, utils.Vec2{X: 28, Y: 42}, utils.Vec2{X: movingplatform.Pos.X - 32 + movingplatform.Vel.X, Y: movingplatform.Pos.Y - 8 + movingplatform.Vel.Y - 4}, utils.Vec2{X: 64, Y: 20}) {
					player.Moving_Platform_Vel.X = movingplatform.Vel.X
					player.Moving_Platform_Vel.Y = movingplatform.Vel.Y
				}
			} else {
				if utils.Collide(utils.Vec2{X: player.Pos.X + 640/2 - 14 + player.Vel.X*player.Speed, Y: player.Pos.Y + 360/2 - 18 + player.Vel.Y}, utils.Vec2{X: 28, Y: 42}, utils.Vec2{X: movingplatform.Pos.X - 32 + movingplatform.Vel.X, Y: movingplatform.Pos.Y - 8 + movingplatform.Vel.Y - 1}, utils.Vec2{X: 64, Y: 18}) {
					player.Moving_Platform_Vel.X = movingplatform.Vel.X
					player.Moving_Platform_Vel.Y = movingplatform.Vel.Y
				}
			}
		}
	}
	collision_x, _ := current_level.CheckCollision(utils.Vec2{X: player.Pos.X + player.Vel.X*player.Speed + player.Moving_Platform_Vel.X + 640/2 - 14, Y: player.Pos.Y + 360/2 - 18}, utils.Vec2{X: 28, Y: 42})
	if collision_x {
		player.Vel.X = 0
		player.Moving_Platform_Vel.X = 0
	}

	for i := range current_level.TriggerTile {
		tile := &current_level.TriggerTile[i]

		if utils.Collide(utils.Vec2{X: player.Pos.X + 640/2 - 14, Y: player.Pos.Y + player.Vel.Y + 360/2 - 18}, utils.Vec2{X: 28, Y: 42}, tile.Pos, utils.Vec2{X: 32, Y: 32}) {
			tile.SendSignal.Active = true
		}
	}

	if utils.Collide(utils.Vec2{X: player.Pos.X + 640/2 - 14, Y: player.Pos.Y + player.Vel.Y + 360/2 - 18}, utils.Vec2{X: 28, Y: 42}, current_level.End_Pos, utils.Vec2{X: 32, Y: 32}) {
		player.Won = true
	}

	camera.Camera.Pos.X = -player.Pos.X
	camera.Camera.Pos.Y = -player.Pos.Y

	player.Pos.X += player.Vel.X*player.Speed + player.Moving_Platform_Vel.X
	player.Pos.Y += player.Vel.Y + player.Moving_Platform_Vel.Y

	player.Gun.Update()

	if ebiten.IsMouseButtonPressed(ebiten.MouseButton0) && player.Gun.CanShoot() {
		player.Charged += 1
	} else if !ebiten.IsMouseButtonPressed(ebiten.MouseButton0) && player.Charged != -1 {
		player.Gun.Shoot(player.Charged)
		player.Charged = -1
	}

	player.Previous_Health = player.Health
}

func (player *PlayerStruct) Draw(screen *ebiten.Image) {
	op := ebiten.DrawImageOptions{}

	if player.Dir {
		op.GeoM.Translate(-16, 0)
		op.GeoM.Scale(-1, 1)
		op.GeoM.Translate(16, 0)
	}

	op.GeoM.Translate(640/2-16, 360/2-24)

	fmt.Println(float64(float64(player.Oxygen) / float64(MAX_PLAYER_OXYGEN)))
	// if player.UnderWater {
	player.AirBubbleImg.SetUniforms(map[string]any{
		"Percent": float64(float64(player.Oxygen) / float64(MAX_PLAYER_OXYGEN)),
	})
	player.AirBubbleImg.Draw(screen, &op)
	// }

	if player.Speed == 1 {
		player.Img.Draw(screen, &op)
	} else {
		player.WithFlyingBootsImage.Draw(screen, &op)
	}

	player.Gun.Draw(screen)

	if player.Charged != -1 {
		op.GeoM.Reset()
		op.GeoM.Translate(640/2-4, 360/2-6)

		player.Charge_UI.SetUniforms(map[string]any{
			"Percent": float64(float64(player.Charged) / float64(player.Gun.GetCharge())),
		})

		player.Charge_UI.Draw(screen, &op)
	}
}

func (player *PlayerStruct) Reset(spawn_pos utils.Vec2) {
	player.Pos = spawn_pos
	player.I_Frames = 0
	player.Health = 100
	player.Won = false
	player.Charged = 0
	player.Speed = 1
	player.Oxygen = MAX_PLAYER_OXYGEN
}

func NewPlayer(pos utils.Vec2) (player PlayerStruct) {
	player.Pos = pos

	player.Img = textures.NewAnimatedTexture("./art/player.png", shaders.Flash_Shader)
	player.WithFlyingBootsImage = textures.NewAnimatedTexture("./art/player_flying_boots.png", shaders.Flash_Shader)
	player.Charge_UI = textures.NewTexture("./art/ui/charge.png", shaders.Fill_Shader)
	player.Gun = gun.CreateFistGun()

	player.Health = MAX_PLAYER_HEALTH
	player.Health_Bar_Img = ebiten.NewImage(250, 24)
	player.Health_Bar_Img.Fill(color.RGBA{255, 50, 50, 255})

	player.Oxygen = MAX_PLAYER_OXYGEN
	player.AirBubbleImg = textures.NewTexture("./art/ui/air_bubble.png", shaders.Air_Bubble_Shader)

	item.PlayerHealth = enemies.Player_Health
	item.PlayerMaxHealth = MAX_PLAYER_HEALTH

	player.Won = false
	player.Speed = 1

	player.DamageMultiplier = 1

	return player
}
