# Task Tracker CLI

Task Tracker CLI is a command-line tool for managing and tracking tasks. It allows users to add, update, delete, mark tasks as in-progress or done, and list tasks by status. Task data is stored in a JSON file locally.

## Project Structure

```
task-cli/
├── Makefile
├── cmd
│   └── main.go
├── go.mod
└── internal
    ├── model
    │   ├── task.go
    │   └── task_test.go
    ├── repository
    │   ├── task_repository.go
    │   └── task_repository_test.go
    └── service
        ├── task_service.go
        └── task_service_test.go
```

## Features

- Add new tasks
- Update task descriptions
- Delete tasks
- Mark tasks as "in-progress" or "done"
- List all tasks or filter by status (todo, in-progress, done)
- Persist task data in a `tasks.json` file

## Dependencies

- Go 1.18 or higher
- No external library dependencies; uses only the Go standard library

## Installation

1. Clone or download the project to your local machine:
   
   ```bash
   git clone <repository-url>
   cd task-cli
   ```

2. Initialize the Go module (if not already initialized):
   
   ```bash
   go mod init task-cli
   ```

3. Install dependencies:
   
   ```bash
   make deps
   ```

4. Build the project:
   
   ```bash
   make build
   ```

## Usage

After building, the executable is named `task-cli`. Below are the supported commands:

### Add a Task

```bash
./task-cli add "Buy groceries"
```

Output: `Task added successfully (ID: 1)`

### Update a Task

```bash
./task-cli update 1 "Buy groceries and cook dinner"
```

Output: `Task updated successfully`

### Delete a Task

```bash
./task-cli delete 1
```

Output: `Task deleted successfully`

### Mark Task Status

- Mark as in-progress:
  
  ```bash
  ./task-cli mark-in-progress 1
  ```
  
  Output: `Task marked as in-progress`

- Mark as done:
  
  ```bash
  ./task-cli mark-done 1
  ```
  
  Output: `Task marked as done`

### List Tasks

- List all tasks:
  
  ```bash
  ./task-cli list
  ```

- Filter by status:
  
  ```bash
  ./task-cli list todo
  ./task-cli list in-progress
  ./task-cli list done
  ```

### Task Properties

Each task includes the following properties, stored in `tasks.json`:

- `id`: Unique identifier
- `description`: Task description
- `status`: Task status (`todo`, `in-progress`, `done`)
- `createdAt`: Creation timestamp
- `updatedAt`: Last updated timestamp

## Testing

The project includes comprehensive unit tests covering the model, repository, and service layers.

Run tests:

```bash
  make test
```

Test files:

- `internal/model/task_test.go`: Tests the task struct
- `internal/repository/task_repository_test.go`: Tests JSON file operations
- `internal/service/task_service_test.go`: Tests task management logic

## Makefile Commands

- `make` or `make all`: Build the project
- `make build`: Compile the `task-cli` binary
- `make run`: Build and run the application
- `make clean`: Remove the binary
- `make deps`: Tidy and download dependencies
- `make fmt`: Format code
- `make test`: Run all tests

## Notes

- Task data is stored in `tasks.json` in the project root directory.
- Commands return error messages for invalid task IDs.
- Ensure the Go environment is properly set up (`go` command available).

## Contributing

Contributions are welcome! Please submit issues or pull requests. Ensure code follows Go conventions and run `make fmt` and `make test` before submitting.

## License

MIT License
