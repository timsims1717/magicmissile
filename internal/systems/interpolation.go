package systems

import (
	"fmt"
	"timsims1717/magicmissile/internal/myecs"
	"timsims1717/magicmissile/pkg/object"
	"timsims1717/magicmissile/pkg/timing"
)

func InterpolationSystem() {
	for _, result := range myecs.Manager.Query(myecs.HasInterpolation) {
		obj, okO := result.Components[myecs.Object].(*object.Object)
		if okO {
			inter := result.Components[myecs.Interpolation]
			if interV, okV := inter.(*object.Interpolation); okV {
				if Interpolate(obj, interV) {
					result.Entity.RemoveComponent(myecs.Interpolation)
				}
			} else if interA, okA := inter.([]*object.Interpolation); okA {
				fin := true
				for _, interV := range interA {
					if !Interpolate(obj, interV) {
						fin = false
					}
				}
				if fin {
					result.Entity.RemoveComponent(myecs.Interpolation)
				}
			}
		}
	}
}

func Interpolate(obj *object.Object, inter *object.Interpolation) bool {
	cur, _, fin := inter.Sequence.Update(timing.DT)
	switch inter.Target {
	case object.InterpolateX:
		obj.Pos.X = cur
	case object.InterpolateY:
		obj.Pos.Y = cur
	case object.InterpolateOffX:
		obj.Offset.X = cur
	case object.InterpolateOffY:
		obj.Offset.Y = cur
	case object.InterpolateRot:
		obj.Rot = cur
	case object.InterpolateSX:
		obj.Sca.X = cur
	case object.InterpolateSY:
		obj.Sca.Y = cur
	case object.InterpolateR:
		obj.Mask.R = uint8(cur)
	case object.InterpolateG:
		obj.Mask.G = uint8(cur)
	case object.InterpolateB:
		obj.Mask.B = uint8(cur)
	case object.InterpolateA:
		obj.Mask.A = uint8(cur)
	case object.InterpolateCustom:
		if inter.Value == nil {
			panic(fmt.Sprintf("interpolate custom on a nil value for object %s", obj.ID))
		}
		*inter.Value = cur
	}
	if fin && inter.OnComplete != nil {
		inter.OnComplete()
	}
	return fin
}
