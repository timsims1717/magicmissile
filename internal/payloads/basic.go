package payloads

import (
	"fmt"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"image/color"
	"math"
	"math/rand"
	"timsims1717/magicmissile/internal/data"
	"timsims1717/magicmissile/internal/figures"
	"timsims1717/magicmissile/internal/myecs"
	"timsims1717/magicmissile/internal/states/game"
	"timsims1717/magicmissile/pkg/img"
	"timsims1717/magicmissile/pkg/object"
	"timsims1717/magicmissile/pkg/sfx"
	"timsims1717/magicmissile/pkg/timing"
)

func BasicMissile(start, target pixel.Vec, speed float64, rgba color.RGBA) {
	obj := object.New()
	obj.Pos = start
	obj.Rot = target.Sub(start).Angle()
	spr := &img.Sprite{
		Key:   "missile",
		Color: rgba,
		Batch: "figures",
	}
	hp := &data.Health{
		HP:   1,
		Team: data.NoTeam,
	}
	hitbox := pixel.R(-16., -3.5, 16, 3.5)
	myecs.Manager.NewEntity().
		AddComponent(myecs.Object, obj).
		AddComponent(myecs.Drawable, spr).
		AddComponent(myecs.Health, hp).
		AddComponent(myecs.Hitbox, &hitbox).
		AddComponent(myecs.Payload, &data.Missile{
			Target: target,
			Speed:  speed,
			//Finish: func(pos pixel.Vec) {
			//	BasicExplosion(obj.Pos, 30., 2., rgba)
			//},
		})
}

func MagicMissile(start, target pixel.Vec, speed float64, rgba color.RGBA) {
	for i := 0; i < 3; i++ {
		pos := target
		s := start
		spd := speed
		col := rgba
		e := myecs.Manager.NewEntity()
		e.AddComponent(myecs.Update, data.NewTimerFunc(func() bool {
			pos.X += rand.Float64()*40. - 20.
			pos.Y += rand.Float64()*40. - 20.
			BasicMissile(s, pos, spd, col)
			myecs.Manager.DisposeEntity(e)
			return false
		}, float64(i)*0.15))
	}
}

func Fireball(start, target pixel.Vec, speed float64) {
	obj := object.New()
	obj.Pos = start
	obj.Rot = target.Sub(start).Angle()
	col := color.RGBA{
		R: 223,
		G: 62,
		B: 35,
		A: 255,
	}
	spr := &img.Sprite{
		Key:   "missile",
		Color: col,
		Batch: "figures",
	}
	hp := &data.Health{
		HP:   1,
		Team: data.NoTeam,
	}
	hitbox := pixel.R(-16., -3.5, 16, 3.5)
	myecs.Manager.NewEntity().
		AddComponent(myecs.Object, obj).
		AddComponent(myecs.Drawable, spr).
		AddComponent(myecs.Health, hp).
		AddComponent(myecs.Hitbox, &hitbox).
		AddComponent(myecs.Payload, &data.Missile{
			Target: target,
			Speed:  speed,
			//Finish: func(pos pixel.Vec) {
			//	BasicExplosion(obj.Pos, 75., 2., col)
			//},
		})
}

