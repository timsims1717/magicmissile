package systems

import (
	"fmt"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
	"image/color"
	"math"
	"timsims1717/magicmissile/internal/data"
	"timsims1717/magicmissile/internal/myecs"
	"timsims1717/magicmissile/pkg/gween64/ease"
	"timsims1717/magicmissile/pkg/img"
	"timsims1717/magicmissile/pkg/object"
	"timsims1717/magicmissile/pkg/timing"
	"timsims1717/magicmissile/pkg/typeface"
	"timsims1717/magicmissile/pkg/util"
	"timsims1717/magicmissile/pkg/viewport"
)

func CreateTowerScrolls() {
	CreateTowerScroll(0)
	CreateTowerScroll(1)
	CreateTowerScroll(2)
}

func CreateTowerScroll(index int) {
	var scroll *data.TowerScroll
	var posStr string
	switch index {
	case 0:
		data.LeftTowerScroll = &data.TowerScroll{}
		scroll = data.LeftTowerScroll
		scroll.Scroll = CreateScroll(data.LeftScrollStart, 101, pixel.V(data.TowerScrollWidth*data.TileSize, data.TileSize), pixel.V(data.TowerScrollWidth*data.TileSize, data.TowerScrollHeight*data.TileSize))
		posStr = "Left"
	case 1:
		data.MidTowerScroll = &data.TowerScroll{}
		scroll = data.MidTowerScroll
		scroll.Scroll = CreateScroll(data.MidScrollStart, 101, pixel.V(data.TowerScrollWidth*data.TileSize, data.TileSize), pixel.V(data.TowerScrollWidth*data.TileSize, data.TowerScrollHeight*data.TileSize))
		posStr = "Middle"
	case 2:
		data.RightTowerScroll = &data.TowerScroll{}
		scroll = data.RightTowerScroll
		scroll.Scroll = CreateScroll(data.RightScrollStart, 101, pixel.V(data.TowerScrollWidth*data.TileSize, data.TileSize), pixel.V(data.TowerScrollWidth*data.TileSize, data.TowerScrollHeight*data.TileSize))
		posStr = "Right"
	default:
		panic(fmt.Sprintf("error: incorrect index %d when creating inventory scroll", index))
	}
	scroll.Scroll.Freeze = true
	scroll.Tower = data.Towers[index]

	rect := img.Batchers[scroll.Tower.Sprite.Batch].GetSprite(scroll.Tower.Sprite.Key).Frame()
	CreateTowerView(scroll, rect)
	CreateTowerText(scroll, posStr, rect)
	CreateSpellList(scroll, index)
}

