package figures

import (
	"github.com/faiface/pixel"
	"timsims1717/magicmissile/pkg/img"
)

func WizardFigure(col pixel.RGBA) *img.Sprite {
	return &img.Sprite{
		Key:   "wizard",
		Color: col,
		Batch: "figures",
	}
}

func FighterFigure(col pixel.RGBA) *img.Sprite {
	return &img.Sprite{
		Key:   "fighter",
		Color: col,
		Batch: "figures",
	}
}

func ZombieFigure(col pixel.RGBA) *img.Sprite {
	return &img.Sprite{
		Key:   "zombie",
		Color: col,
		Batch: "figures",
	}
}
