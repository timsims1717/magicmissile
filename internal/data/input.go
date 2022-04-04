package data

import (
	"github.com/faiface/pixel/pixelgl"
	"timsims1717/magicmissile/pkg/input"
)

var TheInput = &input.Input{
	Buttons: map[string]*input.ButtonSet{
		"moveLeft":  input.NewJoyless(pixelgl.KeyA),
		"moveRight": input.NewJoyless(pixelgl.KeyD),
		"1":         input.NewJoyless(pixelgl.Key1),
		"2":         input.NewJoyless(pixelgl.Key2),
		"click":     input.NewJoyless(pixelgl.MouseButtonLeft),
		"switchR":   {
			Keys:     []pixelgl.Button{pixelgl.KeyE, pixelgl.KeySpace},
			Scroll:   1,
		},
		"switchL":   {
			Keys:     []pixelgl.Button{pixelgl.KeyQ},
			Scroll:   -1,
		},
		"killAll": input.NewJoyless(pixelgl.KeyF5),
		"menuBack": input.NewJoyless(pixelgl.KeyEscape),
	},
	Mode: input.KeyboardMouse,
}
