package inventory

import (
	"fmt"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
	"image/color"
	"math"
	"timsims1717/magicmissile/internal/data"
	"timsims1717/magicmissile/internal/myecs"
	"timsims1717/magicmissile/internal/systems"
	"timsims1717/magicmissile/pkg/gween64/ease"
	"timsims1717/magicmissile/pkg/img"
	"timsims1717/magicmissile/pkg/object"
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
		scroll.Scroll = systems.CreateScroll(data.LeftScrollStart, 101, pixel.V(data.TowerScrollWidth*data.TileSize, data.TileSize), pixel.V(data.TowerScrollWidth*data.TileSize, data.ScrollHeight*data.TileSize))
		posStr = "Left"
	case 1:
		data.MidTowerScroll = &data.TowerScroll{}
		scroll = data.MidTowerScroll
		scroll.Scroll = systems.CreateScroll(data.MidScrollStart, 101, pixel.V(data.TowerScrollWidth*data.TileSize, data.TileSize), pixel.V(data.TowerScrollWidth*data.TileSize, data.ScrollHeight*data.TileSize))
		posStr = "Middle"
	case 2:
		data.RightTowerScroll = &data.TowerScroll{}
		scroll = data.RightTowerScroll
		scroll.Scroll = systems.CreateScroll(data.RightScrollStart, 101, pixel.V(data.TowerScrollWidth*data.TileSize, data.TileSize), pixel.V(data.TowerScrollWidth*data.TileSize, data.ScrollHeight*data.TileSize))
		posStr = "Right"
	default:
		panic(fmt.Sprintf("error: incorrect index %d when creating inventory scroll", index))
	}
	scroll.PosX = scroll.Scroll.Object.Pos.X
	scroll.Scroll.Closed = true
	scroll.Tower = data.Towers[index]

	tvFrame := img.Batchers[scroll.Tower.Sprite.Batch].GetSprite(scroll.Tower.Sprite.Key).Frame()
	CreateTowerView(scroll, tvFrame)
	CreateTowerText(scroll, posStr, tvFrame)
	CreateSaveButton(scroll, tvFrame, index)
	CreateEditButton(scroll, tvFrame, index)
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

