package state

import (
	"fmt"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
	"image/color"
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
	LoadPrc  float64
	ShowLoad bool
}

func New(state State) *AbstractState {
	aState := &AbstractState{
		State: state,
	}
	state.SetAbstract(aState)
	return aState
}

var (
	switchState = false
	currState   = "unknown"
	nextState   = "unknown"

	loading = false
	done    = make(chan struct{})

	states     = map[string]*AbstractState{}
	clearColor color.Color
)

func init() {
	clearColor = colornames.Black
}

func Register(key string, state *AbstractState) {
	if _, ok := states[key]; ok {
		fmt.Printf("error: state '%s' already registered", key)
	} else {
		states[key] = state
	}
}

func SetClearColor(col color.Color) {
	clearColor = col
}

func Update(win *pixelgl.Window) {
	updateState()
	if loading {
		select {
		case <-done:
			loading = false
			currState = nextState
		default:
			//LoadingState.Update(win)
		}
	}
	if cState, ok := states[currState]; ok {
		cState.Update(win)
	}
}

func Draw(win *pixelgl.Window) {
	win.Clear(clearColor)
	cState, ok1 := states[currState]
	nState, ok2 := states[nextState]
	if !ok2 {
		panic(fmt.Sprintf("state %s doesn't exist", nextState))
	}
	if loading && nState.ShowLoad || !ok1 {
		//LoadingState.Draw(win)
	} else {
		cState.Draw(win)
	}
}

func updateState() {
	if !loading && (currState != nextState || switchState) {
		go func() {
			// uninitialize
			if cState, ok := states[currState]; ok {
				cState.Unload()
			}
			// initialize
			if cState, ok := states[nextState]; ok {
				cState.Load()
			}
			done <- struct{}{}
		}()
		loading = true
		switchState = false
	}
}

func SwitchState(s string) {
	if !switchState {
		switchState = true
		nextState = s
	}
}
