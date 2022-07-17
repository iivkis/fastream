package apihv1

import (
	"github.com/gin-gonic/gin"
)

type (
	controllers struct {
		Stream *StreamController
	}

	APIHttpHandlerV1 struct {
		engine      *gin.Engine
		controllers *controllers
	}
)

func NewAPIHttpV1(engine *gin.Engine) *APIHttpHandlerV1 {
	handler := APIHttpHandlerV1{
		engine: engine,
		controllers: &controllers{
			Stream: NewStreamController(),
		},
	}

	handler.init()
	return &handler
}

func (h *APIHttpHandlerV1) init() {
	h.setAPI()
}

func (h *APIHttpHandlerV1) setAPI() {
	api := h.engine.Group("/api/v1")

	//create stream & watch
	{
		api.GET("/ws/stream/create", h.controllers.Stream.WSCreate)
		api.GET("/ws/stream/watch", h.controllers.Stream.WSWatch)
	}
}
