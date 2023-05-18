package payloads

import (
	"timsims1717/magicmissile/internal/data"
	"timsims1717/magicmissile/internal/systems"
	"timsims1717/magicmissile/pkg/object"
	"timsims1717/magicmissile/pkg/util"
)

func Flamewall(mis *data.Missile, obj *object.Object) {
	for i := 0; i < mis.Tier+1; i++ {
		rMis := &data.Missile{
			SprKey:  mis.SprKey,
			Speed:   mis.Speed,
			Colors:  mis.Colors,
			Payload: []data.Payload{{Explosion: mis.Payload[0].Explosion.Copy()}},
		}
		lMis := &data.Missile{
			SprKey:  mis.SprKey,
			Speed:   mis.Speed,
			Colors:  mis.Colors,
			Payload: []data.Payload{{Explosion: mis.Payload[0].Explosion.Copy()}},
		}
		rTarget := obj.Pos
		rTarget.X += 20. * float64(i+1)
		lTarget := obj.Pos
		lTarget.X -= 20. * float64(i+1)
		systems.MakeMissile(rMis, obj.Pos, rTarget)
		systems.MakeMissile(lMis, obj.Pos, lTarget)
	}
}

func Disintegrate(mis *data.Missile, obj *object.Object) {
	nMis := &data.Missile{
		Limit:   20.,
		SprKey:  mis.SprKey,
		Speed:   mis.Speed,
		Colors:  mis.Colors,
		Payload: []data.Payload{{Explosion: mis.Payload[0].Explosion.Copy()}, {Function: Disintegrate}},
		Tier:    mis.Tier,
	}
	tarLen := util.Normalize(mis.Target.Sub(obj.Pos))
	tarLen = tarLen.Scaled(mis.Limit)
	target := mis.Target.Add(tarLen)
	systems.MakeMissile(nMis, obj.Pos, target)
}
