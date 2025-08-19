package service

import (
	"fmt"
	"time"

	"github.com/zevinto/task-cli/internal/model"
	"github.com/zevinto/task-cli/internal/repository"
)

// AddTask creates a new task
func AddTask(description string) error {
	store, err := repository.LoadTasks()
	if err != nil {
		return err
	}

	// Generate new ID
	newID := 1
	if len(store.Tasks) > 0 {
		newID = store.Tasks[len(store.Tasks)-1].ID + 1
	}

	// Create new Task
	newTask := model.Task{
		ID:          newID,
		Description: description,
		Status:      model.StatusTodo,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	store.Tasks = append(store.Tasks, newTask)
	if err := repository.SaveTasks(store); err != nil {
		return err
	}
	fmt.Printf("Task added successfully (ID: %d)\n", newID)
	return nil
}

// UpdateTask modifies an existing task's description
func UpdateTask(id int, description string) error {
	store, err := repository.LoadTasks()
	if err != nil {
		return err
	}
	for i, task := range store.Tasks {
		if task.ID == id {
			store.Tasks[i].Description = description
			store.Tasks[i].UpdatedAt = time.Now()
			return repository.SaveTasks(store)
		}
	}
	return fmt.Errorf("task with ID %d not found", id)
}

// DeleteTask removes a task
func DeleteTask(id int) error {
	store, err := repository.LoadTasks()
	if err != nil {
		return err
	}
	for i, task := range store.Tasks {
		if task.ID == id {
			store.Tasks = append(store.Tasks[:i], store.Tasks[i+1:]...)
			return repository.SaveTasks(store)
		}
	}
	return fmt.Errorf("task with ID %d not found", id)
}

// MarkTaskInProgress marks a task as in-progress
func MarkTaskInProgress(id int) error {
	return markTask(id, model.StatusInProgress)
}

// MarkTaskDone marks a task as done
func MarkTaskDone(id int) error {
	return markTask(id, model.StatusDone)
}

// markTask updates a task's status
func markTask(id int, status string) error {
	store, err := repository.LoadTasks()
	if err != nil {
		return err
	}

	for i, task := range store.Tasks {
		if task.ID == id {
			store.Tasks[i].Status = status
			store.Tasks[i].UpdatedAt = time.Now()
			return repository.SaveTasks(store)
		}
	}
	return fmt.Errorf("task with ID %d not found", id)
}

// ListTasks displays tasks based on status filter
func ListTasks(statusFilter string) error {
	store, err := repository.LoadTasks()
	if err != nil {
		return err
	}
	if len(store.Tasks) == 0 {
		fmt.Println("No tasks found")
		return nil
	}

	for _, task := range store.Tasks {
		if statusFilter == "" || task.Status == statusFilter {
			fmt.Printf("ID: %d\nDescription: %s\nStatus: %s\nCreatedAt: %s\nUpdatedAt: %s\n\n",
				task.ID,
				task.Description,
				task.Status,
				task.CreatedAt.Format(time.DateTime),
				task.UpdatedAt.Format(time.DateTime))
		}
	}
	return nil
}
