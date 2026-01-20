package app

import (
	"bufio"
	"fmt"
	"github.com/AFHH999/ToDo/internal/models"
	"gorm.io/gorm"
	"strconv"
	"strings"
)

// GetInput reads a line from the reader and trims whitespace.
func GetInput(prompt string, reader *bufio.Reader) string {
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func CreateTask(reader *bufio.Reader, db *gorm.DB) {
	fmt.Println("\n--- New Task ---")

	var name string
	for {
		name = GetInput("Enter the task name: ", reader)
		if name != "" {
			break
		}
		fmt.Println("Task name cannot be empty.")
	}

	responsible := GetInput("Enter the responsible person: ", reader)

	var state string
	for {
		state = GetInput("Enter state (To Do, In Progress, Done): ", reader)
		if state == "To Do" || state == "In Progress" || state == "Done" {
			break
		}
		fmt.Println("Invalid state. Please enter exactly: To Do, In Progress, or Done")
	}

	priority := GetInput("Enter priority: ", reader)

	newTask := models.Task{
		Name:        name,
		Responsible: responsible,
		State:       state,
		Priority:    priority,
	}

	result := db.Create(&newTask)
	if result.Error != nil {
		fmt.Println("Failed to create task:", result.Error)
	} else {
		fmt.Printf("Task created successfully! ID: %d\n", newTask.ID)
	}
}

func ListTasks(db *gorm.DB) {
	var tasks []models.Task
	result := db.Find(&tasks)
	if result.Error != nil {
		fmt.Println("Error fetching the tasks: ", result.Error)
		return
	}
	fmt.Println("\n---- Current tasks ----")
	for _, task := range tasks {
		fmt.Printf("[%d] %s | Responsible: %s | State: %s | Priority: %s | Created: %s\n",
			task.ID, task.Name, task.Responsible, task.State, task.Priority, task.CreatedDate)
	}
	fmt.Println("------------------------")
}

func EditTask(db *gorm.DB, reader *bufio.Reader) {
	var task models.Task

	idStr := GetInput("Insert the id of the task to modify: ", reader)
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("Invalid ID")
		return
	}

	if err := db.First(&task, id).Error; err != nil {
		fmt.Println("Task not found!")
		return
	}
	fmt.Printf("Editing task: %s\n", task.Name)

	name := GetInput("Enter new name (press enter to keep current): ", reader)
	if name != "" {
		task.Name = name
	}

	responsible := GetInput("Enter new responsible (press enter to keep current): ", reader)
	if responsible != "" {
		task.Responsible = responsible
	}

	for {
		state := GetInput("Enter state (To Do, In Progress, Done) or keep it: ", reader)
		if state == "" {
			break
		}
		if state == "To Do" || state == "In Progress" || state == "Done" {
			task.State = state
			break
		}
		fmt.Println("Sorry most be To Do, In Progress or Done!")
	}

	priority := GetInput("Enter new priority (press enter to keep current): ", reader)
	if priority != "" {
		task.Priority = priority
	}

	db.Save(&task)
	fmt.Println("Task saved successfully!")
}

func DeleteTask(db *gorm.DB, reader *bufio.Reader) {
	var task models.Task

	idStr := GetInput("Insert the ID of the task to delete: ", reader)
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("Invalid ID")
		return
	}

	if err := db.First(&task, id).Error; err != nil {
		fmt.Println("Task not found!")
		return
	}

	db.Delete(&task)
	fmt.Println("Task deleted successfully!")
}