func ChaosBolt(start, target pixel.Vec, speed float64, count int) {
	obj := object.New()
	obj.Pos = start
	obj.Rot = target.Sub(start).Angle()
	r := false
	g := false
	b := false
	c := rand.Intn(3)
	var R1, G1, B1 uint8
	if c == 0 {
		r = true
		B1 = 255
	} else if c == 1 {
		g = true
		R1 = 255
	} else {
		b = true
		G1 = 255
	}
	col := color.RGBA{
		R: R1,
		G: G1,
		B: B1,
		A: 255,
	}
	spr := &img.Sprite{
		Key:   "missile",
		Color: col,
		Batch: "figures",
	}
	hitbox := pixel.R(-16., -3.5, 16, 3.5)
	e := myecs.Manager.NewEntity().
		AddComponent(myecs.Object, obj).
		AddComponent(myecs.Drawable, spr).
		AddComponent(myecs.Hitbox, &hitbox).
		AddComponent(myecs.Update, data.NewFrameFunc(func() bool {
			add := int(timing.DT * 450.)
			R, G, B := int(spr.Color.R), int(spr.Color.G), int(spr.Color.B)
			if r {
				if R+add > 255 {
					R = 255
					r = false
					g = true
				} else {
					R += add
				}
				if G-add < 0 {
					G = 0
				} else {
					G -= add
				}
				if B-add < 0 {
					B = 0
				} else {
					B -= add
				}
			}
			if g {
				if G+add > 255 {
					G = 255
					g = false
					b = true
				} else {
					G += add
				}
				if R-add < 0 {
					R = 0
				} else {
					R -= add
				}
				if B-add < 0 {
					B = 0
				} else {
					B -= add
				}
			}
			if b {
				if B+add > 255 {
					B = 255
					b = false
					r = true
				} else {
					B += add
				}
				if R-add < 0 {
					R = 0
				} else {
					R -= add
				}
				if G-add < 0 {
					G = 0
				} else {
					G -= add
				}
			}
			spr.Color.R = uint8(R)
			spr.Color.G = uint8(G)
			spr.Color.B = uint8(B)
			return false
		})).
		AddComponent(myecs.Payload, &data.Missile{
			Target: target,
			Speed:  speed,
			//Finish: func(pos pixel.Vec) {
			//	if count < rand.Intn(6)+1 {
			//		tar := pos
			//		tar.X += rand.Float64()*150. - 75.
			//		tar.Y += rand.Float64()*150. - 75.
			//		if tar.Y < game.Frame.Min.Y {
			//			tar.Y = game.Frame.Min.Y
			//		}
			//		ChaosBolt(pos, tar, speed, count+1)
			//		BasicExplosion(obj.Pos, 20.+(2.*float64(count+1)), 2., spr.Color)
			//	} else {
			//		BasicExplosion(obj.Pos, 40., 2., spr.Color)
			//	}
			//},
		})
	if count == 0 {
		hp := &data.Health{
			HP:   1,
			Team: data.NoTeam,
		}
		e.AddComponent(myecs.Health, hp)
	}
}

func BasicExplosion(pos pixel.Vec, radius, expansion float64, rgba color.RGBA) {
	if expansion < 1.5 {
		sfx.SoundPlayer.PlaySound("explosion1", 0.)
	} else {
		sfx.SoundPlayer.PlaySound("explosion2", 0.)
	}
	obj := object.New()
	obj.Pos = pos
	exp := &data.Explosion{
		FullRadius: radius,
		ExpandRate: expansion,
	}
	myecs.Manager.NewEntity().
		AddComponent(myecs.Object, obj).
		AddComponent(myecs.Drawable, data.NewImdFunc("explosions", func(pos pixel.Vec, imd *imdraw.IMDraw) {
			imd.Color = rgba
			imd.Push(pos)
			imd.Circle(exp.CurrRadius, 0.)
			imd.Color = color.RGBA{
				R: 0,
				G: 0,
				B: 0,
				A: 255,
			}
			imd.Push(pos)
			imd.Circle(exp.CurrRadius, 1.)
		})).
		AddComponent(myecs.Payload, exp)
}