func CreateTowerView(scroll *data.TowerScroll, rect pixel.Rect) {
	scroll.TowerViewObj = object.New()
	scroll.TowerViewObj.HideChildren = true
	scroll.TowerViewObj.Offset = scroll.Scroll.TLPos.Add(data.TowerViewPos)
	scroll.TowerViewObj.Offset.Y -= rect.H() * 0.5
	scroll.TowerViewObj.Offset.X += rect.W() * 0.5
	scroll.AddEntity(myecs.Manager.NewEntity().
		AddComponent(myecs.Object, scroll.TowerViewObj).
		AddComponent(myecs.Parent, scroll.Scroll.Object))
	mmObj := object.New()
	mmObj.Sca = pixel.V(rect.W()/data.TileSize, rect.H()/data.TileSize)
	mmObj.Layer = 102
	scroll.AddEntity(myecs.Manager.NewEntity().
		AddComponent(myecs.Object, mmObj).
		AddComponent(myecs.Parent, scroll.TowerViewObj).
		AddComponent(myecs.Drawable, img.NewSprite("scroll_square_mm", data.UIKey)))
	tmObj := object.New()
	tmObj.Offset.Y = rect.H() * 0.5
	tmObj.Sca = pixel.V(rect.W()/data.TileSize, 1.)
	tmObj.Layer = 102
	scroll.AddEntity(myecs.Manager.NewEntity().
		AddComponent(myecs.Object, tmObj).
		AddComponent(myecs.Parent, scroll.TowerViewObj).
		AddComponent(myecs.Drawable, img.NewSprite("scroll_square_tm", data.UIKey)))
	bmObj := object.New()
	bmObj.Offset.Y = rect.H() * -0.5
	bmObj.Sca = pixel.V(rect.W()/data.TileSize, 1.)
	bmObj.Layer = 102
	scroll.AddEntity(myecs.Manager.NewEntity().
		AddComponent(myecs.Object, bmObj).
		AddComponent(myecs.Parent, scroll.TowerViewObj).
		AddComponent(myecs.Drawable, img.NewSprite("scroll_square_bm", data.UIKey)))
	mlObj := object.New()
	mlObj.Offset.X = rect.W() * -0.5
	mlObj.Sca = pixel.V(1., rect.H()/data.TileSize)
	mlObj.Layer = 102
	scroll.AddEntity(myecs.Manager.NewEntity().
		AddComponent(myecs.Object, mlObj).
		AddComponent(myecs.Parent, scroll.TowerViewObj).
		AddComponent(myecs.Drawable, img.NewSprite("scroll_square_ml", data.UIKey)))
	mrObj := object.New()
	mrObj.Offset.X = rect.W() * 0.5
	mrObj.Sca = pixel.V(1., rect.H()/data.TileSize)
	mrObj.Layer = 102
	scroll.AddEntity(myecs.Manager.NewEntity().
		AddComponent(myecs.Object, mrObj).
		AddComponent(myecs.Parent, scroll.TowerViewObj).
		AddComponent(myecs.Drawable, img.NewSprite("scroll_square_mr", data.UIKey)))
	tlObj := object.New()
	tlObj.Offset.Y = rect.H() * 0.5
	tlObj.Offset.X = rect.W() * -0.5
	tlObj.Layer = 102
	scroll.AddEntity(myecs.Manager.NewEntity().
		AddComponent(myecs.Object, tlObj).
		AddComponent(myecs.Parent, scroll.TowerViewObj).
		AddComponent(myecs.Drawable, img.NewSprite("scroll_square_tl", data.UIKey)))
	trObj := object.New()
	trObj.Offset.Y = rect.H() * 0.5
	trObj.Offset.X = rect.W() * 0.5
	trObj.Layer = 102
	scroll.AddEntity(myecs.Manager.NewEntity().
		AddComponent(myecs.Object, trObj).
		AddComponent(myecs.Parent, scroll.TowerViewObj).
		AddComponent(myecs.Drawable, img.NewSprite("scroll_square_tr", data.UIKey)))
	blObj := object.New()
	blObj.Offset.Y = rect.H() * -0.5
	blObj.Offset.X = rect.W() * -0.5
	blObj.Layer = 102
	scroll.AddEntity(myecs.Manager.NewEntity().
		AddComponent(myecs.Object, blObj).
		AddComponent(myecs.Parent, scroll.TowerViewObj).
		AddComponent(myecs.Drawable, img.NewSprite("scroll_square_bl", data.UIKey)))
	brObj := object.New()
	brObj.Offset.Y = rect.H() * -0.5
	brObj.Offset.X = rect.W() * 0.5
	brObj.Layer = 102
	scroll.AddEntity(myecs.Manager.NewEntity().
		AddComponent(myecs.Object, brObj).
		AddComponent(myecs.Parent, scroll.TowerViewObj).
		AddComponent(myecs.Drawable, img.NewSprite("scroll_square_br", data.UIKey)))
	tvObj := object.New()
	tvObj.Layer = 102
	tvObj.Offset.Y = -data.TileSize*0.5 + 5.
	scroll.AddEntity(myecs.Manager.NewEntity().
		AddComponent(myecs.Object, tvObj).
		AddComponent(myecs.Parent, scroll.TowerViewObj).
		AddComponent(myecs.Drawable, scroll.Tower.Sprite))
}

