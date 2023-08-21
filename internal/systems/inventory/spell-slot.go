package inventory

import (
	"fmt"
	"github.com/bytearena/ecs"
	"github.com/faiface/pixel"
	"strconv"
	"timsims1717/magicmissile/internal/data"
	"timsims1717/magicmissile/internal/myecs"
	"timsims1717/magicmissile/pkg/img"
	"timsims1717/magicmissile/pkg/object"
	"timsims1717/magicmissile/pkg/typeface"
	"timsims1717/magicmissile/pkg/util"
	"timsims1717/magicmissile/pkg/viewport"
)

func CreateTowerSpellSlot(slot *data.SpellSlot, layerIndex, listIndex int, entityFn func(*ecs.Entity), vp *viewport.ViewPort) *data.InvSpellSlot {
	invSlot := &data.InvSpellSlot{}
	invSlot.Slot = slot
	invSlot.View = vp
	invSlot.SlotNum = listIndex

	invSlot.NameMObj = object.New()
	invSlot.NameMObj.Layer = 110 + layerIndex
	invSlot.NameMObj.Pos.Y -= data.SquareFrame.H() * float64(listIndex)
	invSlot.NameMObj.Pos.X += data.SquareFrame.W()*1.5 + data.SlotWidth*0.5
	invSlot.NameMObj.Sca = pixel.V(data.SlotWidth/data.TileSize-2, 1.)
	invSlot.NameMObj.SetRect(pixel.R(0., 0., data.SlotWidth, data.SquareFrame.H()))
	invSlot.NameMSpr = img.NewSprite("scroll_square_m", data.UIKey)
	invSlot.Entity = myecs.Manager.NewEntity().
		AddComponent(myecs.Object, invSlot.NameMObj).
		AddComponent(myecs.Drawable, invSlot.NameMSpr).
		AddComponent(myecs.Update, data.NewHoverClickFn(data.TheInput, vp, func(hvc *data.HoverClick) {
			if hvc.Hover && hvc.Input.Get("click").JustPressed() && invSlot.Slot.Spell != "" &&
				data.MovingSpellSlot.Slot == nil && !data.MovingSpellSlot.Moving {
				offset := invSlot.NameMObj.PostPos.Sub(hvc.View.ProjectWorld(data.TheInput.World))
				SetMainMovingSlot(invSlot, data.SlotWidth, offset, false, false, -1)
			}
		}))
	entityFn(invSlot.Entity)
	invSlot.NameLObj = object.New()
	invSlot.NameLObj.Layer = 110 + layerIndex
	invSlot.NameLObj.Offset.X -= (data.SlotWidth - data.TileSize) * 0.5
	invSlot.NameLSpr = img.NewSprite("scroll_square_l", data.UIKey)
	entityFn(myecs.Manager.NewEntity().
		AddComponent(myecs.Object, invSlot.NameLObj).
		AddComponent(myecs.Parent, invSlot.NameMObj).
		AddComponent(myecs.Drawable, invSlot.NameLSpr))
	invSlot.NameRObj = object.New()
	invSlot.NameRObj.Layer = 110 + layerIndex
	invSlot.NameRObj.Offset.X += (data.SlotWidth - data.TileSize) * 0.5
	invSlot.NameRSpr = img.NewSprite("scroll_square_r", data.UIKey)
	entityFn(myecs.Manager.NewEntity().
		AddComponent(myecs.Object, invSlot.NameRObj).
		AddComponent(myecs.Parent, invSlot.NameMObj).
		AddComponent(myecs.Drawable, invSlot.NameRSpr))
	invSlot.NameTxt = typeface.New("main", typeface.NewAlign(typeface.Left, typeface.Center), 1., 0.15, 0., 0.)
	invSlot.NameTxt.Obj.Layer = 113 + layerIndex
	invSlot.NameTxt.Obj.Offset.X -= (data.SlotWidth - data.TileSize*2.) * 0.5
	invSlot.NameTxt.Obj.Offset.Y += 6.
	invSlot.NameTxt.SetColor(data.ScrollText)
	invSlot.NameTxt.SetText(slot.Name)
	entityFn(myecs.Manager.NewEntity().
		AddComponent(myecs.Object, invSlot.NameTxt.Obj).
		AddComponent(myecs.Parent, invSlot.NameMObj).
		AddComponent(myecs.Drawable, invSlot.NameTxt).
		AddComponent(myecs.DrawTarget, vp))

	invSlot.SlotObj = object.New()
	invSlot.SlotObj.Layer = 110 + layerIndex
	invSlot.SlotObj.Offset.X -= data.SquareFrame.W()*1.5 + data.SlotWidth*0.5
	invSlot.SlotObj.SetRect(data.SquareFrame)
	invSlot.SlotSpr = img.NewSprite("scroll_square", data.UIKey)
	entityFn(myecs.Manager.NewEntity().
		AddComponent(myecs.Object, invSlot.SlotObj).
		AddComponent(myecs.Parent, invSlot.NameMObj).
		AddComponent(myecs.Drawable, invSlot.SlotSpr))
	invSlot.SlotTxt = typeface.New("main", typeface.NewAlign(typeface.Center, typeface.Center), 1., 0.15, 0., 0.)
	invSlot.SlotTxt.Obj.Layer = 113 + layerIndex
	invSlot.SlotTxt.Obj.Offset.Y += 6.
	invSlot.SlotTxt.SetColor(data.ScrollText)
	invSlot.SlotTxt.SetText(fmt.Sprintf("%d", listIndex+1))
	entityFn(myecs.Manager.NewEntity().
		AddComponent(myecs.Object, invSlot.SlotTxt.Obj).
		AddComponent(myecs.Parent, invSlot.SlotObj).
		AddComponent(myecs.Drawable, invSlot.SlotTxt).
		AddComponent(myecs.DrawTarget, vp))
	invSlot.TierObj = object.New()
	invSlot.TierObj.Layer = 110 + layerIndex
	invSlot.TierObj.Offset.X -= data.SquareFrame.W()*0.5 + data.SlotWidth*0.5
	invSlot.TierObj.SetRect(data.SquareFrame)
	invSlot.TierSpr = img.NewSprite("scroll_square", data.UIKey)
	entityFn(myecs.Manager.NewEntity().
		AddComponent(myecs.Object, invSlot.TierObj).
		AddComponent(myecs.Parent, invSlot.NameMObj).
		AddComponent(myecs.Drawable, invSlot.TierSpr).
		AddComponent(myecs.Update, data.NewHoverClickFn(data.TheInput, vp, func(hvc *data.HoverClick) {
			if hvc.Hover && hvc.Input.Get("click").JustPressed() &&
				data.MovingSpellSlot.Slot == nil && !data.MovingSpellSlot.Moving {
				offset := invSlot.NameMObj.PostPos.Sub(hvc.View.ProjectWorld(data.TheInput.World))
				SetMainMovingSlot(invSlot, data.SlotWidth, offset, true, false, layerIndex)
			}
		})))
	invSlot.TierTxt = typeface.New("main", typeface.NewAlign(typeface.Center, typeface.Center), 1., 0.15, 0., 0.)
	invSlot.TierTxt.Obj.Layer = 113 + layerIndex
	invSlot.TierTxt.Obj.Offset.Y += 6.
	invSlot.TierTxt.SetColor(data.ScrollText)
	invSlot.TierTxt.SetText(util.RomanNumeral(slot.Tier))
	entityFn(myecs.Manager.NewEntity().
		AddComponent(myecs.Object, invSlot.TierTxt.Obj).
		AddComponent(myecs.Parent, invSlot.TierObj).
		AddComponent(myecs.Drawable, invSlot.TierTxt).
		AddComponent(myecs.DrawTarget, vp))
	return invSlot
}

