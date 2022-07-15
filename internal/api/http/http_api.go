package httpapi

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	apihv1 "github.com/iivkis/fastream/internal/api/http/v1"
)

type APIHttp struct {
	engine *gin.Engine
}

func NewAPIHttp() *APIHttp {
	api := APIHttp{}
	api.engine = gin.Default()

	api.init()
	return &api
}

func (api *APIHttp) init() {
	api.engine.Use(cors.Default())

	apihv1.NewAPIHttpV1(api.engine)
}

func (api *APIHttp) Run(addr string) error {
	return api.engine.Run(addr)
}
