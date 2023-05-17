package systems

import (
	"github.com/faiface/pixel"
	"golang.org/x/image/colornames"
	"image/color"
	"timsims1717/magicmissile/internal/data"
	"timsims1717/magicmissile/internal/myecs"
	"timsims1717/magicmissile/pkg/object"
	"timsims1717/magicmissile/pkg/timing"
)

func MakeExplosion(eFab *data.Explosion, pos pixel.Vec, col color.RGBA) {
	obj := object.New()
	obj.Pos = pos
	exp := &data.Explosion{
		FullRadius: eFab.FullRadius,
		ExpandRate: eFab.ExpandRate,
		Dissipate:  eFab.Dissipate,
		DisRate:    eFab.DisRate,
		StartColor: col,
	}
	myecs.Manager.NewEntity().
		AddComponent(myecs.Object, obj).
		AddComponent(myecs.Explosion, exp)
}

func ExplosionSystem() {
	for _, result := range myecs.Manager.Query(myecs.IsExplosion) {
		_, okO := result.Components[myecs.Object].(*object.Object)
		exp, okE := result.Components[myecs.Explosion].(*data.Explosion)
		if okO && okE {
			if exp.Timer == nil {
				exp.Timer = timing.New(0.)
			}
			exp.Timer.Update()
			t := exp.Timer.Elapsed()
			// expand current radius
			diff := exp.FullRadius - exp.CurrRadius
			exp.CurrRadius += exp.ExpandRate * diff * timing.DT
			if exp.CurrRadius > exp.FullRadius {
				exp.CurrRadius = exp.FullRadius
			}
			if exp.CurrRadius < 0. {
				exp.CurrRadius = 0.
			}
			// if dissipating time is past, expand dissipate radius
			if t > exp.Dissipate {
				exp.DisRadius += exp.DisRate * timing.DT
				if exp.DisRadius > exp.FullRadius {
					// dispose of this entity
					myecs.Manager.DisposeEntity(result.Entity)
				}
				if exp.DisRadius < 0. {
					exp.DisRadius = 0.
				}
			}
		}
	}
}

func DrawExplosionSystem() {
	switch data.ExpDrawType {
	case 0:
		DrawNewExplosionSystem()
	case 1:
		DrawNewExplosionSystem1()
	case 2:
		DrawNewExplosionSystem2()
	}
}

// DrawNewExplosionSystem
// This one looks the best, since each explosion is totally independent.
// Once you get above 15 or so explosions, the performance drops, though.
func DrawNewExplosionSystem() {
	data.ExpView.Canvas.Clear(color.RGBA{})
	data.ExpView1.Canvas.Clear(color.RGBA{})
	data.ExpView1.Canvas.SetComposeMethod(pixel.ComposeXor)
	for _, result := range myecs.Manager.Query(myecs.IsExplosion) {
		obj, okO := result.Components[myecs.Object].(*object.Object)
		exp, okE := result.Components[myecs.Explosion].(*data.Explosion)
		if okO && okE {
			data.ExpView1.Canvas.Clear(color.RGBA{})
			data.GameDraw.Clear()
			data.GameDraw.Color = exp.StartColor
			data.GameDraw.Push(obj.Pos)
			data.GameDraw.Circle(exp.CurrRadius, 0.)
			data.GameDraw.Draw(data.ExpView1.Canvas)
			data.GameDraw.Clear()
			data.GameDraw.Color = colornames.Pink
			data.GameDraw.Push(obj.Pos)
			data.GameDraw.Circle(exp.DisRadius, 0.)
			data.GameDraw.Draw(data.ExpView1.Canvas)
			data.ExpView1.Canvas.Draw(data.GameView.Canvas, data.ExpView1.Mat)
		}
	}
	data.ExpView1.Canvas.SetComposeMethod(pixel.ComposeOver)
	data.ExpView1.Canvas.Draw(data.ExpView.Canvas, data.ExpView1.Mat)
}

// DrawNewExplosionSystem1
// This one doesn't look great, since when an older explosion disappears,
// any explosions behind it reappear.
// No performance issues, though.
func DrawNewExplosionSystem1() {
	data.ExpView1.Canvas.Clear(color.RGBA{})
	data.ExpView.Canvas.Clear(color.RGBA{})
	data.ExpView.Canvas.SetComposeMethod(pixel.ComposeXor)
	data.GameDraw.Clear()
	for _, result := range myecs.Manager.Query(myecs.IsExplosion) {
		obj, okO := result.Components[myecs.Object].(*object.Object)
		exp, okE := result.Components[myecs.Explosion].(*data.Explosion)
		if okO && okE {
			data.GameDraw.Color = exp.StartColor
			data.GameDraw.Push(obj.Pos)
			data.GameDraw.Circle(exp.CurrRadius, 0.)
		}
	}
	data.GameDraw.Draw(data.ExpView1.Canvas)
	data.GameDraw.Clear()
	data.ExpView1.Canvas.Draw(data.ExpView.Canvas, data.ExpView1.Mat)
	data.ExpView1.Canvas.Clear(color.RGBA{})
	//data.ExpView.Canvas.SetComposeMethod(pixel.ComposeOut)
	for _, result := range myecs.Manager.Query(myecs.IsExplosion) {
		obj, okO := result.Components[myecs.Object].(*object.Object)
		exp, okE := result.Components[myecs.Explosion].(*data.Explosion)
		if okO && okE {
			data.GameDraw.Color = colornames.Pink
			data.GameDraw.Push(obj.Pos)
			data.GameDraw.Circle(exp.DisRadius, 0.)
		}
	}
	data.GameDraw.Draw(data.ExpView1.Canvas)
	data.ExpView1.Canvas.Draw(data.ExpView.Canvas, data.ExpView1.Mat)
}

// DrawNewExplosionSystem2
// This one looks good and has decent performance.
func DrawNewExplosionSystem2() {
	data.ExpView.Canvas.Clear(color.RGBA{})
	for _, result := range myecs.Manager.Query(myecs.IsExplosion) {
		obj, okO := result.Components[myecs.Object].(*object.Object)
		exp, okE := result.Components[myecs.Explosion].(*data.Explosion)
		if okO && okE {
			data.ExpView.Canvas.SetComposeMethod(pixel.ComposeOver)
			data.GameDraw.Clear()
			data.GameDraw.Color = exp.StartColor
			data.GameDraw.Push(obj.Pos)
			data.GameDraw.Circle(exp.CurrRadius, 0.)
			data.GameDraw.Draw(data.ExpView.Canvas)
			data.ExpView.Canvas.SetComposeMethod(pixel.ComposeXor)
			data.GameDraw.Clear()
			data.GameDraw.Color = colornames.Pink
			data.GameDraw.Push(obj.Pos)
			data.GameDraw.Circle(exp.DisRadius, 0.)
			data.GameDraw.Draw(data.ExpView.Canvas)
		}
	}
}
