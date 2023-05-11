package data

import (
	"github.com/faiface/pixel/pixelgl"
	pxginput "github.com/timsims1717/pixel-go-input"
)

var TheInput = &pxginput.Input{
	Buttons: map[string]*pxginput.ButtonSet{
		"moveLeft":  pxginput.NewJoyless(pixelgl.KeyA),
		"moveRight": pxginput.NewJoyless(pixelgl.KeyD),
		"1":         pxginput.NewJoyless(pixelgl.Key1),
		"2":         pxginput.NewJoyless(pixelgl.Key2),
		"click":     pxginput.NewJoyless(pixelgl.MouseButtonLeft),
		"switchR": {
			Keys:   []pixelgl.Button{pixelgl.KeyE, pixelgl.KeySpace},
			Scroll: 1,
		},
		"switchL": {
			Keys:   []pixelgl.Button{pixelgl.KeyQ},
			Scroll: -1,
		},
		"killAll":    pxginput.NewJoyless(pixelgl.KeyF10),
		"menuBack":   pxginput.NewJoyless(pixelgl.KeyEscape),
		"fullscreen": pxginput.NewJoyless(pixelgl.KeyF5),
		"debugText":  pxginput.NewJoyless(pixelgl.KeyF4),
		"debugCU":    pxginput.NewJoyless(pixelgl.KeyKP8),
		"debugCD":    pxginput.NewJoyless(pixelgl.KeyKP5),
		"debugCR":    pxginput.NewJoyless(pixelgl.KeyKP6),
		"debugCL":    pxginput.NewJoyless(pixelgl.KeyKP4),
	},
	Mode: pxginput.KeyboardMouse,
}
