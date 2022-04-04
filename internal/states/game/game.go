package game

import (
	"github.com/faiface/pixel"
	"timsims1717/magicmissile/internal/data"
	"timsims1717/magicmissile/pkg/timing"
	"timsims1717/magicmissile/pkg/typeface"
)

var (
	Title    *typeface.Text
	PCs      []*data.PC
	Selected int
	Cursor   pixel.Vec
	Frame    = pixel.R(-725., -185., 725., 425.)
	TownYLvl = -260.
	TownX    = 1200.
	MoveYLvl = -280.
	CharYLvl = -295.
	DeadYLvl = -315.
	MTimer   *timing.Timer
	DTimer   *timing.Timer
	ZTimer   *timing.Timer
	Timer    *timing.Timer
	TimeText *typeface.Text
	MFreq    float64
	MSpd     float64
	BigOn    bool
	ZFreq    float64
	Level    int
	Towns    []*data.Town
	GameOver = false
	OverText *typeface.Text
	WizText  *typeface.Text
	ZSoundT  *timing.Timer
	ZCount   int
	ThunT    *timing.Timer
	MsgTimer *timing.Timer

	StartMsg = []string{
		"You have ANGERED the GODS",
		"Your DOOM is INEVITABLE",
		"This town is DOOMED",
	}
	Msg1 = []string{
		"Your DOOM is INEVITABLE",
		"Your DOOM is INEVITABLE",
		"Your DOOM is INEVITABLE",
		"Your DOOM is INEVITABLE",
		"Your DOOM is INEVITABLE",
		"Feel the WRATH of the GODS",
		"Feel the WRATH of the GODS",
		"Your PUNY spells are USELESS",
		"My MINIONS are UNSTOPPABLE",
	}
	Msg2 = []string{
		"Your, uh, DOOM, is INEVITABLE.",
		"Ahem.",
		"Your DOOM. It's coming to get you!",
		"Your doom is inevitable. Really.",
		"Hey, do we have anything bigger?",
		"Here it comes, I can feel it!",
	}
	GameOver1 = []string{
		"GAME OVER",
		"INEVITABLE",
		"DOOM HAS COME",
	}
	GameOver2 = []string{
		"Oh my SELF that took forever.",
		"You know, they really impressed me!",
		"Inevitable. Whatever.",
		"MOOD HAS COME",
	}
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
	GameOver = townsGone || pcsGone
}