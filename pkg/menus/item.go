package menus

import (
	"github.com/faiface/pixel"
	"timsims1717/magicmissile/pkg/typeface"
)

type Item struct {
	Key  string
	Raw  string
	Hint string
	Text *typeface.Text

	clickFn   func()
	leftFn    func()
	rightFn   func()
	hoverFn   func()
	unHoverFn func()

	Right    bool
	Hovered  bool
	Disabled bool
	NoHover  bool
	Ignore   bool
	NoDraw   bool
	hovered  bool
	disabled bool
	noShowT  bool
	CurrLine int
}

func NewItem(key, raw string, right bool) *Item {
	align := typeface.Left
	if right {
		align = typeface.Right
	}
	tex := typeface.New("main", typeface.NewAlign(typeface.Align(align), typeface.Bottom), 1.5, MenuSize, 0., 0.)
	tex.SetColor(DefaultColor)
	tex.SetText(raw)
	return &Item{
		Key:   key,
		Raw:   raw,
		Text:  tex,
		Right: right,
	}
}

func (i *Item) Update() {
	if i.Disabled && !i.disabled {
		i.disabled = true
		i.hovered = false
		i.Text.SetColor(DisabledColor)
		i.Text.SetSize(MenuSize)
	} else if !i.Disabled && i.disabled {
		i.disabled = false
		i.Text.SetColor(DefaultColor)
		i.Text.SetSize(MenuSize)
	}
	if !i.disabled {
		if i.Hovered && !i.hovered {
			i.hovered = true
			i.Text.SetColor(HoverColor)
			i.Text.SetSize(HoverSize)
		} else if !i.Hovered && i.hovered {
			i.hovered = false
			i.Text.SetColor(DefaultColor)
			i.Text.SetSize(MenuSize)
		}
	}
}

func (i *Item) Draw(target pixel.Target) {
	if i.Text != nil && !i.Ignore && !i.noShowT && !i.NoDraw {
		i.Text.Draw(target)
	}
}

func (i *Item) SetText(raw string) {
	i.Text.SetText(raw)
}

func (i *Item) SetHoverFn(fn func()) {
	i.hoverFn = fn
}

func (i *Item) SetUnhoverFn(fn func()) {
	i.unHoverFn = fn
}

func (i *Item) SetClickFn(fn func()) {
	i.clickFn = fn
}

func (i *Item) SetLeftFn(fn func()) {
	i.leftFn = fn
}

func (i *Item) SetRightFn(fn func()) {
	i.rightFn = fn
}
