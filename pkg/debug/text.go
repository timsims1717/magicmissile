package debug

import (
	"fmt"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
	"strings"
	"timsims1717/magicmissile/pkg/typeface"
)

var (
	debugText *typeface.Text
	lines     = &strings.Builder{}
)

func InitializeText(v *pixel.Vec) {
	col := colornames.Aliceblue
	col.A = 90
	debugText = typeface.New(v, "basic", typeface.NewAlign(typeface.Left, typeface.Top), 1.0, 2.0, 0., 0.)
}

func DrawText(win *pixelgl.Window) {
	debugText.SetText(lines.String())
	debugText.Obj.Pos = pixel.V(5., win.Bounds().H()-5.)
	debugText.Obj.Update()
	debugText.Draw(win)
}

func AddText(s string) {
	lines.WriteString(fmt.Sprintf("%s\n", s))
}
