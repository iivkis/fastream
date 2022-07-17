package launcher

import (
	httpapi "github.com/iivkis/fastream/internal/api/http"
	"github.com/iivkis/fastream/internal/ui"
)

func Launch() {
	go func() {
		APIHttp := httpapi.NewAPIHttp()
		if err := APIHttp.Run(":8080"); err != nil {
			panic(err)
		}
	}()

	appUI := ui.NewUI()
	appUI.Run()
}
