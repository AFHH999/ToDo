package app

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/AFHH999/ToDo/internal/models"
	"github.com/AFHH999/ToDo/internal/repository"
	"strconv"
	"strings"
)

// GetInput reads a line from the reader and trims whitespace.
func GetInput(prompt string, reader *bufio.Reader) string {
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func CreateTask(reader *bufio.Reader, repo repository.TaskRepository) {
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

	err := repo.Create(&newTask)

	if err != nil {
		fmt.Println("Failed to create task:", err)
	} else {
		fmt.Printf("Task created successfully! ID: %d\n", newTask.ID)
	}
}

func ListTasks(repo repository.TaskRepository) {
	var tasks []models.Task
	tasks, err := repo.GetAll()
	if err != nil {
		fmt.Println("Error fetching the tasks: ", err)
		return
	}
	fmt.Println("\n---- Current tasks ----")
	for _, task := range tasks {
		fmt.Printf("[%d] %s | Responsible: %s | State: %s | Priority: %s | Created: %s\n",
			task.ID, task.Name, task.Responsible, task.State, task.Priority, task.CreatedDate)
	}
	fmt.Println("------------------------")
}

func EditTask(repo repository.TaskRepository, reader *bufio.Reader) {

	idStr := GetInput("Insert the id of the task to modify: ", reader)
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("Invalid ID")
		return
	}

	task, err := repo.GetByID(id)
	if err != nil {
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

	if err := repo.Update(task); err != nil {
		fmt.Println("Error saving task:", err)
	}
	fmt.Println("Task saved successfully!")
}

// DeleteTask now properly calls DeleteTaskID internally for the interactive mode
func DeleteTask(repo repository.TaskRepository, reader *bufio.Reader) {
	idStr := GetInput("Insert the ID of the task to delete: ", reader)
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("Invalid ID")
		return
	}
	DeleteTaskID(repo, id)
}

func DeleteTaskID(repo repository.TaskRepository, id int) {
	task, err := repo.GetByID(id)
	if err != nil {
		fmt.Println("Task not found!")
		return
	}

	if err := repo.Delete(task); err != nil {
		fmt.Println("Error deleting task:", err)
	} else {
		fmt.Println("Task deleted successfully!")
	}
}

func CatchFlags(repo repository.TaskRepository) bool {

	var list bool
	var task models.Task
	var deleteID int

	flag.StringVar(&task.Name, "name", "", "Name of the task")
	flag.StringVar(&task.Responsible, "responsible", "Unassigned", "Who is in charge of the task")
	flag.StringVar(&task.State, "state", "To Do", "State of the task (To Do, In Progress, Done)")
	flag.StringVar(&task.Priority, "priority", "Medium", "Priority (High, Medium, Low)")
	flag.BoolVar(&list, "list", false, "List all tasks")

	// Changed from BoolVar to IntVar to accept an ID directly
	flag.IntVar(&deleteID, "delete", 0, "ID of the task to delete it.")

	flag.Parse()

	if list {
		ListTasks(repo)
		return true
	}

	// Check if deleteID was set (not 0)
	if deleteID != 0 {
		DeleteTaskID(repo, deleteID)
		return true
	}

	if task.Name != "" {
		err := repo.Create(&task)
		if err != nil {
			fmt.Println("Error creating the task: ", err)
		} else {
			fmt.Printf("Task '%s' created successfully!\n", task.Name)
		}
		return true
	}
	return false

}
