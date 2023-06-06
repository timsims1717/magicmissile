package systems

import (
	"github.com/faiface/pixel"
	"timsims1717/magicmissile/internal/data"
	"timsims1717/magicmissile/internal/myecs"
	"timsims1717/magicmissile/pkg/img"
	"timsims1717/magicmissile/pkg/object"
	"timsims1717/magicmissile/pkg/timing"
)

func CreateScroll(orig pixel.Vec, baseLayer int, startSize pixel.Vec, fullSize pixel.Vec) *data.Scroll {
	scroll := &data.Scroll{}
	scroll.Object = object.New()
	scroll.Object.Pos = orig
	scroll.CurrDim = startSize
	scroll.FullDim = fullSize
	scroll.TLPos = pixel.V(scroll.FullDim.X*-0.5, scroll.FullDim.Y*0.5)
	scroll.Entity = myecs.Manager.NewEntity().AddComponent(myecs.Object, scroll.Object)
	scroll.AddEntity(scroll.Entity)

	// Scroll Mid
	scroll.MidMid = object.New()
	scroll.MidMid.Offset = pixel.V(0., 0.)
	scroll.MidMid.Sca = pixel.V(startSize.X/data.TileSize, startSize.Y/(data.TileSize*3))
	scroll.MidMid.Layer = baseLayer
	scroll.AddEntity(myecs.Manager.NewEntity().
		AddComponent(myecs.Object, scroll.MidMid).
		AddComponent(myecs.Parent, scroll.Object).
		AddComponent(myecs.Drawable, img.NewSprite("scroll_m", data.UIKey)))
	scroll.MidLeft = object.New()
	scroll.MidLeft.SetRect(img.Batchers[data.UIKey].GetSprite("scroll_ml").Frame())
	scroll.MidLeft.Offset = pixel.V((scroll.MidLeft.Rect.W()+startSize.X)*-0.5, 0.)
	scroll.MidLeft.Sca = pixel.V(1., startSize.Y/(data.TileSize*3))
	scroll.MidLeft.Layer = baseLayer
	scroll.AddEntity(myecs.Manager.NewEntity().
		AddComponent(myecs.Object, scroll.MidLeft).
		AddComponent(myecs.Parent, scroll.Object).
		AddComponent(myecs.Drawable, img.NewSprite("scroll_ml", data.UIKey)))
	scroll.MidRight = object.New()
	scroll.MidRight.SetRect(img.Batchers[data.UIKey].GetSprite("scroll_mr").Frame())
	scroll.MidRight.Offset = pixel.V((scroll.MidRight.Rect.W()+startSize.X)*0.5, 0.)
	scroll.MidRight.Sca = pixel.V(1., startSize.Y/(data.TileSize*3))
	scroll.MidRight.Layer = baseLayer
	scroll.AddEntity(myecs.Manager.NewEntity().
		AddComponent(myecs.Object, scroll.MidRight).
		AddComponent(myecs.Parent, scroll.Object).
		AddComponent(myecs.Drawable, img.NewSprite("scroll_mr", data.UIKey)))

	// Scroll Top
	scroll.TopMid = object.New()
	scroll.TopMid.SetRect(img.Batchers[data.UIKey].GetSprite("scroll_t").Frame())
	scroll.TopMid.Offset = pixel.V(0., (scroll.TopMid.Rect.H()+startSize.Y)*0.5)
	scroll.TopMid.Sca = pixel.V(startSize.X/data.TileSize, 1.)
	scroll.TopMid.Layer = baseLayer
	scroll.AddEntity(myecs.Manager.NewEntity().
		AddComponent(myecs.Object, scroll.TopMid).
		AddComponent(myecs.Parent, scroll.Object).
		AddComponent(myecs.Drawable, img.NewSprite("scroll_t", data.UIKey)))
	scroll.TopLeft = object.New()
	scroll.TopLeft.SetRect(img.Batchers[data.UIKey].GetSprite("scroll_tl").Frame())
	scroll.TopLeft.Offset = pixel.V((scroll.TopLeft.Rect.W()+startSize.X)*-0.5, (scroll.TopLeft.Rect.H()+startSize.Y)*0.5)
	scroll.TopLeft.Layer = baseLayer
	scroll.AddEntity(myecs.Manager.NewEntity().
		AddComponent(myecs.Object, scroll.TopLeft).
		AddComponent(myecs.Parent, scroll.Object).
		AddComponent(myecs.Drawable, img.NewSprite("scroll_tl", data.UIKey)))
	scroll.TopRight = object.New()
	scroll.TopRight.SetRect(img.Batchers[data.UIKey].GetSprite("scroll_tr").Frame())
	scroll.TopRight.Offset = pixel.V((scroll.TopRight.Rect.W()+startSize.X)*0.5, (scroll.TopRight.Rect.H()+startSize.Y)*0.5)
	scroll.TopRight.Layer = baseLayer
	scroll.AddEntity(myecs.Manager.NewEntity().
		AddComponent(myecs.Object, scroll.TopRight).
		AddComponent(myecs.Parent, scroll.Object).
		AddComponent(myecs.Drawable, img.NewSprite("scroll_tr", data.UIKey)))

	// Scroll Bot
	scroll.BotMid = object.New()
	scroll.BotMid.SetRect(img.Batchers[data.UIKey].GetSprite("scroll_b").Frame())
	scroll.BotMid.Offset = pixel.V(0., (scroll.BotMid.Rect.H()+startSize.Y)*-0.5)
	scroll.BotMid.Sca = pixel.V(startSize.X/data.TileSize, 1.)
	scroll.BotMid.Layer = baseLayer
	scroll.AddEntity(myecs.Manager.NewEntity().
		AddComponent(myecs.Object, scroll.BotMid).
		AddComponent(myecs.Parent, scroll.Object).
		AddComponent(myecs.Drawable, img.NewSprite("scroll_b", data.UIKey)))
	scroll.BotLeft = object.New()
	scroll.BotLeft.SetRect(img.Batchers[data.UIKey].GetSprite("scroll_bl").Frame())
	scroll.BotLeft.Offset = pixel.V((scroll.BotLeft.Rect.W()+startSize.X)*-0.5, (scroll.BotLeft.Rect.H()+startSize.Y)*-0.5)
	scroll.BotLeft.Layer = baseLayer
	scroll.AddEntity(myecs.Manager.NewEntity().
		AddComponent(myecs.Object, scroll.BotLeft).
		AddComponent(myecs.Parent, scroll.Object).
		AddComponent(myecs.Drawable, img.NewSprite("scroll_bl", data.UIKey)))
	scroll.BotRight = object.New()
	scroll.BotRight.SetRect(img.Batchers[data.UIKey].GetSprite("scroll_br").Frame())
	scroll.BotRight.Offset = pixel.V((scroll.BotRight.Rect.W()+startSize.X)*0.5, (scroll.BotRight.Rect.H()+startSize.Y)*-0.5)
	scroll.BotRight.Layer = baseLayer
	scroll.AddEntity(myecs.Manager.NewEntity().
		AddComponent(myecs.Object, scroll.BotRight).
		AddComponent(myecs.Parent, scroll.Object).
		AddComponent(myecs.Drawable, img.NewSprite("scroll_br", data.UIKey)))

	return scroll
}

