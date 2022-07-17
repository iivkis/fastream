package containers

import (
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

var _ content = (*HomeContainer)(nil)

type HomeContainer struct {
	BtnCloseApp *widget.Button
}

func (c *HomeContainer) GetContent() *fyne.Container {
	return container.New(
		layout.NewGridLayoutWithRows(2),
		c.BtnCloseApp,
	)
}

func (c *Containers) InitHomeContainer() {
	c.Home.BtnCloseApp = widget.NewButton("закрыть", func() {
		os.Exit(0)
	})
}
