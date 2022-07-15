package apihv1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type (
	v1Controllers struct {
		Stream *StreamController
	}

	HttpApiHandlerV1 struct {
		engine      *gin.Engine
		controllers *v1Controllers
	}
)

func NewAPIHttpV1(engine *gin.Engine) *HttpApiHandlerV1 {
	handler := HttpApiHandlerV1{
		engine: engine,
		controllers: &v1Controllers{
			Stream: NewStreamController(),
		},
	}

	handler.init()
	return &handler
}

func (h *HttpApiHandlerV1) init() {
	h.setHTML()
	h.setAPI()
}

func (h *HttpApiHandlerV1) setHTML() {
	h.engine.LoadHTMLGlob("./web/*.html")

	h.engine.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "stream.html", nil)
	})
}

func (h *HttpApiHandlerV1) setAPI() {
	api := h.engine.Group("/api/v1")

	//create stream & watch
	{
		api.GET("/ws/stream/create", h.controllers.Stream.WSCreate)
		api.GET("/ws/stream/watch", h.controllers.Stream.WSWatch)
	}
}
