package main

import (
	"go-backend-gin/initializers"
	"go-backend-gin/models"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectDB()
}

func main() {

	initializers.DB.AutoMigrate(&models.User{})
	initializers.DB.AutoMigrate(&models.Task{})
}
