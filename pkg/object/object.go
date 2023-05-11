package object

import (
	"fmt"
	"github.com/faiface/pixel"
	"golang.org/x/image/colornames"
	"image/color"
)

var objIndex = uint32(0)

type Object struct {
	ID   string
	Hide bool
	Load bool
	Keep bool
	Kill bool

	Pos  pixel.Vec
	Mat  pixel.Matrix
	Rot  float64
	Sca  pixel.Vec
	Flip bool
	Flop bool
	Rect pixel.Rect

	PostPos pixel.Vec
	LastPos pixel.Vec
	Offset  pixel.Vec

	Mask  color.RGBA
	Layer int

	ILock bool
}

func New() *Object {
	return &Object{
		Sca: pixel.Vec{
			X: 1.,
			Y: 1.,
		},
		Mask: colornames.White,
	}
}

func (obj *Object) WithID(code string) *Object {
	obj.ID = fmt.Sprintf("%s-%d", code, objIndex)
	objIndex++
	return obj
}

func (obj *Object) PointInside(vec pixel.Vec) bool {
	return obj.Rect.Moved(obj.Pos).Moved(pixel.V(-(obj.Rect.W() * 0.5), -(obj.Rect.H() * 0.5))).Contains(vec)
}
