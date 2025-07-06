package level

import (
	"main/enemies"
	"main/utils"
)

type EnemySpawner struct {
	Pos             utils.Vec2
	Enemies         []int
	Responsible_For []enemies.Enemy
	Timer           float64
	Index           int
	Signal          Signal
}

func (spawner *EnemySpawner) Update() {
	if spawner.Timer > 0 {
		spawner.Timer -= 0.1
	} else {
		spawner.Timer = 10

		if spawner.Index < len(spawner.Enemies) {
			spawner.Responsible_For = append(spawner.Responsible_For, enemies.EnemySpawnFuncs[spawner.Enemies[spawner.Index]](spawner.Pos))
			spawner.Responsible_For = append(spawner.Responsible_For, enemies.EnemySpawnFuncs[spawner.Enemies[spawner.Index]](spawner.Pos))
			spawner.Index += 1
		}
	}

	if spawner.Index == len(spawner.Enemies) {
		if len(spawner.Responsible_For) == 0 {
			spawner.Signal.Active = true
		}
	}
}
