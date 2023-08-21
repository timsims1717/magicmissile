package inventory

import (
	"github.com/bytearena/ecs"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"timsims1717/magicmissile/internal/data"
	"timsims1717/magicmissile/internal/myecs"
	"timsims1717/magicmissile/internal/systems"
	"timsims1717/magicmissile/pkg/gween64/ease"
	"timsims1717/magicmissile/pkg/img"
	"timsims1717/magicmissile/pkg/object"
	"timsims1717/magicmissile/pkg/util"
)

func CreateMainMovingSpellSlot() {
	DisposeMovingSpellSlot()
	invSpellSlot := CreateMovingSpellSlot(AddMSSEntity, 2)
	data.MovingSpellSlot.InvSpellSlot = *invSpellSlot
	data.MovingSpellSlot.Entity.AddComponent(myecs.Update, data.NewHoverClickFn(data.TheInput, data.InventoryView, func(hvc *data.HoverClick) {
		if !data.MovingSpellSlot.Moving {
			data.MovingSpellSlot.NameMObj.Pos = data.InventoryView.ProjectWorld(hvc.Input.World)
		}
		if data.MovingSpellSlot.Slot != nil && data.MovingSpellSlot.Slot.Spell != "" &&
			hvc.Input.Get("click").JustReleased() || data.InventoryTrans {
			if data.MovingSpellSlot.TierMoveIndex == -1 {
				hoveredSlot := GetHoveredSlot()
				if hoveredSlot != nil && (hoveredSlot.Slot.Tier == 0 ||
					hoveredSlot.Slot.Tier >= data.MovingSpellSlot.Slot.Tier) {
					// if the hovered slot is a legal place to put the spell
					if data.MovingSpellSlot.PrevSlot != nil && hoveredSlot.Slot.Spell != "" &&
						data.MovingSpellSlot.PrevSlot.Slot.Tier >= data.Missiles[hoveredSlot.Slot.Spell][0].Tier {
						// if we can switch the slots
						MoveNewSlotToSlot(hoveredSlot.View.ProjectedOut(hoveredSlot.NameMObj.Pos), hoveredSlot.Slot, data.MovingSpellSlot.PrevSlot, 0.25)
					} else {
						// put the replaced spell slot into the spell storage
						MoveNewSlotToInventory(hoveredSlot.View.ProjectedOut(hoveredSlot.NameMObj.Pos), hoveredSlot.Slot, 0.25)
					}
					SetMovingSlotToList(hoveredSlot, 0.25)
				} else if data.MovingSpellSlot.PrevSlot != nil {
					// if not, put it back where you got it from
					SetMovingSlotToList(data.MovingSpellSlot.PrevSlot, 0.25)
				} else {
					// if you don't know where it was, put it into the spell storage
					SetMovingSlotToInventory(0.25)
				}
			} else {
				hoveredSlot := GetEmptyTierSlot(data.MovingSpellSlot.TierMoveIndex)
				if hoveredSlot != nil {
					SetMovingSlotToList(hoveredSlot, 0.25)
				} else {
					SetMovingSlotToInventory(0.25)
				}
			}
		}
	}))
}

func SetMovingSpell(moving *data.InvSpellSlot, nWidth float64, offset pixel.Vec) {
	moving.NameMObj.Offset = offset
	moving.NameMObj.Sca = pixel.V(nWidth/data.TileSize-2, 1.)
	moving.NameMObj.SetRect(pixel.R(0., 0., nWidth, data.TileSize*3))
	moving.NameMObj.Hidden = false
	moving.NameMObj.Mask = util.White

	moving.NameTxt.SetText(moving.Slot.Name)
	moving.NameTxt.Obj.Offset.X = -(nWidth - data.TileSize*2.) * 0.5

	moving.NameLObj.Offset.X = -(nWidth - data.TileSize) * 0.5
	moving.NameRObj.Offset.X = (nWidth - data.TileSize) * 0.5

	moving.TierObj.Offset.X = -nWidth*0.5 - data.TileSize*1.5
	moving.TierTxt.Obj.Offset.X = -nWidth*0.5 - data.TileSize*1.5
	moving.TierTxt.SetText(util.RomanNumeral(moving.Slot.Tier))
}

