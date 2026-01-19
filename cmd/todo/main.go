package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/AFHH999/ToDo/internal/db"
	"github.com/AFHH999/ToDo/internal/models"
	"github.com/AFHH999/ToDo/internal/services"
	"gorm.io/gorm"
)

func getInput(prompt string, reader *bufio.Reader) string {
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func handleCreate(reader *bufio.Reader, database *gorm.DB) {
	fmt.Println("\n--- New Task ---")

	var name string
	for {
		name = getInput("Enter the task name: ", reader)
		if name != "" {
			break
		}
		fmt.Println("Task name cannot be empty.")
	}

	responsible := getInput("Enter the responsible person: ", reader)

	var state string
	for {
		state = getInput("Enter state (To Do, In Progress, Done): ", reader)
		if state == "To Do" || state == "In Progress" || state == "Done" {
			break
		}
		fmt.Println("Invalid state. Please enter exactly: To Do, In Progress, or Done")
	}

	priority := getInput("Enter priority: ", reader)

	newTask := models.Task{
		Name:        name,
		Responsible: responsible,
		State:       state,
		Priority:    priority,
	}

	id, err := services.CreateTask(database, newTask)
	if err != nil {
		fmt.Println("Failed to create task:", err)
	} else {
		fmt.Printf("Task created successfully! ID: %d\n", id)
	}
}

func handleList(database *gorm.DB) {
	tasks, err := services.ListTasks(database)
	if err != nil {
		fmt.Println("Error fetching tasks:", err)
		return
	}

	fmt.Println("\n---- Current tasks ----")
	for _, task := range tasks {
		fmt.Printf("[%d] %s | Responsible: %s | State: %s | Priority: %s | Created: %s\n",
			task.ID, task.Name, task.Responsible, task.State, task.Priority, task.CreatedDate)
	}
	fmt.Println("------------------------")
}

func handleEdit(reader *bufio.Reader, database *gorm.DB) {
	idStr := getInput("Insert the id of the task to modify: ", reader)
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("Invalid ID")
		return
	}

	task, err := services.GetTaskByID(database, id)
	if err != nil {
		fmt.Println("Task not found!")
		return
	}
	fmt.Printf("Editing task: %s\n", task.Name)

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

	if err := services.UpdateTask(database, task); err != nil {
		fmt.Println("Error updating task:", err)
	} else {
		fmt.Println("Task saved successfully!")
	}
}

func handleDelete(reader *bufio.Reader, database *gorm.DB) {
	idStr := getInput("Insert the ID of the task to delete: ", reader)
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("Invalid ID")
		return
	}

	if err := services.DeleteTask(database, id); err != nil {
		fmt.Println("Error deleting task:", err)
	} else {
		fmt.Println("Task deleted successfully!")
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	database, err := db.InitDB("test.db")
	if err != nil {
		fmt.Println("Sorry couldn't connect to the database: ", err)
		return
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
			handleCreate(reader, database)
		case 2:
			handleList(database)
		case 3:
			handleEdit(reader, database)
		case 4:
			handleDelete(reader, database)
		case 5:
			fmt.Println("Have a great day!")
			return
		default:
			fmt.Println("Invalid option.")
		}
	}
}
