package payloads

import (
	"github.com/faiface/pixel"
	"golang.org/x/image/colornames"
	"image/color"
	"timsims1717/magicmissile/internal/data"
	"timsims1717/magicmissile/internal/systems"
	"timsims1717/magicmissile/pkg/object"
	"timsims1717/magicmissile/pkg/util"
)

func RainbowSpray(mis *data.Missile, obj *object.Object) {
	var col color.RGBA
	switch mis.Tier {
	case 5:
		col = colornames.Purple
	case 6:
		col = colornames.Indigo
	case 7:
		col = colornames.Blue
	case 8:
		col = colornames.Green
	case 9:
		col = colornames.Yellow
	case 10:
		col = colornames.Orange
	case 11:
		col = colornames.Red
	}
	exp := &data.Explosion{
		FullRadius: float64(mis.Tier) * 6.,
		ExpandRate: 2,
		Dissipate:  2.5,
		DisRate:    50,
		Color:      pixel.ToRGBA(col),
	}
	systems.MakeExplosion(exp, obj.Pos, exp.Color)
	if mis.Tier < 11 {
		nMis := &data.Missile{
			Limit:   float64(mis.Tier) * 4,
			SprKey:  mis.SprKey,
			Speed:   mis.Speed,
			Colors:  mis.Colors,
			Payload: []data.Payload{{Function: RainbowSpray}},
			Tier:    mis.Tier + 1,
		}
		tarLen := util.Normalize(mis.Target.Sub(obj.Pos))
		tarLen = tarLen.Scaled(nMis.Limit)
		target := mis.Target.Add(tarLen)
		systems.MakeMissile(nMis, obj.Pos, target)
	}
}