func SetMainMovingSlot(slot *data.InvSpellSlot, nWidth float64, offset pixel.Vec, incTier, inventory bool, index int) {
	data.MovingSpellSlot.Slot = &data.SpellSlot{
		Tier:  slot.Slot.Tier,
		Spell: slot.Slot.Spell,
		Name:  slot.Slot.Name,
	}
	data.MovingSpellSlot.TierMoveIndex = index
	if !incTier {
		data.MovingSpellSlot.Slot.Tier = data.Missiles[slot.Slot.Spell][0].Tier
	}

	SetMovingSpell(&data.MovingSpellSlot.InvSpellSlot, nWidth, offset)
	data.MovingSpellSlot.NameMObj.Pos = data.InventoryView.ProjectWorld(data.TheInput.World)

	if inventory {
		data.SpellInventory.Spells[slot.Slot.Spell]--
	} else {
		data.MovingSpellSlot.PrevSlot = slot
		if incTier {
			slot.Slot.Tier = 0
		}
		slot.Slot.Spell = ""
		slot.Slot.Name = ""
	}
}

func MoveNewSlotToSlot(orig pixel.Vec, prevSlot *data.SpellSlot, nextSlot *data.InvSpellSlot, dur float64) {
	var entities []*ecs.Entity
	invSpellSlot := CreateMovingSpellSlot(func(e *ecs.Entity) {
		entities = append(entities, e)
	}, 0)
	invSpellSlot.NameMObj.Hidden = false
	invSpellSlot.Slot = &data.SpellSlot{
		Tier:  prevSlot.Tier,
		Spell: prevSlot.Spell,
		Name:  prevSlot.Name,
	}
	if nextSlot.Slot.Tier != 0 {
		invSpellSlot.Slot.Tier = nextSlot.Slot.Tier
	}
	SetMovingSpell(invSpellSlot, data.SlotWidth, pixel.ZV)
	invSpellSlot.NameMObj.Pos = orig
	fPos := nextSlot.View.ProjectedOut(nextSlot.NameMObj.Pos)
	interpolations := []*object.Interpolation{
		object.NewInterpolation(object.InterpolateX).
			AddGween(orig.X, fPos.X, dur, ease.OutCubic).
			SetOnComplete(func() {
				if nextSlot.Slot.Tier == 0 {
					nextSlot.Slot.Tier = invSpellSlot.Slot.Tier
				}
				nextSlot.Slot.Spell = invSpellSlot.Slot.Spell
				nextSlot.Slot.Name = invSpellSlot.Slot.Name
				for _, e := range entities {
					myecs.Manager.DisposeEntity(e)
				}
			}),
		object.NewInterpolation(object.InterpolateY).
			AddGween(orig.Y, fPos.Y, dur, ease.OutCubic),
	}
	if !nextSlot.View.PointInside(nextSlot.NameMObj.Pos) {
		interpolations = append(interpolations, object.NewInterpolation(object.InterpolateCol).
			AddGween(1., 0., dur*0.9, ease.Linear))
	}
	invSpellSlot.Entity.AddComponent(myecs.Interpolation, interpolations)
}

func SetMovingSlotToList(nextSlot *data.InvSpellSlot, dur float64) {
	if nextSlot.Slot.Tier != 0 {
		data.MovingSpellSlot.Slot.Tier = nextSlot.Slot.Tier
		data.MovingSpellSlot.TierTxt.SetText(util.RomanNumeral(data.MovingSpellSlot.Slot.Tier))
	}
	data.MovingSpellSlot.Moving = true
	fPos := nextSlot.View.ProjectedOut(nextSlot.View.Constrain(nextSlot.NameMObj.Pos))
	interpolations := []*object.Interpolation{
		object.NewInterpolation(object.InterpolateX).
			AddGween(data.MovingSpellSlot.NameMObj.Pos.X, fPos.X, dur, ease.OutCubic).
			SetOnComplete(func() {
				if nextSlot.Slot.Tier == 0 {
					nextSlot.Slot.Tier = data.MovingSpellSlot.Slot.Tier
				}
				nextSlot.Slot.Spell = data.MovingSpellSlot.Slot.Spell
				nextSlot.Slot.Name = data.MovingSpellSlot.Slot.Name
				data.MovingSpellSlot.PrevSlot = nil
				data.MovingSpellSlot.Slot = nil
				data.MovingSpellSlot.NameMObj.Hidden = true
				data.MovingSpellSlot.Moving = false
				data.MovingSpellSlot.NameMObj.Mask = util.White
			}),
		object.NewInterpolation(object.InterpolateY).
			AddGween(data.MovingSpellSlot.NameMObj.Pos.Y, fPos.Y, dur, ease.OutCubic),
		object.NewInterpolation(object.InterpolateOffX).
			AddGween(data.MovingSpellSlot.NameMObj.Offset.X, 0., dur, ease.OutCubic),
		object.NewInterpolation(object.InterpolateOffY).
			AddGween(data.MovingSpellSlot.NameMObj.Offset.Y, 0., dur, ease.OutCubic),
	}
	if !nextSlot.View.PointInside(nextSlot.NameMObj.Pos) {
		interpolations = append(interpolations, object.NewInterpolation(object.InterpolateCol).
			AddGween(1., 0., dur*0.9, ease.Linear))
	}
	data.MovingSpellSlot.Entity.AddComponent(myecs.Interpolation, interpolations)
}

