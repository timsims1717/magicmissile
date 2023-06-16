package states

import (
	"fmt"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"math/rand"
	"timsims1717/magicmissile/_archived"
	"timsims1717/magicmissile/internal/data"
	"timsims1717/magicmissile/internal/myecs"
	"timsims1717/magicmissile/internal/states/game"
	"timsims1717/magicmissile/internal/systems"
	"timsims1717/magicmissile/pkg/img"
	"timsims1717/magicmissile/pkg/sfx"
	"timsims1717/magicmissile/pkg/state"
	"timsims1717/magicmissile/pkg/timing"
	"timsims1717/magicmissile/pkg/typeface"
	"timsims1717/magicmissile/pkg/viewport"
)

var OldGameState = &oldGameState{}

type oldGameState struct {
	*state.AbstractState
}

func (s *oldGameState) Unload() {
	systems.ClearSystem()
}

func (s *oldGameState) Load() {
	game.PCs = []*data.PC{}
	game.GameOver = false
	loadUI()
	loadScenery()
	loadTowns()
	loadWizard()
	loadFighter()
	game.MTimer = timing.New(5.)
	game.DTimer = timing.New(15.)
	game.ZTimer = timing.New(8.)
	game.MFreq = 5.
	game.MSpd = 50.
	game.ZFreq = 10.
	game.Level = 0
	game.BigOn = false
	game.Timer = timing.New(0.)
	game.TimeText = typeface.New("main", typeface.NewAlign(typeface.Center, typeface.Center), 1.0, 0.15, 0., 0.)
	game.TimeText.SetPos(pixel.V(0., game.Frame.Max.Y))
	myecs.Manager.NewEntity().AddComponent(myecs.Object, game.TimeText.Obj)
	game.OverText = typeface.New("main", typeface.NewAlign(typeface.Center, typeface.Center), 1.0, 0.25, 0., 0.)
	game.OverText.SetPos(pixel.V(0., 220.))
	game.OverText.SetText(game.StartMsg[rand.Intn(len(game.StartMsg))])
	game.OverText.SetColor(pixel.Alpha(0.))
	game.MsgTimer = timing.New(10.)
	myecs.Manager.NewEntity().AddComponent(myecs.Object, game.OverText.Obj)
	game.WizText = typeface.New("main", typeface.NewAlign(typeface.Center, typeface.Center), 1.0, 0.08, 0., 0.)
	game.WizText.SetPos(pixel.V(0., game.CharYLvl-60.))
	myecs.Manager.NewEntity().AddComponent(myecs.Object, game.WizText.Obj)
}

func (s *oldGameState) Update(win *pixelgl.Window) {
	data.TheInput.Update(win, viewport.MainCamera.Mat)
	if data.TheInput.Get("killAll").JustPressed() {
		for _, town := range game.Towns {
			town.Health.Dead = true
		}
		for _, pc := range game.PCs {
			pc.Char.Health.Dead = true
		}
	}
	if data.TheInput.Get("menuBack").JustPressed() {
		data.TheInput.Get("menuBack").Consume()
		PauseMenu.Open()
	}
	if !PauseMenu.Opened {
		game.Cursor = data.TheInput.World
		if game.Cursor.X < game.Frame.Min.X {
			game.Cursor.X = game.Frame.Min.X
		} else if game.Cursor.X > game.Frame.Max.X {
			game.Cursor.X = game.Frame.Max.X
		}
		if game.Cursor.Y < game.Frame.Min.Y {
			game.Cursor.Y = game.Frame.Min.Y
		} else if game.Cursor.Y > game.Frame.Max.Y {
			game.Cursor.Y = game.Frame.Max.Y
		}
		game.CheckGameOver()
		systems.TemporarySystem()
		systems.FunctionSystem()
		if !game.GameOver {
			systems.ControlSystem()
		}
		systems.PayloadSystem()
		systems.AttackSystem()
		systems.MobSystem()
		systems.HealthSystem()
		systems.AnimationSystem()
		if game.MTimer.UpdateDone() {
			if game.BigOn && rand.Intn(5) == 0 {
				_archived.BigMeteor(game.MSpd * 0.75)
			} else {
				_archived.BasicMeteor(game.MSpd, pixel.ZV)
			}
			game.MTimer = timing.New(rand.Float64()*game.MFreq + 0.5)
		}
		if game.ZTimer.UpdateDone() {
			_archived.BasicZombie()
			game.ZTimer = timing.New(rand.Float64()*game.ZFreq + 0.5)
		}
		if game.DTimer.UpdateDone() {
			game.Level++
			game.ZFreq -= 0.1
			if game.ZFreq < 2. {
				game.ZFreq = 2.
			}
			game.MFreq -= 0.06
			if game.MFreq < 3. {
				game.MFreq = 3.
			}
			game.MSpd += 4.
			if game.MSpd > 120. {
				game.MSpd = 120.
			}
			if game.Level > 8 {
				game.BigOn = true
			}
			game.DTimer = timing.New(10.)
		}
		if game.ZSoundT.UpdateDone() {
			game.ZSoundT = timing.New(4. + rand.Float64()*8.)
			if game.ZCount > 0 {
				sfx.SoundPlayer.PlaySound("zombie", 0.)
			}
		}
		if game.ThunT.UpdateDone() {
			game.ThunT = timing.New(40. + rand.Float64()*240.)
			sfx.SoundPlayer.PlaySound(fmt.Sprintf("thunder%d", rand.Intn(2)+1), 0.)
		}
		if game.Timer != nil && !game.GameOver {
			game.Timer.Update()
			elapsed := int(game.Timer.Elapsed())
			h := elapsed / 3600
			m := elapsed / 60 % 60
			sec := elapsed % 60
			if h > 0 {
				game.TimeText.SetText(fmt.Sprintf("%d:%02d:%02d", h, m, sec))
			} else {
				game.TimeText.SetText(fmt.Sprintf("%02d:%02d", m, sec))
			}
			if game.MsgTimer.UpdateDone() {
				if game.OverText.NoShow {
					if elapsed > 300. {
						game.OverText.SetText(game.Msg2[rand.Intn(len(game.Msg2))])
					} else {
						game.OverText.SetText(game.Msg1[rand.Intn(len(game.Msg1))])
					}
					game.OverText.NoShow = false
					game.MsgTimer = timing.New(10.)
				} else {
					game.MsgTimer = timing.New(50.)
					game.OverText.NoShow = true
				}
			}
		}
		if game.GameOver && !GameOverMenu.Opened {
			if game.Timer.Elapsed() > 300. {
				game.OverText.SetText(game.GameOver2[rand.Intn(len(game.GameOver2))])
			} else {
				game.OverText.SetText(game.GameOver1[rand.Intn(len(game.GameOver1))])
			}
			GameOverMenu.Open()
			game.OverText.NoShow = false
		}
	}
	UpdateMenus(data.TheInput)
}

func (s *oldGameState) Draw(win *pixelgl.Window) {
	img.Clear()
	systems.DrawSystem(win, 0)
	img.Draw(win)
	game.TimeText.Draw(win)
	game.OverText.Draw(win)
	game.WizText.Draw(win)
	DrawMenus(win)
}

func (s *oldGameState) SetAbstract(aState *state.AbstractState) {
	s.AbstractState = aState
}
