package routes

import (
	"customer-management/controllers"
	"github.com/gin-gonic/gin"
)

func ConfigurarRotas(router *gin.Engine) {
	clientGroup := router.Group("/clients")
	{
		clientGroup.POST("/", controllers.CreateClient)
		clientGroup.GET("/", controllers.ListClient)
		clientGroup.GET("/:id", controllers.SearchClientByID)
		clientGroup.PUT("/:id", controllers.UpdateClient)
		clientGroup.DELETE("/:id", controllers.DeleteClient)
	}
}
