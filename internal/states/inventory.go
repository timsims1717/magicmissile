package states

import (
	"fmt"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
	"timsims1717/magicmissile/internal/data"
	"timsims1717/magicmissile/internal/myecs"
	"timsims1717/magicmissile/internal/systems"
	"timsims1717/magicmissile/internal/systems/inventory"
	"timsims1717/magicmissile/pkg/debug"
	"timsims1717/magicmissile/pkg/img"
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
	data.InventoryView.PortPos = viewport.MainCamera.CamPos
	data.InventoryView.Update()
	systems.CreateTowersNoBG()
	data.SquareFrame = img.Batchers[data.UIKey].GetSprite("scroll_square").Frame()
	inventory.CreateTowerScrolls()
	inventory.CreateMainMovingSpellSlot()
	inventory.CreateSpellInventory()
}

func (s *inventoryState) Update(win *pixelgl.Window) {
	debug.AddTruthText(fmt.Sprintf("Inventory State (%d) - Transition", data.InventoryState), data.InventoryTrans)
	debug.AddIntCoords("Main Camera Pos", int(viewport.MainCamera.CamPos.X), int(viewport.MainCamera.CamPos.Y))
	debug.AddIntCoords("World", int(data.TheInput.World.X), int(data.TheInput.World.Y))
	inPos := data.InventoryView.ProjectWorld(data.TheInput.World)
	debug.AddIntCoords("Inventory View In", int(inPos.X), int(inPos.Y))
	inside, edge := data.InventoryView.WorldInside(data.TheInput.World)
	debug.AddTruthText("Inventory Point Inside", inside)
	debug.AddIntCoords("Inventory Edge", int(edge.X), int(edge.Y))

	if options.Updated {
		s.UpdateViews()
	}

	if data.TheInput.Get("debugCU").Pressed() {
		data.LeftTowerScroll.InvSlots[0].SlotObj.Pos.Y += timing.DT * 50.
	} else if data.TheInput.Get("debugCD").Pressed() {
		data.LeftTowerScroll.InvSlots[0].SlotObj.Pos.Y -= timing.DT * 50.
	}
	if data.TheInput.Get("debugCR").Pressed() {
		data.LeftTowerScroll.InvSlots[0].SlotObj.Pos.X += timing.DT * 50.
	} else if data.TheInput.Get("debugCL").Pressed() {
		data.LeftTowerScroll.ListView.CamPos.X -= timing.DT * 50.
	}
	if data.InventoryTrans {
		switch data.InventoryState {
		case -1:
			if data.LeftTowerScroll.Scroll.Closed &&
				data.MidTowerScroll.Scroll.Closed &&
				data.RightTowerScroll.Scroll.Closed {
				data.InventoryTrans = false
			}
		case 0, 1, 2:
			count := 0
			if data.LeftTowerScroll.Scroll.Closed {
				count++
			}
			if data.MidTowerScroll.Scroll.Closed {
				count++
			}
			if data.RightTowerScroll.Scroll.Closed {
				count++
			}
			if count >= 2 {
				data.InventoryTrans = false
			}
		case 3:
			if data.LeftTowerScroll.Scroll.Opened &&
				data.MidTowerScroll.Scroll.Opened &&
				data.RightTowerScroll.Scroll.Opened {
				data.InventoryTrans = false
			}
		}
	}
	if data.TheInput.Get("showInventory").JustPressed() && !data.InventoryTrans {
		if data.InventoryState == -1 {
			inventory.ShowTowerScroll(data.LeftTowerScroll)
			inventory.ShowTowerScroll(data.MidTowerScroll)
			inventory.ShowTowerScroll(data.RightTowerScroll)
			data.InventoryState = 3
		} else {
			inventory.HideTowerScroll(data.LeftTowerScroll)
			inventory.HideTowerScroll(data.MidTowerScroll)
			inventory.HideTowerScroll(data.RightTowerScroll)
			data.InventoryState = -1
		}
		data.InventoryTrans = true
	}

	systems.FunctionSystem()
	inventory.ScrollSystem()
	systems.InterpolationSystem()
	systems.ParentSystem()
	systems.ObjectSystem()
	inventory.UpdateListViews()
	inventory.UpdateSpellInventory(false)
	data.InventoryView.Update()
	systems.TemporarySystem()
	myecs.UpdateManager()
	debug.AddText(fmt.Sprintf("Entity Count: %d", myecs.FullCount))
}

func (s *inventoryState) Draw(win *pixelgl.Window) {
	data.InventoryView.Canvas.Clear(colornames.Pink)
	inventory.DrawTowerScrollSystem(win)
	inventory.DrawSpellInventory(win)
	inventory.DrawMovingSpellSlots(win)
	data.InventoryView.Draw(viewport.MainCamera.Canvas)
}

func (s *inventoryState) SetAbstract(aState *state.AbstractState) {
	s.AbstractState = aState
}

func (s *inventoryState) UpdateViews() {
	ratio := viewport.MainCamera.Rect.W() / data.BaseWidth
	data.InventoryView.PortSize = pixel.V(ratio, ratio)
	data.InventoryView.PortPos = viewport.MainCamera.CamPos
}
