package systems

import (
	"timsims1717/magicmissile/internal/data"
	"timsims1717/magicmissile/internal/myecs"
	"timsims1717/magicmissile/internal/states/game"
	"timsims1717/magicmissile/pkg/object"
	"timsims1717/magicmissile/pkg/timing"
)

func ControlSystem() {
	for i, pc := range game.PCs {
		if data.TheInput.Get(pc.Move.Key).JustPressed() && !pc.Char.Health.Dead {
			game.Selected = i
			break
		}
	}
	if data.TheInput.Get("switchR").JustPressed() || game.PCs[game.Selected].Char.Health.Dead {
		selected := game.Selected
		first := true
		for selected != game.Selected || first {
			if first {
				first = false
			} else {
				pc := game.PCs[selected]
				if !pc.Char.Health.Dead {
					game.Selected = selected
					break
				}
			}
			selected++
			if selected > len(game.PCs)-1 {
				selected = 0
			}
		}
	} else if data.TheInput.Get("switchL").JustPressed() {
		selected := game.Selected
		first := true
		for selected != game.Selected || first {
			if first {
				first = false
			} else {
				pc := game.PCs[selected]
				if !pc.Char.Health.Dead {
					game.Selected = selected
					break
				}
			}
			selected--
			if selected < 0 {
				selected = len(game.PCs)-1
			}
		}
	}
	for i, pc := range game.PCs {
		pc.Move.Selected = game.Selected == i
	}
	for _, result := range myecs.Manager.Query(myecs.PlayerChar) {
		obj, ok := result.Components[myecs.Object].(*object.Object)
		move, okM := result.Components[myecs.Movable].(*data.Moving)
		if ok && okM {
			if move.Selected && !move.Wait {
				moving := false
				if data.TheInput.Get("moveLeft").Pressed() {
					obj.Pos.X -= move.Speed * timing.DT
					if obj.Pos.X < game.Frame.Min.X {
						obj.Pos.X = game.Frame.Min.X
					}
					obj.Flip = true
					moving = true
				} else if data.TheInput.Get("moveRight").Pressed() {
					obj.Pos.X += move.Speed * timing.DT
					if obj.Pos.X > game.Frame.Max.X {
						obj.Pos.X = game.Frame.Max.X
					}
					obj.Flip = false
					moving = true
				}
				if moving || obj.Pos.Y != game.CharYLvl {
					if moving && move.Up {
						if obj.Pos.Y < game.MoveYLvl {
							obj.Pos.Y += 150. * timing.DT
							if obj.Pos.Y > game.MoveYLvl {
								obj.Pos.Y = game.MoveYLvl
								move.Up = false
							}
						}
					} else if obj.Pos.Y > game.CharYLvl {
						obj.Pos.Y -= 150. * timing.DT
						if obj.Pos.Y < game.CharYLvl {
							obj.Pos.Y = game.CharYLvl
							move.Up = moving
						}
					} else {
						move.Up = moving
					}
				}
			} else if obj.Pos.Y > game.CharYLvl {
				obj.Pos.Y -= 150. * timing.DT
				if obj.Pos.Y < game.CharYLvl {
					obj.Pos.Y = game.CharYLvl
				}
			}
		}
	}
}