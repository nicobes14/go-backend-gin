package services

import (
	"go-backend-gin/initializers"
	"go-backend-gin/models"

	"gorm.io/gorm"
)

func GetTasks() (*gorm.DB, []models.Task) {
	var tasks []models.Task
	return initializers.DB.Find(&tasks), tasks
}

func CreateTask(task models.Task) (*gorm.DB, models.Task) {
	return initializers.DB.Create(&task), task
}
