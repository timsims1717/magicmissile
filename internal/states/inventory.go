package states

import (
	"fmt"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
	"timsims1717/magicmissile/internal/data"
	"timsims1717/magicmissile/internal/myecs"
	"timsims1717/magicmissile/internal/systems"
	"timsims1717/magicmissile/pkg/debug"
	"timsims1717/magicmissile/pkg/options"
	"timsims1717/magicmissile/pkg/state"
	"timsims1717/magicmissile/pkg/timing"
	"timsims1717/magicmissile/pkg/viewport"
)

var InventoryState = &inventoryState{}

type inventoryState struct {
	*state.AbstractState
}

func (s *inventoryState) Unload() {
	data.InventoryView = nil
	data.LeftTowerScroll = nil
}

func (s *inventoryState) Load() {
	data.InventoryView = viewport.New(nil)
	data.InventoryView.SetRect(pixel.R(0, 0, data.BaseWidth, data.BaseHeight))
	data.InventoryView.CamPos = pixel.ZV
	data.InventoryView.PortPos = viewport.MainCamera.PostCamPos
	systems.CreateTowersNoBG()
	systems.CreateTowerScrolls()
}

func (s *inventoryState) Update(win *pixelgl.Window) {
	debug.AddText("Inventory State")
	debug.AddIntCoords("World", int(data.TheInput.World.X), int(data.TheInput.World.Y))
	inPos := data.InventoryView.Projected(data.TheInput.World)
	debug.AddIntCoords("Inventory View In", int(inPos.X), int(inPos.Y))
	debug.AddIntCoords("Left Tower Pos", int(data.LeftTowerScroll.Scroll.Object.Pos.X), int(data.LeftTowerScroll.Scroll.Object.Pos.Y))

	if options.Updated {
		s.UpdateViews()
	}

	if data.TheInput.Get("debugCU").Pressed() {
		data.InventoryView.CamPos.Y += timing.DT * 50.
	} else if data.TheInput.Get("debugCD").Pressed() {
		data.InventoryView.CamPos.Y -= timing.DT * 50.
	}
	if data.TheInput.Get("debugCR").Pressed() {
		data.InventoryView.CamPos.X += timing.DT * 50.
	} else if data.TheInput.Get("debugCL").Pressed() {
		data.InventoryView.CamPos.X -= timing.DT * 50.
	}
	if data.TheInput.Get("showInventory").JustPressed() {
		systems.ShowTowerScroll(data.LeftTowerScroll)
		systems.ShowTowerScroll(data.MidTowerScroll)
		systems.ShowTowerScroll(data.RightTowerScroll)
	}

	systems.FunctionSystem()
	systems.ScrollSystem()
	systems.InterpolationSystem()
	systems.ParentSystem()
	systems.ObjectSystem()
	data.InventoryView.Update()
	systems.TemporarySystem()
	myecs.UpdateManager()
	debug.AddText(fmt.Sprintf("Entity Count: %d", myecs.FullCount))
}

func (s *inventoryState) Draw(win *pixelgl.Window) {
	data.InventoryView.Canvas.Clear(colornames.Pink)
	systems.DrawScrollSystem(win)
	data.InventoryView.Draw(win)
}

func (s *inventoryState) SetAbstract(aState *state.AbstractState) {
	s.AbstractState = aState
}

func (s *inventoryState) UpdateViews() {
	data.InventoryView.SetRect(pixel.R(0, 0, viewport.MainCamera.Rect.W(), viewport.MainCamera.Rect.H()))
	data.InventoryView.SetZoom(viewport.MainCamera.Rect.W() / data.BaseWidth)
}
