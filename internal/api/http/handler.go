package httph

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	apictrlv1 "github.com/iivkis/fastream/internal/api/http/v1"
)

type HTTPHandler struct {
	engine *gin.Engine
}

func NewHTTPHandler() *HTTPHandler {
	hadnler := HTTPHandler{engine: gin.Default()}
	hadnler.init()
	return &hadnler
}

func (h *HTTPHandler) init() {
	h.engine.Use(cors.Default())
	apictrlv1.SetControllers(h.engine)
}

func (h *HTTPHandler) Run(addr string) error {
	return h.engine.Run(addr)
}
