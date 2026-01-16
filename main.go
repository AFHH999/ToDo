package main

import (
	"fmt"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Name        string
	Responsible string
	State       bool
	Priority    string
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		fmt.Println("Could not migrate the database:", err)
		return
	}
	err = db.AutoMigrate(&Task{})
	if err != nil {
		fmt.Println("Could not migrate the database: ", err)
		return
	}
}
