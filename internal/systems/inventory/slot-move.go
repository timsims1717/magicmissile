package inventory

import (
	"github.com/bytearena/ecs"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"timsims1717/magicmissile/internal/data"
	"timsims1717/magicmissile/internal/myecs"
	"timsims1717/magicmissile/internal/systems"
	"timsims1717/magicmissile/pkg/img"
	"timsims1717/magicmissile/pkg/object"
	"timsims1717/magicmissile/pkg/typeface"
	"timsims1717/magicmissile/pkg/util"
)

func CreateMovingSpellSlot() {
	DisposeMovingSpellSlot()
	data.MovingSpellSlot = &data.InvSpellSlot{}
	data.MovingSpellSlot.NameMObj = object.New()
	data.MovingSpellSlot.NameMObj.Layer = 120
	data.MovingSpellSlot.NameMSpr = img.NewSprite("scroll_square_m", data.UIKey)
	data.MovingSpellSlot.NameMObj.HideChildren = true
	data.MovingSpellSlot.NameMObj.Hidden = true
	AddMSSEntity(myecs.Manager.NewEntity().
		AddComponent(myecs.Object, data.MovingSpellSlot.NameMObj).
		AddComponent(myecs.Drawable, data.MovingSpellSlot.NameMSpr).
		AddComponent(myecs.Update, data.NewHoverClickFn(data.TheInput, data.InventoryView, func(hvc *data.HoverClick) {
			data.MovingSpellSlot.NameMObj.Pos = data.InventoryView.Projected(hvc.Input.World)
			if data.MovingSpellSlot.Slot != nil && data.MovingSpellSlot.Slot.Spell != "" {
				hoveredSlot := GetHoveredSlot()
				if hvc.Input.Get("click").Pressed() && hoveredSlot != nil {

				} else if hvc.Input.Get("click").JustReleased() {
					if hoveredSlot != nil && hoveredSlot.Slot.Tier >= data.MovingSpellSlot.Slot.Tier {
						// if the hovered slot is a legal place to put the spell
						if data.MovingSpellSlot.PrevSlot != nil && hoveredSlot.Slot.Spell != "" &&
							data.MovingSpellSlot.PrevSlot.Tier >= hoveredSlot.Slot.Tier {
							// if we can switch the slots
							data.MovingSpellSlot.PrevSlot.Spell = hoveredSlot.Slot.Spell
							data.MovingSpellSlot.PrevSlot.Name = hoveredSlot.Slot.Name
						} else {
							// put the replaced spell slot into the spell storage
						}
						hoveredSlot.Slot.Spell = data.MovingSpellSlot.Slot.Spell
						hoveredSlot.Slot.Name = data.MovingSpellSlot.Slot.Name
					} else if data.MovingSpellSlot.PrevSlot != nil {
						// if not, put it back where you got it from
						data.MovingSpellSlot.PrevSlot.Spell = data.MovingSpellSlot.Slot.Spell
						data.MovingSpellSlot.PrevSlot.Name = data.MovingSpellSlot.Slot.Name
					} else {
						// if you don't know where it was, put it into the spell storage
					}
					data.MovingSpellSlot.PrevSlot = nil
					data.MovingSpellSlot.Slot = nil
					data.MovingSpellSlot.NameMObj.Hidden = true
				}
			}
		})))
	data.MovingSpellSlot.NameLObj = object.New()
	data.MovingSpellSlot.NameLObj.Layer = 120
	data.MovingSpellSlot.NameLSpr = img.NewSprite("scroll_square_l", data.UIKey)
	AddMSSEntity(myecs.Manager.NewEntity().
		AddComponent(myecs.Object, data.MovingSpellSlot.NameLObj).
		AddComponent(myecs.Parent, data.MovingSpellSlot.NameMObj).
		AddComponent(myecs.Drawable, data.MovingSpellSlot.NameLSpr))
	data.MovingSpellSlot.NameRObj = object.New()
	data.MovingSpellSlot.NameRObj.Layer = 120
	data.MovingSpellSlot.NameRSpr = img.NewSprite("scroll_square_r", data.UIKey)
	AddMSSEntity(myecs.Manager.NewEntity().
		AddComponent(myecs.Object, data.MovingSpellSlot.NameRObj).
		AddComponent(myecs.Parent, data.MovingSpellSlot.NameMObj).
		AddComponent(myecs.Drawable, data.MovingSpellSlot.NameRSpr))
	data.MovingSpellSlot.NameTxt = typeface.New("main", typeface.NewAlign(typeface.Left, typeface.Center), 1., 0.15, 0., 0.)
	data.MovingSpellSlot.NameTxt.Obj.Layer = 121
	data.MovingSpellSlot.NameTxt.Obj.Offset.Y += 6.
	data.MovingSpellSlot.NameTxt.SetColor(data.ScrollText)
	AddMSSEntity(myecs.Manager.NewEntity().
		AddComponent(myecs.Object, data.MovingSpellSlot.NameTxt.Obj).
		AddComponent(myecs.Parent, data.MovingSpellSlot.NameMObj).
		AddComponent(myecs.Drawable, data.MovingSpellSlot.NameTxt).
		AddComponent(myecs.DrawTarget, data.InventoryView))
	data.MovingSpellSlot.TierObj = object.New()
	data.MovingSpellSlot.TierObj.Layer = 120
	data.MovingSpellSlot.TierSpr = img.NewSprite("scroll_square", data.UIKey)
	AddMSSEntity(myecs.Manager.NewEntity().
		AddComponent(myecs.Object, data.MovingSpellSlot.TierObj).
		AddComponent(myecs.Parent, data.MovingSpellSlot.NameMObj).
		AddComponent(myecs.Drawable, data.MovingSpellSlot.TierSpr))
	data.MovingSpellSlot.TierTxt = typeface.New("main", typeface.NewAlign(typeface.Center, typeface.Center), 1., 0.15, 0., 0.)
	data.MovingSpellSlot.TierTxt.Obj.Layer = 121
	data.MovingSpellSlot.TierTxt.Obj.Offset.Y += 6.
	data.MovingSpellSlot.TierTxt.SetColor(data.ScrollText)
	AddMSSEntity(myecs.Manager.NewEntity().
		AddComponent(myecs.Object, data.MovingSpellSlot.TierTxt.Obj).
		AddComponent(myecs.Parent, data.MovingSpellSlot.NameMObj).
		AddComponent(myecs.Drawable, data.MovingSpellSlot.TierTxt).
		AddComponent(myecs.DrawTarget, data.InventoryView))
}

