package main

import (
	"bufio"
	"fmt"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"os"
	"strconv"
	"strings"
	"time"
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

func getInput(prompt string, reader *bufio.Reader) string { // This is how to get a personalized prompt in Go
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func createTask(reader *bufio.Reader) Task {
	fmt.Println("\n--- New Task ---")

	// Loop 1: Name (Required)
	var name string
	for {
		name = getInput("Enter the task name: ", reader)
		if name != "" {
			break
		}
		fmt.Println("Task name cannot be empty.")
	}

	// Responsible (Optional)
	responsible := getInput("Enter the responsible person: ", reader)

	// Loop 2: State (Strict Validation)
	var state string
	for {
		state = getInput("Enter state (To Do, In Progress, Done): ", reader)
		// strictly check the allowed values
		if state == "To Do" || state == "In Progress" || state == "Done" {
			break
		}
		fmt.Println("Invalid state. Please enter exactly: To Do, In Progress, or Done")
	}

	// Priority (Optional)
	priority := getInput("Enter priority: ", reader)

	return Task{
		Name:        name,
		Responsible: responsible,
		State:       state,
		Priority:    priority,
	}
}

func listTasks(db *gorm.DB) {
	var tasks []Task
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

func editTask(db *gorm.DB, reader *bufio.Reader) {
	var task Task

	idStr := getInput("Insert the id of the task to modify: ", reader)
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

	// Update logic using helper
	name := getInput("Enter new name (press enter to keep current): ", reader)
	if name != "" {
		task.Name = name
	}

	responsible := getInput("Enter new responsible (press enter to keep current): ", reader)
	if responsible != "" {
		task.Responsible = responsible
	}

	state := getInput("Enter new state (press enter to keep current): ", reader)
	if state != "" {
		task.State = state
	}

	priority := getInput("Enter new priority (press enter to keep current): ", reader)
	if priority != "" {
		task.Priority = priority
	}

	db.Save(&task)
	fmt.Println("Task saved successfully!")
}

func deleteTask(db *gorm.DB, reader *bufio.Reader) {
	var task Task

	idStr := getInput("Insert the ID of the task to delete: ", reader)
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

func main() {
	reader := bufio.NewReader(os.Stdin)

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		fmt.Println("Sorry couldn't connect to the database: ", err)
	}
	if err := db.AutoMigrate(&Task{}); err != nil {
		fmt.Println("Migration failed: ", err)
	}

	for {
		fmt.Println("\nPlease what do you want to do?")
		fmt.Println("1- To add a new task")
		fmt.Println("2- To list all the tasks")
		fmt.Println("3- To edit the task")
		fmt.Println("4- To delete the task")
		fmt.Println("5- To exit")

		input := getInput("", reader)
		menu, err := strconv.Atoi(input)

		if err != nil {
			fmt.Println("Please enter a number.")
			continue
		}

		switch menu {
		case 1:
			task := createTask(reader)
			result := db.Create(&task)
			if result.Error != nil {
				fmt.Println("Failed to create task:", result.Error)
			} else {
				fmt.Printf("Task created successfully! ID: %d\n", task.ID)
			}
		case 2:
			listTasks(db)
		case 3:
			editTask(db, reader)
		case 4:
			deleteTask(db, reader)
		case 5:
			fmt.Println("Have a great day!")
			return
		default:
			fmt.Println("Invalid option.")
		}
	}
}
