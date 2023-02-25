package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Title       string `gorm:"unique"`
	Description string
	Done        bool `gorm:"default:false"`
}
