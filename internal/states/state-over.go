package states

import (
	"github.com/faiface/pixel/pixelgl"
	"timsims1717/magicmissile/internal/data"
	"timsims1717/magicmissile/internal/states/game"
	"timsims1717/magicmissile/internal/systems"
	"timsims1717/magicmissile/pkg/camera"
	"timsims1717/magicmissile/pkg/img"
	"timsims1717/magicmissile/pkg/state"
	"timsims1717/magicmissile/pkg/timing"
)

var OverState = &overState{}

type overState struct {
	*state.AbstractState
}

func (s *overState) Unload() {
	systems.ClearSystem()
}

func (s *overState) Load(done chan struct{}) {
	game.MTimer = timing.New(5.)
	done <- struct{}{}
}

func (s *overState) Update(win *pixelgl.Window) {
	data.TheInput.Update(win)
	game.CheckGameOver()
	systems.TemporarySystem()
	systems.FunctionSystem()
	systems.PayloadSystem()
	systems.HealthSystem()
	systems.FullTransformSystem()
	systems.AnimationSystem()
	if game.MTimer.UpdateDone() {
		// Game Over message, menu
	}
	camera.Cam.Update(win)
}

func (s *overState) Draw(win *pixelgl.Window) {
	img.Clear()
	systems.DrawSystem(win)
	img.Draw(win)
}

func (s *overState) SetAbstract(aState *state.AbstractState) {
	s.AbstractState = aState
}