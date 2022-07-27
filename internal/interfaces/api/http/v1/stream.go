package restful

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/iivkis/fastream/internal/service/v1"
)

type streamController struct {
	service  *service.Service
	upgrader *websocket.Upgrader
}

func newStreamController(service *service.Service) *streamController {
	upgrader := &websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	return &streamController{
		service:  service,
		upgrader: upgrader,
	}
}

func (c *streamController) Create(ctx *gin.Context) {
	conn, err := c.upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	c.service.Stream.Create(ctx.Request.Context(), newStreamConn(conn))
}

func (c *streamController) Watch(ctx *gin.Context) {
	conn, err := c.upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, NewResponse(nil, err))
		return
	}
	defer conn.Close()

	c.service.Stream.Watch(ctx.Request.Context(), newStreamConn(conn))
}
