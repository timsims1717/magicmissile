package game

import (
	"timsims1717/magicmissile/internal/data"
	"timsims1717/magicmissile/pkg/timing"
)

var (
	PCs     []*data.PC
	MTimer  *timing.Timer
	Towns   []*data.Town
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