func CreateMovingSpellSlot(entityFn func(*ecs.Entity), layerIndex int) *data.InvSpellSlot {
	invSlot := &data.InvSpellSlot{}
	invSlot.NameMObj = object.New()
	invSlot.NameMObj.Layer = 120 + layerIndex
	invSlot.NameMSpr = img.NewSprite("scroll_square_m", data.UIKey)
	invSlot.NameMObj.HideChildren = true
	invSlot.NameMObj.Hidden = true
	invSlot.Entity = myecs.Manager.NewEntity().
		AddComponent(myecs.Object, invSlot.NameMObj).
		AddComponent(myecs.Drawable, invSlot.NameMSpr)
	entityFn(invSlot.Entity)
	invSlot.NameLObj = object.New()
	invSlot.NameLObj.Layer = 120 + layerIndex
	invSlot.NameLSpr = img.NewSprite("scroll_square_l", data.UIKey)
	entityFn(myecs.Manager.NewEntity().
		AddComponent(myecs.Object, invSlot.NameLObj).
		AddComponent(myecs.Parent, invSlot.NameMObj).
		AddComponent(myecs.Drawable, invSlot.NameLSpr))
	invSlot.NameRObj = object.New()
	invSlot.NameRObj.Layer = 120 + layerIndex
	invSlot.NameRSpr = img.NewSprite("scroll_square_r", data.UIKey)
	entityFn(myecs.Manager.NewEntity().
		AddComponent(myecs.Object, invSlot.NameRObj).
		AddComponent(myecs.Parent, invSlot.NameMObj).
		AddComponent(myecs.Drawable, invSlot.NameRSpr))
	invSlot.NameTxt = typeface.New("main", typeface.NewAlign(typeface.Left, typeface.Center), 1., 0.15, 0., 0.)
	invSlot.NameTxt.Obj.Layer = 121 + layerIndex
	invSlot.NameTxt.Obj.Offset.Y += 6.
	invSlot.NameTxt.SetColor(data.ScrollText)
	entityFn(myecs.Manager.NewEntity().
		AddComponent(myecs.Object, invSlot.NameTxt.Obj).
		AddComponent(myecs.Parent, invSlot.NameMObj).
		AddComponent(myecs.Drawable, invSlot.NameTxt).
		AddComponent(myecs.DrawTarget, data.InventoryView))
	invSlot.TierObj = object.New()
	invSlot.TierObj.Layer = 120 + layerIndex
	invSlot.TierSpr = img.NewSprite("scroll_square", data.UIKey)
	entityFn(myecs.Manager.NewEntity().
		AddComponent(myecs.Object, invSlot.TierObj).
		AddComponent(myecs.Parent, invSlot.NameMObj).
		AddComponent(myecs.Drawable, invSlot.TierSpr))
	invSlot.TierTxt = typeface.New("main", typeface.NewAlign(typeface.Center, typeface.Center), 1., 0.15, 0., 0.)
	invSlot.TierTxt.Obj.Layer = 121 + layerIndex
	invSlot.TierTxt.Obj.Offset.Y += 6.
	invSlot.TierTxt.SetColor(data.ScrollText)
	entityFn(myecs.Manager.NewEntity().
		AddComponent(myecs.Object, invSlot.TierTxt.Obj).
		AddComponent(myecs.Parent, invSlot.NameMObj).
		AddComponent(myecs.Drawable, invSlot.TierTxt).
		AddComponent(myecs.DrawTarget, data.InventoryView))
	return invSlot
}

