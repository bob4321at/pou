package enemies

import (
	"main/utils"

	"github.com/hajimehoshi/ebiten/v2"
)

type Enemy interface {
	Draw(screen *ebiten.Image)
	Update()
	Hit(damage int)
	HitPlayer()
	GetPosition() utils.Vec2
	GetSize() utils.Vec2
	GetHealth() int
}

type EnemyProjectile interface {
	Draw(screen *ebiten.Image)
	Update()
	Hit()
	GetSize() utils.Vec2
	GetPos() utils.Vec2
	ShouldRemove() bool
}

var EnemySpawnFuncs = map[int]func(pos utils.Vec2) Enemy{
	1: NewTarget,
	2: NewBlueGhost,
	3: NewOrangeGuy,
	4: NewDummy,
	5: NewRedGhost,
	6: NewRedGuy,
	7: NewTurret,
	8: NewBlueGuy,
	9: NewGreenGuy,
}

var Level_Hitbox []utils.Vec2
var Breakable_Tile_Hitboxes []utils.Vec2
var AllEnemies []Enemy

var Player_Pos *utils.Vec2
var Player_Health *int

var Enemies_To_Add []Enemy
var Projectiles []EnemyProjectile

func ManageProjectiles() {
	for projectile_index := len(Projectiles); projectile_index > 0; projectile_index++ {
		projectile := Projectiles[projectile_index-1]
		if projectile.ShouldRemove() {
			utils.RemoveArrayElement(projectile_index, &Projectiles)
		}
	}
}

func SpawnProjectile(projectile EnemyProjectile) {
	Projectiles = append(Projectiles, projectile)
}

func CheckLevelCollision(pos utils.Vec2, size utils.Vec2) (bool, utils.Vec2) {
	hit := false
	tile_pos := utils.Vec2{}

	for _, tile := range Level_Hitbox {
		check := utils.Collide(pos, size, utils.Vec2{X: float64(tile.X), Y: float64(tile.Y)}, utils.Vec2{X: 32, Y: 32})

		if check {
			hit = true
			tile_pos = utils.Vec2{X: float64(tile.X), Y: float64(tile.Y)}
		}
	}

	for _, breakable_tile := range Breakable_Tile_Hitboxes {
		check := utils.Collide(pos, size, utils.Vec2{X: float64(breakable_tile.X), Y: float64(breakable_tile.Y)}, utils.Vec2{X: 32, Y: 32})

		if check {
			hit = true
			tile_pos = utils.Vec2{X: float64(breakable_tile.X), Y: float64(breakable_tile.Y)}
		}
	}

	return hit, tile_pos
}
