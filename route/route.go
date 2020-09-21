package route

import (
	"ginswagger/controller"
	"github.com/gin-gonic/gin"
)

func InitRoute(engine *gin.Engine)*gin.Engine {
	route := engine.Group("/hello")
	{
		route.GET("/world",controller.World)
	}
	return engine
}
