package typeface

import (
	"github.com/faiface/pixel"
	"timsims1717/magicmissile/pkg/object"
)

var (
	theSymbols = map[string]symbol{}
)

type symbol struct {
	spr *pixel.Sprite
	sca float64
}

type symbolHandle struct {
	symbol symbol
	trans  *object.Object
}

func RegisterSymbol(key string, spr *pixel.Sprite, scalar float64) {
	theSymbols[key] = symbol{
		spr: spr,
		sca: scalar,
	}
}