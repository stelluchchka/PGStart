package api

import (
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/")
	{
		api.POST("/commands", CreateCommand)
		api.GET("/commands", GetCommands)
		api.GET("/commands/:id", GetCommand)
	}
}
