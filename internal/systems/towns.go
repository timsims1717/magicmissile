package systems

import (
	"math/rand"
	"timsims1717/magicmissile/internal/data"
	"timsims1717/magicmissile/internal/myecs"
	"timsims1717/magicmissile/pkg/img"
	"timsims1717/magicmissile/pkg/object"
)

func CreateTowns() {
	if data.CurrBackground != nil {
		bg := data.CurrBackground.Backgrounds[data.TownLayer]
		p := bg.Perlin
		div := float64(data.BaseWidth / 10)
		data.Towns = []*data.Town{}
		for i := 0; i < 7; i++ {
			if i != 3 {
				spr := img.NewSprite("town_1", data.ObjectKey)
				obj := object.New()
				x := div * float64(-3+i)
				obj.Pos.X = x
				obj.Pos.Y = p.Noise1D(x/bg.Layer.WaveLength)*bg.Layer.Scale + data.ForeOffset + bg.Layer.VOffset(x)
				obj.Layer = data.TownLayer
				obj.Flip = rand.Intn(2) == 0
				obj.Offset.Y += img.Batchers[data.ObjectKey].GetSprite(spr.Key).Frame().H()*0.5 - 10.
				obj.Pos.Y += 6.
				e := myecs.Manager.NewEntity()
				e.AddComponent(myecs.Object, obj).
					AddComponent(myecs.Drawable, spr)
				data.Towns = append(data.Towns, &data.Town{
					Health: nil,
					Object: obj,
					Sprite: spr,
					Entity: e,
				})
			}
		}
	} else {
		panic("towns can't be created without a background")
	}
}

func DisposeTowns() {
	for _, town := range data.Towns {
		myecs.Manager.DisposeEntity(town.Entity)
	}
	data.Towns = []*data.Town{}
}
