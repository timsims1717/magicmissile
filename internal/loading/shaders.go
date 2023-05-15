package loading

import (
	"timsims1717/magicmissile/internal/data"
	"timsims1717/magicmissile/pkg/shaders"
)

func LoadShaders() {
	sh, err := shaders.LoadFileToString("assets/shaders/bg-shader.frag.glsl")
	if err != nil {
		panic(err)
	}
	data.BGShader = sh
}
