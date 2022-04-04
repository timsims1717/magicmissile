package systems

import (
	"math"
	"math/rand"
	"timsims1717/magicmissile/internal/data"
	"timsims1717/magicmissile/internal/myecs"
	"timsims1717/magicmissile/internal/states/game"
	"timsims1717/magicmissile/pkg/object"
	"timsims1717/magicmissile/pkg/timing"
)

func MobSystem() {
	for _, result := range myecs.Manager.Query(myecs.IsMob) {
		if obj, ok := result.Components[myecs.Object].(*object.Object); ok {
			if mob, ok0 := result.Components[myecs.Mob].(*data.Mob); ok0 {
				if !mob.Char.Health.Dead {
					if !mob.Attack.Attacking {
						if mob.Target == nil || mob.Target.Health.Dead {
							town := game.Towns[rand.Intn(len(game.Towns))]
							if !town.Health.Dead {
								mob.Target = town
							} else {
								mob.Target = nil
							}
						}
						if mob.Target != nil {
							xDist := mob.Target.Obj.Pos.X - obj.Pos.X
							if math.Abs(xDist) < 30. {
								obj.Pos.Y += mob.Speed * timing.DT
								if obj.Pos.Y > mob.Target.Obj.Pos.Y {
									obj.Pos.Y = mob.Target.Obj.Pos.Y
								}
							} else {
								if obj.Pos.Y > game.CharYLvl {
									obj.Pos.Y -= mob.Speed * timing.DT
									if obj.Pos.Y < game.CharYLvl {
										obj.Pos.Y = game.CharYLvl
									}
								} else if obj.Pos.Y < game.CharYLvl {
									obj.Pos.Y += mob.Speed * timing.DT
									if obj.Pos.Y > game.CharYLvl {
										obj.Pos.Y = game.CharYLvl
									}
								}
							}
							if math.Abs(mob.Target.Obj.Pos.X - obj.Pos.X) > 1. {
								if xDist > 0. {
									obj.Pos.X += mob.Speed * timing.DT
									mob.Char.Obj.Flip = false
								} else {
									obj.Pos.X -= mob.Speed * timing.DT
	 								mob.Char.Obj.Flip = true
								}
							}
						}
					}
				} else {
					myecs.Manager.DisposeEntity(result)
				}
			}
		}
	}
}