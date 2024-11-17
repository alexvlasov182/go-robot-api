package routes

import (
	"go-robot-api/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRobotRoutes(router *gin.Engine) {
	robotGroup := router.Group("/api/robots")
	{
		robotGroup.POST("/", controllers.CreateRobot)
		robotGroup.GET("/", controllers.GetAllRobots)
	}
}
