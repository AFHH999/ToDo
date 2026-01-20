package db

import (
	"github.com/AFHH999/ToDo/internal/models"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

// Init opens a connection to the SQLite database and runs migrations.
func Init(dbPath string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// AutoMigrate will create the tables based on the struct
	if err := db.AutoMigrate(&models.Task{}); err != nil {
		return nil, err
	}

	return db, nil
}

