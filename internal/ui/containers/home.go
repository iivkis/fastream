package containers

import (
	"fmt"
	"os/exec"
	"runtime"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

var _ content = (*HomeContainer)(nil)

type HomeContainer struct {
	BtnStartStream  *widget.Button
	BtnCopyAddrLink *widget.Button
	BtnCloseApp     *widget.Button
}

func (c *HomeContainer) GetContent() *fyne.Container {
	return container.New(
		layout.NewVBoxLayout(),

		container.New(
			layout.NewGridLayoutWithRows(4),
			c.BtnStartStream,
			c.BtnCopyAddrLink,
			c.BtnCloseApp,
		),
	)
}

func (c *Containers) InitHomeContainer() {
	c.Home.BtnStartStream = widget.NewButton("начать стрим", func() {
		url := "http://localhost:8080"

		switch runtime.GOOS {
		case "linux":
			exec.Command("xdg-open", url).Start()
		case "windows":
			exec.Command("explorer", url).Start()
		default:
			dialog.ShowError(fmt.Errorf("платформа не поддерживается"), c.win)
		}

		time.Sleep(time.Second)
		c.win.RequestFocus()
	})

	c.Home.BtnCopyAddrLink = widget.NewButton("копировать ссылку на стрим", func() {
		c.win.Clipboard().SetContent("http://localhost:8080")

		go func() {
			c.Home.BtnCopyAddrLink.Disable()
			time.Sleep(time.Second)
			c.Home.BtnCopyAddrLink.Enable()
		}()
	})

	c.Home.BtnCloseApp = widget.NewButton("остановить и закрыть", func() {
		c.win.Close()
	})
}
