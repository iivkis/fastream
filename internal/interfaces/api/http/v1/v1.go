package restfulv1

import (
	"github.com/gin-gonic/gin"
	servicev1 "github.com/iivkis/fastream/internal/service/v1"
)

type Controllers struct {
	engine *gin.Engine

	Stream *streamController
	Utils  *utilsController
}

func SetControllers(engine *gin.Engine, service *servicev1.Service) {
	handler := Controllers{
		Stream: newStreamController(service),
		Utils:  newUtilsController(),
	}

	handler.engine = engine
	handler.init()
}

func (h *Controllers) init() {
	h.setAPI()
}

func (h *Controllers) setAPI() {
	api := h.engine.Group("/api/v1")

	//create stream & watch
	{
		api.GET("/ws/stream/create", h.Stream.Create)
		api.GET("/ws/stream/watch", h.Stream.Watch)
	}

	//utils
	{
		api.GET("/utils/local_ip", h.Utils.GetLocalIP)
	}
}
