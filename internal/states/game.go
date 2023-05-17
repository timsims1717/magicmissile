package states

import (
	"fmt"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"image/color"
	"timsims1717/magicmissile/internal/data"
	"timsims1717/magicmissile/internal/myecs"
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
	//data.GameView.CamPos = pixel.V(data.BaseWidth*0.5, data.BaseHeight*0.5)
	data.GameView.CamPos = pixel.ZV
	data.GameView.PortPos = viewport.MainCamera.PostCamPos
	data.ExpView = viewport.New(nil)
	data.ExpView.SetRect(pixel.R(0, 0, data.BaseWidth, data.BaseHeight))
	data.ExpView.CamPos = pixel.ZV
	data.ExpView.PortPos = viewport.MainCamera.PostCamPos
	data.ExpView1 = viewport.New(nil)
	data.ExpView1.SetRect(pixel.R(0, 0, data.BaseWidth, data.BaseHeight))
	data.ExpView1.CamPos = pixel.ZV
	data.ExpView1.PortPos = viewport.MainCamera.PostCamPos
	data.GameDraw = imdraw.New(nil)
	systems.GenerateRandomBackground("ForestValley")
	systems.UpdateBackgrounds()
	systems.CreateTowns()
	systems.CreateTowers()
}

func (s *backgroundTestState) Update(win *pixelgl.Window) {
	debug.AddText("Game State")
	inPos := data.GameView.Projected(data.TheInput.World)
	debug.AddText(fmt.Sprintf("World: (%d,%d)", int(data.TheInput.World.X), int(data.TheInput.World.Y)))
	debug.AddText(fmt.Sprintf("GameView: (%d,%d)", int(inPos.X), int(inPos.Y)))

	if options.Updated {
		s.UpdateViews()
	}

	if data.TheInput.Get("changeBackground").JustPressed() {
		systems.GenerateRandomBackground("ForestValley")
		systems.UpdateBackgrounds()
	}
	if data.TheInput.Get("debugSpellTier").JustPressed() {
		data.TierTest++
		data.TierTest %= 5
	}
	if data.TheInput.Get("debugSpellName").JustPressed() {
		data.SpellTest++
		data.SpellTest %= len(data.SpellKeys)
	}
	if data.TheInput.Get("debugCU").Pressed() {
		data.CurrBackground.Backgrounds[0].View.PortPos.Y += timing.DT * 50.
	} else if data.TheInput.Get("debugCD").Pressed() {
		data.CurrBackground.Backgrounds[0].View.PortPos.Y -= timing.DT * 50.
	}
	if data.TheInput.Get("debugCR").Pressed() {
		data.CurrBackground.Backgrounds[0].View.PortPos.X += timing.DT * 50.
	} else if data.TheInput.Get("debugCL").Pressed() {
		data.CurrBackground.Backgrounds[0].View.PortPos.X -= timing.DT * 50.
	}
	if data.TheInput.Get("fireLeft").JustPressed() {
		tower := data.Towers[0]
		target := inPos
		target.X -= data.BaseWidth * 0.5
		target.Y -= data.BaseHeight * 0.5
		missile := data.Missiles[data.SpellKeys[data.SpellTest]][data.TierTest]
		systems.FireFromTower(missile, tower, target)
	}
	if data.TheInput.Get("fireMid").JustPressed() {
		tower := data.Towers[1]
		target := inPos
		target.X -= data.BaseWidth * 0.5
		target.Y -= data.BaseHeight * 0.5
		missile := data.Missiles[data.SpellKeys[data.SpellTest]][data.TierTest]
		systems.FireFromTower(missile, tower, target)
	}
	if data.TheInput.Get("fireRight").JustPressed() {
		tower := data.Towers[2]
		target := inPos
		target.X -= data.BaseWidth * 0.5
		target.Y -= data.BaseHeight * 0.5
		missile := data.Missiles[data.SpellKeys[data.SpellTest]][data.TierTest]
		systems.FireFromTower(missile, tower, target)
	}

	systems.FunctionSystem()
	systems.MissileSystem()
	systems.ExplosionSystem()
	//systems.HealthSystem()
	systems.ParentSystem()
	systems.ObjectSystem()
	//systems.UpdateBackgrounds()
	for _, bg := range data.CurrBackground.Backgrounds {
		bg.View.Update()
	}
	data.GameView.Update()
	data.ExpView.Update()
	data.ExpView1.Update()
	systems.TemporarySystem()
	myecs.UpdateManager()
	debug.AddText(fmt.Sprintf("Entity Count: %d", myecs.FullCount))
}

func (s *backgroundTestState) Draw(win *pixelgl.Window) {
	systems.DrawExplosionSystem()
	data.GameView.Canvas.Clear(data.CurrBackground.BackCol)
	for i, bg := range data.CurrBackground.Backgrounds {
		bg.View.Canvas.Clear(color.RGBA{})
		img.Clear()
		systems.DrawSystem(win, i)
		img.Batchers[data.ObjectKey].Draw(bg.View.Canvas)
		bg.IMDraw.Draw(bg.View.Canvas)
		bg.View.Canvas.Draw(data.GameView.Canvas, bg.View.Mat)
	}
	data.ExpView.Canvas.Draw(data.GameView.Canvas, data.ExpView.Mat)
	systems.DrawSystem(win, 10)
	img.Batchers[data.ParticleKey].Draw(data.GameView.Canvas)
	data.GameView.Canvas.Draw(win, data.GameView.Mat)
}

func (s *backgroundTestState) SetAbstract(aState *state.AbstractState) {
	s.AbstractState = aState
}

func (s *backgroundTestState) UpdateViews() {
	data.GameView.SetRect(pixel.R(0, 0, viewport.MainCamera.Rect.W(), viewport.MainCamera.Rect.H()))
	data.GameView.SetZoom(viewport.MainCamera.Rect.W() / data.BaseWidth)
}
