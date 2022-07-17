package containers

import "fyne.io/fyne/v2"

type content interface {
	GetContent() *fyne.Container
}

type Containers struct {
	win  fyne.Window
	Home HomeContainer
}

func New(win fyne.Window) {
	c := &Containers{win: win}
	c.init()

	c.win.SetContent(c.Home.GetContent())
}

func (c *Containers) init() {
	c.InitHomeContainer()
}
