package systems

import (
	"github.com/aquilax/go-perlin"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"math/rand"
	"timsims1717/magicmissile/internal/data"
	"timsims1717/magicmissile/pkg/viewport"
)

func GenerateRandomBackground(code string) {
	realm := data.AllRealms[code]
	data.CurrBackground = realm.Backgrounds[rand.Intn(len(realm.Backgrounds))]
	GenerateBackground()
}

func GenerateBackground() {
	layers := len(data.CurrBackground.Layers)
	layerD := (data.BackOffset - data.ForeOffset) / float64(layers-1)
	data.CurrBackground.Backgrounds = []*data.BackgroundLayer{}
	for i, layer := range data.CurrBackground.Layers {
		var seed = rand.Int63()
		p := perlin.NewPerlin(data.Alpha, data.Beta, data.N, seed)
		vp := viewport.New(nil)
		vp.SetRect(pixel.R(0, 0, data.BaseWidth, data.BaseHeight))
		vp.CamPos = pixel.ZV
		vp.PortPos = viewport.MainCamera.PostCamPos
		vp.Canvas.SetUniform("uRed", float32(layer.Color.R)/255.)
		vp.Canvas.SetUniform("uGreen", float32(layer.Color.G)/255.)
		vp.Canvas.SetUniform("uBlue", float32(layer.Color.B)/255.)
		vp.Canvas.SetFragmentShader(data.BGShader)
		offset := data.BackOffset - layerD*float64(i)
		if i == layers-1 {
			offset = data.ForeOffset
		}
		data.CurrBackground.Backgrounds = append(data.CurrBackground.Backgrounds, &data.BackgroundLayer{
			Layer:  layer,
			Offset: offset,
			Perlin: p,
			Color:  layer.Color,
			IMDraw: imdraw.New(nil),
			View:   vp,
		})
	}
}

func UpdateBackgrounds() {
	for _, bg := range data.CurrBackground.Backgrounds {
		bg.View.Update()
		bg.IMDraw.Clear()
		bg.IMDraw.Intensity = 1.

		bg.IMDraw.Color = bg.Color
		for x := -data.BaseWidth; x < data.BaseWidth; x++ {
			xf := float64(x)
			y := bg.Perlin.Noise1D(xf/bg.Layer.WaveLength)*bg.Layer.Scale + bg.Offset + bg.Layer.VOffset(xf)
			renderTexturedLine(xf, y, bg.IMDraw)
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