func MoveNewSlotToInventory(orig pixel.Vec, prevSlot *data.SpellSlot, dur float64) {
	var entities []*ecs.Entity
	invSpellSlot := CreateMovingSpellSlot(func(e *ecs.Entity) {
		entities = append(entities, e)
	}, 0)
	invSpellSlot.NameMObj.Hidden = false
	invSpellSlot.Slot = &data.SpellSlot{
		Tier:  prevSlot.Tier,
		Spell: prevSlot.Spell,
		Name:  prevSlot.Name,
	}
	SetMovingSpell(invSpellSlot, data.SlotWidth, pixel.ZV)
	invSpellSlot.NameMObj.Pos = orig
	fPos := data.SpellInventory.ListViewObj.Pos
	for _, slot := range data.SpellInventory.Slots {
		if slot.Slot.Spell == invSpellSlot.Slot.Spell && slot.SlotNum != 0 {
			fPos = slot.View.ProjectedOut(slot.View.Constrain(slot.NameMObj.Pos))
			break
		}
	}
	invSpellSlot.Entity.AddComponent(myecs.Interpolation, []*object.Interpolation{
		object.NewInterpolation(object.InterpolateX).
			AddGween(orig.X, fPos.X, dur, ease.OutCubic).
			SetOnComplete(func() {
				data.SpellInventory.Spells[data.MovingSpellSlot.Slot.Spell]++
				for _, e := range entities {
					myecs.Manager.DisposeEntity(e)
				}
			}),
		object.NewInterpolation(object.InterpolateY).
			AddGween(orig.Y, fPos.Y, dur, ease.OutCubic),
	})
}

func SetMovingSlotToInventory(dur float64) {
	data.MovingSpellSlot.Moving = true
	fPos := data.SpellInventory.ListViewObj.Pos
	for _, slot := range data.SpellInventory.Slots {
		if slot.Slot.Spell == data.MovingSpellSlot.Slot.Spell && slot.SlotNum != 0 {
			fPos = slot.View.ProjectedOut(slot.View.Constrain(slot.NameMObj.Pos))
			break
		}
	}
	interpolations := []*object.Interpolation{
		object.NewInterpolation(object.InterpolateX).
			AddGween(data.MovingSpellSlot.NameMObj.Pos.X, fPos.X, dur, ease.OutCubic).
			SetOnComplete(func() {
				data.SpellInventory.Spells[data.MovingSpellSlot.Slot.Spell]++
				data.MovingSpellSlot.PrevSlot = nil
				data.MovingSpellSlot.Slot = nil
				data.MovingSpellSlot.NameMObj.Hidden = true
				data.MovingSpellSlot.Moving = false
				data.MovingSpellSlot.NameMObj.Mask = util.White
			}),
		object.NewInterpolation(object.InterpolateY).
			AddGween(data.MovingSpellSlot.NameMObj.Pos.Y, fPos.Y, dur, ease.OutCubic),
		object.NewInterpolation(object.InterpolateOffX).
			AddGween(data.MovingSpellSlot.NameMObj.Offset.X, 0., dur, ease.OutCubic),
		object.NewInterpolation(object.InterpolateOffY).
			AddGween(data.MovingSpellSlot.NameMObj.Offset.Y, 0., dur, ease.OutCubic),
	}
	data.MovingSpellSlot.Entity.AddComponent(myecs.Interpolation, interpolations)
}

func DrawMovingSpellSlots(win *pixelgl.Window) {
	img.Clear()
	systems.DrawSystem(win, 120)
	img.Batchers[data.UIKey].Draw(data.InventoryView.Canvas)
	systems.DrawSystem(win, 121)
	img.Clear()
	systems.DrawSystem(win, 122)
	img.Batchers[data.UIKey].Draw(data.InventoryView.Canvas)
	systems.DrawSystem(win, 123)
}

func AddMSSEntity(e *ecs.Entity) {
	data.MSSEntities = append(data.MSSEntities, e)
}

func DisposeMovingSpellSlot() {
	for _, e := range data.MSSEntities {
		myecs.Manager.DisposeEntity(e)
	}
	data.NewMovingSpellSlot()
}
