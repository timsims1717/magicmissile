package data

import (
	"github.com/aquilax/go-perlin"
	"github.com/faiface/pixel/imdraw"
	"image/color"
	"timsims1717/magicmissile/pkg/viewport"
)

type Background struct {
	Layer  int
	Perlin *perlin.Perlin
	Color  color.Color
	View   *viewport.ViewPort
	IMDraw *imdraw.IMDraw
}

var (
	Scale            = 100.
	WaveLength       = 100.
	Alpha            = 2.
	Beta             = 2.
	N                = int32(3)
	MaximumSeedValue = int64(100)
	VerticalOffset   = float64(BaseHeight) / 4.

	Backgrounds []*Background
)
