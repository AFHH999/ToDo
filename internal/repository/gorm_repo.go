package repository

import (
	"github.com/AFHH999/ToDo/internal/models"
	"gorm.io/gorm"
)

// GormRepository implements TaskRepository using GORM.
type GormRepository struct {
	db *gorm.DB
}

// NewGormRepository creates a new instance of GormRepository.
func NewGormRepository(db *gorm.DB) *GormRepository {
	return &GormRepository{db: db}
}

func (r *GormRepository) Create(task *models.Task) error {
	return r.db.Create(task).Error
}

func (r *GormRepository) GetAll() ([]models.Task, error) {
	var tasks []models.Task
	result := r.db.Find(&tasks)
	return tasks, result.Error
}

func (r *GormRepository) GetByID(id int) (*models.Task, error) {
	var task models.Task
	result := r.db.First(&task, id)
	return &task, result.Error
}

func (r *GormRepository) Update(task *models.Task) error {
	return r.db.Save(task).Error
}

func (r *GormRepository) Delete(task *models.Task) error {
	return r.db.Delete(task).Error
}
