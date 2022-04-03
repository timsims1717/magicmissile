package systems

import (
	"timsims1717/magicmissile/internal/data"
	"timsims1717/magicmissile/internal/myecs"
	"timsims1717/magicmissile/internal/states/game"
	"timsims1717/magicmissile/pkg/object"
	"timsims1717/magicmissile/pkg/timing"
)

func ControlSystem() {
	getNext := false
	alive := -1
	for i, pc := range game.PCs {
		if pc.Char.Health.Dead && pc.Move.Selected {
			pc.Move.Selected = false
			getNext = true
		} else if getNext {
			pc.Move.Selected = true
			getNext = false
			break
		} else if !pc.Char.Health.Dead && alive < 0 {
			alive = i
		}
	}
	if getNext && alive > -1 {
		game.PCs[alive].Move.Selected = true
	}
	for i, pc := range game.PCs {
		if data.TheInput.Get(pc.Move.Key).JustPressed() &&
			!pc.Char.Health.Dead && !pc.Move.Selected {
			pc.Move.Selected = true
			for j, pc2 := range game.PCs {
				if j != i {
					pc2.Move.Selected = false
				}
			}
			break
		}
	}
	for _, result := range myecs.Manager.Query(myecs.PlayerChar) {
		obj, ok := result.Components[myecs.Object].(*object.Object)
		move, okM := result.Components[myecs.Movable].(*data.Moving)
		if ok && okM && move.Selected {
			if data.TheInput.Get("moveLeft").Pressed() {
				obj.Pos.X -= move.Speed * timing.DT
				if obj.Pos.X < -750. {
					obj.Pos.X = -750.
				}
				obj.Flip = true
			} else if data.TheInput.Get("moveRight").Pressed() {
				obj.Pos.X += move.Speed * timing.DT
				if obj.Pos.X > 750. {
					obj.Pos.X = 750.
				}
				obj.Flip = false
			}
		}
	}
}