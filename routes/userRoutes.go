package routes

import (
	"go-backend-gin/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	router.POST("/users", controllers.CreateTask)
	router.GET("/users", controllers.GetTasks)
	router.GET("/users/:id", controllers.GetTaskInfo)
	router.DELETE("/users/:id", controllers.DeleteTask)
}
