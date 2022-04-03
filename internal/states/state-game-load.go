package states

import (
	"github.com/faiface/pixel"
	"image/color"
	"math"
	"math/rand"
	"timsims1717/magicmissile/internal/data"
	"timsims1717/magicmissile/internal/figures"
	"timsims1717/magicmissile/internal/myecs"
	"timsims1717/magicmissile/internal/payloads"
	"timsims1717/magicmissile/internal/states/game"
	"timsims1717/magicmissile/pkg/img"
	"timsims1717/magicmissile/pkg/object"
	"timsims1717/magicmissile/pkg/timing"
)

func loadUI() {
	game.Cursor = data.TheInput.World
	spr := &img.Sprite{
		Key:   "cursor",
		Color: white,
		Batch: "figures",
	}
	myecs.Manager.NewEntity().
		AddComponent(myecs.Object, object.New()).
		AddComponent(myecs.Parent, &game.Cursor).
		AddComponent(myecs.Drawable, spr)
	selObj := object.New()
	selObj.Pos.Y = game.TownYLvl + 50.
	selSpr := &img.Sprite{
		Key:   "selected",
		Color: white,
		Batch: "figures",
	}
	myecs.Manager.NewEntity().
		AddComponent(myecs.Object, selObj).
		AddComponent(myecs.Drawable, selSpr).
		AddComponent(myecs.Update, data.NewFrameFunc(func() bool {
			selObj.Pos.X = game.PCs[game.Selected].Char.Obj.Pos.X
			selSpr.Color = game.PCs[game.Selected].Char.Spr.Color
			return false
		}))
}

func loadWizard() {
	wizCol := color.RGBA{
		R: 121,
		G: 58,
		B: 128,
		A: 255,
	}
	wPos := pixel.V(-90.,game.CharYLvl)
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
	spell := rand.Intn(3)
	e := myecs.Manager.NewEntity()
	wandArm := figures.WandArm(wizCol)
	var waitTimer *timing.Timer
	arm := myecs.Manager.NewEntity().
		AddComponent(myecs.Parent, pc.Char.Obj).
		AddComponent(myecs.Object, wandArm.Obj).
		AddComponent(myecs.Drawable, wandArm.Spr).
		AddComponent(myecs.Update, data.NewFrameFunc(func() bool {
			if waitTimer.UpdateDone() {
				pc.Move.Wait = false
			}
			angle := game.Cursor.Sub(pc.Char.Obj.Pos).Angle()
			if pc.Char.Obj.Flip {
				wandArm.Obj.Rot = math.Pi - angle
			} else {
				wandArm.Obj.Rot = angle
			}
			if data.TheInput.Get("click").JustPressed() && pc.Move.Selected {
				switch spell {
				case 0:
					payloads.MagicMissile(pc.Char.Obj.Pos, data.TheInput.World, 500., wizCol)
				case 1:
					payloads.ChaosBolt(pc.Char.Obj.Pos, data.TheInput.World, 500., 0)
				case 2:
					payloads.Fireball(pc.Char.Obj.Pos, data.TheInput.World, 500.)
				}
				spell = rand.Intn(3)
				waitTimer = timing.New(0.)
				pc.Move.Wait = true
			}
			return false
		}))
	hitbox := pixel.R(-16., -32., 16., 32.)
	e.AddComponent(myecs.Object, pc.Char.Obj).
		AddComponent(myecs.Health, pc.Char.Health).
		AddComponent(myecs.Hitbox, &hitbox).
		AddComponent(myecs.Movable, pc.Move).
		AddComponent(myecs.Drawable, pc.Char.Spr).
		AddComponent(myecs.Update, data.NewFrameFunc(func() bool {
			if pc.Char.Health.Dead || game.GameOver {
				if pc.Char.Obj.Flip {
					wandArm.Obj.Rot = math.Pi
				} else {
					wandArm.Obj.Rot = 0.
				}
				if pc.Char.Obj.Rot < math.Pi * 0.5 {
					pc.Char.Obj.Rot += 8. * timing.DT
					if pc.Char.Obj.Rot > math.Pi * 0.5 {
						pc.Char.Obj.Rot = math.Pi * 0.5
					}
				}
				if pc.Char.Obj.Pos.Y > game.DeadYLvl {
					pc.Char.Obj.Pos.Y -= 160. * timing.DT
					if pc.Char.Obj.Pos.Y < game.DeadYLvl {
						pc.Char.Obj.Pos.Y = game.DeadYLvl
					}
				}
				e.RemoveComponent(myecs.Movable)
				arm.RemoveComponent(myecs.Update)
			}
			return false
		}))
	game.PCs = append(game.PCs, pc)
}

func loadFighter() {
	fightCol := color.RGBA{
		R: 155,
		G: 23,
		B: 45,
		A: 255,
	}
	fPos := pixel.V(0.,game.CharYLvl)
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
		Range:     60.,
		Team:      data.Player,
	}
	e := myecs.Manager.NewEntity()
	axeArm := figures.AxeArm(fightCol)
	arm := myecs.Manager.NewEntity().
		AddComponent(myecs.Parent, pc.Char.Obj).
		AddComponent(myecs.Object, axeArm.Obj).
		AddComponent(myecs.Drawable, axeArm.Spr).
		AddComponent(myecs.Update, data.NewFrameFunc(func() bool {
			pc.Move.Wait = atk.Attacking
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
	e.AddComponent(myecs.Object, pc.Char.Obj).
		AddComponent(myecs.Health, pc.Char.Health).
		AddComponent(myecs.Hitbox, &hitbox).
		AddComponent(myecs.Drawable, pc.Char.Spr).
		AddComponent(myecs.Movable, pc.Move).
		AddComponent(myecs.Attack, atk).
		AddComponent(myecs.Update, data.NewFrameFunc(func() bool {
			if pc.Char.Health.Dead || game.GameOver {
				if pc.Char.Obj.Flip {
					axeArm.Obj.Rot = math.Pi
				} else {
					axeArm.Obj.Rot = 0.
				}
				if pc.Char.Obj.Rot < math.Pi * 0.5 {
					pc.Char.Obj.Rot += 8. * timing.DT
					if pc.Char.Obj.Rot > math.Pi * 0.5 {
						pc.Char.Obj.Rot = math.Pi * 0.5
					}
				}
				if pc.Char.Obj.Pos.Y > game.DeadYLvl {
					pc.Char.Obj.Pos.Y -= 160. * timing.DT
					if pc.Char.Obj.Pos.Y < game.DeadYLvl {
						pc.Char.Obj.Pos.Y = game.DeadYLvl
					}
				}
				arm.RemoveComponent(myecs.Update)
				e.RemoveComponent(myecs.Movable)
				e.RemoveComponent(myecs.Attack)
			}
			return false
		}))
	game.PCs = append(game.PCs, pc)
}

func loadTowns() {
	spr := &img.Sprite{
		Key:    "house1",
		Color:  white,
		Batch:  "stuff",
	}
	sprD := &img.Sprite{
		Key:    "house1dead",
		Color:  white,
		Batch:  "stuff",
	}
	for i := 0; i < 8; i++ {
		x := game.TownX * -0.5 + float64(i) * (game.TownX / 7.)
		y := game.TownYLvl
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
					e.RemoveComponent(myecs.Health)
					e.RemoveComponent(myecs.Hitbox)
					e.AddComponent(myecs.Drawable, sprD)
				}
				return false
			}))
		game.Towns = append(game.Towns, town)
	}
}