package state

import (
	"fmt"
	"github.com/faiface/pixel/pixelgl"
)

type State interface {
	Unload()
	Load()
	Update(*pixelgl.Window)
	Draw(*pixelgl.Window)
	SetAbstract(*AbstractState)
}

type AbstractState struct {
	State
}

func New(state State) *AbstractState {
	aState := &AbstractState{
		State:   state,
	}
	state.SetAbstract(aState)
	return aState
}

var (
	switchState = false
	currState   = "unknown"
	nextState   = "unknown"

	States = map[string]*AbstractState{}
)

func Register(key string, state *AbstractState) {
	if _, ok := States[key]; ok {
		fmt.Printf("error: state '%s' already registered", key)
	} else {
		States[key] = state
	}
}

func Update(win *pixelgl.Window) {
	updateState()
	if cState, ok := States[currState]; ok {
		cState.Update(win)
	}
}

func Draw(win *pixelgl.Window) {
	if cState, ok := States[currState]; ok {
		cState.Draw(win)
	}
}

func updateState() {
	if currState != nextState || switchState {
		// uninitialize
		if cState, ok := States[currState]; ok {
			cState.Unload()
		}
		// initialize
		if cState, ok := States[nextState]; ok {
			cState.Load()
		}
		currState = nextState
		switchState = false
	}
}

func SwitchState(s string) {
	if !switchState {
		switchState = true
		nextState = s
	}
}
