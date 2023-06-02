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
	"timsims1717/magicmissile/pkg/img"
	"timsims1717/magicmissile/pkg/object"
	"timsims1717/magicmissile/pkg/typeface"
	"timsims1717/magicmissile/pkg/viewport"
)

func CreateInventoryScrolls() {
	CreateInventoryScroll(0)
	CreateInventoryScroll(1)
	CreateInventoryScroll(2)
}

func CreateInventoryScroll(index int) {
	obj := object.New()
	var scroll *data.TowerScroll
	var posStr string
	switch index {
	case 0:
		data.LeftTowerScroll = &data.TowerScroll{}
		scroll = data.LeftTowerScroll
		obj.Pos = data.LeftScrollPos
		posStr = "Left"
	case 1:
		data.MidTowerScroll = &data.TowerScroll{}
		scroll = data.MidTowerScroll
		obj.Pos = data.MidScrollPos
		posStr = "Middle"
	case 2:
		data.RightTowerScroll = &data.TowerScroll{}
		scroll = data.RightTowerScroll
		obj.Pos = data.RightScrollPos
		posStr = "Right"
	default:
		panic(fmt.Sprintf("error: incorrect index %d when creating inventory scroll", index))
	}
	scroll.Tower = data.Towers[index]
	scroll.Object = obj
	scroll.AddEntity(myecs.Manager.NewEntity().AddComponent(myecs.Object, obj))
	leftObj := object.New()
	leftObj.Layer = 101
	scroll.AddEntity(myecs.Manager.NewEntity().
		AddComponent(myecs.Object, leftObj).
		AddComponent(myecs.Parent, obj).
		AddComponent(myecs.Drawable, data.ScrollTop))
	midObj := object.New()
	midObj.Layer = 101
	midObj.Offset.Y -= data.TowerScrollHeight * data.TileSize * 0.5
	midObj.Sca = pixel.V(1, data.ScrollScale)
	scroll.AddEntity(myecs.Manager.NewEntity().
		AddComponent(myecs.Object, midObj).
		AddComponent(myecs.Parent, obj).
		AddComponent(myecs.Drawable, data.ScrollMid))
	botObj := object.New()
	botObj.Layer = 101
	botObj.Offset.Y -= data.TowerScrollHeight * data.TileSize
	scroll.AddEntity(myecs.Manager.NewEntity().
		AddComponent(myecs.Object, botObj).
		AddComponent(myecs.Parent, obj).
		AddComponent(myecs.Drawable, data.ScrollBot))
	rect := img.Batchers[scroll.Tower.Sprite.Batch].GetSprite(scroll.Tower.Sprite.Key).Frame()
	CreateTowerView(scroll, rect)
	CreateTowerText(scroll, posStr, rect)
	CreateSpellList(scroll, index)
}

