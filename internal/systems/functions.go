package systems

import (
	"timsims1717/magicmissile/internal/data"
	"timsims1717/magicmissile/internal/myecs"
)

func FunctionSystem() {
	for _, result := range myecs.Manager.Query(myecs.HasUpdate) {
		fnA := result.Components[myecs.Update]
		if fnT, ok := fnA.(*data.TimerFunc); ok {
			if fnT.Timer.UpdateDone() {
				if fnT.Func() {
					result.Entity.RemoveComponent(myecs.Update)
				} else {
					fnT.Timer.Reset()
				}
			}
		} else if fnF, ok := fnA.(*data.FrameFunc); ok {
			if fnF.Func() {
				result.Entity.RemoveComponent(myecs.Update)
			}
		}
	}
}