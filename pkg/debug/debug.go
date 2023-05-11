package debug

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"strings"
)

var (
	Debug = false
	Text  = false
)

func Initialize(txt, fps *pixel.Vec) {
	InitializeLines()
	InitializeText(txt)
	InitializeFPS(fps)
}

func Draw(win *pixelgl.Window) {
	if Text {
		DrawText(win)
	}
	DrawFPS(win)
}

func Clear() {
	imd.Clear()
	lines = &strings.Builder{}
}
