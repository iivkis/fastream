package containers

import "fyne.io/fyne/v2"

type content interface {
	GetContent() *fyne.Container
}

type Containers struct {
	win  fyne.Window
	Home HomeContainer
}

func newContainers(win fyne.Window) *Containers {
	return &Containers{win: win}
}

func (c *Containers) init() {
	c.InitHomeContainer()
	c.win.SetContent(c.Home.GetContent())
}

func Setup(win fyne.Window) {
	c := newContainers(win)

	win.SetFixedSize(true)
	win.CenterOnScreen()

	c.init()
}
