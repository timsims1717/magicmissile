package object

import (
	"github.com/faiface/pixel"
	"github.com/google/uuid"
)

type Object struct {
	ID   uuid.UUID
	Hide bool
	Gone bool

	Pos     pixel.Vec
	Rect    pixel.Rect
	Mat     pixel.Matrix
	Offset  pixel.Vec
	APos    pixel.Vec
	LastPos pixel.Vec
	Rot     float64
	RotArnd pixel.Vec
	Scalar  pixel.Vec
	Flip    bool
	Flop    bool
}

func New() *Object {
	return &Object{
		ID:     uuid.New(),
		Scalar: pixel.Vec{
			X: 1.,
			Y: 1.,
		},
	}
}
