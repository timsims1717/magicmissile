package figures

import (
	"image/color"
	"timsims1717/magicmissile/pkg/img"
)

func WizardFigure(col color.RGBA) *img.Sprite {
	return &img.Sprite{
		Key:   "wizard",
		Color: col,
		Batch: "figures",
	}
}

func FighterFigure(col color.RGBA) *img.Sprite {
	return &img.Sprite{
		Key:   "fighter",
		Color: col,
		Batch: "figures",
	}
}

func ZombieFigure(col color.RGBA) *img.Sprite {
	return &img.Sprite{
		Key:   "zombie",
		Color: col,
		Batch: "figures",
	}
}