func BasicMeteor(spd float64, pos pixel.Vec) {
	obj := object.New()
	if pos == pixel.ZV {
		obj.Pos.X = float64(rand.Intn(1520) - 760)
		obj.Pos.Y = 460.
	} else {
		obj.Pos = pos
		obj.Pos.X += rand.Float64()*48. - 24.
		obj.Pos.Y += rand.Float64()*48. - 24.
	}
	obj.Rot = math.Pi*rand.Float64()*2. - 1.
	spr := &img.Sprite{
		Key: fmt.Sprintf("meteor_sm_%d", rand.Intn(2)),
		Color: color.RGBA{
			R: 255,
			G: 255,
			B: 255,
			A: 255,
		},
		Batch: "stuff",
	}
	var target pixel.Vec
	try := 0
	for {
		town := game.Towns[rand.Intn(len(game.Towns))]
		if !town.Health.Dead {
			target = town.Object.Pos
			break
		}
		try++
		if try > 12 {
			target = pixel.V(float64(rand.Intn(1520)-760), game.TownYLvl)
			break
		}
	}
	speed := spd
	hp := &data.Health{
		HP: 1,
	}
	hitbox := pixel.C(pixel.ZV, 16.)
	obj.Rot = math.Pi*rand.Float64()*2. - 1.
	rSpd := rand.Float64()*2. - 1.
	myecs.Manager.NewEntity().
		AddComponent(myecs.Object, obj).
		AddComponent(myecs.Drawable, spr).
		AddComponent(myecs.Health, hp).
		AddComponent(myecs.Hitbox, &hitbox).
		AddComponent(myecs.Update, data.NewFrameFunc(func() bool {
			obj.Rot += rSpd * timing.DT
			if obj.Rot > math.Pi {
				obj.Rot -= math.Pi * 2.
			} else if obj.Rot < -math.Pi {
				obj.Rot += math.Pi * 2.
			}
			return false
		})).
		AddComponent(myecs.Payload, &data.Missile{
			Target: target,
			Speed:  speed,
			//Finish: func(pos pixel.Vec) {
			//	BasicExplosion(obj.Pos, 60., 1., color.RGBA{
			//		R: 223,
			//		G: 62,
			//		B: 35,
			//		A: 255,
			//	})
			//},
		})
}

func BigMeteor(spd float64) {
	obj := object.New()
	obj.Pos.X = float64(rand.Intn(1520) - 760)
	obj.Pos.Y = 460.
	obj.Rot = math.Pi*rand.Float64()*2. - 1.
	spr := &img.Sprite{
		Key: "meteor_lrg",
		Color: color.RGBA{
			R: 255,
			G: 255,
			B: 255,
			A: 255,
		},
		Batch: "stuff",
	}
	var target pixel.Vec
	try := 0
	for {
		town := game.Towns[rand.Intn(len(game.Towns))]
		if !town.Health.Dead {
			target = town.Object.Pos
			break
		}
		try++
		if try > 12 {
			target = pixel.V(float64(rand.Intn(1520)-760), game.TownYLvl)
			break
		}
	}
	speed := spd
	hp := &data.Health{
		HP: 1,
	}
	hitbox := pixel.C(pixel.ZV, 32.)
	obj.Rot = math.Pi*rand.Float64()*2. - 1.
	rSpd := rand.Float64()*2. - 1.
	var breakUp *timing.Timer
	if rand.Intn(2) == 0 {
		breakUp = timing.New(rand.Float64()*3. + 2.)
	}
	e := myecs.Manager.NewEntity()
	e.AddComponent(myecs.Object, obj).
		AddComponent(myecs.Drawable, spr).
		AddComponent(myecs.Health, hp).
		AddComponent(myecs.Hitbox, &hitbox).
		AddComponent(myecs.Update, data.NewFrameFunc(func() bool {
			obj.Rot += rSpd * timing.DT
			if obj.Rot > math.Pi {
				obj.Rot -= math.Pi * 2.
			} else if obj.Rot < -math.Pi {
				obj.Rot += math.Pi * 2.
			}
			twnCnt := 0
			for _, twn := range game.Towns {
				if !twn.Health.Dead {
					twnCnt++
				}
			}
			if breakUp != nil && breakUp.UpdateDone() && twnCnt > 2 {
				cnt := rand.Intn(2) + 3
				for i := 0; i < cnt; i++ {
					BasicMeteor(spd, obj.Pos)
				}
				myecs.Manager.DisposeEntity(e)
			}
			return false
		})).
		AddComponent(myecs.Payload, &data.Missile{
			Target: target,
			Speed:  speed,
			//Finish: func(pos pixel.Vec) {
			//	BasicExplosion(obj.Pos, 90., 1., color.RGBA{
			//		R: 223,
			//		G: 62,
			//		B: 35,
			//		A: 255,
			//	})
			//},
		})
}

