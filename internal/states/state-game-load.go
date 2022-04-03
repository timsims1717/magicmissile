package states

import (
	"github.com/faiface/pixel"
	"image/color"
	"math"
	"timsims1717/magicmissile/internal/data"
	"timsims1717/magicmissile/internal/figures"
	"timsims1717/magicmissile/internal/myecs"
	"timsims1717/magicmissile/internal/payloads"
	"timsims1717/magicmissile/internal/states/game"
	"timsims1717/magicmissile/pkg/img"
	"timsims1717/magicmissile/pkg/object"
	"timsims1717/magicmissile/pkg/timing"
)

func loadWizard() {
	wizCol := color.RGBA{
		R: 200,
		G: 0,
		B: 200,
		A: 255,
	}
	wPos := pixel.V(-90.,-375.)
	wObj := object.New()
	wObj.Pos = wPos
	pc := &data.PC{
		Char: &data.Character{
			Obj:    wObj,
			Spr:    figures.WizardFigure(wizCol),
			Health: &data.Health{
				HP:   4,
				Team: data.Player,
			},
		},
		Move: &data.Moving{
			Selected: true,
			Speed:    125.,
			Key:      "1",
		},
	}
	wandArm := figures.WandArm(wizCol)
	arm := myecs.Manager.NewEntity().
		AddComponent(myecs.Parent, pc.Char.Obj).
		AddComponent(myecs.Object, wandArm.Obj).
		AddComponent(myecs.Drawable, wandArm.Spr).
		AddComponent(myecs.Update, data.NewFrameFunc(func() bool {
			angle := data.TheInput.World.Sub(pc.Char.Obj.Pos).Angle()
			if pc.Char.Obj.Flip {
				wandArm.Obj.Rot = math.Pi - angle
			} else {
				wandArm.Obj.Rot = angle
			}
			if data.TheInput.Get("click").JustPressed() && pc.Move.Selected {
				payloads.BasicMissile(pc.Char.Obj.Pos, data.TheInput.World, 500., wizCol)
			}
			return false
		}))
	hitbox := pixel.R(-16., -32., 16., 32.)
	e := myecs.Manager.NewEntity()
	e.AddComponent(myecs.Object, pc.Char.Obj).
		AddComponent(myecs.Health, pc.Char.Health).
		AddComponent(myecs.Hitbox, &hitbox).
		AddComponent(myecs.Movable, pc.Move).
		AddComponent(myecs.Drawable, pc.Char.Spr).
		AddComponent(myecs.Update, data.NewFrameFunc(func() bool {
			if pc.Char.Health.Dead || game.GameOver {
				e.RemoveComponent(myecs.Movable)
				arm.RemoveComponent(myecs.Update)
			}
			return false
		}))
	game.PCs = append(game.PCs, pc)
}

func loadFighter() {
	fightCol := color.RGBA{
		R: 200,
		G: 180,
		B: 0,
		A: 255,
	}
	fPos := pixel.V(0.,-375.)
	fObj := object.New()
	fObj.Pos = fPos
	pc := &data.PC{
		Char: &data.Character{
			Obj:    fObj,
			Spr:    figures.FighterFigure(fightCol),
			Health: &data.Health{
				HP:   7,
				Team: data.Player,
			},
		},
		Move: &data.Moving{
			Selected: false,
			Speed:    125.,
			Key:      "2",
		},
	}
	atk := &data.Attack{
		WindUp:    0.2,
		WindDown:  0.5,
		Recover:   0.5,
		Damage:    1,
		Range:     40.,
		Team:      data.Player,
	}
	axeArm := figures.AxeArm(fightCol)
	myecs.Manager.NewEntity().
		AddComponent(myecs.Parent, pc.Char.Obj).
		AddComponent(myecs.Object, axeArm.Obj).
		AddComponent(myecs.Drawable, axeArm.Spr).
		AddComponent(myecs.Update, data.NewFrameFunc(func() bool {
			if atk.Attacking {
				if atk.Target != nil {
					axeArm.Obj.Rot += 8. * timing.DT
					if axeArm.Obj.Rot > axeArm.WindUp {
						axeArm.Obj.Rot = axeArm.WindUp
					}
				} else {
					axeArm.Obj.Rot -= 35. * timing.DT
					if axeArm.Obj.Rot < axeArm.Strike {
						axeArm.Obj.Rot = axeArm.Strike
					}
				}
			} else {
				if axeArm.Obj.Rot < axeArm.Resting {
					axeArm.Obj.Rot += 4. * timing.DT
					if axeArm.Obj.Rot > axeArm.Resting {
						axeArm.Obj.Rot = axeArm.Resting
					}
				} else if axeArm.Obj.Rot > axeArm.Resting {
					axeArm.Obj.Rot -= 4. * timing.DT
					if axeArm.Obj.Rot < axeArm.Resting {
						axeArm.Obj.Rot = axeArm.Resting
					}
				}
			}
			return false
		}))
	//game.Fighter.Base = base
	//game.Fighter.Arm = arm
	hitbox := pixel.R(-16., -32., 16., 32.)
	e := myecs.Manager.NewEntity()
	e.AddComponent(myecs.Object, pc.Char.Obj).
		AddComponent(myecs.Health, pc.Char.Health).
		AddComponent(myecs.Hitbox, &hitbox).
		AddComponent(myecs.Drawable, pc.Char.Spr).
		AddComponent(myecs.Movable, pc.Move).
		AddComponent(myecs.Attack, atk).
		AddComponent(myecs.Update, data.NewFrameFunc(func() bool {
			if pc.Char.Health.Dead || game.GameOver {
				e.RemoveComponent(myecs.Movable)
				e.RemoveComponent(myecs.Attack)
			}
			return false
		}))
	game.PCs = append(game.PCs, pc)
}

func loadTowns() {
	spr := &img.Sprite{
		Key:    "town",
		Color:  white,
		Batch:  "test",
	}
	for i := 0; i < 8; i++ {
		x := -700. + float64(i) * (1400. / 7.)
		y := -325.
		obj := object.New()
		obj.Pos = pixel.V(x, y)
		hp := &data.Health{
			HP:   5,
			Team: data.Player,
		}
		town := &data.Town{
			Health: hp,
			Obj:    obj,
		}
		hitbox := pixel.R(-2., -2., 2., 2.)
		e := myecs.Manager.NewEntity()
		e.AddComponent(myecs.Object, obj).
			AddComponent(myecs.Health, hp).
			AddComponent(myecs.Hitbox, &hitbox).
			AddComponent(myecs.Drawable, spr).
			AddComponent(myecs.Update, data.NewFrameFunc(func() bool {
				if hp.Dead {
					myecs.Manager.DisposeEntity(e)
				}
				return false
			}))
		game.Towns = append(game.Towns, town)
	}
}