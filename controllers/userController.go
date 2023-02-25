package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUserInfo(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "aaaa",
	})
}

func GetUsers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "aaaa",
	})
}

func DeleteUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "aaaa",
	})
}

func UpdateUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "aaaa",
	})
}
