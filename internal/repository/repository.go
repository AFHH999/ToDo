package repository

import "github.com/AFHH999/ToDo/internal/models"

// TaskRepository defines the interface for database operations.
type TaskRepository interface {
	Create(task *models.Task) error
	GetAll() ([]models.Task, error)
	GetByID(id int) (*models.Task, error)
	Update(task *models.Task) error
	Delete(task *models.Task) error
}
