package distributor

import (
	"github.com/gin-gonic/gin"
)

type DistributorSPAConfig struct {
	PathToIndexFile string
	PathToStaticDir string
}

type DistributorSPA struct {
	engine *gin.Engine
	config DistributorSPAConfig
}

func NewDistributorSPA(dconfig DistributorSPAConfig) *DistributorSPA {
	d := &DistributorSPA{config: dconfig}

	d.engine = NewEngine()
	d.init()

	return d
}

func (d *DistributorSPA) init() {
	d.engine.LoadHTMLGlob(d.config.PathToIndexFile)
	d.engine.Static("/assets/", d.config.PathToStaticDir)

	d.engine.NoRoute(func(ctx *gin.Context) {
		ctx.File(d.config.PathToIndexFile)
	})
}

func (d *DistributorSPA) Run(addr string) error {
	return d.engine.Run(addr)
}