func CreateTowerText(scroll *data.TowerScroll, pos string, rect pixel.Rect) {
	scroll.TitleText = typeface.New(nil, "main", typeface.NewAlign(typeface.Left, typeface.Top), 1., 0.15, 300., 0.)
	scroll.TitleText.Obj.Layer = 103
	scroll.TitleText.SetOffset(scroll.Scroll.TLPos.Add(data.TowerTitlePos))
	scroll.TitleText.SetColor(data.ScrollText)
	scroll.TitleText.SetText(fmt.Sprintf("%s Wizard's Tower", pos))
	scroll.AddEntity(myecs.Manager.NewEntity().
		AddComponent(myecs.Object, scroll.TitleText.Obj).
		AddComponent(myecs.Parent, scroll.Scroll.Object).
		AddComponent(myecs.Drawable, scroll.TitleText).
		AddComponent(myecs.DrawTarget, data.InventoryView.Canvas))

	scroll.SubclassText = typeface.New(nil, "main", typeface.NewAlign(typeface.Left, typeface.Top), 1., 0.12, 0., 0.)
	scroll.SubclassText.Obj.Layer = 103
	scroll.SubclassText.SetOffset(scroll.Scroll.TLPos.Add(data.SubtitlePos))
	scroll.SubclassText.Obj.Offset.Y -= scroll.TitleText.Height
	scroll.SubclassText.SetColor(data.ScrollText)
	scroll.SubclassText.SetText("Specialization: Arcanist")
	scroll.AddEntity(myecs.Manager.NewEntity().
		AddComponent(myecs.Object, scroll.SubclassText.Obj).
		AddComponent(myecs.Parent, scroll.Scroll.Object).
		AddComponent(myecs.Drawable, scroll.SubclassText).
		AddComponent(myecs.DrawTarget, data.InventoryView.Canvas))

	scroll.LevelText = typeface.New(nil, "main", typeface.NewAlign(typeface.Left, typeface.Top), 1., 0.12, 0., 0.)
	scroll.LevelText.Obj.Layer = 103
	scroll.LevelText.SetOffset(scroll.Scroll.TLPos.Add(data.SubtitlePos))
	scroll.LevelText.Obj.Offset.Y = scroll.SubclassText.Obj.Offset.Y - scroll.SubclassText.Height
	scroll.LevelText.SetColor(data.ScrollText)
	scroll.LevelText.SetText("Level III")
	scroll.AddEntity(myecs.Manager.NewEntity().
		AddComponent(myecs.Object, scroll.LevelText.Obj).
		AddComponent(myecs.Parent, scroll.Scroll.Object).
		AddComponent(myecs.Drawable, scroll.LevelText).
		AddComponent(myecs.DrawTarget, data.InventoryView.Canvas))

	scroll.TableHeadY = math.Min(scroll.LevelText.Obj.Offset.Y-scroll.LevelText.Height, scroll.TowerViewObj.Offset.Y-rect.H()*0.5) - 15.

	scroll.SlotHead = typeface.New(nil, "main", typeface.NewAlign(typeface.Left, typeface.Top), 1., 0.10, 0., 0.)
	scroll.SlotHead.Obj.Layer = 103
	scroll.SlotHead.SetOffset(pixel.V(scroll.Scroll.TLPos.X+data.SlotsHeadX, scroll.TableHeadY))
	scroll.SlotHead.SetColor(data.ScrollText)
	scroll.SlotHead.SetText("Slot")
	scroll.AddEntity(myecs.Manager.NewEntity().
		AddComponent(myecs.Object, scroll.SlotHead.Obj).
		AddComponent(myecs.Parent, scroll.Scroll.Object).
		AddComponent(myecs.Drawable, scroll.SlotHead).
		AddComponent(myecs.DrawTarget, data.InventoryView.Canvas))

	scroll.LvlHead = typeface.New(nil, "main", typeface.NewAlign(typeface.Left, typeface.Top), 1., 0.10, 0., 0.)
	scroll.LvlHead.Obj.Layer = 103
	scroll.LvlHead.SetOffset(pixel.V(scroll.Scroll.TLPos.X+data.LevelHeadX, scroll.TableHeadY))
	scroll.LvlHead.SetColor(data.ScrollText)
	scroll.LvlHead.SetText("Lvl")
	scroll.AddEntity(myecs.Manager.NewEntity().
		AddComponent(myecs.Object, scroll.LvlHead.Obj).
		AddComponent(myecs.Parent, scroll.Scroll.Object).
		AddComponent(myecs.Drawable, scroll.LvlHead).
		AddComponent(myecs.DrawTarget, data.InventoryView.Canvas))

	scroll.NameHead = typeface.New(nil, "main", typeface.NewAlign(typeface.Left, typeface.Top), 1., 0.10, 0., 0.)
	scroll.NameHead.Obj.Layer = 103
	scroll.NameHead.SetOffset(pixel.V(scroll.Scroll.TLPos.X+data.NameHeadX, scroll.TableHeadY))
	scroll.NameHead.SetColor(data.ScrollText)
	scroll.NameHead.SetText("Spell")
	scroll.AddEntity(myecs.Manager.NewEntity().
		AddComponent(myecs.Object, scroll.NameHead.Obj).
		AddComponent(myecs.Parent, scroll.Scroll.Object).
		AddComponent(myecs.Drawable, scroll.NameHead).
		AddComponent(myecs.DrawTarget, data.InventoryView.Canvas))
}

