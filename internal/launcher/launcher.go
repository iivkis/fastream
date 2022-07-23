package launcher

import (
	"fmt"

	"github.com/iivkis/fastream/internal/config"
	"github.com/iivkis/fastream/internal/distributor"
	httph "github.com/iivkis/fastream/internal/interfaces/api/http"
	"github.com/iivkis/fastream/internal/interfaces/ui"
)

func Launch() {
	var (
		addrAPI = fmt.Sprintf("%s:%s", config.Env.APP_HOST, config.Env.API_PORT)
		addrSPA = fmt.Sprintf("%s:%s", config.Env.APP_HOST, config.Env.SPA_PORT)
	)

	//run API
	go func() {
		HTTPHandler := httph.NewHTTPHandler()

		if err := HTTPHandler.Run(addrAPI); err != nil {
			panic(err)
		}
	}()

	//run SPA
	go func() {
		HTTPDist := distributor.NewDistributorSPA(distributor.DistributorSPAConfig{
			PathToIndexFile: "./web/dist/index.html",
			PathToStaticDir: "./web/dist/assets",
		})

		if err := HTTPDist.Run(addrSPA); err != nil {
			panic(err)
		}
	}()

	//run UI
	appUI := ui.NewUI()
	appUI.Run()
}
