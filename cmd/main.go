package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/zevinto/task-cli/internal/model"
	"github.com/zevinto/task-cli/internal/service"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: task-cli <command> [arguments]")
		fmt.Println("Commands:")
		fmt.Println("  add <description>")
		fmt.Println("  update <id> <description>")
		fmt.Println("  delete <id>")
		fmt.Println("  mark-in-progress <id>")
		fmt.Println("  mark-done <id>")
		fmt.Println("  list [done|todo|in-progress]")
		os.Exit(1)
	}

	command := os.Args[1]

	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Usage: task-cli add <description>")
			os.Exit(1)
		}
		if err := service.AddTask(os.Args[2]); err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("Task added successfully")
	case "update":
		if len(os.Args) < 4 {
			fmt.Println("Usage: task-cli update <id> <description>")
			os.Exit(1)
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Error: Invalid ID format")
			os.Exit(1)
		}
		if err := service.UpdateTask(id, os.Args[3]); err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("Task updated successfully")
	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Usage: task-cli delete <id>")
			os.Exit(1)
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Error: Invalid ID format")
			os.Exit(1)
		}
		if err := service.DeleteTask(id); err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("Task deleted successfully")
	case "mark-in-progress":
		if len(os.Args) < 3 {
			fmt.Println("Usage: task-cli mark-in-progress <id>")
			os.Exit(1)
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Error: Invalid ID format")
			os.Exit(1)
		}
		if err := service.MarkTaskInProgress(id); err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("Task marked in-progress successfully")
	case "mark-done":
		if len(os.Args) < 3 {
			fmt.Println("Usage: task-cli mark-done <id>")
			os.Exit(1)
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Error: Invalid ID format")
			os.Exit(1)
		}
		if err := service.MarkTaskDone(id); err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("Task marked done successfully")
	case "list":
		var status string
		if len(os.Args) < 2 {
			status = os.Args[2]
			if status != model.StatusTodo && status != model.StatusDone && status != model.StatusInProgress {
				fmt.Printf("Error: Invalid status. Use 'todo', 'in-progress', or 'done'\n")
				os.Exit(1)
			}
		}
		if err := service.ListTasks(status); err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
	default:
		fmt.Printf("Unknown command: %s\n", command)
		fmt.Println("Available commands: add, update, delete, mark-in-progress, mark-done, list")
		os.Exit(1)
	}
}
