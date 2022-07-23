package restfulv1

import (
	"github.com/gin-gonic/gin"
)

type Controllers struct {
	engine *gin.Engine

	Stream *StreamController
	Utils  *UtilsController
}

func SetControllers(engine *gin.Engine) {
	handler := Controllers{
		Stream: NewStreamController(),
		Utils:  NewUtilsController(),
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
		api.GET("/ws/stream/create", h.Stream.WSCreate)
		api.GET("/ws/stream/watch", h.Stream.WSWatch)
	}

	//utils
	{
		api.GET("/utils/local_ip", h.Utils.GetLocalIP)
	}
}
