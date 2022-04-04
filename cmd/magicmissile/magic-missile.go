package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"image/color"
	"math/rand"
	"time"
	"timsims1717/magicmissile/internal/states"
	"timsims1717/magicmissile/pkg/camera"
	"timsims1717/magicmissile/pkg/img"
	"timsims1717/magicmissile/pkg/sfx"
	"timsims1717/magicmissile/pkg/state"
	"timsims1717/magicmissile/pkg/timing"
	"timsims1717/magicmissile/pkg/typeface"
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

	// sfx
	sfx.SoundPlayer.RegisterSound("assets/click.wav", "click")
	sfx.SoundPlayer.RegisterSound("assets/explosion1.wav", "explosion1")
	sfx.SoundPlayer.RegisterSound("assets/explosion2.wav", "explosion2")
	sfx.SoundPlayer.RegisterSound("assets/smash.wav", "smash")
	sfx.SoundPlayer.RegisterSound("assets/thunder1.wav", "thunder1")
	sfx.SoundPlayer.RegisterSound("assets/thunder2.wav", "thunder2")
	sfx.SoundPlayer.RegisterSound("assets/zombie.wav", "zombie")
	sfx.SoundPlayer.RegisterSound("assets/zombie-hit.wav", "zombie-hit")

	// music
	sfx.MusicPlayer.RegisterMusicTrack("assets/wind1.wav", "wind1")
	sfx.MusicPlayer.RegisterMusicTrack("assets/wind2.wav", "wind2")
	sfx.MusicPlayer.NewSet("ambience", []string{"wind1", "wind2"}, sfx.Random, 0., 4.)

	mainFont, err := typeface.LoadTTF("assets/FR73PixD.ttf", 200.)
	typeface.Atlases["main"] = text.NewAtlas(mainFont, text.ASCII)

	titleFont, err := typeface.LoadTTF("assets/KumarOne.ttf", 200.)
	typeface.Atlases["title"] = text.NewAtlas(titleFont, text.ASCII)

	states.InitMenus(win)

	stuffSheet, err := img.LoadSpriteSheet("assets/stuff.json")
	if err != nil {
		panic(err)
	}
	img.AddBatcher("stuff", stuffSheet, true, true)
	scenerySheet, err := img.LoadSpriteSheet("assets/scenery.json")
	if err != nil {
		panic(err)
	}
	img.AddBatcher("sceneryfg", scenerySheet, true, true)
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
	img.AddIMDrawer("health", true, true)

	state.Register("menu", state.New(states.MenuState))
	state.Register("game", state.New(states.GameState))
	state.SwitchState("menu")

	timing.Reset()
	for !win.Closed() {
		timing.Update()

		state.Update(win)
		camera.Cam.Update(win)

		win.Clear(color.RGBA{
			R: 0,
			G: 0,
			B: 0,
			A: 255,
		})

		state.Draw(win)

		sfx.MusicPlayer.Update()
		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}