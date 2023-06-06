package data

import (
	"github.com/faiface/pixel/pixelgl"
	pxginput "github.com/timsims1717/pixel-go-input"
)

var TheInput = &pxginput.Input{
	Buttons: map[string]*pxginput.ButtonSet{
		"fireLeft":  pxginput.NewJoyless(pixelgl.KeyA),
		"fireMid":   pxginput.NewJoyless(pixelgl.KeyS),
		"fireRight": pxginput.NewJoyless(pixelgl.KeyD),
		"click":     pxginput.NewJoyless(pixelgl.MouseButtonLeft),
		"switchR": {
			Keys:   []pixelgl.Button{pixelgl.KeyE, pixelgl.KeySpace},
			Scroll: 1,
		},
		"switchL": {
			Keys:   []pixelgl.Button{pixelgl.KeyQ},
			Scroll: -1,
		},
		"showInventory":    pxginput.NewJoyless(pixelgl.KeyI),
		"killAll":          pxginput.NewJoyless(pixelgl.KeyF10),
		"menuBack":         pxginput.NewJoyless(pixelgl.KeyEscape),
		"fullscreen":       pxginput.NewJoyless(pixelgl.KeyF5),
		"fuzzy":            pxginput.NewJoyless(pixelgl.KeyF6),
		"debugText":        pxginput.NewJoyless(pixelgl.KeyF4),
		"debugCU":          pxginput.NewJoyless(pixelgl.KeyKP8),
		"debugCD":          pxginput.NewJoyless(pixelgl.KeyKP5),
		"debugCR":          pxginput.NewJoyless(pixelgl.KeyKP6),
		"debugCL":          pxginput.NewJoyless(pixelgl.KeyKP4),
		"changeBackground": pxginput.NewJoyless(pixelgl.KeyF2),
		"debugExpDrawType": pxginput.NewJoyless(pixelgl.KeyF11),
		"debugSpellTier":   pxginput.NewJoyless(pixelgl.KeyKPAdd),
		"debugSpellName":   pxginput.NewJoyless(pixelgl.KeyKPSubtract),
	},
	Mode: pxginput.KeyboardMouse,
}
