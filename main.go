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

func create_task() Task {

	//Task
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the task: ")

	for {
		name, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Sorry could not get that input")
			continue

		}
		Name := strings.TrimSpace(name)

		//Responsible
		fmt.Println("Enter the resposible for the task: ")
		responsible, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Sorry couldn't get the input: ")
			continue

		}
		Responsible := strings.TrimSpace(responsible)

		//State
		fmt.Println("Enter the state of the task: ")
		state, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Sorry couldn't get the input of the state: ", err)
			continue

		}
		State := strings.TrimSpace(state)

		//Priority
		fmt.Println("Enter the priority of the task: ")
		priority, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Sorry couldn't get the input: ", err)
			continue

		}
		Priority := strings.TrimSpace(priority)

		newTask := Task{
			Name:        Name,
			Responsible: Responsible,
			State:       State,
			Priority:    Priority,
		}

		return newTask
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

func editTask(db *gorm.DB) {
	var task Task

	//Getting input from user
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Insert the id of the task to modify: ")
	id_Str, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Couldn't read that string: ", err)
		return
	}
	id_Str = strings.TrimSpace(id_Str)
	ID, err := strconv.Atoi(id_Str)
	if err != nil {
		fmt.Println("Invalid ID", err)
		return
	}
	if err := db.First(&task, ID).Error; err != nil {
		fmt.Println("Task not found!")
		return
	}
	fmt.Printf("Editing task: %s\n", task.Name)

	//Changing name
	fmt.Println("Enter the new task, or press enter to keep it!")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)
	if name != "" {
		task.Name = name
	}

	// Changing Responsible
	fmt.Println("Enter the name of the responsible for the task or leave it to keep it: ")
	responsible, _ := reader.ReadString('\n')
	responsible = strings.TrimSpace(responsible)
	if responsible != "" {
		task.Responsible = responsible
	}

	//Changing state
	fmt.Println("Enter the state of the task, or leave it to keep it: ")
	state, _ := reader.ReadString('\n')
	state = strings.TrimSpace(state)
	if state != "" {
		task.State = state
	}

	//Changing priority
	fmt.Println("Enter the priority, or leave it to keep it: ")
	priority, _ := reader.ReadString('\n')
	priority = strings.TrimSpace(priority)
	if priority != "" {
		task.Priority = priority
	}

	//Save the new task
	db.Save(&task)
	fmt.Println("Task save it successfully!")
}

func deleteTask(db *gorm.DB) {

	var task Task

	// Getting the ID from standard input
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Insert the ID of the task to delete: ")
	id_Str, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Some error just happened: ", err)
		return
	}
	id_Str = strings.TrimSpace(id_Str)
	ID, err := strconv.Atoi(id_Str)
	if err != nil {
		fmt.Println("Sorry couldn't parse that string: ", err)
		return
	}
	if err := db.First(&task, ID).Error; err != nil {
		fmt.Println("Task not found!")
		return
	}

	db.Delete(&task)
	fmt.Println("Task delete successfully!")
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
		fmt.Println("Please what do you want to do?")
		fmt.Println("1- To add a new task")
		fmt.Println("2- To list all the tasks")
		fmt.Println("3- To edit the task")
		fmt.Println("4- To delete the task")
		fmt.Println("5- To exit")

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Couldn't get the input")
		}

		input = strings.TrimSpace(input)

		Input := strings.TrimSpace(input)
		menu, err := strconv.ParseInt(Input, 10, 32)

		if err != nil {
			fmt.Println("Sorry wrong input: ", err)
		}

		if menu == 1 {
			task := create_task()
			result := db.Create(&task)
			if result.Error != nil {
				fmt.Println("Failed to create task:", result.Error)
			} else {
				fmt.Printf("Task created successfully! ID: %d\n", task.ID)
			}
		} else if menu == 2 {
			listTasks(db)
		} else if menu == 3 {
			editTask(db)
		} else if menu == 4 {
			deleteTask(db)
		} else if menu == 5 {
			fmt.Println("Have a great day!")
			break
		}
	}
}
