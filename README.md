# ToDo CLI Application

A simple and efficient command-line interface (CLI) for managing tasks, built with **Go** and **GORM**. This application uses a lo
persistently.

## Features

- **Create Tasks**: Add new tasks with a name, responsible person, state, and priority.
- **List Tasks**: View all your current tasks in a formatted list.
- **Edit Tasks**: Update existing tasks by their ID.
- **Delete Tasks**: Remove tasks you no longer need.
- **Persistent Storage**: All data is saved in a local `test.db` file.
- **Input Validation**: Ensures task names are not empty and states are valid (To Do, In Progress, Done).

## Prerequisites

Before running this project, ensure you have the following installed:

- [Go](https://go.dev/doc/install) (version 1.21 or higher recommended)

## Installation

1. Clone this repository to your local machine:

   ```bash
   git clone https://github.com/AFHH999/ToDo.git
   cd ToDo
   ```

2. Install the necessary dependencies:

   ```bash
   go mod tidy                                                                                                                    
   ```

## Usage

To start the application, run:

```bash
go run cmd/todo/main.go
```

Follow the on-screen menu to manage your tasks:

1. **Add a new task**: Prompts for task details.
2. **List all tasks**: Displays ID, Name, Responsible, State, Priority, and Creation Date.
3. **Edit a task**: Allows updating specific fields of a task by its ID.
4. **Delete a task**: Removes a task from the database by its ID.
5. **Exit**: Closes the application.

Also to use the app you can use the flag functionality

### CLI Flags

You can also use command-line flags to perform actions directly without entering the interactive menu.

To see all available flags:
```bash
go run cmd/todo/main.go -help
```

**Available Flags:**

- `-list`: List all tasks.
- `-name "Task Name"`: Create a new task with the specified name.
- `-responsible "Name"`: Specify the responsible person (used with `-name`). Default: "Unassigned".
- `-state "State"`: Specify the task state (To Do, In Progress, Done) (used with `-name`). Default: "To Do".
- `-priority "Level"`: Specify the priority (High, Medium, Low) (used with `-name`). Default: "Medium".
- `-delete <ID>`: Delete a task by its ID.

**Examples:**

*List all tasks:*
```bash
go run cmd/todo/main.go -list
```

*Create a new task:*
```bash
go run cmd/todo/main.go -name "Fix Bug #42" -responsible "Alice" -priority "High"
```

*Delete a task:*
```bash
go run cmd/todo/main.go -delete 5
```

## Database

The application automatically creates a `test.db` file in the root directory upon first run using GORM's `AutoMigrate` feature.

## Author

- **Felipe** - [AFHH999](https://github.com/AFHH999)
