package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"image/color"
	"math/rand"
	"time"
	"timsims1717/magicmissile/internal/states"
	"timsims1717/magicmissile/pkg/camera"
	"timsims1717/magicmissile/pkg/img"
	"timsims1717/magicmissile/pkg/state"
	"timsims1717/magicmissile/pkg/timing"
)

func run() {
	rand.Seed(time.Now().Unix())
	conf := pixelgl.WindowConfig{
		Title:  "Magic Missile",
		Bounds: pixel.R(0, 0, 1600, 900),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(conf)
	if err != nil {
		panic(err)
	}
	win.SetSmooth(true)

	camera.Cam = camera.New(true)
	camera.Cam.Opt.WindowScale = 900.
	camera.Cam.SetZoom(1.)
	camera.Cam.SetSize(1600., 900.)

	testSheet, err := img.LoadSpriteSheet("assets/test.json")
	if err != nil {
		panic(err)
	}
	img.AddBatcher("test", testSheet, true, true)
	figuresSheet, err := img.LoadSpriteSheet("assets/figures.json")
	if err != nil {
		panic(err)
	}
	img.AddBatcher("figures", figuresSheet, true, true)
	img.AddIMDrawer("explosions", true, true)

	state.Register("game", state.New(states.GameState))
	state.Register("over", state.New(states.OverState))
	state.SwitchState("game")

	timing.Reset()
	for !win.Closed() {
		timing.Update()

		state.Update(win)
		camera.Cam.Update(win)

		win.Clear(color.RGBA{
			R: 100,
			G: 100,
			B: 100,
			A: 255,
		})

		state.Draw(win)

		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}