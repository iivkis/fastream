package launcher

import (
	"fmt"

	httph "github.com/iivkis/fastream/internal/api/http"
	"github.com/iivkis/fastream/internal/config"
	"github.com/iivkis/fastream/internal/distributor"
	"github.com/iivkis/fastream/internal/ui"
)

func Launch() {
	var (
		addrAPI = fmt.Sprintf("%s:%s", config.Env.API_HOST, config.Env.API_PORT)
		addrSPA = fmt.Sprintf("%s:%s", config.Env.SPA_HOST, config.Env.SPA_PORT)
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