func CreateSpellList(scroll *data.TowerScroll, index int) {
	rect := pixel.R(0, 0, (data.TowerScrollWidth-3)*data.TileSize, (data.TowerScrollHeight-10)*data.TileSize-math.Abs(scroll.TableHeadY))
	scroll.ListViewObj = object.New()
	scroll.ListViewObj.HideChildren = true
	scroll.ListViewObj.Offset = pixel.V(scroll.Scroll.TLPos.X+rect.W()*0.5+12., scroll.TableHeadY-40.-rect.H()*0.5)
	scroll.AddEntity(myecs.Manager.NewEntity().
		AddComponent(myecs.Object, scroll.ListViewObj).
		AddComponent(myecs.Parent, scroll.Scroll.Object))
	mmObj := object.New()
	mmObj.Sca = pixel.V(rect.W()/data.TileSize, rect.H()/data.TileSize)
	mmObj.Layer = 102
	scroll.AddEntity(myecs.Manager.NewEntity().
		AddComponent(myecs.Object, mmObj).
		AddComponent(myecs.Parent, scroll.ListViewObj).
		AddComponent(myecs.Drawable, img.NewSprite("list_bg_m", data.UIKey)))
	tmObj := object.New()
	tmObj.Offset.Y = rect.H()*0.5 + 3.
	tmObj.Sca = pixel.V(rect.W()/data.TileSize, 1.)
	tmObj.Layer = 102
	scroll.AddEntity(myecs.Manager.NewEntity().
		AddComponent(myecs.Object, tmObj).
		AddComponent(myecs.Parent, scroll.ListViewObj).
		AddComponent(myecs.Drawable, img.NewSprite("list_bg_sv", data.UIKey)))
	bmObj := object.New()
	bmObj.Offset.Y = rect.H()*-0.5 - 3.
	bmObj.Sca = pixel.V(rect.W()/data.TileSize, 1.)
	bmObj.Layer = 102
	bmObj.Flop = true
	scroll.AddEntity(myecs.Manager.NewEntity().
		AddComponent(myecs.Object, bmObj).
		AddComponent(myecs.Parent, scroll.ListViewObj).
		AddComponent(myecs.Drawable, img.NewSprite("list_bg_sv", data.UIKey)))
	mlObj := object.New()
	mlObj.Offset.X = rect.W()*-0.5 - 3.
	mlObj.Sca = pixel.V(1., rect.H()/data.TileSize)
	mlObj.Layer = 102
	scroll.AddEntity(myecs.Manager.NewEntity().
		AddComponent(myecs.Object, mlObj).
		AddComponent(myecs.Parent, scroll.ListViewObj).
		AddComponent(myecs.Drawable, img.NewSprite("list_bg_sh", data.UIKey)))
	mrObj := object.New()
	mrObj.Offset.X = rect.W()*0.5 + 3.
	mrObj.Sca = pixel.V(1., rect.H()/data.TileSize)
	mrObj.Layer = 102
	mrObj.Flip = true
	scroll.AddEntity(myecs.Manager.NewEntity().
		AddComponent(myecs.Object, mrObj).
		AddComponent(myecs.Parent, scroll.ListViewObj).
		AddComponent(myecs.Drawable, img.NewSprite("list_bg_sh", data.UIKey)))
	tlObj := object.New()
	tlObj.Offset.Y = rect.H()*0.5 + 3.
	tlObj.Offset.X = rect.W()*-0.5 - 3.
	tlObj.Layer = 102
	scroll.AddEntity(myecs.Manager.NewEntity().
		AddComponent(myecs.Object, tlObj).
		AddComponent(myecs.Parent, scroll.ListViewObj).
		AddComponent(myecs.Drawable, img.NewSprite("list_bg_c", data.UIKey)))
	trObj := object.New()
	trObj.Offset.Y = rect.H()*0.5 + 3.
	trObj.Offset.X = rect.W()*0.5 + 3.
	trObj.Layer = 102
	trObj.Flip = true
	scroll.AddEntity(myecs.Manager.NewEntity().
		AddComponent(myecs.Object, trObj).
		AddComponent(myecs.Parent, scroll.ListViewObj).
		AddComponent(myecs.Drawable, img.NewSprite("list_bg_c", data.UIKey)))
	blObj := object.New()
	blObj.Offset.Y = rect.H()*-0.5 - 3.
	blObj.Offset.X = rect.W()*-0.5 - 3.
	blObj.Layer = 102
	blObj.Flop = true
	scroll.AddEntity(myecs.Manager.NewEntity().
		AddComponent(myecs.Object, blObj).
		AddComponent(myecs.Parent, scroll.ListViewObj).
		AddComponent(myecs.Drawable, img.NewSprite("list_bg_c", data.UIKey)))
	brObj := object.New()
	brObj.Offset.Y = rect.H()*-0.5 - 3.
	brObj.Offset.X = rect.W()*0.5 + 3.
	brObj.Layer = 102
	brObj.Flip = true
	brObj.Flop = true
	scroll.AddEntity(myecs.Manager.NewEntity().
		AddComponent(myecs.Object, brObj).
		AddComponent(myecs.Parent, scroll.ListViewObj).
		AddComponent(myecs.Drawable, img.NewSprite("list_bg_c", data.UIKey)))

	scroll.ListView = viewport.New(nil)
	scroll.ListView.SetRect(pixel.R(0., 0., rect.W()-data.TileSize*1.5-2., rect.H()+data.TileSize+6.))
	scroll.ListView.CamPos = pixel.V(data.SlotListViewX, data.SlotListViewY)
	scroll.ListView.PortPos = scroll.Scroll.Object.Pos.Add(scroll.ListViewObj.Offset)
	scroll.ListView.PortPos.X -= data.TileSize * 1.5

	frame := img.Batchers[data.UIKey].GetSprite("scroll_square").Frame()
	nWidth := scroll.ListView.Rect.W() - frame.W()*2.
	var origSlot *object.Object
	for i, slot := range scroll.Tower.Slots {
		slotObj := object.New()
		if i == 0 {
			origSlot = slotObj
		}
		slotObj.Layer = 110 + index
		slotObj.Offset.Y -= frame.H() * float64(i)
		scroll.AddEntity(myecs.Manager.NewEntity().
			AddComponent(myecs.Object, slotObj).
			AddComponent(myecs.Parent, origSlot).
			AddComponent(myecs.Drawable, img.NewSprite("scroll_square", data.UIKey)))
		slotTxt := typeface.New(nil, "main", typeface.NewAlign(typeface.Center, typeface.Center), 1., 0.15, 0., 0.)
		slotTxt.Obj.Layer = 113 + index
		slotTxt.SetOffset(slotObj.Offset)
		slotTxt.Obj.Offset.Y += 6.
		slotTxt.SetColor(data.ScrollText)
		slotTxt.SetText(fmt.Sprintf("%d", i+1))
		scroll.AddEntity(myecs.Manager.NewEntity().
			AddComponent(myecs.Object, slotTxt.Obj).
			AddComponent(myecs.Parent, origSlot).
			AddComponent(myecs.Drawable, slotTxt).
			AddComponent(myecs.DrawTarget, scroll.ListView.Canvas))
		lvlObj := object.New()
		lvlObj.Layer = 110 + index
		lvlObj.Offset.Y -= frame.H() * float64(i)
		lvlObj.Offset.X += frame.W()
		scroll.AddEntity(myecs.Manager.NewEntity().
			AddComponent(myecs.Object, lvlObj).
			AddComponent(myecs.Parent, origSlot).
			AddComponent(myecs.Drawable, img.NewSprite("scroll_square", data.UIKey)))
		lvlTxt := typeface.New(nil, "main", typeface.NewAlign(typeface.Center, typeface.Center), 1., 0.15, 0., 0.)
		lvlTxt.Obj.Layer = 113 + index
		lvlTxt.SetOffset(lvlObj.Offset)
		lvlTxt.Obj.Offset.Y += 6.
		lvlTxt.SetColor(data.ScrollText)
		lvlTxt.SetText(util.RomanNumeral(slot.Tier))
		scroll.AddEntity(myecs.Manager.NewEntity().
			AddComponent(myecs.Object, lvlTxt.Obj).
			AddComponent(myecs.Parent, origSlot).
			AddComponent(myecs.Drawable, lvlTxt).
			AddComponent(myecs.DrawTarget, scroll.ListView.Canvas))
		nameMObj := object.New()
		nameMObj.Layer = 110 + index
		nameMObj.Offset.Y -= frame.H() * float64(i)
		nameMObj.Offset.X += frame.W()*1.5 + nWidth*0.5
		nameMObj.Sca = pixel.V(nWidth/data.TileSize-2, 1.)
		scroll.AddEntity(myecs.Manager.NewEntity().
			AddComponent(myecs.Object, nameMObj).
			AddComponent(myecs.Parent, origSlot).
			AddComponent(myecs.Drawable, img.NewSprite("scroll_square_m", data.UIKey)))
		nameLObj := object.New()
		nameLObj.Layer = 110 + index
		nameLObj.Offset.Y -= frame.H() * float64(i)
		nameLObj.Offset.X += frame.W()*2. - data.TileSize
		scroll.AddEntity(myecs.Manager.NewEntity().
			AddComponent(myecs.Object, nameLObj).
			AddComponent(myecs.Parent, origSlot).
			AddComponent(myecs.Drawable, img.NewSprite("scroll_square_l", data.UIKey)))
		nameRObj := object.New()
		nameRObj.Layer = 110 + index
		nameRObj.Offset.Y -= frame.H() * float64(i)
		nameRObj.Offset.X += frame.W()*1.5 + nWidth - data.TileSize*0.5
		scroll.AddEntity(myecs.Manager.NewEntity().
			AddComponent(myecs.Object, nameRObj).
			AddComponent(myecs.Parent, origSlot).
			AddComponent(myecs.Drawable, img.NewSprite("scroll_square_r", data.UIKey)))
		nameTxt := typeface.New(nil, "main", typeface.NewAlign(typeface.Left, typeface.Center), 1., 0.15, 0., 0.)
		nameTxt.Obj.Layer = 113 + index
		nameTxt.SetOffset(nameLObj.Offset)
		nameTxt.Obj.Offset.Y += 6.
		nameTxt.SetColor(data.ScrollText)
		nameTxt.SetText(slot.Name)
		scroll.AddEntity(myecs.Manager.NewEntity().
			AddComponent(myecs.Object, nameTxt.Obj).
			AddComponent(myecs.Parent, origSlot).
			AddComponent(myecs.Drawable, nameTxt).
			AddComponent(myecs.DrawTarget, scroll.ListView.Canvas))
		if i == len(scroll.Tower.Slots)-1 {
			// set listView YLim
			scroll.ListView.SetYLim(math.Min(scroll.ListView.Rect.H()*-0.5+frame.H()*0.5, frame.H()*-(float64(i)+0.5)+scroll.ListView.Rect.H()*0.5), scroll.ListView.Rect.H()*-0.5+frame.H()*0.5)
		}
	}

	arrowRect := img.Batchers[data.UIKey].GetSprite("scroll_arrow_up").Frame()
	upArrowSpr := img.NewSprite("scroll_arrow_up", data.UIKey)
	dwnArrowSpr := img.NewSprite("scroll_arrow_dwn", data.UIKey)
	var upArrowTimer *timing.Timer
	var dwnArrowTimer *timing.Timer
	upArrowClick := false
	dwnArrowClick := false
	arrowUpObj := object.New()
	arrowUpObj.Offset.Y = rect.H()*0.5 - 5.
	arrowUpObj.Offset.X = rect.W()*0.5 - 13.
	arrowUpObj.Layer = 102
	arrowUpObj.SetRect(arrowRect)
	upE := myecs.Manager.NewEntity().
		AddComponent(myecs.Object, arrowUpObj).
		AddComponent(myecs.Parent, scroll.ListViewObj).
		AddComponent(myecs.Drawable, upArrowSpr).
		AddComponent(myecs.Update, data.NewHoverClickFn(data.TheInput, data.InventoryView, func(click *data.HoverClick) {
			upArrowSpr.Key = "scroll_arrow_up"
			if click.Hover {
				if click.Input.Get("click").JustPressed() {
					upArrowClick = true
					if upArrowTimer == nil {
						upArrowTimer = timing.New(0.35)
					}
					newPos := scroll.ListView.CamPos
					newPos.Y += frame.H()
					scroll.ListView.MoveTo(newPos, 0.2, false)
				} else if click.Input.Get("click").Pressed() && upArrowClick {
					if upArrowTimer.UpdateDone() {
						vel := pixel.V(0., 300.)
						scroll.ListView.SetVel(vel)
					}
				} else {
					upArrowClick = false
					upArrowTimer = nil
				}
				if upArrowClick {
					upArrowSpr.Key = "scroll_arrow_up_pressed"
				}
			} else if !click.Input.Get("click").Pressed() {
				upArrowClick = false
			}
		}))
	scroll.AddEntity(upE)
	arrowDwnObj := object.New()
	arrowDwnObj.Offset.Y = rect.H()*-0.5 + 5.
	arrowDwnObj.Offset.X = rect.W()*0.5 - 13.
	arrowDwnObj.Layer = 102
	arrowDwnObj.SetRect(arrowRect)
	dwnE := myecs.Manager.NewEntity().
		AddComponent(myecs.Object, arrowDwnObj).
		AddComponent(myecs.Parent, scroll.ListViewObj).
		AddComponent(myecs.Drawable, dwnArrowSpr).
		AddComponent(myecs.Update, data.NewHoverClickFn(data.TheInput, data.InventoryView, func(click *data.HoverClick) {
			dwnArrowSpr.Key = "scroll_arrow_dwn"
			if click.Hover {
				if click.Input.Get("click").JustPressed() {
					dwnArrowClick = true
					if dwnArrowTimer == nil {
						dwnArrowTimer = timing.New(0.35)
					}
					newPos := scroll.ListView.CamPos
					newPos.Y -= frame.H()
					scroll.ListView.MoveTo(newPos, 0.2, false)
				} else if click.Input.Get("click").Pressed() && dwnArrowClick {
					if dwnArrowTimer.UpdateDone() {
						vel := pixel.V(0., -300.)
						scroll.ListView.SetVel(vel)
					}
				} else {
					dwnArrowClick = false
					dwnArrowTimer = nil
				}
				if dwnArrowClick {
					dwnArrowSpr.Key = "scroll_arrow_dwn_pressed"
				}
			} else if !click.Input.Get("click").Pressed() {
				dwnArrowClick = false
			}
		}))
	scroll.AddEntity(dwnE)

	// scroll bar
	barClick := false
	barObj := object.New()
	barObj.SetRect(img.Batchers[data.UIKey].GetSprite("scroll_bar").Frame())
	barBot := rect.H()*-0.5 + (barObj.Rect.H()+arrowRect.H())*0.5 + 5.
	barTop := rect.H()*0.5 - (barObj.Rect.H()+arrowRect.H())*0.5 - 5.
	barObj.Offset.Y = barTop
	barObj.Offset.X = rect.W()*0.5 - 13.
	barObj.Layer = 102
	barOffset := 0.
	barE := myecs.Manager.NewEntity().
		AddComponent(myecs.Object, barObj).
		AddComponent(myecs.Parent, scroll.ListViewObj).
		AddComponent(myecs.Drawable, img.NewSprite("scroll_bar", data.UIKey)).
		AddComponent(myecs.Update, data.NewHoverClickFn(data.TheInput, data.InventoryView, func(click *data.HoverClick) {
			if click.Hover {
				if click.Input.Get("click").JustPressed() {
					barClick = true
					barOffset = barObj.PostPos.Y - click.View.Projected(click.Input.World).Y
				}
			}
			if !click.Input.Get("click").Pressed() {
				barClick = false
			}
			listCamTop, listCamBot := scroll.ListView.GetLimY()
			if barClick {
				inPos := click.View.Projected(click.Input.World)
				barObj.Offset.Y -= barObj.PostPos.Y - inPos.Y - barOffset
				if barObj.Offset.Y > barTop {
					barObj.Offset.Y = barTop
				} else if barObj.Offset.Y < barBot {
					barObj.Offset.Y = barBot
				}
				barRatio := (barTop - barObj.Offset.Y) / (barTop - barBot)
				scroll.ListView.CamPos.Y = -(barRatio * listCamTop) + (barRatio * listCamBot) + listCamTop
			}
			camRatio := (listCamTop - scroll.ListView.CamPos.Y) / (listCamTop - listCamBot)
			barObj.Offset.Y = -(camRatio * barTop) + (camRatio * barBot) + barTop
		}))
	scroll.AddEntity(barE)
}

