package systems

import (
	"github.com/aquilax/go-perlin"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"golang.org/x/image/colornames"
	"image/color"
	"math/rand"
	"timsims1717/magicmissile/internal/data"
	"timsims1717/magicmissile/internal/shaders"
	"timsims1717/magicmissile/pkg/viewport"
)

func GenerateBackground() {
	baseColor := colornames.Darkmagenta
	data.Backgrounds = []*data.Background{}
	for i := 0; i < 3; i++ {
		var seed = rand.Int63n(data.MaximumSeedValue)
		p := perlin.NewPerlin(data.Alpha, data.Beta, data.N, seed)
		l := uint8(i * 10)
		r := baseColor.R + l
		g := baseColor.G + l
		b := baseColor.B + l
		if r > 255 {
			r = 255
		}
		if g > 255 {
			g = 255
		}
		if b > 255 {
			b = 255
		}
		col := color.RGBA{
			R: r,
			G: g,
			B: b,
			A: 255,
		}
		vp := viewport.New(nil)
		vp.SetRect(pixel.R(0, 0, data.BaseWidth, data.BaseHeight))
		vp.CamPos = pixel.ZV
		vp.PortPos = viewport.MainCamera.PostCamPos
		r1 := float32(col.R) / 255.
		vp.Canvas.SetUniform("uRed", r1)
		vp.Canvas.SetUniform("uGreen", float32(col.G)/255.)
		vp.Canvas.SetUniform("uBlue", float32(col.B)/255.)
		vp.Canvas.SetFragmentShader(shaders.BGShader)
		data.Backgrounds = append(data.Backgrounds, &data.Background{
			Layer:  i,
			Perlin: p,
			Color:  col,
			IMDraw: imdraw.New(nil),
			View:   vp,
		})
	}
}

func UpdateBackgrounds() {
	for _, bg := range data.Backgrounds {
		bg.View.Update()
		bg.IMDraw.Clear()
		bg.IMDraw.Intensity = 1.
		bg.IMDraw.Color = bg.Color
		for x := -data.BaseWidth; x < data.BaseWidth; x++ {
			y := bg.Perlin.Noise1D(float64(x)/data.WaveLength)*data.Scale + data.VerticalOffset - float64(bg.Layer*50)
			renderTexturedLine(float64(x), y, bg.IMDraw)
		}
	}
}

func renderTexturedLine(x, y float64, imd *imdraw.IMDraw) {
	//imd.Picture = pixel.V(x, 0)
	imd.Push(pixel.V(x, -data.BaseHeight))
	//imd.Picture = pixel.V(x+1, y)
	imd.Push(pixel.V(x+1, y))
	imd.Rectangle(0)
}
