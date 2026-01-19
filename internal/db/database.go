package db

import (
	"github.com/AFHH999/ToDo/internal/models"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func InitDB(dbName string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// AutoMigrate the schema
	if err := db.AutoMigrate(&models.Task{}); err != nil {
		return nil, err
	}

	return db, nil
}
