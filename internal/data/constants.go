package data

import "image/color"

const (
	// batchers
	ObjectKey   = "objects"
	ParticleKey = "particles"
	UIKey       = "ui"

	BaseWidth  = 1600
	BaseHeight = 900
)

var (
	Highlight = color.RGBA{
		R: 255,
		G: 0,
		B: 175,
		A: 255,
	}
	ScrollText = color.RGBA{
		R: 61,
		G: 53,
		B: 40,
		A: 255,
	}
)
