package systems

import (
	"timsims1717/magicmissile/internal/data"
	"timsims1717/magicmissile/internal/myecs"
	"timsims1717/magicmissile/pkg/img"
	"timsims1717/magicmissile/pkg/object"
	"timsims1717/magicmissile/pkg/timing"
	"timsims1717/magicmissile/pkg/util"
)

func MissileSystem() {
	for _, result := range myecs.Manager.Query(myecs.IsMissile) {
		obj, ok := result.Components[myecs.Object].(*object.Object)
		missile, okM := result.Components[myecs.Missile].(*data.Missile)
		if ok && okM {
			l := missile.Target.X > obj.Pos.X
			b := missile.Target.Y > obj.Pos.Y
			norm := util.Normalize(missile.Target.Sub(obj.Pos))
			obj.Pos.X += norm.X * missile.Speed * timing.DT
			obj.Pos.Y += norm.Y * missile.Speed * timing.DT
			if obj.Kill || l == (missile.Target.X < obj.Pos.X) || b == (missile.Target.Y < obj.Pos.Y) {
				// missile reached target or was destroyed
				for _, f := range missile.Finish {
					if m, ok1 := f.(*data.Missile); ok1 {
						obj1 := object.New()
						obj1.Pos = obj.Pos
						obj1.Layer = obj.Layer
						obj1.Rect = img.Batchers[data.ObjectKey].GetSprite(m.Sprite.Key).Frame()
						m.Target = obj.Pos.Add(m.Target)
						m.Object = obj1
						myecs.Manager.NewEntity().
							AddComponent(myecs.Object, obj).
							AddComponent(myecs.Drawable, m.Sprite).
							AddComponent(myecs.Missile, m)
					} else if e, ok2 := f.(*data.Explosion); ok2 {
						myecs.Manager.NewEntity().
							AddComponent(myecs.Object, obj).
							AddComponent(myecs.Explosion, e)
					}
				}
				myecs.Manager.DisposeEntity(result)
			}
		}
	}
}
