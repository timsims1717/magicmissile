package systems

import (
	"github.com/bytearena/ecs"
	"github.com/faiface/pixel"
	"timsims1717/magicmissile/internal/data"
	"timsims1717/magicmissile/internal/myecs"
	"timsims1717/magicmissile/pkg/img"
	"timsims1717/magicmissile/pkg/object"
	"timsims1717/magicmissile/pkg/timing"
	"timsims1717/magicmissile/pkg/viewport"
)

func CreateScrollBar(AddEntity func(*ecs.Entity), parentObj *object.Object, parentView *viewport.ViewPort, view *viewport.ViewPort) {
	arrowRect := img.Batchers[data.UIKey].GetSprite("scroll_arrow_up").Frame()
	upArrowSpr := img.NewSprite("scroll_arrow_up", data.UIKey)
	dwnArrowSpr := img.NewSprite("scroll_arrow_dwn", data.UIKey)
	var upArrowTimer *timing.Timer
	var dwnArrowTimer *timing.Timer
	upArrowClick := false
	dwnArrowClick := false
	arrowUpObj := object.New()
	arrowUpObj.Offset.Y = parentObj.Rect.H()*0.5 - 5.
	arrowUpObj.Offset.X = parentObj.Rect.W()*0.5 - 13.
	arrowUpObj.Layer = 102
	arrowUpObj.SetRect(arrowRect)
	upE := myecs.Manager.NewEntity().
		AddComponent(myecs.Object, arrowUpObj).
		AddComponent(myecs.Parent, parentObj).
		AddComponent(myecs.Drawable, upArrowSpr).
		AddComponent(myecs.Update, data.NewHoverClickFn(data.TheInput, parentView, func(click *data.HoverClick) {
			upArrowSpr.Key = "scroll_arrow_up"
			if click.Hover && !data.MovingSpellSlot.Moving {
				if click.Input.Get("click").JustPressed() {
					upArrowClick = true
					if upArrowTimer == nil {
						upArrowTimer = timing.New(0.35)
					}
					newPos := view.CamPos
					newPos.Y += data.SquareFrame.H()
					view.MoveTo(newPos, 0.15, false)
				} else if click.Input.Get("click").Pressed() && upArrowClick {
					if upArrowTimer.UpdateDone() {
						vel := pixel.V(0., 300.)
						view.SetVel(vel)
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
	AddEntity(upE)
	arrowDwnObj := object.New()
	arrowDwnObj.Offset.Y = parentObj.Rect.H()*-0.5 + 5.
	arrowDwnObj.Offset.X = parentObj.Rect.W()*0.5 - 13.
	arrowDwnObj.Layer = 102
	arrowDwnObj.SetRect(arrowRect)
	dwnE := myecs.Manager.NewEntity().
		AddComponent(myecs.Object, arrowDwnObj).
		AddComponent(myecs.Parent, parentObj).
		AddComponent(myecs.Drawable, dwnArrowSpr).
		AddComponent(myecs.Update, data.NewHoverClickFn(data.TheInput, parentView, func(click *data.HoverClick) {
			dwnArrowSpr.Key = "scroll_arrow_dwn"
			if click.Hover && !data.MovingSpellSlot.Moving {
				if click.Input.Get("click").JustPressed() {
					dwnArrowClick = true
					if dwnArrowTimer == nil {
						dwnArrowTimer = timing.New(0.35)
					}
					newPos := view.CamPos
					newPos.Y -= data.SquareFrame.H()
					view.MoveTo(newPos, 0.15, false)
				} else if click.Input.Get("click").Pressed() && dwnArrowClick {
					if dwnArrowTimer.UpdateDone() {
						vel := pixel.V(0., -300.)
						view.SetVel(vel)
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
	AddEntity(dwnE)

	// scroll bar
	barClick := false
	barObj := object.New()
	barObj.SetRect(img.Batchers[data.UIKey].GetSprite("scroll_bar").Frame())
	barBot := parentObj.Rect.H()*-0.5 + (barObj.Rect.H()+arrowRect.H())*0.5 + 5.
	barTop := parentObj.Rect.H()*0.5 - (barObj.Rect.H()+arrowRect.H())*0.5 - 5.
	barObj.Offset.Y = barTop
	barObj.Offset.X = parentObj.Rect.W()*0.5 - 13.
	barObj.Layer = 102
	barOffset := 0.
	barE := myecs.Manager.NewEntity().
		AddComponent(myecs.Object, barObj).
		AddComponent(myecs.Parent, parentObj).
		AddComponent(myecs.Drawable, img.NewSprite("scroll_bar", data.UIKey)).
		AddComponent(myecs.Update, data.NewHoverClickFn(data.TheInput, parentView, func(hvc *data.HoverClick) {
			if hvc.Hover {
				if hvc.Input.Get("click").JustPressed() {
					barClick = true
					barOffset = barObj.PostPos.Y - hvc.View.ProjectWorld(hvc.Input.World).Y
				}
			}
			if !hvc.Input.Get("click").Pressed() {
				barClick = false
			}
			listCamTop, listCamBot := view.GetLimY()
			if barClick && !data.MovingSpellSlot.Moving {
				inPos := hvc.View.ProjectWorld(hvc.Input.World)
				barObj.Offset.Y -= barObj.PostPos.Y - inPos.Y - barOffset
				if barObj.Offset.Y > barTop {
					barObj.Offset.Y = barTop
				} else if barObj.Offset.Y < barBot {
					barObj.Offset.Y = barBot
				}
				barRatio := (barTop - barObj.Offset.Y) / (barTop - barBot)
				view.CamPos.Y = -(barRatio * listCamTop) + (barRatio * listCamBot) + listCamTop
			}
			camRatio := (listCamTop - view.CamPos.Y) / (listCamTop - listCamBot)
			barObj.Offset.Y = -(camRatio * barTop) + (camRatio * barBot) + barTop
		}))
	AddEntity(barE)
}
