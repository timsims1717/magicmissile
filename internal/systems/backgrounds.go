package systems

import (
	"github.com/aquilax/go-perlin"
	"github.com/bytearena/ecs"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"math/rand"
	"timsims1717/magicmissile/internal/data"
	"timsims1717/magicmissile/internal/myecs"
	"timsims1717/magicmissile/pkg/img"
	"timsims1717/magicmissile/pkg/object"
	"timsims1717/magicmissile/pkg/viewport"
)

func GenerateRandomBackground(code string) {
	DisposeBackground()
	realm := data.AllRealms[code]
	data.CurrBackground = realm.Backgrounds[rand.Intn(len(realm.Backgrounds))]
	GenerateBackground()
}

func GenerateBackground() {
	layers := len(data.CurrBackground.Layers)
	data.TownLayer = layers - 1
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
		var sprs []*ecs.Entity
		for _, sprite := range layer.Sprites {
			r := sprite.Freq / 8
			count := sprite.Freq
			if r > 0 {
				count += rand.Intn(r) - r/2
			}
			for j := 0; j < count; j++ {
				x := rand.Float64()*data.BaseWidth - data.BaseWidth*0.5
				y := p.Noise1D(x/layer.WaveLength)*layer.Scale + offset + layer.VOffset(x)
				spr := img.NewSprite(sprite.Key, data.ObjectKey)
				layerSca := float64(i) / float64(layers)
				obj := object.New()
				obj.Pos = pixel.V(x, y)
				obj.Layer = i
				obj.Offset.Y += img.Batchers[data.ObjectKey].Sprites[sprite.Key].Frame().H() * rand.Float64() * 0.5 * layerSca
				obj.Flip = rand.Intn(2) == 0
				obj.Sca.X *= rand.Float64()*0.5*layerSca + 0.7
				obj.Sca.Y *= rand.Float64()*0.5*layerSca + 0.7
				e := myecs.Manager.NewEntity().
					AddComponent(myecs.Object, obj).
					AddComponent(myecs.Drawable, spr)
				sprs = append(sprs, e)
			}
		}
		data.CurrBackground.Backgrounds = append(data.CurrBackground.Backgrounds, &data.BackgroundLayer{
			Layer:   layer,
			Offset:  offset,
			Perlin:  p,
			Color:   layer.Color,
			IMDraw:  imdraw.New(nil),
			View:    vp,
			Sprites: sprs,
		})
	}
}

func DisposeBackground() {
	if data.CurrBackground != nil {
		for _, bg := range data.CurrBackground.Backgrounds {
			for _, sprite := range bg.Sprites {
				myecs.Manager.DisposeEntity(sprite)
			}
			bg.IMDraw = nil
			bg.Perlin = nil
		}
		data.CurrBackground.Backgrounds = []*data.BackgroundLayer{}
	}
	data.CurrBackground = nil
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
