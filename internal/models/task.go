package models

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	ID          uint `gorm:"primarykey"`
	Name        string
	Responsible string
	State       string
	Priority    string
	CreatedDate string
}

// BeforeCreate is a GORM hook that runs before inserting a record
func (t *Task) BeforeCreate(tx *gorm.DB) (err error) {
	t.CreatedDate = time.Now().Format("2006-01-02 15:04")
	return
}
