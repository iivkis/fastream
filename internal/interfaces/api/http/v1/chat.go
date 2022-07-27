package restful

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/iivkis/fastream/internal/service/v1"
)

type chatController struct {
	service  *service.Service
	upgrader websocket.Upgrader
}

func newChatController(service *service.Service) *chatController {
	upgrader := websocket.Upgrader{
		CheckOrigin: func(_ *http.Request) bool {
			return true
		},
	}

	return &chatController{
		service:  service,
		upgrader: upgrader,
	}
}

func (c *chatController) Connect(ctx *gin.Context) {
	conn, err := c.upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	c.service.Chat.Connect(ctx.Request.Context(), newChatConnection(conn))
}
