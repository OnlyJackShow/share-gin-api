package controller

import (
	"github.com/gin-gonic/gin"
)

// @Description 方法描述
// @Accept  json
// @Produce  json
// @Param   some_id     path    string  false     "Some ID"
// @Param   offset     query    int     true        "Offset"
// @Param   limit      query    int     true        "Offset"
// @Success 200 {string} string	"ok"
// @Failure 400 {string} string "We need ID!!"
// @Failure 404 {string} string "Can not find ID"
// @Router /testapi/get [get]
func World(g *gin.Context) {
	g.JSON(200, "hello")
}
