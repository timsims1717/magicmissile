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

	TownLayer = 0
	Towns     []*Town
	Towers    []*Tower
)
