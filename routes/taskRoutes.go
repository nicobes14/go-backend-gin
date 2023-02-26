package routes

import (
	"go-backend-gin/controllers"

	"github.com/gin-gonic/gin"
)

func TasksRoutes(router *gin.Engine) {
	router.POST("/tasks", controllers.CreateTask)
	router.GET("/tasks", controllers.GetTasks)
	router.GET("/tasks/:id", controllers.GetTaskInfo)
	router.DELETE("/tasks/:id", controllers.DeleteTask)
}
