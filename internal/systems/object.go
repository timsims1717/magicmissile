package systems

import (
	"github.com/faiface/pixel"
	"timsims1717/magicmissile/internal/myecs"
	"timsims1717/magicmissile/pkg/object"
)

func FullTransformSystem() {
	for _, result := range myecs.Manager.Query(myecs.IsObject) {
		if trans, ok := result.Components[myecs.Object].(*object.Object); ok {
			trans.Mat = pixel.IM
		}
	}
	PreParentSystem()
	TransformSystem()
	PostParentSystem()
}

func PreParentSystem() {
	for _, result := range myecs.Manager.Query(myecs.HasParent) {
		trans, okT := result.Components[myecs.Object].(*object.Object)
		parent, okP := result.Components[myecs.Parent].(*object.Object)
		if okT && okP {
			if parent.Flip != trans.Flip {
				trans.Flip = parent.Flip
				trans.Pos.X *= -1.
			}
			if parent.Flop != trans.Flop {
				trans.Flop = parent.Flop
				trans.Pos.Y *= -1.
			}
		}
	}
}

func TransformSystem() {
	for _, result := range myecs.Manager.Query(myecs.IsObject) {
		if trans, ok := result.Components[myecs.Object].(*object.Object); ok {
			trans.PostPos = trans.Pos.Add(trans.Offset)
			//trans.APos.X = math.Round(trans.APos.X)
			//trans.APos.Y = math.Round(trans.APos.Y)
			trans.Mat = trans.Mat.ScaledXY(pixel.ZV, trans.Sca)
			//trans.Mat = trans.Mat.Rotated(trans.RotArnd, trans.Rot)
			if trans.Flip && trans.Flop {
				trans.Mat = trans.Mat.Scaled(pixel.ZV, -1.)
			} else if trans.Flip {
				trans.Mat = trans.Mat.ScaledXY(pixel.ZV, pixel.V(-1., 1.))
			} else if trans.Flop {
				trans.Mat = trans.Mat.ScaledXY(pixel.ZV, pixel.V(1., -1.))
			}
			trans.Mat = trans.Mat.Moved(trans.PostPos)
		}
	}
}

func PostParentSystem() {
	for _, result := range myecs.Manager.Query(myecs.HasParent) {
		trans, okT := result.Components[myecs.Object].(*object.Object)
		parent := result.Components[myecs.Parent]
		if okT {
			if pos, ok := parent.(*pixel.Vec); ok {
				trans.Mat = trans.Mat.Moved(*pos)
			} else if pos, ok := parent.(pixel.Vec); ok {
				// using a non-pointer to a pixel.Vec will freeze the item forever
				trans.Mat = trans.Mat.Moved(pos)
			} else if par, ok := parent.(*object.Object); ok {
				trans.Mat = trans.Mat.Moved(par.Pos)
			}
		}
	}
}
