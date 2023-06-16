package inventory

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"image/color"
	"math"
	"math/rand"
	"strconv"
	"timsims1717/magicmissile/internal/data"
	"timsims1717/magicmissile/internal/myecs"
	"timsims1717/magicmissile/internal/systems"
	"timsims1717/magicmissile/pkg/img"
	"timsims1717/magicmissile/pkg/object"
	"timsims1717/magicmissile/pkg/typeface"
	"timsims1717/magicmissile/pkg/viewport"
)

func CreateSpellInventory() {
	DisposeSpellInventory()
	data.SpellInventory.Spells = make(map[string]int)
	for key, spell := range data.Missiles {
		if key == "arcaneshot" {
			data.SpellInventory.Spells[key] = -1
		} else {
			cnt := rand.Intn(5) - spell[0].Tier
			if cnt < 0 {
				cnt = 0
			}
			data.SpellInventory.Spells[key] = cnt + 1
		}
	}
	data.SpellInventory.Scroll = systems.CreateScroll(data.SpellInventoryPos, 101, pixel.V(data.SpellInventoryWidth*data.TileSize, data.ScrollHeight*data.TileSize), pixel.V(data.SpellInventoryWidth*data.TileSize, data.ScrollHeight*data.TileSize))
	data.SpellInventory.TitleText = typeface.New("main", typeface.NewAlign(typeface.Left, typeface.Top), 1., 0.15, 300., 0.)
	data.SpellInventory.TitleText.Obj.Layer = 103
	data.SpellInventory.TitleText.SetOffset(data.SpellInventory.Scroll.TLPos.Add(data.InventoryTitlePos))
	data.SpellInventory.TitleText.SetColor(data.ScrollText)
	data.SpellInventory.TitleText.SetText("Spell Inventory")
	data.SpellInventory.AddEntity(myecs.Manager.NewEntity().
		AddComponent(myecs.Object, data.SpellInventory.TitleText.Obj).
		AddComponent(myecs.Parent, data.SpellInventory.Scroll.Object).
		AddComponent(myecs.Drawable, data.SpellInventory.TitleText).
		AddComponent(myecs.DrawTarget, data.InventoryView.Canvas))

	rect := pixel.R(0, 0, (data.SlotWidth+data.SquareFrame.W()*2)*2.+data.SquareFrame.W()*0.5+2., (data.ScrollHeight-6)*data.TileSize)
	data.SpellInventory.ListViewObj = object.New()
	data.SpellInventory.ListViewObj.SetRect(rect)
	data.SpellInventory.ListViewObj.HideChildren = true
	data.SpellInventory.ListViewObj.Offset = pixel.V(data.SpellInventory.Scroll.TLPos.X+rect.W()*0.5+12., data.SpellInventory.Scroll.TLPos.Y-rect.H()*0.5-data.SpellInventory.TitleText.Height-30.)
	data.SpellInventory.AddEntity(myecs.Manager.NewEntity().
		AddComponent(myecs.Object, data.SpellInventory.ListViewObj).
		AddComponent(myecs.Parent, data.SpellInventory.Scroll.Object))

	mmObj := object.New()
	mmObj.Sca = pixel.V(rect.W()/data.TileSize, rect.H()/data.TileSize)
	mmObj.Layer = 102
	data.SpellInventory.AddEntity(myecs.Manager.NewEntity().
		AddComponent(myecs.Object, mmObj).
		AddComponent(myecs.Parent, data.SpellInventory.ListViewObj).
		AddComponent(myecs.Drawable, img.NewSprite("list_bg_m", data.UIKey)))
	tmObj := object.New()
	tmObj.Offset.Y = rect.H()*0.5 + 3.
	tmObj.Sca = pixel.V(rect.W()/data.TileSize, 1.)
	tmObj.Layer = 102
	data.SpellInventory.AddEntity(myecs.Manager.NewEntity().
		AddComponent(myecs.Object, tmObj).
		AddComponent(myecs.Parent, data.SpellInventory.ListViewObj).
		AddComponent(myecs.Drawable, img.NewSprite("list_bg_sv", data.UIKey)))
	bmObj := object.New()
	bmObj.Offset.Y = rect.H()*-0.5 - 3.
	bmObj.Sca = pixel.V(rect.W()/data.TileSize, 1.)
	bmObj.Layer = 102
	bmObj.Flop = true
	data.SpellInventory.AddEntity(myecs.Manager.NewEntity().
		AddComponent(myecs.Object, bmObj).
		AddComponent(myecs.Parent, data.SpellInventory.ListViewObj).
		AddComponent(myecs.Drawable, img.NewSprite("list_bg_sv", data.UIKey)))
	mlObj := object.New()
	mlObj.Offset.X = rect.W()*-0.5 - 3.
	mlObj.Sca = pixel.V(1., rect.H()/data.TileSize)
	mlObj.Layer = 102
	data.SpellInventory.AddEntity(myecs.Manager.NewEntity().
		AddComponent(myecs.Object, mlObj).
		AddComponent(myecs.Parent, data.SpellInventory.ListViewObj).
		AddComponent(myecs.Drawable, img.NewSprite("list_bg_sh", data.UIKey)))
	mrObj := object.New()
	mrObj.Offset.X = rect.W()*0.5 + 3.
	mrObj.Sca = pixel.V(1., rect.H()/data.TileSize)
	mrObj.Layer = 102
	mrObj.Flip = true
	data.SpellInventory.AddEntity(myecs.Manager.NewEntity().
		AddComponent(myecs.Object, mrObj).
		AddComponent(myecs.Parent, data.SpellInventory.ListViewObj).
		AddComponent(myecs.Drawable, img.NewSprite("list_bg_sh", data.UIKey)))
	tlObj := object.New()
	tlObj.Offset.Y = rect.H()*0.5 + 3.
	tlObj.Offset.X = rect.W()*-0.5 - 3.
	tlObj.Layer = 102
	data.SpellInventory.AddEntity(myecs.Manager.NewEntity().
		AddComponent(myecs.Object, tlObj).
		AddComponent(myecs.Parent, data.SpellInventory.ListViewObj).
		AddComponent(myecs.Drawable, img.NewSprite("list_bg_c", data.UIKey)))
	trObj := object.New()
	trObj.Offset.Y = rect.H()*0.5 + 3.
	trObj.Offset.X = rect.W()*0.5 + 3.
	trObj.Layer = 102
	trObj.Flip = true
	data.SpellInventory.AddEntity(myecs.Manager.NewEntity().
		AddComponent(myecs.Object, trObj).
		AddComponent(myecs.Parent, data.SpellInventory.ListViewObj).
		AddComponent(myecs.Drawable, img.NewSprite("list_bg_c", data.UIKey)))
	blObj := object.New()
	blObj.Offset.Y = rect.H()*-0.5 - 3.
	blObj.Offset.X = rect.W()*-0.5 - 3.
	blObj.Layer = 102
	blObj.Flop = true
	data.SpellInventory.AddEntity(myecs.Manager.NewEntity().
		AddComponent(myecs.Object, blObj).
		AddComponent(myecs.Parent, data.SpellInventory.ListViewObj).
		AddComponent(myecs.Drawable, img.NewSprite("list_bg_c", data.UIKey)))
	brObj := object.New()
	brObj.Offset.Y = rect.H()*-0.5 - 3.
	brObj.Offset.X = rect.W()*0.5 + 3.
	brObj.Layer = 102
	brObj.Flip = true
	brObj.Flop = true
	data.SpellInventory.AddEntity(myecs.Manager.NewEntity().
		AddComponent(myecs.Object, brObj).
		AddComponent(myecs.Parent, data.SpellInventory.ListViewObj).
		AddComponent(myecs.Drawable, img.NewSprite("list_bg_c", data.UIKey)))

	data.SpellInventory.ListView = viewport.New(nil)
	data.SpellInventory.ListView.SetRect(pixel.R(0., 0., rect.W()-data.TileSize*1.5-2., rect.H()+data.TileSize+6.))
	data.SpellInventory.ListView.ParentView = data.InventoryView
	data.SpellInventory.ListView.CamPos = pixel.V(-(data.SquareFrame.W()*2.+data.SlotWidth*0.5)+data.SpellInventory.ListView.Rect.W()*0.5, data.SlotListViewY)
	data.SpellInventory.ListView.PortPos = data.SpellInventory.Scroll.Object.Pos.Add(data.SpellInventory.ListViewObj.Offset)
	data.SpellInventory.ListView.PortPos.X -= data.TileSize * 1.5

	systems.CreateScrollBar(data.SpellInventory.AddEntity, data.SpellInventory.ListViewObj, data.InventoryView, data.SpellInventory.ListView)
	UpdateSpellInventory(true)
}

