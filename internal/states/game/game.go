package game

import (
	"github.com/faiface/pixel"
	"timsims1717/magicmissile/internal/data"
	"timsims1717/magicmissile/pkg/timing"
)

var (
	PCs      []*data.PC
	Selected int
	Cursor   pixel.Vec
	Frame    = pixel.R(-725., -190., 725., 425.)
	TownYLvl = -260.
	TownX    = 1200.
	MoveYLvl = -280.
	CharYLvl = -295.
	DeadYLvl = -315.
	MTimer   *timing.Timer
	DTimer   *timing.Timer
	ZTimer   *timing.Timer
	MFreq    float64
	MSpd     float64
	BigOn    bool
	ZFreq    float64
	Level    int
	Towns    []*data.Town
	GameOver = false
)

func CheckGameOver() {
	pcsGone := true
	for _, pc := range PCs {
		if !pc.Char.Health.Dead {
			pcsGone = false
		}
	}
	townsGone := true
	for _, town := range Towns {
		if !town.Health.Dead {
			townsGone = false
		}
	}
	if townsGone || pcsGone {
		GameOver = true
	}
}