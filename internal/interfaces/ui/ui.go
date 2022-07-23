package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/iivkis/fastream/internal/interfaces/ui/containers"
)

type UI struct {
	app fyne.App
	win fyne.Window
}

func NewUI() *UI {
	ui := new(UI)
	ui.init()
	return ui
}

func (ui *UI) init() {
	ui.app = app.New()
	ui.win = ui.app.NewWindow("Fastream")
	containers.Add(ui.win)
}

func (ui *UI) Run() {
	ui.win.ShowAndRun()
}
