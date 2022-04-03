package systems

import (
	"timsims1717/magicmissile/internal/myecs"
	"timsims1717/magicmissile/pkg/object"
	"timsims1717/magicmissile/pkg/timing"
)

func TemporarySystem() {
	for _, result := range myecs.Manager.Query(myecs.IsTemp) {
		temp := result.Components[myecs.Temp]
		trans, okT := result.Components[myecs.Object].(*object.Object)
		if okT {
			if timer, ok := temp.(*timing.Timer); ok {
				if timer.UpdateDone() {
					trans.Hide = true
					trans.Gone = true
					myecs.Manager.DisposeEntity(result.Entity)
				}
			} else if check, ok := temp.(myecs.ClearFlag); ok {
				if check {
					trans.Hide = true
					trans.Gone = true
					myecs.Manager.DisposeEntity(result.Entity)
				}
			}
		}
	}
}

func ClearSystem() {
	for _, result := range myecs.Manager.Query(myecs.IsTemp) {
		trans, ok := result.Components[myecs.Object].(*object.Object)
		if ok {
			trans.Hide = true
			trans.Gone = true
		}
		myecs.Manager.DisposeEntity(result.Entity)
	}
}
