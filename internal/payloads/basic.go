package payloads

import (
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
	"timsims1717/magicmissile/pkg/timing"
)

func BasicMissile(start, target pixel.Vec, speed float64, rgba color.RGBA) {
	obj := object.New()
	obj.Pos = start
	obj.Rot = target.Sub(start).Angle()
	spr := &img.Sprite{
		Key:    "missile",
		Color:  rgba,
		Batch:  "test",
	}
	hp := &data.Health{
		HP:   1,
		Team: data.NoTeam,
	}
	hitbox := pixel.R(-7.5, -2.5, 7.5, 2.5)
	myecs.Manager.NewEntity().
		AddComponent(myecs.Object, obj).
		AddComponent(myecs.Drawable, spr).
		AddComponent(myecs.Health, hp).
		AddComponent(myecs.Hitbox, &hitbox).
		AddComponent(myecs.Payload, &data.Missile{
			Target: target,
			Speed:  speed,
			Finish: func(pos pixel.Vec) {
				BasicExplosion(obj.Pos, 30., 2., rgba)
			},
		})
}

func BasicExplosion(pos pixel.Vec, radius, expansion float64, rgba color.RGBA) {
	obj := object.New()
	obj.Pos = pos
	exp := &data.Explosion{
		Radius:    radius,
		Expansion: expansion,
	}
	myecs.Manager.NewEntity().
		AddComponent(myecs.Object, obj).
		AddComponent(myecs.Drawable, data.NewImdFunc("explosions", func(pos pixel.Vec, imd *imdraw.IMDraw) {
			imd.Color = rgba
			imd.Push(pos)
			imd.Circle(exp.CurrRadius, 0.)
		})).
		AddComponent(myecs.Payload, exp)
}

func BasicMeteor() {
	obj := object.New()
	obj.Pos.X = float64(rand.Intn(1520) - 760)
	obj.Pos.Y = 460.
	obj.Rot = math.Pi * rand.Float64() * 2. - 1.
	spr := &img.Sprite{
		Key:    "meteor",
		Color:  color.RGBA{
			R: 255,
			G: 255,
			B: 255,
			A: 255,
		},
		Batch:  "test",
	}
	var target pixel.Vec
	try := 0
	for {
		town := game.Towns[rand.Intn(len(game.Towns))]
		if !town.Health.Dead {
			target = town.Obj.Pos
			break
		}
		try++
		if try > 12 {
			target = pixel.V(float64(rand.Intn(1520)-760), -325.)
			break
		}
	}
	speed := 50.
	hp := &data.Health{
		HP: 1,
	}
	hitbox := pixel.C(pixel.ZV, 8.5)
	myecs.Manager.NewEntity().
		AddComponent(myecs.Object, obj).
		AddComponent(myecs.Drawable, spr).
		AddComponent(myecs.Health, hp).
		AddComponent(myecs.Hitbox, &hitbox).
		AddComponent(myecs.Payload, &data.Missile{
			Target: target,
			Speed:  speed,
			Finish: func(pos pixel.Vec) {
				BasicExplosion(obj.Pos, 50., 2., color.RGBA{
					R: 200,
					G: 100,
					B: 0,
					A: 255,
				})
			},
		})
}

func BasicZombie() {
	col := color.RGBA{
		R: 60,
		G: 200,
		B: 40,
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
		obj.Pos.X = twn.Obj.Pos.X
		for math.Abs(obj.Pos.X-twn.Obj.Pos.X) < 250. {
			obj.Pos.X = float64(rand.Intn(1520)-760)
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
		WindUp:    1.,
		WindDown:  0.5,
		Recover:   2.,
		Damage:    1,
		Range:     20.,
		Team:      data.Enemy,
	}
	hitbox := pixel.R(-16., -32., 16., 32.)
	zombieArm := figures.ZombieArm(col)
	wobbleTimer := timing.New(0.8)
	arm := myecs.Manager.NewEntity()
	arm.AddComponent(myecs.Parent, mob.Char.Obj).
		AddComponent(myecs.Object, zombieArm.Obj).
		AddComponent(myecs.Drawable, zombieArm.Spr).
		AddComponent(myecs.Update, data.NewFrameFunc(func() bool {
			if mob.Char.Health.Dead {
				myecs.Manager.DisposeEntity(arm)
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
	myecs.Manager.NewEntity().
		AddComponent(myecs.Object, mob.Char.Obj).
		AddComponent(myecs.Mob, mob).
		AddComponent(myecs.Health, mob.Char.Health).
		AddComponent(myecs.Hitbox, &hitbox).
		AddComponent(myecs.Attack, mob.Attack).
		AddComponent(myecs.Drawable, figures.ZombieFigure(col))
}