package typeface

import (
	"fmt"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/text"
	"image/color"
	"timsims1717/magicmissile/pkg/object"
)

type Text struct {
	Raw     string
	Text    *text.Text
	Color   color.RGBA
	Align   Alignment
	Symbols []symbolHandle
	NoShow  bool

	Increment bool
	CurrPos   int
	Width     float64
	Height    float64
	MaxWidth  float64
	MaxHeight float64
	MaxLines  int

	Parent       *pixel.Vec
	RelativeSize float64
	SymbolSize   float64
	Obj          *object.Object

	rawLines   []string
	lineWidths []float64
	fullHeight float64
}

func New(parent *pixel.Vec, atlas string, align Alignment, lineHeight, relativeSize, maxWidth, maxHeight float64) *Text {
	tex := text.New(pixel.ZV, Atlases[atlas])
	tex.LineHeight *= lineHeight
	obj := object.New()
	obj.Sca = pixel.V(relativeSize, relativeSize)
	return &Text{
		Text:  tex,
		Align: align,
		Color: color.RGBA{
			R: 255,
			G: 255,
			B: 255,
			A: 255,
		},
		Width:        maxWidth,
		Height:       maxHeight,
		MaxWidth:     maxWidth,
		MaxHeight:    maxHeight,
		MaxLines:     int(maxHeight / (tex.LineHeight * relativeSize)),
		Parent:       parent,
		RelativeSize: relativeSize,
		SymbolSize:   1.,
		Obj:          obj,
	}
}

func (item *Text) Draw(target pixel.Target) {
	if !item.NoShow {
		item.Text.Draw(target, item.Obj.Mat)
	}
}

func (item *Text) SetWidth(width float64) {
	item.MaxWidth = width
	item.SetText(item.Raw)
}

func (item *Text) SetHeight(height float64) {
	item.MaxHeight = height
	item.SetText(item.Raw)
}

func (item *Text) SetColor(col color.RGBA) {
	item.Color = col
	item.updateText()
}

func (item *Text) SetSize(size float64) {
	item.RelativeSize = size
	item.SetText(item.Raw)
}

func (item *Text) SetPos(pos pixel.Vec) {
	item.Obj.Pos = pos
	item.updateText()
}

func (item *Text) IncrementTextPos() {
	if item.Increment {

	}
}

func (item *Text) SkipIncrement() {
	if item.Increment {

	}
}

func (item *Text) PrintLines() {
	for _, line := range item.rawLines {
		fmt.Println(line)
	}
}