func UpdateSpellInventory(doUpdate bool) {
	update := false
spellLoop:
	for spell, count := range data.SpellInventory.Spells {
		for _, slot := range data.SpellInventory.Slots {
			if slot.Slot.Spell == spell {
				if slot.SlotNum != count {
					slot.SlotTxt.SetText(strconv.Itoa(count))
					slot.SlotNum = count
					if count == 0 {
						slot.NameMObj.Hidden = true
						update = true
					}
				}
				continue spellLoop
			}
		}
		if count != 0 {
			// create a new spell slot in the inventory
			base := data.Missiles[spell][0]
			data.SpellInventory.Slots = append(data.SpellInventory.Slots, CreateInventorySpellSlot(&data.SpellSlot{
				Tier:  base.Tier,
				Spell: base.Key,
				Name:  base.Name,
			}, count, data.SpellInventory.AddEntity, data.SpellInventory.ListView))
			update = true
		}
	}
	if update || doUpdate {
		// order the list
		for i := 1; i < len(data.SpellInventory.Slots); i++ {
			el := data.SpellInventory.Slots[i]
			j := i - 1
			jel := data.SpellInventory.Slots[j]
			for j >= 0 && jel.Slot.Tier > el.Slot.Tier {
				data.SpellInventory.Slots[j+1] = jel
				j--
				if j > -1 {
					jel = data.SpellInventory.Slots[j]
				}
			}
			data.SpellInventory.Slots[j+1] = el
		}
		rowCnt := 0
		currTierCnt := -1
		tier := 1
		bottom := 0
		for _, slot := range data.SpellInventory.Slots {
			if slot.NameMObj.Hidden {
				continue
			}
			if tier == slot.Slot.Tier {
				currTierCnt++
				if currTierCnt%2 == 0 {
					rowCnt++
				}
			} else {
				currTierCnt = 0
				tier = slot.Slot.Tier
				rowCnt++
			}
			slot.NameMObj.Pos.Y = -float64(rowCnt) * data.SquareFrame.H()
			if currTierCnt%2 != 0 {
				slot.NameMObj.Pos.X = data.SlotWidth + (data.SquareFrame.W() * 2)
			} else {
				//slot.NameMObj.Pos.X = data.SlotWidth * -0.5
			}
			bottom = rowCnt
		}
		// set listView YLim
		data.SpellInventory.ListView.SetYLim(math.Min(data.SpellInventory.ListView.Rect.H()*-0.5+data.SquareFrame.H()*0.5, data.SquareFrame.H()*-(float64(bottom)+0.5)+data.SpellInventory.ListView.Rect.H()*0.5), data.SpellInventory.ListView.Rect.H()*-0.5+data.SquareFrame.H()*0.5)
	}
	// update list position
	data.SpellInventory.ListView.PortPos = data.SpellInventory.Scroll.Object.Pos.Add(data.SpellInventory.ListViewObj.Offset)
	data.SpellInventory.ListView.PortPos.X -= data.TileSize * 1.5
	// update list view
	data.SpellInventory.ListView.Update()
}

func DrawSpellInventory(win *pixelgl.Window) {
	data.SpellInventory.ListView.Canvas.Clear(color.RGBA{})
	img.Clear()
	systems.DrawSystem(win, 116)
	img.Batchers[data.UIKey].Draw(data.SpellInventory.ListView.Canvas)
	//img.Batchers[data.UIKey].Draw(data.InventoryView.Canvas)
	// draw the text
	systems.DrawSystem(win, 117)
	// draw to the canvas
	data.SpellInventory.ListView.Draw(data.InventoryView.Canvas)
}

func DisposeSpellInventory() {
	for _, e := range data.SpellInventory.Entities {
		myecs.Manager.DisposeEntity(e)
	}
	if data.SpellInventory.Scroll != nil {
		systems.DisposeScroll(data.SpellInventory.Scroll)
	}
	data.SpellInventory.TitleText = nil
	data.SpellInventory.ListViewObj = nil
	data.SpellInventory.ListView = nil
}