func CreateTowerView(scroll *data.TowerScroll, rect pixel.Rect) {
	scroll.TowerViewObj = object.New()
	scroll.TowerViewObj.Offset = data.TowerViewPos
	scroll.TowerViewObj.Offset.Y -= rect.H() * 0.5
	scroll.TowerViewObj.Offset.X += rect.W() * 0.5
	scroll.AddEntity(myecs.Manager.NewEntity().
		AddComponent(myecs.Object, scroll.TowerViewObj).
		AddComponent(myecs.Parent, scroll.Object))
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
	scroll.TitleText.SetOffset(data.TowerTitlePos)
	scroll.TitleText.SetColor(colornames.Black)
	scroll.TitleText.SetText(fmt.Sprintf("%s Wizard's Tower", pos))
	scroll.AddEntity(myecs.Manager.NewEntity().
		AddComponent(myecs.Object, scroll.TitleText.Obj).
		AddComponent(myecs.Parent, scroll.Object).
		AddComponent(myecs.Drawable, scroll.TitleText).
		AddComponent(myecs.DrawTarget, data.InventoryView.Canvas))

	scroll.SubclassText = typeface.New(nil, "main", typeface.NewAlign(typeface.Left, typeface.Top), 1., 0.12, 0., 0.)
	scroll.SubclassText.Obj.Layer = 103
	scroll.SubclassText.SetOffset(pixel.V(data.SubtitlePosX, scroll.TitleText.Obj.Offset.Y-scroll.TitleText.Height))
	scroll.SubclassText.SetColor(colornames.Black)
	scroll.SubclassText.SetText("Specialization: Arcanist")
	scroll.AddEntity(myecs.Manager.NewEntity().
		AddComponent(myecs.Object, scroll.SubclassText.Obj).
		AddComponent(myecs.Parent, scroll.Object).
		AddComponent(myecs.Drawable, scroll.SubclassText).
		AddComponent(myecs.DrawTarget, data.InventoryView.Canvas))

	scroll.LevelText = typeface.New(nil, "main", typeface.NewAlign(typeface.Left, typeface.Top), 1., 0.12, 0., 0.)
	scroll.LevelText.Obj.Layer = 103
	scroll.LevelText.SetOffset(pixel.V(data.SubtitlePosX, scroll.SubclassText.Obj.Offset.Y-scroll.SubclassText.Height))
	scroll.LevelText.SetColor(colornames.Black)
	scroll.LevelText.SetText("Level III")
	scroll.AddEntity(myecs.Manager.NewEntity().
		AddComponent(myecs.Object, scroll.LevelText.Obj).
		AddComponent(myecs.Parent, scroll.Object).
		AddComponent(myecs.Drawable, scroll.LevelText).
		AddComponent(myecs.DrawTarget, data.InventoryView.Canvas))

	scroll.TableHeadY = math.Min(scroll.LevelText.Obj.Offset.Y-scroll.LevelText.Height, scroll.TowerViewObj.Offset.Y-rect.H()*0.5) - 15.

	slotTxt := typeface.New(nil, "main", typeface.NewAlign(typeface.Left, typeface.Top), 1., 0.10, 0., 0.)
	slotTxt.Obj.Layer = 103
	slotTxt.SetOffset(pixel.V(data.SlotsHeadX, scroll.TableHeadY))
	slotTxt.SetColor(colornames.Black)
	slotTxt.SetText("Slot")
	scroll.AddEntity(myecs.Manager.NewEntity().
		AddComponent(myecs.Object, slotTxt.Obj).
		AddComponent(myecs.Parent, scroll.Object).
		AddComponent(myecs.Drawable, slotTxt).
		AddComponent(myecs.DrawTarget, data.InventoryView.Canvas))

	lvlTxt := typeface.New(nil, "main", typeface.NewAlign(typeface.Left, typeface.Top), 1., 0.10, 0., 0.)
	lvlTxt.Obj.Layer = 103
	lvlTxt.SetOffset(pixel.V(data.LevelHeadX, scroll.TableHeadY))
	lvlTxt.SetColor(colornames.Black)
	lvlTxt.SetText("Lvl")
	scroll.AddEntity(myecs.Manager.NewEntity().
		AddComponent(myecs.Object, lvlTxt.Obj).
		AddComponent(myecs.Parent, scroll.Object).
		AddComponent(myecs.Drawable, lvlTxt).
		AddComponent(myecs.DrawTarget, data.InventoryView.Canvas))

	nameTxt := typeface.New(nil, "main", typeface.NewAlign(typeface.Left, typeface.Top), 1., 0.10, 0., 0.)
	nameTxt.Obj.Layer = 103
	nameTxt.SetOffset(pixel.V(data.NameHeadX, scroll.TableHeadY))
	nameTxt.SetColor(colornames.Black)
	nameTxt.SetText("Spell")
	scroll.AddEntity(myecs.Manager.NewEntity().
		AddComponent(myecs.Object, nameTxt.Obj).
		AddComponent(myecs.Parent, scroll.Object).
		AddComponent(myecs.Drawable, nameTxt).
		AddComponent(myecs.DrawTarget, data.InventoryView.Canvas))
}

