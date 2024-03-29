package main

import (
	"fmt"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"math/rand"
	"time"
	"timsims1717/magicmissile/internal/data"
	"timsims1717/magicmissile/internal/loading"
	"timsims1717/magicmissile/internal/states"
	"timsims1717/magicmissile/pkg/debug"
	"timsims1717/magicmissile/pkg/img"
	"timsims1717/magicmissile/pkg/options"
	"timsims1717/magicmissile/pkg/sfx"
	"timsims1717/magicmissile/pkg/state"
	"timsims1717/magicmissile/pkg/timing"
	"timsims1717/magicmissile/pkg/typeface"
	"timsims1717/magicmissile/pkg/viewport"
)

func run() {
	rand.New(rand.NewSource(time.Now().Unix()))
	conf := pixelgl.WindowConfig{
		Title:     "Magic Missile",
		Bounds:    pixel.R(0, 0, 1600, 900),
		VSync:     true,
		Invisible: true,
	}
	win, err := pixelgl.NewWindow(conf)
	if err != nil {
		panic(err)
	}
	win.SetSmooth(true)

	viewport.MainCamera = viewport.New(win.Canvas())
	viewport.MainCamera.SetRect(pixel.R(0, 0, 1600, 900))
	viewport.MainCamera.CamPos = pixel.V(data.BaseWidth*0.5, data.BaseHeight*0.5)

	options.VSync = true
	options.BilinearFilter = true

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

	mainFont, err := typeface.LoadTTF("assets/JunicodeTwoBeta-Bold.ttf", 200.)
	typeface.Atlases["main"] = text.NewAtlas(mainFont, text.ASCII)

	titleFont, err := typeface.LoadTTF("assets/Enchanted Land.ttf", 200.)
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

	state.Register("game", state.New(states.GameState))
	state.Register("inventory", state.New(states.InventoryState))
	state.Register("old-menu", state.New(states.OldMenuState))
	state.Register("old-game", state.New(states.OldGameState))
	state.PushState("inventory")

	debug.Initialize(&viewport.MainCamera.PostCamPos)
	debug.Text = true

	loading.LoadShaders()
	loading.LoadImg()
	loading.LoadTileMaps()

	err = loading.LoadRealms("assets/data/realms.json")
	if err != nil {
		panic(err)
	}
	err = loading.LoadSpells("assets/data/spells.json")
	if err != nil {
		panic(err)
	}

	win.Show()
	win.Canvas()
	timing.Reset()
	for !win.Closed() {
		timing.Update()
		debug.Clear()
		options.WindowUpdate(win)
		if options.Updated {
			viewport.MainCamera.CamPos = pixel.V(viewport.MainCamera.Rect.W()*0.5, viewport.MainCamera.Rect.H()*0.5)
		}

		data.TheInput.Update(win, viewport.MainCamera.Mat)
		if data.TheInput.Get("debugPause").JustPressed() {
			fmt.Println("BREAKPOINT")
		}
		if data.TheInput.Get("fullscreen").JustPressed() {
			options.FullScreen = !options.FullScreen
		}
		if data.TheInput.Get("fuzzy").JustPressed() {
			options.BilinearFilter = !options.BilinearFilter
		}
		if data.TheInput.Get("debugText").JustPressed() {
			debug.Text = !debug.Text
		}
		if data.TheInput.Get("debugExpDrawType").JustPressed() {
			data.ExpDrawType++
			data.ExpDrawType %= 4
		}

		state.Update(win)
		viewport.MainCamera.Update()

		state.Draw(win)
		win.SetSmooth(false)
		debug.Draw(win)
		win.SetSmooth(true)

		sfx.MusicPlayer.Update()
		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
