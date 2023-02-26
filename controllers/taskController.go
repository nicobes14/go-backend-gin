package controllers

import (
	"go-backend-gin/models"
	"go-backend-gin/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTaskInfo(c *gin.Context) {
	// obtener el task ID desde los params
	id := c.Param("id")
	// obtener los datos desde la base de datos
	result, task := services.GetTask(id)

	if result.Error != nil {
		c.Status(404222)
		return
	}
	// enviar los datos al cliente
	c.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "Tasks ",
		"task":    task,
	})
}

func GetTasks(c *gin.Context) {
	// obtener toda las tareas de la bd
	result, tasks := services.GetTasks()

	// si hay error
	if result.Error != nil {
		c.Status(400)
		return
	}
	// normal
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

	var body models.CreateTask
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// crear la tarea
	result, task := services.CreateTask(models.Task{Title: body.Title, Description: body.Description})

	if result.Error != nil {
		c.JSON(400, gin.H{"error": result.Error})
		return
	}

	// enviar status final
	c.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "Task created",
		"task":    task,
	})
}
