package data

import (
	"github.com/faiface/pixel"
	"golang.org/x/image/colornames"
	"image/color"
	"timsims1717/magicmissile/pkg/util"
)

const (
	// batchers
	ObjectKey   = "objects"
	ParticleKey = "particles"
	UIKey       = "ui"

	BaseWidth  = 1600
	BaseHeight = 900
)

var (
	Highlight = pixel.ToRGBA(color.RGBA{
		R: 255,
		G: 0,
		B: 175,
		A: 255,
	})
	ScrollText = pixel.ToRGBA(color.RGBA{
		R: 61,
		G: 53,
		B: 40,
		A: 255,
	})
	WhiteText = util.White
	Red       = pixel.ToRGBA(colornames.Red)
	Yellow    = pixel.ToRGBA(colornames.Yellow)
	Gray      = pixel.ToRGBA(colornames.Lightgrey)
	Green     = pixel.ToRGBA(colornames.Lightgreen)
)
