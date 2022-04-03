package systems

import (
	"github.com/faiface/pixel"
	"timsims1717/magicmissile/internal/data"
	"timsims1717/magicmissile/internal/myecs"
	"timsims1717/magicmissile/pkg/object"
	"timsims1717/magicmissile/pkg/timing"
)

func AttackSystem() {
	for _, result := range myecs.Manager.Query(myecs.CanAttack) {
		obj, ok := result.Components[myecs.Object].(*object.Object)
		atk, okA := result.Components[myecs.Attack].(*data.Attack)
		if ok && okA {
			atk.Timer.Update()
			if !atk.Attacking && atk.Timer.Done() {
				for _, resultH := range myecs.Manager.Query(myecs.HasHealth) {
					tarObj, ok1 := resultH.Components[myecs.Object].(*object.Object)
					hp, ok2 := resultH.Components[myecs.Health].(*data.Health)
					hitbox := resultH.Components[myecs.Hitbox]
					if ok1 && ok2 && (atk.Team == 0 || (atk.Team != hp.Team && hp.Team != 0)) {
						c, okC := hitbox.(*pixel.Circle)
						r, okR := hitbox.(*pixel.Rect)
						rng := pixel.C(obj.Pos, atk.Range)
						if (okC && rng.Intersect(c.Moved(tarObj.Pos)).Radius > 0.) ||
							(okR && rng.IntersectRect(r.Moved(tarObj.Pos)) != pixel.ZV) {
							atk.Attacking = true
							atk.Target = resultH.Entity
							atk.Timer = timing.New(atk.WindUp)
						}
					}
				}
			} else if atk.Attacking && atk.Target != nil && atk.Timer.Done() {
				tarObjC, ok1 := atk.Target.GetComponentData(myecs.Object)
				hpC, ok2 := atk.Target.GetComponentData(myecs.Health)
				hitbox, ok3 := atk.Target.GetComponentData(myecs.Hitbox)
				if ok1 && ok2 && ok3 {
					tarObj, ok1 := tarObjC.(*object.Object)
					hp, ok2 := hpC.(*data.Health)
					if ok1 && ok2 && (atk.Team == 0 || (atk.Team != hp.Team && hp.Team != 0)) {
						c, okC := hitbox.(*pixel.Circle)
						r, okR := hitbox.(*pixel.Rect)
						rng := pixel.C(obj.Pos, atk.Range)
						if (okC && rng.Intersect(c.Moved(tarObj.Pos)).Radius > 0.) ||
							(okR && rng.IntersectRect(r.Moved(tarObj.Pos)) != pixel.ZV) {
							hp.HP -= atk.Damage
						}
					}
				}
				atk.Target = nil
				atk.Timer = timing.New(atk.WindDown)
			} else if atk.Attacking && atk.Target == nil && atk.Timer.Done() {
				atk.Attacking = false
				atk.Timer = timing.New(atk.Recover)
			}
			if atk.Target != nil {
				if tarObjC, ok1 := atk.Target.GetComponentData(myecs.Object); ok1 {
					if tarObj, ok1 := tarObjC.(*object.Object); ok1 {
						obj.Flip = tarObj.Pos.X < obj.Pos.X
					}
				}
			}
		}
	}
}