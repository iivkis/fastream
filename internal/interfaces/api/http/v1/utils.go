package restful

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iivkis/fastream/internal/service/v1"
)

type utilsController struct {
	service *service.Service
}

func newUtilsController(service *service.Service) *utilsController {
	return &utilsController{
		service: service,
	}
}

type utilsGetLocalIPResponse struct {
	LocalIP string `json:"local_ip"`
}

func newUtilsGetLocalIPResponse(localIP string) *utilsGetLocalIPResponse {
	return &utilsGetLocalIPResponse{
		LocalIP: localIP,
	}
}

func (c *utilsController) GetLocalIP(ctx *gin.Context) {
	localIP, err := c.service.Utils.GetLocalIP()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, NewResponse(nil, err))
		return
	}

	resp := newUtilsGetLocalIPResponse(localIP)

	ctx.JSON(http.StatusOK, NewResponse(resp, nil))
}
