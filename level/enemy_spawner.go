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
	SendSignal      Signal
	ReceiveSignal   Signal
	Spawn           bool
}

func (spawner *EnemySpawner) Update(level *Level) {
	if spawner.Spawn == true {
		if spawner.Timer > 0 {
			spawner.Timer -= 0.1
		} else {
			spawner.Timer = 10

			if spawner.Index < len(spawner.Enemies) {
				spawner.Responsible_For = append(spawner.Responsible_For, enemies.EnemySpawnFuncs[spawner.Enemies[spawner.Index]](spawner.Pos))
				spawner.Index += 1
			}
		}

		if spawner.Index == len(spawner.Enemies) {
			if len(spawner.Responsible_For) == 0 {
				spawner.SendSignal.Active = true
			}
		}
	} else {
		for _, spawner := range level.Enemy_Spawners {
			if spawner.SendSignal.Id == spawner.ReceiveSignal.Id {
				if spawner.SendSignal.Active {
					spawner.ReceiveSignal.Active = true
				}
			}
		}
		for _, trigger := range level.TriggerTile {
			if trigger.Signal == spawner.ReceiveSignal.Id {
				if trigger.Active {
					spawner.ReceiveSignal.Active = true
				}
			}
		}

		if spawner.ReceiveSignal.Active {
			spawner.Spawn = true
		}
	}
}
