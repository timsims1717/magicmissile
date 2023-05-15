package data

import (
	"github.com/faiface/pixel/imdraw"
	"timsims1717/magicmissile/pkg/viewport"
)

var (
	GameView *viewport.ViewPort
	ExpView  *viewport.ViewPort
	ExpView1 *viewport.ViewPort
	GameDraw *imdraw.IMDraw

	ExpDrawType = 2
	ExpTestNum  = 1
	ExpTexture  []uint8
)

const (
	BaseWidth  = 1600
	BaseHeight = 900
)
