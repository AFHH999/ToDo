package services

import (
	"github.com/AFHH999/ToDo/internal/models"
	"gorm.io/gorm"
)

// CreateTask saves a new task to the database
func CreateTask(db *gorm.DB, task models.Task) (uint, error) {
	result := db.Create(&task)
	return task.ID, result.Error
}

// ListTasks returns all tasks from the database
func ListTasks(db *gorm.DB) ([]models.Task, error) {
	var tasks []models.Task
	result := db.Find(&tasks)
	return tasks, result.Error
}

// GetTaskByID finds a single task
func GetTaskByID(db *gorm.DB, id int) (models.Task, error) {
	var task models.Task
	result := db.First(&task, id)
	return task, result.Error
}

// UpdateTask saves changes to an existing task
func UpdateTask(db *gorm.DB, task models.Task) error {
	return db.Save(&task).Error
}

// DeleteTask removes a task by ID
func DeleteTask(db *gorm.DB, id int) error {
	var task models.Task
	if err := db.First(&task, id).Error; err != nil {
		return err
	}
	return db.Delete(&task).Error
}
