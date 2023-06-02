package systems

import (
	"timsims1717/magicmissile/internal/data"
	"timsims1717/magicmissile/internal/myecs"
	"timsims1717/magicmissile/pkg/object"
	"timsims1717/magicmissile/pkg/timing"
)

func ObjectSystem() {
	for _, result := range myecs.Manager.Query(myecs.IsObject) {
		if obj, ok := result.Components[myecs.Object].(*object.Object); ok {
			if obj.Kill {
				myecs.Manager.DisposeEntity(result)
			} else {
				obj.Update()
			}
		}
	}
}

func ParentSystem() {
	for _, result := range myecs.Manager.Query(myecs.HasParent) {
		tran, okT := result.Components[myecs.Object].(*object.Object)
		parent, okP := result.Components[myecs.Parent].(*object.Object)
		if okT && okP {
			if parent.Kill {
				myecs.Manager.DisposeEntity(result)
			} else {
				tran.Pos = parent.Pos.Add(parent.Offset)
				tran.Hide = parent.Hide
			}
		}
	}
}

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
		} else if fnU, ok := fnA.(*data.Funky); ok {
			fnU.Fn()
		}
	}
}

func TemporarySystem() {
	for _, result := range myecs.Manager.Query(myecs.IsTemp) {
		temp := result.Components[myecs.Temp]
		obj, okT := result.Components[myecs.Object].(*object.Object)
		if okT {
			if timer, ok := temp.(*timing.Timer); ok {
				if timer.UpdateDone() {
					obj.Hide = true
					obj.Kill = true
					myecs.Manager.DisposeEntity(result.Entity)
				}
			} else if check, ok := temp.(myecs.ClearFlag); ok {
				if check {
					obj.Hide = true
					obj.Kill = true
					myecs.Manager.DisposeEntity(result.Entity)
				}
			}
		}
	}
}

func ClearSystem() {
	for _, result := range myecs.Manager.Query(myecs.IsObject) {
		obj, ok := result.Components[myecs.Object].(*object.Object)
		if ok {
			obj.Hide = true
			obj.Kill = true
		}
		myecs.Manager.DisposeEntity(result.Entity)
	}
}
