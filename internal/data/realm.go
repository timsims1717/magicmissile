package data

import (
	"github.com/aquilax/go-perlin"
	"github.com/bytearena/ecs"
	"github.com/faiface/pixel/imdraw"
	"image/color"
	"timsims1717/magicmissile/pkg/viewport"
)

type Realm struct {
	Name        string        `json:"name"`
	Code        string        `json:"code"`
	Backgrounds []*Background `json:"backgrounds"`
}

type Background struct {
	Layers      []*LayerGenerator  `json:"layers"`
	BackCol     color.RGBA         `json:"sky"`
	Backgrounds []*BackgroundLayer `json:"-"`
}

type LayerGenerator struct {
	Scale      float64               `json:"scale"`
	WaveLength float64               `json:"waveLength"`
	VOffset    func(float64) float64 `json:"-"`
	Color      color.RGBA            `json:"color"`
	VFnCode    VFnCode               `json:"vCode"`
	Sprites    []*BackgroundSprite   `json:"sprites"`
}

type BackgroundLayer struct {
	Layer   *LayerGenerator
	Offset  float64
	Perlin  *perlin.Perlin
	Color   color.Color
	View    *viewport.ViewPort
	IMDraw  *imdraw.IMDraw
	Sprites []*ecs.Entity
}

type BackgroundSprite struct {
	Key  string `json:"key"`
	Freq int    `json:"freq"`
}

var (
	ForeOffset = -BaseHeight * 0.4
	BackOffset = BaseHeight * 0.35
	Alpha      = 2.
	Beta       = 2.
	N          = int32(3)

	CurrBackground *Background
	AllRealms      map[string]*Realm

	BGShader string
)
