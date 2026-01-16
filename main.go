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
	State       string
	Priority    string
}

func create_task() Task {
	fmt.Println("Enter the name, Responsible, state and Priority of the task: ")

	var Name string
	var Responsible string
	var State string
	var Priority string

	n, err := fmt.Scanln(&Name, &Responsible, &State, &Priority)
	if err != nil {
		fmt.Println("something went wrong there: ", err)
	}

	newTask := Task{
		Name:        Name,
		Responsible: Responsible,
		State:       State,
		Priority:    Priority,
	}

	fmt.Println(n)
	return newTask
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		fmt.Println("Sorry couldn't connect to the database: ", err)
	}

	if err := db.AutoMigrate(&Task{}); err != nil {
		fmt.Println("Migration failed: ", err)
	}

	task := create_task()
	db.Create(&task)
}
