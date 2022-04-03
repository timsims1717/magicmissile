package data

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"timsims1717/magicmissile/pkg/timing"
)

type TimerFunc struct {
	Timer *timing.Timer
	Func  func() bool
}

func NewTimerFunc(fn func() bool, dur float64) *TimerFunc {
	return &TimerFunc{
		Timer: timing.New(dur),
		Func:  fn,
	}
}

type FrameFunc struct {
	Func func() bool
}

func NewFrameFunc(fn func() bool) *FrameFunc {
	return &FrameFunc{Func: fn}
}

type ImdFunc struct {
	Key  string
	Func func(pixel.Vec, *imdraw.IMDraw)
}

func NewImdFunc(key string, fn func(pixel.Vec, *imdraw.IMDraw)) *ImdFunc {
	return &ImdFunc{
		Key: key,
		Func: fn,
	}
}