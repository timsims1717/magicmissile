package systems

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"timsims1717/magicmissile/internal/data"
	"timsims1717/magicmissile/internal/myecs"
	"timsims1717/magicmissile/pkg/img"
	"timsims1717/magicmissile/pkg/object"
	"timsims1717/magicmissile/pkg/reanimator"
)

func AnimationSystem() {
	for _, result := range myecs.Manager.Query(myecs.HasAnim) {
		obj, okO := result.Components[myecs.Object].(*object.Object)
		anim, ok := result.Components[myecs.Animated].(*reanimator.Tree)
		if okO && ok && !obj.Hide {
			anim.Update()
		}
	}
}

func DrawSystem(win *pixelgl.Window) {
	for _, result := range myecs.Manager.Query(myecs.IsDrawable) {
		obj, okO := result.Components[myecs.Object].(*object.Object)
		if okO {
			draw := result.Components[myecs.Drawable]
			if spr, ok0 := draw.(*pixel.Sprite); ok0 {
				spr.Draw(win, obj.Mat)
			} else if sprH, ok1 := draw.(*img.Sprite); ok1 {
				if batch, okB := img.Batchers[sprH.Batch]; okB {
					batch.DrawSpriteColor(sprH.Key, obj.Mat, sprH.Color)
				}
			} else if anim, ok2 := draw.(*reanimator.Tree); ok2 {
				if batch, okB := img.Batchers[anim.Batch]; okB {
					anim.Draw(batch.Batch(), obj.Mat)
				}
			} else if fn, okF := draw.(*data.ImdFunc); okF {
				fn.Func(obj.Pos, img.IMDrawers[fn.Key].IMD())
			}
		}
	}
}