package systems

import (
	"github.com/faiface/pixel"
	"timsims1717/magicmissile/internal/data"
	"timsims1717/magicmissile/internal/myecs"
	"timsims1717/magicmissile/pkg/object"
	"timsims1717/magicmissile/pkg/timing"
	"timsims1717/magicmissile/pkg/util"
)

func PayloadSystem() {
	for _, result := range myecs.Manager.Query(myecs.HasPayload) {
		if obj, ok := result.Components[myecs.Object].(*object.Object); ok {
			if missile, okM := result.Components[myecs.Payload].(*data.Missile); okM {
				l := missile.Target.X > obj.Pos.X
				b := missile.Target.Y > obj.Pos.Y
				norm := util.Normalize(missile.Target.Sub(obj.Pos))
				obj.Pos.X += norm.X * missile.Speed * timing.DT
				obj.Pos.Y += norm.Y * missile.Speed * timing.DT
				if obj.Kill || l == (missile.Target.X < obj.Pos.X) || b == (missile.Target.Y < obj.Pos.Y) {
					// missile reached target or was destroyed
					if missile.Finish != nil {
						missile.Finish(obj.Pos)
					}
					myecs.Manager.DisposeEntity(result)
				}
			} else if explosion, okE := result.Components[myecs.Payload].(*data.Explosion); okE {
				if explosion.ExpandRate > 0. && explosion.FullRadius == explosion.CurrRadius {
					// explosion is maximized
					explosion.ExpandRate *= -1.
				} else if explosion.ExpandRate < 0. && explosion.CurrRadius == 0. {
					// explosion is done
					myecs.Manager.DisposeEntity(result)
					continue
				} else {
					explosion.CurrRadius += explosion.ExpandRate * explosion.FullRadius * timing.DT
					if explosion.CurrRadius > explosion.FullRadius {
						explosion.CurrRadius = explosion.FullRadius
					}
					if explosion.CurrRadius < 0. {
						explosion.CurrRadius = 0.
					}
				}
				for _, hpResult := range myecs.Manager.Query(myecs.HasHealth) {
					if objHP, okHP := hpResult.Components[myecs.Object].(*object.Object); okHP {
						hitbox := hpResult.Components[myecs.Hitbox]
						c, okC := hitbox.(*pixel.Circle)
						r, okR := hitbox.(*pixel.Rect)
						rng := pixel.C(obj.Pos, explosion.CurrRadius)
						if (okC && rng.Intersect(c.Moved(objHP.Pos)).Radius > 0.) ||
							(okR && rng.IntersectRect(r.Moved(objHP.Pos)) != pixel.ZV) {
							if hp, okB := hpResult.Components[myecs.Health].(*data.Health); okB {
								hp.Dead = true
							}
						}
					}
				}
			}
		}
	}
}
