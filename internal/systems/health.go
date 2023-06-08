package systems

import (
	"timsims1717/magicmissile/internal/data"
	"timsims1717/magicmissile/internal/myecs"
	"timsims1717/magicmissile/pkg/object"
)

func HealthSystem() {
	for _, result := range myecs.Manager.Query(myecs.HasHealth) {
		if obj, ok := result.Components[myecs.Object].(*object.Object); ok {
			if hp, okB := result.Components[myecs.Health].(*data.Health); okB {
				if hp.HP == 0 {
					hp.Dead = true
				}
				if hp.Dead {
					obj.Killed = true
				}
			}
		}
	}
}
