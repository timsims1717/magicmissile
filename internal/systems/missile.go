package systems

import (
	"fmt"
	"github.com/faiface/pixel"
	"golang.org/x/image/colornames"
	"image/color"
	"math/rand"
	"timsims1717/magicmissile/internal/data"
	"timsims1717/magicmissile/internal/myecs"
	"timsims1717/magicmissile/pkg/img"
	"timsims1717/magicmissile/pkg/object"
	"timsims1717/magicmissile/pkg/timing"
	"timsims1717/magicmissile/pkg/util"
)

func FireFromTower(mFab *data.Missile, tower *data.Tower, target pixel.Vec) {
	if mFab != nil {
		FireSpell(mFab, tower.Object.Pos.Add(tower.Origin), target)
	} else {
		fmt.Println("Warning: no missile to fire")
	}
}

func FireSpell(mFab *data.Missile, origin, target pixel.Vec) {
	if mFab != nil {
		count := 1
		if mFab.Count > 1 {
			count = mFab.Count
		}
		for i := 0; i < count; i++ {
			if mFab.Delay > 0 {
				e := myecs.Manager.NewEntity()
				e.AddComponent(myecs.Update, data.NewTimerFunc(func() bool {
					MakeMissile(mFab, origin, target)
					myecs.Manager.DisposeEntity(e)
					return false
				}, float64(i)*mFab.Delay))
			} else {
				MakeMissile(mFab, origin, target)
			}
		}
	} else {
		fmt.Println("Warning: no missile to fire")
	}
}

func MakeMissile(mFab *data.Missile, origin, target pixel.Vec) {
	if mFab != nil {
		var col color.RGBA
		if len(mFab.Colors) > 0 {
			col, _ = util.ParseHexColorFast(mFab.Colors[rand.Intn(len(mFab.Colors))])
		} else {
			col = colornames.Black
		}
		spr := &img.Sprite{
			Key:   mFab.SprKey,
			Batch: data.ParticleKey,
			Color: col,
		}
		obj := object.New()
		obj.Pos = origin
		obj.Rot = target.Sub(obj.Pos).Angle()
		obj.Layer = 10
		obj.Rect = img.Batchers[data.ParticleKey].GetSprite(spr.Key).Frame()
		if mFab.Spread > 0 {
			// todo: change to circle instead of square
			target.X += rand.Float64()*mFab.Spread*2 - mFab.Spread
			target.Y += rand.Float64()*mFab.Spread*2 - mFab.Spread
		}
		m := &data.Missile{
			Object:  obj,
			Sprite:  spr,
			Target:  target,
			Speed:   mFab.Speed,
			Payload: mFab.Payload,
		}
		myecs.Manager.NewEntity().
			AddComponent(myecs.Object, obj).
			AddComponent(myecs.Drawable, m.Sprite).
			AddComponent(myecs.Missile, m)
	} else {
		fmt.Println("Warning: no missile to fire")
	}
}

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
				for _, f := range missile.Payload {
					if f.Missile != nil {
						target := obj.Pos.Add(f.Missile.Target)
						MakeMissile(f.Missile, obj.Pos, target)
					}
					if f.Explosion != nil {
						MakeExplosion(f.Explosion, obj.Pos, missile.Sprite.Color)
					}
					if f.Spell != nil && *f.Spell != "" {
						if mSet, okS := data.Missiles[*f.Spell]; okS {
							for _, ms := range mSet {
								if ms != nil {
									target := obj.Pos.Add(ms.Target)
									MakeMissile(ms, obj.Pos, target)
									break
								}
							}
						} else {
							fmt.Printf("Warning: spell %s not found\n", *f.Spell)
						}
					}
					if f.Function != nil {
						f.Function()
					}
				}
				myecs.Manager.DisposeEntity(result)
			}
		}
	}
}
