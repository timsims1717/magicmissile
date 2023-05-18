package payloads

import (
	"github.com/faiface/pixel"
	"math/rand"
	"timsims1717/magicmissile/internal/data"
	"timsims1717/magicmissile/internal/systems"
	"timsims1717/magicmissile/pkg/object"
	"timsims1717/magicmissile/pkg/util"
)

func IceShatter(mis *data.Missile, obj *object.Object) {
	count := mis.Tier*8 + rand.Intn(8) + rand.Intn(8)
	for i := 0; i < count; i++ {
		nMis := &data.Missile{
			SprKey:  mis.SprKey,
			Speed:   mis.Speed,
			Colors:  mis.Colors,
			Payload: []data.Payload{{Explosion: mis.Payload[0].Explosion.Copy()}},
		}
		offset := pixel.ZV
		offset.X = rand.Float64()*2. - 1.
		offset.Y = rand.Float64()*2. - 1.
		offset = util.Normalize(offset)
		offset = offset.Scaled(float64(25 + rand.Intn(mis.Tier*30)))
		target := obj.Pos.Add(offset)
		systems.MakeMissile(nMis, obj.Pos, target)
	}
}
