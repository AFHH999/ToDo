package main

import (
	"bufio"
	"fmt"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"os"
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

// Delete the Scan function, doesn't do shit!
func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		fmt.Println("Sorry couldn't connect to the database: ", err)
	}
	if err := db.AutoMigrate(&Task{}); err != nil {
		fmt.Println("Migration failed: ", err)
	}

	var menu int

	for {
		fmt.Println("Please what do you want to do?")
		fmt.Println("1- To add a new task")
		fmt.Println("2- To list all the tasks")
		fmt.Println("3- To exit")
		_, err := fmt.Scan(&menu)
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
			break
		}
	}
}
