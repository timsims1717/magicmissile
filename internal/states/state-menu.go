package states

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"image/color"
	"timsims1717/magicmissile/internal/data"
	"timsims1717/magicmissile/internal/myecs"
	"timsims1717/magicmissile/internal/states/game"
	"timsims1717/magicmissile/internal/systems"
	"timsims1717/magicmissile/pkg/img"
	"timsims1717/magicmissile/pkg/sfx"
	"timsims1717/magicmissile/pkg/state"
	"timsims1717/magicmissile/pkg/typeface"
	"timsims1717/magicmissile/pkg/viewport"
)

var OldMenuState = &oldMenuState{}

type oldMenuState struct {
	*state.AbstractState
}

func (s *oldMenuState) Unload() {
	systems.ClearSystem()
}

func (s *oldMenuState) Load() {
	MainMenu.Open()
	game.Title = typeface.New("title", typeface.NewAlign(typeface.Center, typeface.Center), 1.0, 1.0, 0., 0.)
	game.Title.SetColor(color.RGBA{
		R: 223,
		G: 62,
		B: 35,
		A: 255,
	})
	game.Title.SetPos(pixel.V(0., 250.))
	game.Title.SetText("MagicMissile")
	myecs.Manager.NewEntity().AddComponent(myecs.Object, game.Title.Obj)
	loadTowns()
	loadScenery()
	sfx.MusicPlayer.PlayMusic("ambience")
}

func (s *oldMenuState) Update(win *pixelgl.Window) {
	data.TheInput.Update(win, viewport.MainCamera.Mat)
	systems.TemporarySystem()
	systems.FunctionSystem()
	UpdateMenus(data.TheInput)
}

func (s *oldMenuState) Draw(win *pixelgl.Window) {
	img.Clear()
	systems.DrawSystem(win, 0)
	img.Draw(win)
	DrawMenus(win)
	game.Title.Draw(win)
}

func (s *oldMenuState) SetAbstract(aState *state.AbstractState) {
	s.AbstractState = aState
}
