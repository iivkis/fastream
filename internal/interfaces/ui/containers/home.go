package containers

import (
	"fmt"
	"os/exec"
	"runtime"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/iivkis/fastream/internal/config"
)

var _ content = (*HomeContainer)(nil)

type HomeContainer struct {
	BtnStartStream *widget.Button
	BtnCloseApp    *widget.Button
}

func (c *HomeContainer) GetContent() *fyne.Container {
	return container.New(
		layout.NewVBoxLayout(),

		container.New(
			layout.NewGridLayoutWithRows(4),
			c.BtnStartStream,
			c.BtnCloseApp,
		),
	)
}

func (c *Containers) InitHomeContainer() {
	webUI := fmt.Sprintf("http://%s:%s", "localhost", config.Env.SPA_PORT)

	c.Home.BtnStartStream = widget.NewButton("начать стрим", func() {
		switch runtime.GOOS {
		case "linux":
			exec.Command("xdg-open", webUI).Start()
		case "windows":
			exec.Command("explorer", webUI).Start()
		default:
			dialog.ShowError(fmt.Errorf("платформа не поддерживается"), c.win)
		}
	})

	c.Home.BtnCloseApp = widget.NewButton("остановить и закрыть", func() {
		c.win.Close()
	})
}
