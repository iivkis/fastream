package containers

import "fyne.io/fyne/v2"

type content interface {
	GetContent() *fyne.Container
}

type Containers struct {
	win  fyne.Window
	Home HomeContainer
}

func Setup(win fyne.Window) {
	c := &Containers{win: win}

	win.SetFixedSize(true)
	win.CenterOnScreen()

	c.init()
}

func (c *Containers) init() {
	c.InitHomeContainer()
	c.win.SetContent(c.Home.GetContent())
}