func CreateInventorySpellSlot(slot *data.SpellSlot, count int, entityFn func(*ecs.Entity), vp *viewport.ViewPort) *data.InvSpellSlot {
	invSlot := &data.InvSpellSlot{}
	invSlot.Slot = slot
	invSlot.View = vp
	invSlot.SlotNum = count

	invSlot.NameMObj = object.New()
	invSlot.NameMObj.Layer = 116
	invSlot.NameMObj.Pos.X = 0
	invSlot.NameMObj.Sca = pixel.V(data.SlotWidth/data.TileSize-2, 1.)
	invSlot.NameMObj.SetRect(pixel.R(0., 0., data.SlotWidth, data.SquareFrame.H()))
	invSlot.NameMObj.HideChildren = true
	invSlot.NameMSpr = img.NewSprite("scroll_square_m", data.UIKey)
	invSlot.Entity = myecs.Manager.NewEntity().
		AddComponent(myecs.Object, invSlot.NameMObj).
		AddComponent(myecs.Drawable, invSlot.NameMSpr).
		AddComponent(myecs.Update, data.NewHoverClickFn(data.TheInput, vp, func(hvc *data.HoverClick) {
			if hvc.Hover && hvc.Input.Get("click").JustPressed() && invSlot.Slot.Spell != "" &&
				data.MovingSpellSlot.Slot == nil && !data.MovingSpellSlot.Moving &&
				data.SpellInventory.Spells[invSlot.Slot.Spell] != 0 {
				offset := invSlot.NameMObj.PostPos.Sub(hvc.View.ProjectWorld(data.TheInput.World))
				SetMainMovingSlot(invSlot, data.SlotWidth, offset, false, true, -1)
			}
		}))
	entityFn(invSlot.Entity)
	invSlot.NameLObj = object.New()
	invSlot.NameLObj.Layer = 116
	invSlot.NameLObj.Offset.X -= (data.SlotWidth - data.TileSize) * 0.5
	invSlot.NameLSpr = img.NewSprite("scroll_square_l", data.UIKey)
	entityFn(myecs.Manager.NewEntity().
		AddComponent(myecs.Object, invSlot.NameLObj).
		AddComponent(myecs.Parent, invSlot.NameMObj).
		AddComponent(myecs.Drawable, invSlot.NameLSpr))
	invSlot.NameRObj = object.New()
	invSlot.NameRObj.Layer = 116
	invSlot.NameRObj.Offset.X += (data.SlotWidth - data.TileSize) * 0.5
	invSlot.NameRSpr = img.NewSprite("scroll_square_r", data.UIKey)
	entityFn(myecs.Manager.NewEntity().
		AddComponent(myecs.Object, invSlot.NameRObj).
		AddComponent(myecs.Parent, invSlot.NameMObj).
		AddComponent(myecs.Drawable, invSlot.NameRSpr))
	invSlot.NameTxt = typeface.New("main", typeface.NewAlign(typeface.Left, typeface.Center), 1., 0.15, 0., 0.)
	invSlot.NameTxt.Obj.Layer = 117
	invSlot.NameTxt.Obj.Offset.X -= (data.SlotWidth - data.TileSize*2.) * 0.5
	invSlot.NameTxt.Obj.Offset.Y += 6.
	invSlot.NameTxt.SetColor(data.ScrollText)
	invSlot.NameTxt.SetText(slot.Name)
	entityFn(myecs.Manager.NewEntity().
		AddComponent(myecs.Object, invSlot.NameTxt.Obj).
		AddComponent(myecs.Parent, invSlot.NameMObj).
		AddComponent(myecs.Drawable, invSlot.NameTxt).
		AddComponent(myecs.DrawTarget, vp))

	invSlot.SlotObj = object.New()
	invSlot.SlotObj.Layer = 116
	invSlot.SlotObj.Offset.X -= data.SquareFrame.W()*1.5 + data.SlotWidth*0.5
	invSlot.SlotObj.SetRect(data.SquareFrame)
	invSlot.SlotSpr = img.NewSprite("scroll_square", data.UIKey)
	entityFn(myecs.Manager.NewEntity().
		AddComponent(myecs.Object, invSlot.SlotObj).
		AddComponent(myecs.Parent, invSlot.NameMObj).
		AddComponent(myecs.Drawable, invSlot.SlotSpr))
	invSlot.SlotTxt = typeface.New("main", typeface.NewAlign(typeface.Center, typeface.Center), 1., 0.15, 0., 0.)
	invSlot.SlotTxt.Obj.Layer = 117
	invSlot.SlotTxt.Obj.Offset.X -= data.SquareFrame.W()*1.5 + data.SlotWidth*0.5
	invSlot.SlotTxt.Obj.Offset.Y += 6.
	invSlot.SlotTxt.SetColor(data.ScrollText)
	invSlot.SlotTxt.SetText(strconv.Itoa(count))
	entityFn(myecs.Manager.NewEntity().
		AddComponent(myecs.Object, invSlot.SlotTxt.Obj).
		AddComponent(myecs.Parent, invSlot.NameMObj).
		AddComponent(myecs.Drawable, invSlot.SlotTxt).
		AddComponent(myecs.DrawTarget, vp))
	invSlot.TierObj = object.New()
	invSlot.TierObj.Layer = 116
	invSlot.TierObj.Offset.X -= data.SquareFrame.W()*0.5 + data.SlotWidth*0.5
	invSlot.TierObj.SetRect(data.SquareFrame)
	invSlot.TierSpr = img.NewSprite("scroll_square", data.UIKey)
	entityFn(myecs.Manager.NewEntity().
		AddComponent(myecs.Object, invSlot.TierObj).
		AddComponent(myecs.Parent, invSlot.NameMObj).
		AddComponent(myecs.Drawable, invSlot.TierSpr))
	invSlot.TierTxt = typeface.New("main", typeface.NewAlign(typeface.Center, typeface.Center), 1., 0.15, 0., 0.)
	invSlot.TierTxt.Obj.Layer = 117
	invSlot.TierTxt.Obj.Offset.X -= data.SquareFrame.W()*0.5 + data.SlotWidth*0.5
	invSlot.TierTxt.Obj.Offset.Y += 6.
	invSlot.TierTxt.SetColor(data.ScrollText)
	invSlot.TierTxt.SetText(util.RomanNumeral(slot.Tier))
	entityFn(myecs.Manager.NewEntity().
		AddComponent(myecs.Object, invSlot.TierTxt.Obj).
		AddComponent(myecs.Parent, invSlot.NameMObj).
		AddComponent(myecs.Drawable, invSlot.TierTxt).
		AddComponent(myecs.DrawTarget, vp))
	return invSlot
}
