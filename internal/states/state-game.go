package states

import (
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
	loadTowns()
	loadWizard()
	loadFighter()
	game.MTimer = timing.New(5.)
	done <- struct{}{}
}

func (s *gameState) Update(win *pixelgl.Window) {
	data.TheInput.Update(win)
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
		if rand.Intn(4) > 0 {
			payloads.BasicMeteor()
			game.MTimer = timing.New(rand.Float64() * 4.)
		} else {
			payloads.BasicZombie()
			game.MTimer = timing.New(rand.Float64() * 4.)
		}
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