package containers

import "fyne.io/fyne/v2"

type content interface {
	GetContent() *fyne.Container
}

type Containers struct {
	win  fyne.Window
	Home HomeContainer
}

func Add(win fyne.Window) {
	c := &Containers{win: win}

	win.Resize(fyne.NewSize(300, 400))
	win.SetFixedSize(true)
	win.CenterOnScreen()

	c.init()
}

func (c *Containers) init() {
	c.InitHomeContainer()
	c.win.SetContent(c.Home.GetContent())
}