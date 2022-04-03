package states

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"image/color"
	"math/rand"
	"timsims1717/magicmissile/internal/data"
	"timsims1717/magicmissile/internal/payloads"
	"timsims1717/magicmissile/internal/states/game"
	"timsims1717/magicmissile/internal/systems"
	"timsims1717/magicmissile/pkg/camera"
	"timsims1717/magicmissile/pkg/img"
	"timsims1717/magicmissile/pkg/state"
	"timsims1717/magicmissile/pkg/timing"
)

var white = color.RGBA{
	R: 255,
	G: 255,
	B: 255,
	A: 255,
}

var GameState = &gameState{}

type gameState struct {
	*state.AbstractState
}

func (s *gameState) Unload() {
	systems.ClearSystem()
}

func (s *gameState) Load(done chan struct{}) {
	loadUI()
	loadTowns()
	loadWizard()
	loadFighter()
	game.MTimer = timing.New(5.)
	game.DTimer = timing.New(15.)
	game.ZTimer = timing.New(8.)
	game.MFreq = 5.
	game.MSpd = 50.
	game.ZFreq = 10.
	game.Level = 0
	game.BigOn = false
	done <- struct{}{}
}

func (s *gameState) Update(win *pixelgl.Window) {
	data.TheInput.Update(win)
	game.Cursor = data.TheInput.World
	if game.Cursor.X < game.Frame.Min.X {
		game.Cursor.X = game.Frame.Min.X
	} else if game.Cursor.X > game.Frame.Max.X {
		game.Cursor.X = game.Frame.Max.X
	}
	if game.Cursor.Y < game.Frame.Min.Y {
		game.Cursor.Y = game.Frame.Min.Y
	} else if game.Cursor.Y > game.Frame.Max.Y {
		game.Cursor.Y = game.Frame.Max.Y
	}
	game.CheckGameOver()
	if game.GameOver {
		state.SwitchState("over")
	}
	systems.TemporarySystem()
	systems.FunctionSystem()
	systems.ControlSystem()
	systems.PayloadSystem()
	systems.AttackSystem()
	systems.MobSystem()
	systems.HealthSystem()
	systems.FullTransformSystem()
	systems.AnimationSystem()
	if game.MTimer.UpdateDone() {
		if game.BigOn && rand.Intn(5) == 0 {
			payloads.BigMeteor(game.MSpd * 0.75)
		} else {
			payloads.BasicMeteor(game.MSpd, pixel.ZV)
		}
		game.MTimer = timing.New(rand.Float64() * game.MFreq + 0.5)
	}
	if game.ZTimer.UpdateDone() {
		payloads.BasicZombie()
		game.ZTimer = timing.New(rand.Float64() * game.ZFreq + 0.5)
	}
	if game.DTimer.UpdateDone() {
		game.Level++
		game.ZFreq -= 0.1
		if game.ZFreq < 2. {
			game.ZFreq = 2.
		}
		game.MFreq -= 0.06
		if game.MFreq < 3. {
			game.MFreq = 3.
		}
		game.MSpd += 4.
		if game.MSpd > 120. {
			game.MSpd = 120.
		}
		if game.Level > 8 {
			game.BigOn = true
		}
		game.DTimer = timing.New(10.)
	}
	camera.Cam.Update(win)
}

func (s *gameState) Draw(win *pixelgl.Window) {
	img.Clear()
	systems.DrawSystem(win)
	img.Draw(win)
}

func (s *gameState) SetAbstract(aState *state.AbstractState) {
	s.AbstractState = aState
}