func UpdateScroll(scroll *data.Scroll) {
	if !scroll.Freeze {
		diff := 1500. * timing.DT
		if diff > scroll.FullDim.Y*0.2 {
			diff = scroll.FullDim.Y * 0.2
		}
		if scroll.CurrDim.X < scroll.FullDim.X {
			scroll.CurrDim.X += diff
			if scroll.CurrDim.X > scroll.FullDim.X {
				scroll.CurrDim.X = scroll.FullDim.X
			}
		}
		if scroll.CurrDim.Y < scroll.FullDim.Y {
			scroll.CurrDim.Y += diff
			if scroll.CurrDim.Y > scroll.FullDim.Y {
				scroll.CurrDim.Y = scroll.FullDim.Y
			}
		}
		// Scroll Mid
		scroll.MidMid.Sca = pixel.V(scroll.CurrDim.X/data.TileSize, scroll.CurrDim.Y/(data.TileSize*3))
		scroll.MidLeft.Offset = pixel.V((scroll.MidLeft.Rect.W()+scroll.CurrDim.X)*-0.5, 0.)
		scroll.MidLeft.Sca = pixel.V(1., scroll.CurrDim.Y/(data.TileSize*3))
		scroll.MidRight.Offset = pixel.V((scroll.MidRight.Rect.W()+scroll.CurrDim.X)*0.5, 0.)
		scroll.MidRight.Sca = pixel.V(1., scroll.CurrDim.Y/(data.TileSize*3))
		// Scroll Top
		scroll.TopMid.Offset = pixel.V(0., (scroll.TopMid.Rect.H()+scroll.CurrDim.Y)*0.5)
		scroll.TopMid.Sca = pixel.V(scroll.CurrDim.X/data.TileSize, 1.)
		scroll.TopLeft.Offset = pixel.V((scroll.TopLeft.Rect.W()+scroll.CurrDim.X)*-0.5, (scroll.TopLeft.Rect.H()+scroll.CurrDim.Y)*0.5)
		scroll.TopRight.Offset = pixel.V((scroll.TopRight.Rect.W()+scroll.CurrDim.X)*0.5, (scroll.TopRight.Rect.H()+scroll.CurrDim.Y)*0.5)
		// Scroll Bot
		scroll.BotMid.Offset = pixel.V(0., (scroll.BotMid.Rect.H()+scroll.CurrDim.Y)*-0.5)
		scroll.BotMid.Sca = pixel.V(scroll.CurrDim.X/data.TileSize, 1.)
		scroll.BotLeft.Offset = pixel.V((scroll.BotLeft.Rect.W()+scroll.CurrDim.X)*-0.5, (scroll.BotLeft.Rect.H()+scroll.CurrDim.Y)*-0.5)
		scroll.BotRight.Offset = pixel.V((scroll.BotRight.Rect.W()+scroll.CurrDim.X)*0.5, (scroll.BotRight.Rect.H()+scroll.CurrDim.Y)*-0.5)
	}
	scroll.Opened = scroll.CurrDim == scroll.FullDim
}