func BasicZombie() {
	game.ZCount++
	col := color.RGBA{
		R: 91,
		G: 149,
		B: 56,
		A: 255,
	}
	mob := &data.Mob{
		Speed: 50.,
		Char:  &data.Character{},
	}
	var twn *data.Town
	try := 0
	for {
		town := game.Towns[rand.Intn(len(game.Towns))]
		if !town.Health.Dead {
			twn = town
			break
		}
		try++
		if try > 12 {
			break
		}
	}
	obj := object.New()
	if twn != nil {
		obj.Pos.X = twn.Object.Pos.X
		for math.Abs(obj.Pos.X-twn.Object.Pos.X) < 250. {
			obj.Pos.X = float64(rand.Intn(1520) - 760)
		}
		mob.Target = twn
	} else if rand.Intn(2) == 0 {
		obj.Pos.X = 810.
	} else {
		obj.Pos.X = -810.
	}
	obj.Pos.Y = -475.
	mob.Char.Health = &data.Health{
		HP:   1,
		Team: data.Enemy,
	}
	mob.Char.Obj = obj
	mob.Attack = &data.Attack{
		WindUp:   1.,
		WindDown: 0.5,
		Recover:  2.,
		Damage:   1,
		Range:    20.,
		Team:     data.Enemy,
	}
	hitbox := pixel.R(-16., -32., 16., 32.)
	zombieArm := figures.ZombieArm(col)
	wobbleTimer := timing.New(0.8)
	dead := false
	var deadTimer *timing.Timer
	arm := myecs.Manager.NewEntity()
	e := myecs.Manager.NewEntity()
	arm.AddComponent(myecs.Parent, mob.Char.Obj).
		AddComponent(myecs.Object, zombieArm.Obj).
		AddComponent(myecs.Drawable, zombieArm.Spr).
		AddComponent(myecs.Update, data.NewFrameFunc(func() bool {
			if mob.Char.Health.Dead {
				if !dead {
					game.ZCount--
					sfx.SoundPlayer.PlaySound("zombie-hit", -1.)
					dead = true
					deadTimer = timing.New(1.)
				}
				if mob.Char.Obj.Rot < math.Pi*0.5 {
					mob.Char.Obj.Rot += 8. * timing.DT
					if mob.Char.Obj.Rot > math.Pi*0.5 {
						mob.Char.Obj.Rot = math.Pi * 0.5
					}
				}
				if mob.Char.Obj.Pos.Y > game.DeadYLvl {
					mob.Char.Obj.Pos.Y -= 160. * timing.DT
					if mob.Char.Obj.Pos.Y < game.DeadYLvl {
						mob.Char.Obj.Pos.Y = game.DeadYLvl
					}
				}
				if deadTimer.UpdateDone() {
					myecs.Manager.DisposeEntity(e)
					myecs.Manager.DisposeEntity(arm)
				}
				return false
			}
			if mob.Attack.Attacking {
				if mob.Attack.Target != nil {
					zombieArm.Obj.Rot += 2. * timing.DT
					if zombieArm.Obj.Rot > zombieArm.WindUp {
						zombieArm.Obj.Rot = zombieArm.WindUp
					}
				} else {
					zombieArm.Obj.Rot -= 8. * timing.DT
					if zombieArm.Obj.Rot < zombieArm.Strike {
						zombieArm.Obj.Rot = zombieArm.Strike
					}
				}
			} else {
				if zombieArm.Obj.Rot < zombieArm.Resting {
					zombieArm.Obj.Rot += 0.5 * timing.DT
					if zombieArm.Obj.Rot > zombieArm.Resting {
						zombieArm.Obj.Rot = zombieArm.Resting
					}
				} else if zombieArm.Obj.Rot > zombieArm.Resting {
					zombieArm.Obj.Rot -= 0.5 * timing.DT
					if zombieArm.Obj.Rot < zombieArm.Resting {
						zombieArm.Obj.Rot = zombieArm.Resting
					}
				}
			}
			if wobbleTimer.UpdateDone() {
				zombieArm.Resting *= -1.
				wobbleTimer.Reset()
			}
			return false
		}))
	e.AddComponent(myecs.Object, mob.Char.Obj).
		AddComponent(myecs.Mob, mob).
		AddComponent(myecs.Health, mob.Char.Health).
		AddComponent(myecs.Hitbox, &hitbox).
		AddComponent(myecs.Attack, mob.Attack).
		AddComponent(myecs.Drawable, figures.ZombieFigure(col))
}