func ScrollSystem() {
	UpdateTowerScroll(data.LeftTowerScroll)
	UpdateTowerScroll(data.MidTowerScroll)
	UpdateTowerScroll(data.RightTowerScroll)
}

func ShowTowerScroll(scroll *data.TowerScroll) {
	scroll.Scroll.Entity.AddComponent(myecs.Interpolation, object.NewInterpolation(object.InterpolateY).
		SetGween(scroll.Scroll.Object.Pos.Y, data.TowerScrollY, 1.5, ease.OutCubic))
}

func UpdateTowerScroll(scroll *data.TowerScroll) {
	// update scroll
	UpdateScroll(scroll.Scroll)
	scroll.Scroll.Freeze = !data.InventoryView.PointInside(scroll.Scroll.Object.Pos)
	// update tower view
	scroll.TowerViewObj.Hide = !scroll.Scroll.Opened
	scroll.TitleText.Obj.Hide = !scroll.Scroll.Opened
	scroll.SubclassText.Obj.Hide = !scroll.Scroll.Opened
	scroll.LevelText.Obj.Hide = !scroll.Scroll.Opened
	scroll.SlotHead.Obj.Hide = !scroll.Scroll.Opened
	scroll.LvlHead.Obj.Hide = !scroll.Scroll.Opened
	scroll.NameHead.Obj.Hide = !scroll.Scroll.Opened
	scroll.ListViewObj.Hide = !scroll.Scroll.Opened
	if scroll.Scroll.Opened {
		scroll.ListView.Mask = colornames.White
	} else {
		scroll.ListView.Mask = color.RGBA{}
	}
	// update titles
	// update list
	// update list position
	scroll.ListView.PortPos = scroll.Scroll.Object.Pos.Add(scroll.ListViewObj.Offset)
	scroll.ListView.PortPos.X -= data.TileSize * 1.5
	// update list view
	scroll.ListView.Update()
}

