package controllers

import (
	"go-backend-gin/initializers"
	"go-backend-gin/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTaskInfo(c *gin.Context) {
	// obtener el task ID desde los params

	// obtener los datos desde la base de datos

	// enviar los datos al cliente
	c.JSON(http.StatusOK, gin.H{
		"message": "aaaa",
	})
}

func GetTasks(c *gin.Context) {
	// obtener toda las tareas de la bd
	var tasks []models.Task
	result := initializers.DB.Find(&tasks)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "Tasks ",
		"task":    tasks,
	})
}

func DeleteTask(c *gin.Context) {
	// obtener el task ID desde los params

	// borrar la tarea de la bd

	// responder el status final
	c.JSON(http.StatusOK, gin.H{
		"message": "aaaa",
	})
}

func UpdateTask(c *gin.Context) {
	// obtener el task ID desde los params

	// actualizar los datos

	// enviar status final
	c.JSON(http.StatusOK, gin.H{
		"message": "aaaa",
	})
}

func CreateTask(c *gin.Context) {
	// obtener la task desde el body
	var body struct {
		Title       string `form:"title" json:"title" xml:"title"  binding:"required"`
		Description string `form:"description" json:"description" xml:"description"  binding:"required"`
	}
	// verificamos el body
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// crear la tarea
	task := models.Task{Title: body.Title, Description: body.Description}
	result := initializers.DB.Create(&task)

	if result.Error != nil {
		c.Status(400)
		return
	}

	// enviar status final
	c.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "Task created",
		"task":    task,
	})
}
