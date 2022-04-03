package state

import (
	"fmt"
	"github.com/faiface/pixel/pixelgl"
)

type State interface {
	Unload()
	Load(chan struct{})
	Update(*pixelgl.Window)
	Draw(*pixelgl.Window)
	SetAbstract(*AbstractState)
}

type AbstractState struct {
	State
	LoadPrc float64
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
	loading     = false
	loadingDone = false
	done        = make(chan struct{})

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
	if loading {
		select{
		case <-done:
			loading = false
			loadingDone = true
		default:
			if LoadingScreen.Update != nil {
				LoadingScreen.Update(win)
			}
		}
	} else {
		if cState, ok := States[currState]; ok {
			cState.Update(win)
		}
	}
}

func Draw(win *pixelgl.Window) {
	if loading {
		if LoadingScreen.Draw != nil {
			LoadingScreen.Draw(win)
		}
	} else if !loadingDone {
		if cState, ok := States[currState]; ok {
			cState.Draw(win)
		}
	} else {
		loadingDone = false
	}
}

func updateState() {
	if !loading && (currState != nextState || switchState) {
		// uninitialize
		if cState, ok := States[currState]; ok {
			go cState.Unload()
		}
		// initialize
		if cState, ok := States[nextState]; ok {
			go cState.Load(done)
			loading = true
			loadingDone = false
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
