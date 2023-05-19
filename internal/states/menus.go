package states

import (
	"github.com/faiface/pixel/pixelgl"
	pxginput "github.com/timsims1717/pixel-go-input"
	"timsims1717/magicmissile/internal/myecs"
	"timsims1717/magicmissile/pkg/menus"
	"timsims1717/magicmissile/pkg/sfx"
	"timsims1717/magicmissile/pkg/state"
)

var (
	MainMenu     *menus.Menu
	PauseMenu    *menus.Menu
	GameOverMenu *menus.Menu
)

func InitMenus(win *pixelgl.Window) {
	InitMainMenu(win)
	InitPauseMenu(win)
	InitGameOverMenu(win)
}

func InitMainMenu(win *pixelgl.Window) {
	MainMenu = menus.New("main", nil)
	start := MainMenu.AddItem("start", "Start", false)
	quit := MainMenu.AddItem("quit", "Quit", false)
	myecs.Manager.NewEntity().AddComponent(myecs.Object, start.Text.Obj)
	myecs.Manager.NewEntity().AddComponent(myecs.Object, quit.Text.Obj)

	start.SetClickFn(func() {
		state.PushState("game")
		sfx.SoundPlayer.PlaySound("click", 0.0)
		MainMenu.Close()
	})
	quit.SetClickFn(func() {
		sfx.SoundPlayer.PlaySound("click", 0.0)
		win.SetClosed(true)
	})
}

func InitPauseMenu(win *pixelgl.Window) {
	PauseMenu = menus.New("main", nil)
	resume := PauseMenu.AddItem("resume", "Resume", false)
	restart := PauseMenu.AddItem("restart", "Restart", false)
	quit := PauseMenu.AddItem("quit", "Quit", false)
	myecs.Manager.NewEntity().AddComponent(myecs.Object, resume.Text.Obj)
	myecs.Manager.NewEntity().AddComponent(myecs.Object, restart.Text.Obj)
	myecs.Manager.NewEntity().AddComponent(myecs.Object, quit.Text.Obj)

	resume.SetClickFn(func() {
		sfx.SoundPlayer.PlaySound("click", 0.0)
		PauseMenu.Close()
	})
	restart.SetClickFn(func() {
		state.PushState("game")
		sfx.SoundPlayer.PlaySound("click", 0.0)
		PauseMenu.Close()
	})
	quit.SetClickFn(func() {
		sfx.SoundPlayer.PlaySound("click", 0.0)
		win.SetClosed(true)
	})
}

func InitGameOverMenu(win *pixelgl.Window) {
	GameOverMenu = menus.New("gameover", nil)
	restart := GameOverMenu.AddItem("restart", "Restart", false)
	quit := GameOverMenu.AddItem("quit", "Quit", false)
	myecs.Manager.NewEntity().AddComponent(myecs.Object, restart.Text.Obj)
	myecs.Manager.NewEntity().AddComponent(myecs.Object, quit.Text.Obj)

	restart.SetClickFn(func() {
		state.PushState("game")
		sfx.SoundPlayer.PlaySound("click", 0.0)
		GameOverMenu.Close()
	})
	quit.SetClickFn(func() {
		sfx.SoundPlayer.PlaySound("click", 0.0)
		win.SetClosed(true)
	})
}

func UpdateMenus(in *pxginput.Input) {
	MainMenu.Update(in)
	PauseMenu.Update(in)
	GameOverMenu.Update(in)
}

func DrawMenus(win *pixelgl.Window) {
	MainMenu.Draw(win)
	PauseMenu.Draw(win)
	GameOverMenu.Draw(win)
}
