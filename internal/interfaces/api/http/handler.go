package httph

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	restfulv1 "github.com/iivkis/fastream/internal/interfaces/api/http/v1"
	servicev1 "github.com/iivkis/fastream/internal/service/v1"
)

type HTTPHandler struct {
	engine    *gin.Engine
	serviceV1 *servicev1.Service
}

func NewHTTPHandler(serviceV1 *servicev1.Service) *HTTPHandler {
	hadnler := HTTPHandler{
		engine:    gin.Default(),
		serviceV1: serviceV1,
	}
	hadnler.init()
	return &hadnler
}

func (h *HTTPHandler) init() {
	h.engine.Use(cors.Default())
	restfulv1.SetControllers(h.engine, h.serviceV1)
}

func (h *HTTPHandler) Run(addr string) error {
	return h.engine.Run(addr)
}
