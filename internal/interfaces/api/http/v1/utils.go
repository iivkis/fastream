package restfulv1

import (
	"net"
	"net/http"

	"github.com/gin-gonic/gin"
)

type utilsController struct{}

func newUtilsController() *utilsController {
	return &utilsController{}
}

func (c *utilsController) GetLocalIP(ctx *gin.Context) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, NewResponse(nil, err))
		return
	}
	defer conn.Close()

	localIP := conn.LocalAddr().(*net.UDPAddr).IP

	ctx.JSON(http.StatusOK, NewResponse(gin.H{
		"local_ip": localIP,
	}, nil))
}