func CreateTowerText(scroll *data.TowerScroll, pos string, tvFrame pixel.Rect) {
	scroll.TitleText = typeface.New("main", typeface.NewAlign(typeface.Left, typeface.Top), 1., 0.15, 300., 0.)
	scroll.TitleText.Obj.Layer = 103
	scroll.TitleText.SetOffset(scroll.Scroll.TLPos.Add(data.TowerTitlePos))
	scroll.TitleText.Obj.Offset.X += tvFrame.W()
	scroll.TitleText.SetColor(data.ScrollText)
	scroll.TitleText.SetText(fmt.Sprintf("%s Wizard's Tower", pos))
	scroll.AddEntity(myecs.Manager.NewEntity().
		AddComponent(myecs.Object, scroll.TitleText.Obj).
		AddComponent(myecs.Parent, scroll.Scroll.Object).
		AddComponent(myecs.Drawable, scroll.TitleText).
		AddComponent(myecs.DrawTarget, data.InventoryView.Canvas))

	scroll.SubclassText = typeface.New("main", typeface.NewAlign(typeface.Left, typeface.Top), 1., 0.12, 0., 0.)
	scroll.SubclassText.Obj.Layer = 103
	scroll.SubclassText.SetOffset(scroll.Scroll.TLPos.Add(data.SubtitlePos))
	scroll.SubclassText.Obj.Offset.X += tvFrame.W()
	scroll.SubclassText.Obj.Offset.Y -= scroll.TitleText.Height
	scroll.SubclassText.SetColor(data.ScrollText)
	scroll.SubclassText.SetText("Specialization: Arcanist")
	scroll.AddEntity(myecs.Manager.NewEntity().
		AddComponent(myecs.Object, scroll.SubclassText.Obj).
		AddComponent(myecs.Parent, scroll.Scroll.Object).
		AddComponent(myecs.Drawable, scroll.SubclassText).
		AddComponent(myecs.DrawTarget, data.InventoryView.Canvas))

	scroll.LevelText = typeface.New("main", typeface.NewAlign(typeface.Left, typeface.Top), 1., 0.12, 0., 0.)
	scroll.LevelText.Obj.Layer = 103
	scroll.LevelText.SetOffset(scroll.Scroll.TLPos.Add(data.SubtitlePos))
	scroll.LevelText.Obj.Offset.X += tvFrame.W()
	scroll.LevelText.Obj.Offset.Y = scroll.SubclassText.Obj.Offset.Y - scroll.SubclassText.Height
	scroll.LevelText.SetColor(data.ScrollText)
	scroll.LevelText.SetText("Level III")
	scroll.AddEntity(myecs.Manager.NewEntity().
		AddComponent(myecs.Object, scroll.LevelText.Obj).
		AddComponent(myecs.Parent, scroll.Scroll.Object).
		AddComponent(myecs.Drawable, scroll.LevelText).
		AddComponent(myecs.DrawTarget, data.InventoryView.Canvas))

	scroll.TableHeadY = math.Min(scroll.LevelText.Obj.Offset.Y-scroll.LevelText.Height, scroll.TowerViewObj.Offset.Y-tvFrame.H()*0.5) - 15.

	scroll.SlotHead = typeface.New("main", typeface.NewAlign(typeface.Left, typeface.Top), 1., 0.10, 0., 0.)
	scroll.SlotHead.Obj.Layer = 103
	scroll.SlotHead.SetOffset(pixel.V(scroll.Scroll.TLPos.X+data.SlotsHeadX, scroll.TableHeadY))
	scroll.SlotHead.SetColor(data.ScrollText)
	scroll.SlotHead.SetText("Slot")
	scroll.AddEntity(myecs.Manager.NewEntity().
		AddComponent(myecs.Object, scroll.SlotHead.Obj).
		AddComponent(myecs.Parent, scroll.Scroll.Object).
		AddComponent(myecs.Drawable, scroll.SlotHead).
		AddComponent(myecs.DrawTarget, data.InventoryView.Canvas))

	scroll.LvlHead = typeface.New("main", typeface.NewAlign(typeface.Left, typeface.Top), 1., 0.10, 0., 0.)
	scroll.LvlHead.Obj.Layer = 103
	scroll.LvlHead.SetOffset(pixel.V(scroll.Scroll.TLPos.X+data.LevelHeadX, scroll.TableHeadY))
	scroll.LvlHead.SetColor(data.ScrollText)
	scroll.LvlHead.SetText("Lvl")
	scroll.AddEntity(myecs.Manager.NewEntity().
		AddComponent(myecs.Object, scroll.LvlHead.Obj).
		AddComponent(myecs.Parent, scroll.Scroll.Object).
		AddComponent(myecs.Drawable, scroll.LvlHead).
		AddComponent(myecs.DrawTarget, data.InventoryView.Canvas))

	scroll.NameHead = typeface.New("main", typeface.NewAlign(typeface.Left, typeface.Top), 1., 0.10, 0., 0.)
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

func CreateEditButton(scroll *data.TowerScroll, tvFrame pixel.Rect, index int) {
	editTxt := typeface.New("main", typeface.NewAlign(typeface.Center, typeface.Center), 1., 0.15, 0., 0.)
	editTxt.Obj.Layer = 103
	editTxt.Obj.Offset.Y += 6.
	editTxt.SetColor(data.ScrollText)
	editTxt.SetText("Edit")

	scroll.EditButton = object.New()
	scroll.EditButton.Layer = 102
	scroll.EditButton.HideChildren = true
	scroll.EditButton.Offset = scroll.Scroll.TLPos.Add(data.EditButtonPos)
	scroll.EditButton.Offset.X += tvFrame.W() + editTxt.Width*0.5 + data.TileSize
	scroll.EditButton.Offset.Y = scroll.LevelText.Obj.Offset.Y - scroll.LevelText.Height - data.TileSize*2.
	scroll.EditButton.Sca = pixel.V(math.Ceil(editTxt.Width/data.TileSize), 1.)
	scroll.EditButton.SetRect(pixel.R(0., 0., editTxt.Width+data.TileSize*2., data.TileSize*3.))
	editMSpr := img.NewSprite("scroll_square_m", data.UIKey)
	edit := myecs.Manager.NewEntity().
		AddComponent(myecs.Object, scroll.EditButton).
		AddComponent(myecs.Parent, scroll.Scroll.Object).
		AddComponent(myecs.Drawable, editMSpr)
	editLObj := object.New()
	editLObj.Layer = 102
	editLObj.Offset.X -= 0.5 * (data.TileSize + editTxt.Width)
	editLSpr := img.NewSprite("scroll_square_l", data.UIKey)
	scroll.AddEntity(myecs.Manager.NewEntity().
		AddComponent(myecs.Object, editLObj).
		AddComponent(myecs.Parent, scroll.EditButton).
		AddComponent(myecs.Drawable, editLSpr))
	editRObj := object.New()
	editRObj.Layer = 102
	editRObj.Offset.X += 0.5 * (data.TileSize + editTxt.Width)
	editRSpr := img.NewSprite("scroll_square_r", data.UIKey)
	scroll.AddEntity(myecs.Manager.NewEntity().
		AddComponent(myecs.Object, editRObj).
		AddComponent(myecs.Parent, scroll.EditButton).
		AddComponent(myecs.Drawable, editRSpr))

	editClicked := false
	edit.AddComponent(myecs.Update, data.NewHoverClickFn(data.TheInput, data.InventoryView, func(click *data.HoverClick) {
		editClickInner := false
		if click.Hover && data.InventoryState == 3 && !data.InventoryTrans {
			if click.Input.Get("click").JustPressed() {
				editClicked = true
			} else if click.Input.Get("click").JustReleased() && editClicked {
				// trigger
				if index != 0 {
					HideTowerScroll(data.LeftTowerScroll)
				}
				if index != 1 {
					HideTowerScroll(data.MidTowerScroll)
				}
				if index != 2 {
					HideTowerScroll(data.RightTowerScroll)
				}
				MoveTowerScrollX(scroll, data.LeftScrollX)
				data.InventoryState = index
				data.InventoryTrans = true
				editClicked = false
			}
			if editClicked {
				editClickInner = true
			}
		}
		if !click.Input.Get("click").Pressed() {
			editClicked = false
		}
		if editClickInner {
			editMSpr.Key = "scroll_square_m_pressed"
			editLSpr.Key = "scroll_square_l_pressed"
			editRSpr.Key = "scroll_square_r_pressed"
		} else {
			editMSpr.Key = "scroll_square_m"
			editLSpr.Key = "scroll_square_l"
			editRSpr.Key = "scroll_square_r"
		}
	}))
	scroll.AddEntity(edit)

	scroll.AddEntity(myecs.Manager.NewEntity().
		AddComponent(myecs.Object, editTxt.Obj).
		AddComponent(myecs.Parent, scroll.EditButton).
		AddComponent(myecs.Drawable, editTxt).
		AddComponent(myecs.DrawTarget, data.InventoryView.Canvas))
}

func CreateSaveButton(scroll *data.TowerScroll, tvFrame pixel.Rect, index int) {
	saveTxt := typeface.New("main", typeface.NewAlign(typeface.Center, typeface.Center), 1., 0.15, 0., 0.)
	saveTxt.Obj.Layer = 103
	saveTxt.Obj.Offset.Y += 6.
	saveTxt.SetColor(data.ScrollText)
	saveTxt.SetText("Save")

	scroll.SaveButton = object.New()
	scroll.SaveButton.Layer = 102
	scroll.SaveButton.HideChildren = true
	scroll.SaveButton.Offset = scroll.Scroll.TLPos.Add(data.EditButtonPos)
	scroll.SaveButton.Offset.X += tvFrame.W() + saveTxt.Width*0.5 + data.TileSize
	scroll.SaveButton.Offset.Y = scroll.LevelText.Obj.Offset.Y - scroll.LevelText.Height - data.TileSize*2.
	scroll.SaveButton.Sca = pixel.V(math.Ceil(saveTxt.Width/data.TileSize), 1.)
	scroll.SaveButton.SetRect(pixel.R(0., 0., saveTxt.Width+data.TileSize*2., data.TileSize*3.))
	saveMSpr := img.NewSprite("scroll_square_m", data.UIKey)
	save := myecs.Manager.NewEntity().
		AddComponent(myecs.Object, scroll.SaveButton).
		AddComponent(myecs.Parent, scroll.Scroll.Object).
		AddComponent(myecs.Drawable, saveMSpr)
	saveLObj := object.New()
	saveLObj.Layer = 102
	saveLObj.Offset.X -= 0.5 * (data.TileSize + saveTxt.Width)
	saveLSpr := img.NewSprite("scroll_square_l", data.UIKey)
	scroll.AddEntity(myecs.Manager.NewEntity().
		AddComponent(myecs.Object, saveLObj).
		AddComponent(myecs.Parent, scroll.SaveButton).
		AddComponent(myecs.Drawable, saveLSpr))
	saveRObj := object.New()
	saveRObj.Layer = 102
	saveRObj.Offset.X += 0.5 * (data.TileSize + saveTxt.Width)
	saveRSpr := img.NewSprite("scroll_square_r", data.UIKey)
	scroll.AddEntity(myecs.Manager.NewEntity().
		AddComponent(myecs.Object, saveRObj).
		AddComponent(myecs.Parent, scroll.SaveButton).
		AddComponent(myecs.Drawable, saveRSpr))

	saveClicked := false
	save.AddComponent(myecs.Update, data.NewHoverClickFn(data.TheInput, data.InventoryView, func(click *data.HoverClick) {
		saveClickInner := false
		if click.Hover && data.InventoryState == index && !data.InventoryTrans {
			if click.Input.Get("click").JustPressed() {
				saveClicked = true
			} else if click.Input.Get("click").JustReleased() && saveClicked {
				// trigger
				if index != 0 {
					ShowTowerScroll(data.LeftTowerScroll)
				}
				if index != 1 {
					ShowTowerScroll(data.MidTowerScroll)
				}
				if index != 2 {
					ShowTowerScroll(data.RightTowerScroll)
				}
				MoveTowerScrollX(scroll, scroll.PosX)
				data.InventoryState = 3
				data.InventoryTrans = true
				saveClicked = false
			}
			if saveClicked {
				saveClickInner = true
			}
		}
		if !click.Input.Get("click").Pressed() {
			saveClicked = false
		}
		if saveClickInner {
			saveMSpr.Key = "scroll_square_m_pressed"
			saveLSpr.Key = "scroll_square_l_pressed"
			saveRSpr.Key = "scroll_square_r_pressed"
		} else {
			saveMSpr.Key = "scroll_square_m"
			saveLSpr.Key = "scroll_square_l"
			saveRSpr.Key = "scroll_square_r"
		}
	}))
	scroll.AddEntity(save)

	scroll.AddEntity(myecs.Manager.NewEntity().
		AddComponent(myecs.Object, saveTxt.Obj).
		AddComponent(myecs.Parent, scroll.SaveButton).
		AddComponent(myecs.Drawable, saveTxt).
		AddComponent(myecs.DrawTarget, data.InventoryView.Canvas))
}

func CreateSpellList(scroll *data.TowerScroll, index int) {
	rect := pixel.R(0, 0, (data.TowerScrollWidth-3)*data.TileSize, (data.ScrollHeight-10)*data.TileSize-math.Abs(scroll.TableHeadY))
	scroll.ListViewObj = object.New()
	scroll.ListViewObj.SetRect(rect)
	scroll.ListViewObj.HideChildren = true
	scroll.ListViewObj.Offset = pixel.V(scroll.Scroll.TLPos.X+rect.W()*0.5+12., scroll.TableHeadY-40.-rect.H()*0.5)
	scroll.AddEntity(myecs.Manager.NewEntity().
		AddComponent(myecs.Object, scroll.ListViewObj).
		AddComponent(myecs.Parent, scroll.Scroll.Object).
		AddComponent(myecs.Update, data.NewHoverClickFn(data.TheInput, data.InventoryView, func(hvc *data.HoverClick) {
			if hvc.Hover {
				if hvc.Input.Get("scrollUp").Pressed() && !data.MovingSpellSlot.Moving {
					newPos := pixel.V(0., data.SquareFrame.H())
					scroll.ListView.AddMove(newPos, 0.05)
				} else if hvc.Input.Get("scrollDwn").Pressed() && !data.MovingSpellSlot.Moving {
					newPos := pixel.V(0., -data.SquareFrame.H())
					scroll.ListView.AddMove(newPos, 0.05)
				} else if data.MovingSpellSlot.Slot != nil && !data.MovingSpellSlot.Moving {
					if inView, edge := scroll.ListView.WorldInside(data.TheInput.World); inView {
						if edge.Y < 0 && edge.Y > -25. {
							vel := pixel.V(0., 300.)
							scroll.ListView.SetVel(vel)
						} else if edge.Y > 0 && edge.Y < 25. {
							vel := pixel.V(0., -300.)
							scroll.ListView.SetVel(vel)
						}
					}
				}
			}
		})))
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
	scroll.ListView.ParentView = data.InventoryView
	scroll.ListView.CamPos = pixel.V(data.SlotListViewX, data.SlotListViewY)
	scroll.ListView.PortPos = scroll.Scroll.Object.Pos.Add(scroll.ListViewObj.Offset)
	scroll.ListView.PortPos.X -= data.TileSize * 1.5

	data.SlotWidth = scroll.ListView.Rect.W() - data.SquareFrame.W()*2.
	for i, slot := range scroll.Tower.Slots {
		towerSlot := CreateTowerSpellSlot(slot, index, i, scroll.AddEntity, scroll.ListView)

		scroll.InvSlots = append(scroll.InvSlots, towerSlot)
		if i == len(scroll.Tower.Slots)-1 {
			// set listView YLim
			scroll.ListView.SetYLim(math.Min(scroll.ListView.Rect.H()*-0.5+data.SquareFrame.H()*0.5, data.SquareFrame.H()*-(float64(i)+0.5)+scroll.ListView.Rect.H()*0.5), scroll.ListView.Rect.H()*-0.5+data.SquareFrame.H()*0.5)
		}
	}

	systems.CreateScrollBar(scroll.AddEntity, scroll.ListViewObj, data.InventoryView, scroll.ListView)
}

func MoveTowerScrollX(scroll *data.TowerScroll, posX float64) {
	scroll.Scroll.Inters = append(scroll.Scroll.Inters, object.NewInterpolation(object.InterpolateX).
		SetGween(scroll.Scroll.Object.Pos.X, posX, 1.5, ease.OutQuint))
	scroll.Scroll.Entity.AddComponent(myecs.Interpolation, scroll.Scroll.Inters)
}

func ShowTowerScroll(scroll *data.TowerScroll) {
	scroll.Scroll.Inters = []*object.Interpolation{
		object.NewInterpolation(object.InterpolateY).
			SetGween(scroll.Scroll.Object.Pos.Y, data.TowerScrollY, 1.5, ease.OutCubic),
	}
	scroll.Scroll.Entity.AddComponent(myecs.Interpolation, scroll.Scroll.Inters)
	scroll.Scroll.Object.Pos.X = scroll.PosX
}

func HideTowerScroll(scroll *data.TowerScroll) {
	scroll.Scroll.Inters = []*object.Interpolation{
		object.NewInterpolation(object.InterpolateY).
			SetGween(scroll.Scroll.Object.Pos.Y, data.TowerScrollYStart, 0.5, ease.Linear),
		object.NewInterpolation(object.InterpolateCustom).
			SetValue(&scroll.Scroll.CurrDim.Y).
			SetGween(scroll.Scroll.CurrDim.Y, 1., 0.75, ease.Linear).
			SetOnComplete(func() {
				scroll.Scroll.Closed = true
			}),
	}
	scroll.Scroll.Entity.AddComponent(myecs.Interpolation, scroll.Scroll.Inters)
}

func ScrollSystem() {
	UpdateTowerScroll(data.LeftTowerScroll, 0)
	UpdateTowerScroll(data.MidTowerScroll, 1)
	UpdateTowerScroll(data.RightTowerScroll, 2)
}

func UpdateTowerScroll(scroll *data.TowerScroll, index int) {
	// update scroll
	systems.UpdateScroll(scroll.Scroll)
	if scroll.Scroll.Closed && data.InventoryView.PointInside(scroll.Scroll.Object.Pos) {
		scroll.Scroll.Inters = append(scroll.Scroll.Inters, object.NewInterpolation(object.InterpolateCustom).
			SetValue(&scroll.Scroll.CurrDim.Y).
			SetGween(scroll.Scroll.CurrDim.Y, scroll.Scroll.FullDim.Y, 0.75, ease.Linear))
		scroll.Scroll.Entity.AddComponent(myecs.Interpolation, scroll.Scroll.Inters)
		scroll.Scroll.Closed = false
	}
	// update tower view
	scroll.TowerViewObj.Hidden = !scroll.Scroll.Opened
	scroll.TitleText.Obj.Hidden = !scroll.Scroll.Opened
	scroll.SubclassText.Obj.Hidden = !scroll.Scroll.Opened
	scroll.LevelText.Obj.Hidden = !scroll.Scroll.Opened
	scroll.EditButton.Hidden = !scroll.Scroll.Opened || data.InventoryState == index
	scroll.SaveButton.Hidden = !scroll.Scroll.Opened || data.InventoryState != index
	scroll.SlotHead.Obj.Hidden = !scroll.Scroll.Opened
	scroll.LvlHead.Obj.Hidden = !scroll.Scroll.Opened
	scroll.NameHead.Obj.Hidden = !scroll.Scroll.Opened
	scroll.ListViewObj.Hidden = !scroll.Scroll.Opened
	if scroll.Scroll.Opened {
		scroll.ListView.Mask = colornames.White
	} else {
		scroll.ListView.Mask = color.RGBA{}
	}
	// update titles

}

func UpdateListViews() {
	UpdateListView(data.LeftTowerScroll, 0)
	UpdateListView(data.MidTowerScroll, 1)
	UpdateListView(data.RightTowerScroll, 2)
}

func UpdateListView(scroll *data.TowerScroll, index int) {
	// update list
	inPos := scroll.ListView.ProjectWorld(data.TheInput.World)
	for i, slot := range scroll.InvSlots {
		tier := slot.Slot.Tier
		roman := util.RomanNumeral(tier)
		if slot.TierTxt.Raw != roman {
			slot.TierTxt.SetText(roman)
			if roman == "" {
				slot.SlotSpr.Key = "scroll_square_pressed"
			} else {
				slot.SlotSpr.Key = "scroll_square"
			}
		}
		name := slot.Slot.Name
		if slot.NameTxt.Raw != name {
			slot.NameTxt.SetText(name)
			if name == "" {
				slot.NameMSpr.Key = "scroll_square_m_pressed"
				slot.NameLSpr.Key = "scroll_square_l_pressed"
				slot.NameRSpr.Key = "scroll_square_r_pressed"
			} else {
				slot.NameMSpr.Key = "scroll_square_m"
				slot.NameLSpr.Key = "scroll_square_l"
				slot.NameRSpr.Key = "scroll_square_r"
			}
		}
		// if the slot is being hovered over and data.MovingSpellSlot.Slot != nil
		//  change the mask of the slot (red for too low of a tier, blue for the current selection)
		maskCode := data.MaskWhite
		if !data.InventoryTrans && !data.MovingSpellSlot.Moving &&
			data.MovingSpellSlot.Slot != nil && data.MovingSpellSlot.TierMoveIndex == -1 {
			if data.MovingSpellSlot.Slot.Tier > slot.Slot.Tier {
				maskCode = data.MaskRed
			} else if scroll.ListView.PointInside(inPos) && slot.NameMObj.PointInside(inPos) {
				maskCode = data.MaskYellow
			}
		}
		if data.MaskCode(maskCode) != slot.MaskCode {
			switch maskCode {
			case data.MaskWhite:
				slot.SlotSpr.Color = util.White
				slot.TierSpr.Color = util.White
				slot.NameLSpr.Color = util.White
				slot.NameMSpr.Color = util.White
				slot.NameRSpr.Color = util.White
			case data.MaskRed:
				slot.TierSpr.Color = data.Red
				slot.NameLSpr.Color = data.Gray
				slot.NameMSpr.Color = data.Gray
				slot.NameRSpr.Color = data.Gray
			case data.MaskYellow:
				slot.SlotSpr.Color = data.Green
				slot.TierSpr.Color = data.Green
				slot.NameLSpr.Color = data.Green
				slot.NameMSpr.Color = data.Green
				slot.NameRSpr.Color = data.Green
			}
			slot.MaskCode = data.MaskCode(maskCode)
		}
		if !data.InventoryTrans && !data.MovingSpellSlot.Moving &&
			data.MovingSpellSlot.Slot != nil && data.MovingSpellSlot.TierMoveIndex == index {
			inY := scroll.ListView.ProjectWorld(data.TheInput.World).Y
			if inY < slot.NameMObj.Pos.Y && i != 0 {
				prevSlot := scroll.InvSlots[i-1]
				if prevSlot.Slot.Tier == 0 {
					slot.Slot.Tier, prevSlot.Slot.Tier = prevSlot.Slot.Tier, slot.Slot.Tier
					slot.Slot.Name, prevSlot.Slot.Name = prevSlot.Slot.Name, slot.Slot.Name
					slot.Slot.Spell, prevSlot.Slot.Spell = prevSlot.Slot.Spell, slot.Slot.Spell
				}
			} else if inY > slot.NameMObj.Pos.Y && i != len(scroll.InvSlots)-1 {
				nextSlot := scroll.InvSlots[i+1]
				if nextSlot.Slot.Tier == 0 {
					slot.Slot.Tier, nextSlot.Slot.Tier = nextSlot.Slot.Tier, slot.Slot.Tier
					slot.Slot.Name, nextSlot.Slot.Name = nextSlot.Slot.Name, slot.Slot.Name
					slot.Slot.Spell, nextSlot.Slot.Spell = nextSlot.Slot.Spell, slot.Slot.Spell
				}
			}
		}
	}
	// update list position
	scroll.ListView.PortPos = scroll.Scroll.Object.Pos.Add(scroll.ListViewObj.Offset)
	scroll.ListView.PortPos.X -= data.TileSize * 1.5
	// update list view
	scroll.ListView.Update()
}

func DrawTowerScrollSystem(win *pixelgl.Window) {
	img.Clear()
	// draw the scrolls and the tower view
	systems.DrawSystem(win, 101)
	systems.DrawSystem(win, 102)
	img.Batchers[data.UIKey].Draw(data.InventoryView.Canvas)
	img.Batchers[data.ObjectKey].Draw(data.InventoryView.Canvas)
	// draw the text
	systems.DrawSystem(win, 103)
	DrawListView(win, data.LeftTowerScroll, 0)
	DrawListView(win, data.MidTowerScroll, 1)
	DrawListView(win, data.RightTowerScroll, 2)
}

func DrawListView(win *pixelgl.Window, scroll *data.TowerScroll, index int) {
	// draw the slot list in the middle
	scroll.ListView.Canvas.Clear(color.RGBA{})
	img.Clear()
	systems.DrawSystem(win, 110+index)
	img.Batchers[data.UIKey].Draw(scroll.ListView.Canvas)
	//img.Batchers[data.UIKey].Draw(data.InventoryView.Canvas)
	// draw the text
	systems.DrawSystem(win, 113+index)
	// draw to the canvas
	scroll.ListView.Draw(data.InventoryView.Canvas)
}

func DisposeTowerScrolls() {
	DisposeTowerScroll(data.LeftTowerScroll)
	DisposeTowerScroll(data.MidTowerScroll)
	DisposeTowerScroll(data.RightTowerScroll)
}

func DisposeTowerScroll(scroll *data.TowerScroll) {
	for _, e := range scroll.Entities {
		myecs.Manager.DisposeEntity(e)
	}
	if scroll.Scroll != nil {
		systems.DisposeScroll(scroll.Scroll)
	}
	scroll.ListView = nil
	scroll = nil
}

func GetEmptyTierSlot(index int) *data.InvSpellSlot {
	var scroll *data.TowerScroll
	switch index {
	case 0:
		scroll = data.LeftTowerScroll
	case 1:
		scroll = data.MidTowerScroll
	case 2:
		scroll = data.RightTowerScroll
	default:
		panic(fmt.Sprintf("GetEmptyTierSlot: bad tower index %d", index))
	}
	for _, slot := range scroll.InvSlots {
		if slot.Slot.Tier == 0 {
			return slot
		}
	}
	return nil
}

func GetHoveredSlot() *data.InvSpellSlot {
	inPos := data.LeftTowerScroll.ListView.ProjectWorld(data.TheInput.World)
	if data.LeftTowerScroll.ListView.PointInside(inPos) {
		for _, slot := range data.LeftTowerScroll.InvSlots {
			if slot.NameMObj.PointInside(inPos) ||
				slot.TierObj.PointInside(inPos) ||
				slot.SlotObj.PointInside(inPos) {
				return slot
			}
		}
	}
	inPos = data.MidTowerScroll.ListView.ProjectWorld(data.TheInput.World)
	if data.MidTowerScroll.ListView.PointInside(inPos) {
		for _, slot := range data.MidTowerScroll.InvSlots {
			if slot.NameMObj.PointInside(inPos) ||
				slot.TierObj.PointInside(inPos) ||
				slot.SlotObj.PointInside(inPos) {
				return slot
			}
		}
	}
	inPos = data.RightTowerScroll.ListView.ProjectWorld(data.TheInput.World)
	if data.RightTowerScroll.ListView.PointInside(inPos) {
		for _, slot := range data.RightTowerScroll.InvSlots {
			if slot.NameMObj.PointInside(inPos) ||
				slot.TierObj.PointInside(inPos) ||
				slot.SlotObj.PointInside(inPos) {
				return slot
			}
		}
	}
	return nil
}
