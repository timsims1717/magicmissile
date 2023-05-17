package loading

import (
	"timsims1717/magicmissile/internal/data"
	"timsims1717/magicmissile/pkg/img"
)

func LoadImg() {
	objSheet, err := img.LoadSpriteSheet("assets/img/objects.json")
	if err != nil {
		panic(err)
	}
	img.AddBatcher(data.ObjectKey, objSheet, true, true)

	partSheet, err := img.LoadSpriteSheet("assets/img/particles.json")
	if err != nil {
		panic(err)
	}
	img.AddBatcher(data.ParticleKey, partSheet, true, true)
}
