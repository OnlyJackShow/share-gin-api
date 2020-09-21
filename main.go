package main

import (
	_ "ginswagger/docs"
	"ginswagger/route"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

func main() {
	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	route := route.InitRoute(r)
	route.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	route.Run(":2001") // listen and serve on 0.0.0.0:8080
}