func CreateSpellList(scroll *data.TowerScroll, index int) {
	rect := pixel.R(0, 0, (data.TowerScrollWidth-5)*data.TileSize, math.Abs(scroll.TableHeadY+(data.TowerScrollHeight-7)*data.TileSize))
	scroll.ListViewObj = object.New()
	scroll.ListViewObj.Offset.Y = scroll.TableHeadY - 40. - rect.H()*0.5
	scroll.ListViewObj.Offset.X = rect.W()*0.5 + 12.
	scroll.AddEntity(myecs.Manager.NewEntity().
		AddComponent(myecs.Object, scroll.ListViewObj).
		AddComponent(myecs.Parent, scroll.Object))
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

	arrowUpObj := object.New()
	arrowUpObj.Offset.Y = rect.H()*0.5 - 5.
	arrowUpObj.Offset.X = rect.W()*0.5 - 13.
	arrowUpObj.Layer = 102
	scroll.AddEntity(myecs.Manager.NewEntity().
		AddComponent(myecs.Object, arrowUpObj).
		AddComponent(myecs.Parent, scroll.ListViewObj).
		AddComponent(myecs.Drawable, img.NewSprite("scroll_arrow_up", data.UIKey)))
	arrowDwnObj := object.New()
	arrowDwnObj.Offset.Y = rect.H()*-0.5 + 5.
	arrowDwnObj.Offset.X = rect.W()*0.5 - 13.
	arrowDwnObj.Layer = 102
	scroll.AddEntity(myecs.Manager.NewEntity().
		AddComponent(myecs.Object, arrowDwnObj).
		AddComponent(myecs.Parent, scroll.ListViewObj).
		AddComponent(myecs.Drawable, img.NewSprite("scroll_arrow_dwn", data.UIKey)))

	scroll.ListView = viewport.New(nil)
	scroll.ListView.SetRect(pixel.R(0., 0., rect.W()-data.TileSize*1.5-2., rect.H()+data.TileSize+6.))
	scroll.ListView.CamPos = pixel.V(data.SlotListViewX, data.SlotListViewY)
	scroll.ListView.PortPos = scroll.Object.Pos.Add(scroll.ListViewObj.Offset)
	scroll.ListView.PortPos.X -= data.TileSize * 1.5

	frame := img.Batchers[data.UIKey].GetSprite("scroll_square").Frame()
	nWidth := scroll.ListView.Rect.W() - frame.W()*2.
	var origSlot *object.Object
	for i, _ := range scroll.Tower.Slots {
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
		lvlObj := object.New()
		lvlObj.Layer = 110 + index
		lvlObj.Offset.Y -= frame.H() * float64(i)
		lvlObj.Offset.X += frame.W()
		scroll.AddEntity(myecs.Manager.NewEntity().
			AddComponent(myecs.Object, lvlObj).
			AddComponent(myecs.Parent, origSlot).
			AddComponent(myecs.Drawable, img.NewSprite("scroll_square", data.UIKey)))
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
		if i == 0 {
			// set topoflistY
		} else if i == len(scroll.Tower.Slots)-1 {
			// set botoflistY
		}
	}
}

func ScrollSystem() {
	UpdateScroll(data.LeftTowerScroll)
	UpdateScroll(data.MidTowerScroll)
	UpdateScroll(data.RightTowerScroll)
}

func UpdateScroll(scroll *data.TowerScroll) {
	// update position
	// update tower view
	// update titles
	// update list
	// update list position
	scroll.ListView.PortPos = scroll.Object.Pos.Add(scroll.ListViewObj.Offset)
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
	scroll.ListView.Canvas.Draw(data.InventoryView.Canvas, scroll.ListView.Mat)
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
