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

var GameState = &gameState{}

type gameState struct {
	*state.AbstractState
}

func (s *gameState) Unload() {
	data.GameView = nil
	data.ExpView = nil
	data.ExpView1 = nil
	data.GameDraw = nil
	systems.DisposeBackground()
	systems.DisposeTowns()
	systems.DisposeTowers()
}

func (s *gameState) Load() {
	data.GameView = viewport.New(nil)
	data.GameView.SetRect(pixel.R(0, 0, data.BaseWidth, data.BaseHeight))
	//data.GameView.CamPos = pixel.V(data.BaseWidth*0.5, data.BaseHeight*0.5)
	data.GameView.CamPos = pixel.ZV
	data.GameView.PortPos = viewport.MainCamera.CamPos
	data.ExpView = viewport.New(nil)
	data.ExpView.SetRect(pixel.R(0, 0, data.BaseWidth, data.BaseHeight))
	data.ExpView.CamPos = pixel.ZV
	data.ExpView.PortPos = data.GameView.CamPos
	data.ExpView1 = viewport.New(nil)
	data.ExpView1.SetRect(pixel.R(0, 0, data.BaseWidth, data.BaseHeight))
	data.ExpView1.CamPos = pixel.ZV
	data.ExpView1.PortPos = data.GameView.CamPos
	data.GameDraw = imdraw.New(nil)
	systems.GenerateRandomBackground("ForestValley")
	systems.UpdateBackgrounds()
	systems.CreateTowns()
	systems.CreateTowers()
}

func (s *gameState) Update(win *pixelgl.Window) {
	debug.AddText("Game State")
	debug.AddIntCoords("World", int(data.TheInput.World.X), int(data.TheInput.World.Y))
	inPos := data.GameView.Projected(data.TheInput.World)
	debug.AddIntCoords("Game View In", int(inPos.X), int(inPos.Y))

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
	debug.AddText(fmt.Sprintf("Spell Tier: %d", data.TierTest+1))
	debug.AddText(fmt.Sprintf("Spell Name: %s", data.SpellKeys[data.SpellTest]))
	if data.TheInput.Get("debugCU").Pressed() {
		data.GameView.PortPos.Y += timing.DT * 50.
	} else if data.TheInput.Get("debugCD").Pressed() {
		data.GameView.PortPos.Y -= timing.DT * 50.
	}
	if data.TheInput.Get("debugCR").Pressed() {
		data.GameView.PortPos.X += timing.DT * 50.
	} else if data.TheInput.Get("debugCL").Pressed() {
		data.GameView.PortPos.X -= timing.DT * 50.
	}
	if data.TheInput.Get("fireLeft").JustPressed() {
		tower := data.Towers[0]
		target := inPos
		systems.FireNextFromTower(tower, target)
	}
	if data.TheInput.Get("fireMid").JustPressed() {
		tower := data.Towers[1]
		target := inPos
		systems.FireNextFromTower(tower, target)
	}
	if data.TheInput.Get("fireRight").JustPressed() {
		tower := data.Towers[2]
		target := inPos
		systems.FireNextFromTower(tower, target)
	}

	systems.FunctionSystem()
	systems.MissileSystem()
	systems.ExplosionSystem()
	//systems.HealthSystem()
	systems.InterpolationSystem()
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

func (s *gameState) Draw(win *pixelgl.Window) {
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
	data.ExpView.Draw(data.GameView.Canvas)
	systems.DrawSystem(win, 10)
	img.Batchers[data.ParticleKey].Draw(data.GameView.Canvas)
	data.GameView.Draw(win)
}

func (s *gameState) SetAbstract(aState *state.AbstractState) {
	s.AbstractState = aState
}

func (s *gameState) UpdateViews() {
	ratio := viewport.MainCamera.Rect.W() / data.BaseWidth
	data.GameView.PortSize = pixel.V(ratio, ratio)
	data.GameView.PortPos = viewport.MainCamera.CamPos
}
