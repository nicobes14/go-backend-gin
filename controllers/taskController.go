package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTaskInfo(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "aaaa",
	})
}

func GetTasks(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "aaaa",
	})
}

func DeleteTask(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "aaaa",
	})
}

func UpdateTask(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "aaaa",
	})
}
