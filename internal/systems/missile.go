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
		origin := tower.Object.Pos.Add(tower.Origin)
		if mFab.Target != pixel.ZV {
			target = origin.Add(mFab.Target)
		}
		FireSpell(mFab, origin, target)
	} else {
		fmt.Println("Warning: no missile to fire")
	}
}

func FireNextFromTower(tower *data.Tower, target pixel.Vec) {
	if tower.CurrSlot < data.SpellSlotNum {
		slot := tower.Slots[tower.CurrSlot]
		mSet := data.Missiles[slot.Spell]
		var mFab *data.Missile
		for _, mf := range mSet {
			if mf.Tier == slot.Tier {
				mFab = mf
			}
		}
		tower.CurrSlot++
		if mFab != nil {
			origin := tower.Object.Pos.Add(tower.Origin)
			if mFab.Target != pixel.ZV {
				target = origin.Add(mFab.Target)
			}
			FireSpell(mFab, origin, target)
		} else {
			fmt.Println("Warning: no missile to fire")
		}
	} else {
		fmt.Println("tower out of spells")
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
			} else if mFab.Arc > 0 {
				ray := target.Sub(origin)
				r := util.Magnitude(ray)
				if r < mFab.Arc {
					r = mFab.Arc
				}
				t := mFab.Arc / r
				tSeg := t * 2 / float64(count-1)
				t -= tSeg * float64(i)
				n := util.Normalize(ray)
				rot := n.Rotated(t)
				nTarget := rot.Scaled(r).Add(origin)
				MakeMissile(mFab, origin, nTarget)
			} else if mFab.Angle > 0 {
				ray := target.Sub(origin)
				r := util.Magnitude(ray)
				t := mFab.Angle
				tSeg := t * 2 / float64(count-1)
				t -= tSeg * float64(i)
				n := util.Normalize(ray)
				rot := n.Rotated(t)
				nTarget := rot.Scaled(r).Add(origin)
				MakeMissile(mFab, origin, nTarget)
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
		//if mFab.Trans != 0 && mFab.Trans != 255 {
		//	col.A = mFab.Trans
		//}
		spr := &img.Sprite{
			Key:   mFab.SprKey,
			Batch: data.ParticleKey,
			Color: col,
		}
		obj := object.New()
		obj.Pos = origin
		obj.Rot = target.Sub(obj.Pos).Angle()
		obj.Layer = 10
		if spr.Key != "" {
			obj.SetRect(img.Batchers[data.ParticleKey].GetSprite(spr.Key).Frame())
		}
		if mFab.Spread > 0 {
			// todo: change to circle instead of square
			target.X += rand.Float64()*mFab.Spread*2 - mFab.Spread
			target.Y += rand.Float64()*mFab.Spread*2 - mFab.Spread
		}
		m := &data.Missile{
			SprKey:  mFab.SprKey,
			Object:  obj,
			Sprite:  spr,
			Colors:  mFab.Colors,
			Tier:    mFab.Tier,
			Limit:   mFab.Limit,
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
			if missile.Speed == 0. {
				obj.Pos = missile.Target
			} else {
				norm := util.Normalize(missile.Target.Sub(obj.Pos))
				obj.Pos.X += norm.X * missile.Speed * timing.DT
				obj.Pos.Y += norm.Y * missile.Speed * timing.DT
				missile.Travel += util.Magnitude(norm) * missile.Speed * timing.DT
			}
			if data.GameView.PointInside(obj.Pos.Scaled(0.98)) {
				if obj.Killed || (missile.Limit > 0 && missile.Travel > missile.Limit) || missile.Speed == 0. ||
					l == (missile.Target.X <= obj.Pos.X) || b == (missile.Target.Y <= obj.Pos.Y) {
					// missile reached target or was destroyed
					for _, f := range missile.Payload {
						if f.Missile != nil {
							target := obj.Pos.Add(f.Missile.Target)
							MakeMissile(f.Missile, obj.Pos, target)
						}
						if f.Explosion != nil {
							MakeExplosion(f.Explosion, obj.Pos, missile.Sprite.Color)
						}
						if f.Function != nil {
							f.Function(missile, obj)
						}
					}
					myecs.Manager.DisposeEntity(result)
				}
			} else {
				myecs.Manager.DisposeEntity(result)
			}
		}
	}
}