func SetMovingSpell(slot *data.SpellSlot, nWidth float64, offset pixel.Vec, incTier bool) {
	data.MovingSpellSlot.Slot = &data.SpellSlot{
		Tier:  slot.Tier,
		Spell: slot.Spell,
		Name:  slot.Name,
	}
	if !incTier {
		data.MovingSpellSlot.Slot.Tier = data.Missiles[slot.Spell][0].Tier
	}
	data.MovingSpellSlot.PrevSlot = slot

	data.MovingSpellSlot.NameMObj.Offset = offset
	data.MovingSpellSlot.NameMObj.Sca = pixel.V(nWidth/data.TileSize-2, 1.)
	data.MovingSpellSlot.NameMObj.SetRect(pixel.R(0., 0., nWidth, data.TileSize*3))
	data.MovingSpellSlot.NameMObj.Hidden = false

	data.MovingSpellSlot.NameTxt.SetText(data.MovingSpellSlot.Slot.Name)
	data.MovingSpellSlot.NameTxt.Obj.Offset.X = -(nWidth - data.TileSize*2.) * 0.5

	data.MovingSpellSlot.NameLObj.Offset.X = -(nWidth - data.TileSize) * 0.5
	data.MovingSpellSlot.NameRObj.Offset.X = (nWidth - data.TileSize) * 0.5

	data.MovingSpellSlot.TierObj.Offset.X = -nWidth*0.5 - data.TileSize*1.5
	data.MovingSpellSlot.TierTxt.Obj.Offset.X = -nWidth*0.5 - data.TileSize*1.5
	data.MovingSpellSlot.TierTxt.SetText(util.RomanNumeral(data.MovingSpellSlot.Slot.Tier))

	if incTier {
		slot.Tier = 0
	}
	slot.Spell = ""
	slot.Name = ""
}

func DrawMovingSpellSlot(win *pixelgl.Window) {
	img.Clear()
	systems.DrawSystem(win, 120)
	img.Batchers[data.UIKey].Draw(data.InventoryView.Canvas)
	systems.DrawSystem(win, 121)
}

func AddMSSEntity(e *ecs.Entity) {
	data.MSSEntities = append(data.MSSEntities, e)
}

func DisposeMovingSpellSlot() {
	for _, e := range data.MSSEntities {
		myecs.Manager.DisposeEntity(e)
	}
	data.MovingSpellSlot = nil
}