func DrawScrollSystem(win *pixelgl.Window) {
	img.Clear()
	// draw the scrolls and the tower view
	DrawSystem(win, 101)
	DrawSystem(win, 102)
	img.Batchers[data.UIKey].Draw(data.InventoryView.Canvas)
	img.Batchers[data.ObjectKey].Draw(data.InventoryView.Canvas)
	// draw the text
	DrawSystem(win, 103)
	DrawScroll(win, data.LeftTowerScroll, 0)
	DrawScroll(win, data.MidTowerScroll, 1)
	DrawScroll(win, data.RightTowerScroll, 2)
}

func DrawScroll(win *pixelgl.Window, scroll *data.TowerScroll, index int) {
	// draw the slot list in the middle
	scroll.ListView.Canvas.Clear(color.RGBA{})
	img.Clear()
	DrawSystem(win, 110+index)
	img.Batchers[data.UIKey].Draw(scroll.ListView.Canvas)
	// draw the text
	DrawSystem(win, 113+index)
	// draw to the canvas
	scroll.ListView.Draw(data.InventoryView.Canvas)
}

func DisposeScrolls() {
	DisposeScroll(data.LeftTowerScroll)
	DisposeScroll(data.MidTowerScroll)
	DisposeScroll(data.RightTowerScroll)
}

func DisposeScroll(scroll *data.TowerScroll) {
	for _, e := range scroll.Entities {
		myecs.Manager.DisposeEntity(e)
	}
	scroll.ListView = nil
	scroll = nil
}
