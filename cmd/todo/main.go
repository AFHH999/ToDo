package main

import (
	"bufio"
	"fmt"
	"github.com/AFHH999/ToDo/internal/app"
	"github.com/AFHH999/ToDo/internal/db"
	"os"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	dbPath := os.Getenv("TODO_DB_PATH") // This works because it ask the OS where is the database
	if dbPath == "" {
		dbPath = "test.db" // If theres is none then it crate it.
	}

	database, err := db.Init(dbPath)
	if err != nil {
		fmt.Println("Sorry couldn't connect to the database: ", err)
		return
	}

	if app.CatchFlags(database) {
		return
	}

	for {
		fmt.Println("\nPlease what do you want to do?")
		fmt.Println("1- To add a new task")
		fmt.Println("2- To list all the tasks")
		fmt.Println("3- To edit the task")
		fmt.Println("4- To delete the task")
		fmt.Println("5- To exit")

		input := app.GetInput("", reader)
		menu, err := strconv.Atoi(input)

		if err != nil {
			fmt.Println("Please enter a number.")
			continue
		}

		switch menu {
		case 1:
			app.CreateTask(reader, database)
		case 2:
			app.ListTasks(database)
		case 3:
			app.EditTask(database, reader)
		case 4:
			app.DeleteTask(database, reader)
		case 5:
			fmt.Println("Have a great day!")
			return
		default:
			fmt.Println("Invalid option.")
		}
	}
}
