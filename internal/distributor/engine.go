package distributor

import "github.com/gin-gonic/gin"

func NewEngine() *gin.Engine {
	engine := gin.Default()
	return engine
}
