package systems

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"timsims1717/magicmissile/internal/myecs"
	"timsims1717/magicmissile/pkg/img"
	"timsims1717/magicmissile/pkg/object"
	"timsims1717/magicmissile/pkg/reanimator"
)

func AnimationSystem() {
	for _, result := range myecs.Manager.Query(myecs.HasAnimation) {
		obj, okO := result.Components[myecs.Object].(*object.Object)
		theAnim := result.Components[myecs.Animated]
		if okO && !obj.Hide {
			if theAnim == nil {
				continue
			} else if anims, okS := theAnim.([]*reanimator.Tree); okS {
				for _, anim := range anims {
					anim.Update()
				}
			} else if anim, okA := theAnim.(*reanimator.Tree); okA {
				anim.Update()
			}
		}
	}
}

func DrawSystem(win *pixelgl.Window, layer int) {
	count := 0
	for _, result := range myecs.Manager.Query(myecs.IsDrawable) {
		obj, okO := result.Components[myecs.Object].(*object.Object)
		if okO && obj.Layer == layer && !obj.Hide {
			draw := result.Components[myecs.Drawable]
			if draw == nil {
				continue
			} else if draws, okD := draw.([]*img.Sprite); okD {
				for _, d := range draws {
					DrawThing(d, obj, win)
					count++
				}
			} else if anims, okA := draw.([]*reanimator.Tree); okA {
				for _, d := range anims {
					DrawThing(d, obj, win)
					count++
				}
			} else {
				DrawThing(draw, obj, win)
				count++
			}
		}
	}
	//debug.AddText(fmt.Sprintf("Layer %d: %d entities", layer, count))
}

func DrawThing(draw interface{}, obj *object.Object, win *pixelgl.Window) {
	if spr, ok0 := draw.(*pixel.Sprite); ok0 {
		spr.Draw(win, obj.Mat)
	} else if sprH, ok1 := draw.(*img.Sprite); ok1 {
		if batch, okB := img.Batchers[sprH.Batch]; okB {
			batch.DrawSpriteColor(sprH.Key, obj.Mat.Moved(sprH.Offset), sprH.Color)
		}
	} else if anim, ok2 := draw.(*reanimator.Tree); ok2 {
		res := anim.CurrentSprite()
		if res != nil {
			if _, okB := img.Batchers[res.Batch]; okB {
				res.Spr.DrawColorMask(img.Batchers[res.Batch].Batch(), obj.Mat.Moved(res.Off), res.Col)
			}
		}
	}
}
