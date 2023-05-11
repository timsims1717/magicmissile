package states

import (
	"fmt"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"image/color"
	"timsims1717/magicmissile/internal/data"
	"timsims1717/magicmissile/internal/systems"
	"timsims1717/magicmissile/pkg/debug"
	"timsims1717/magicmissile/pkg/img"
	"timsims1717/magicmissile/pkg/options"
	"timsims1717/magicmissile/pkg/state"
	"timsims1717/magicmissile/pkg/timing"
	"timsims1717/magicmissile/pkg/viewport"
)

var BGTestState = &backgroundTestState{}

type backgroundTestState struct {
	*state.AbstractState
}

func (s *backgroundTestState) Unload() {
	data.GameView = nil
}

func (s *backgroundTestState) Load() {
	data.GameView = viewport.New(nil)
	data.GameView.SetRect(pixel.R(0, 0, data.BaseWidth, data.BaseHeight))
	data.GameView.CamPos = pixel.ZV
	data.GameView.PortPos = viewport.MainCamera.PostCamPos
	//data.GameView.Canvas.SetUniform("uM", float32(0.5))
	//data.GameView.Canvas.SetUniform("uB", float32(0.3))
	//data.GameView.Canvas.SetFragmentShader(shaders.BGShader)
	systems.GenerateBackground()
}

func (s *backgroundTestState) Update(win *pixelgl.Window) {
	debug.AddText("Background Test")

	if options.Updated {
		s.UpdateViews()
	}

	if data.TheInput.Get("debugCU").Pressed() {
		data.GameView.CamPos.Y += timing.DT * 50.
	} else if data.TheInput.Get("debugCD").Pressed() {
		data.GameView.CamPos.Y -= timing.DT * 50.
	}
	if data.TheInput.Get("debugCR").Pressed() {
		data.GameView.CamPos.X += timing.DT * 50.
	} else if data.TheInput.Get("debugCL").Pressed() {
		data.GameView.CamPos.X -= timing.DT * 50.
	}

	systems.UpdateBackgrounds()
	data.GameView.Update()
	debug.AddText(fmt.Sprintf("GameView Pos: (%d,%d)", int(data.GameView.CamPos.X), int(data.GameView.CamPos.Y)))
}

func (s *backgroundTestState) Draw(win *pixelgl.Window) {
	data.GameView.Canvas.Clear(color.RGBA{})
	for _, bg := range data.Backgrounds {
		bg.View.Canvas.Clear(color.RGBA{})
		//spr := &img.Sprite{
		//	Key:    "house1",
		//	Color:  white,
		//	Batch:  "stuff",
		//}
		img.Batchers["stuff"].DrawSprite("house1", pixel.IM.Moved(pixel.V(bg.View.Rect.W()*-0.5, bg.View.Rect.H()*-0.5)))
		img.Batchers["stuff"].DrawSprite("house1", pixel.IM.Moved(pixel.V(bg.View.Rect.W()*0.5, bg.View.Rect.H()*-0.5)))
		img.Batchers["stuff"].DrawSprite("house1", pixel.IM.Moved(pixel.V(bg.View.Rect.W()*0.5, bg.View.Rect.H()*0.5)))
		img.Batchers["stuff"].DrawSprite("house1", pixel.IM.Moved(pixel.V(bg.View.Rect.W()*-0.5, bg.View.Rect.H()*0.5)))
		img.Batchers["stuff"].Draw(bg.View.Canvas)
		bg.IMDraw.Draw(bg.View.Canvas)
		bg.View.Canvas.Draw(data.GameView.Canvas, bg.View.Mat)
	}
	data.GameView.Canvas.Draw(win, data.GameView.Mat)
}

func (s *backgroundTestState) SetAbstract(aState *state.AbstractState) {
	s.AbstractState = aState
}

func (s *backgroundTestState) UpdateViews() {
	data.GameView.SetRect(pixel.R(0, 0, viewport.MainCamera.Rect.W(), viewport.MainCamera.Rect.H()))
	data.GameView.SetZoom(viewport.MainCamera.Rect.W() / data.BaseWidth)
}
