package app

import (
	"bufio"
	"strings"
	"testing"

	"github.com/AFHH999/ToDo/internal/models"
	"github.com/AFHH999/ToDo/internal/repository"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func setup_TestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to connect to test the database: %v", err)
	}

	err = db.AutoMigrate(&models.Task{})
	if err != nil {
		t.Fatalf("Failed to migrate test database")
	}

	return db
}

func TestGetInput(t *testing.T) {
	input := "Hello, World!\n"
	reader := bufio.NewReader(strings.NewReader(input))
	result := GetInput("Prompt: ", reader)
	if result != "Hello, World!" {
		t.Errorf("Expected: 'Hello, World!', got '%s'", result)
	}
}

func TestCreatedTask(t *testing.T) {
	db := setup_TestDB(t)
	repo := repository.NewGormRepository(db)

	input := "Test task\nFelipe\nIn Progress\nHigh\n"
	reader := bufio.NewReader(strings.NewReader(input))

	CreateTask(reader, repo)

	var task models.Task
	result := db.First(&task)

	if result.Error != nil {
		t.Errorf("task was not found in database: %v", result.Error)
	}

	if task.Name != "Test task" {
		t.Errorf("Expected: 'Test task', got '%s'", task.Name)
	}
}

func TestEditTask(t *testing.T) {
	db := setup_TestDB(t)
	repo := repository.NewGormRepository(db)

	// 1. Create a task
	inputCreate := "Original Task\nFelipe\nIn Progress\nHigh\n"
	readerCreate := bufio.NewReader(strings.NewReader(inputCreate))
	CreateTask(readerCreate, repo)

	// 2. Edit the task (Fresh reader to avoid buffer issues)
	// Sequence: ID(1) -> New Name -> New Responsible -> New State -> New Priority
	inputEdit := "1\nUpdated Task\nHugo\nDone\nLow\n"
	readerEdit := bufio.NewReader(strings.NewReader(inputEdit))

	EditTask(repo, readerEdit)

	// 3. Verify
	var task models.Task
	db.First(&task, 1)

	if task.Name != "Updated Task" {
		t.Errorf("Expected 'Updated Task', got '%s'", task.Name)
	}
	if task.Responsible != "Hugo" {
		t.Errorf("Expected 'Hugo', got '%s'", task.Responsible)
	}
	if task.State != "Done" {
		t.Errorf("Expected 'Done', got '%s'", task.State)
	}

	if task.Priority != "Low" {
		t.Errorf("Expected 'Low', got '%s'", task.Priority)
	}
}

// func TestListTask() {}
