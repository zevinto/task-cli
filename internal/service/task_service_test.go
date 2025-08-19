package service

import (
	"os"
	"testing"

	"github.com/zevinto/task-cli/internal/model"
	"github.com/zevinto/task-cli/internal/repository"
)

// Setup creates a clean environment for each test
func setup(t *testing.T) {
	os.Remove("tasks.json")
}

// TestAddTask tests adding a new task
func TestAddTask(t *testing.T) {
	setup(t)
	defer os.Remove("tasks.json")

	err := AddTask("Buy groceries")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	store, err := repository.LoadTasks()
	if err != nil {
		t.Errorf("Expected no error loading tasks, got %v", err)
	}
	if len(store.Tasks) != 1 {
		t.Errorf("Expected 1 task, got %d", len(store.Tasks))
	}
	if store.Tasks[0].Description != "Buy groceries" {
		t.Errorf("Expected description 'Buy groceries', got %s", store.Tasks[0].Description)
	}
	if store.Tasks[0].Status != model.StatusTodo {
		t.Errorf("Expected status %s, got %s", model.StatusTodo, store.Tasks[0].Status)
	}
}

// TestUpdateTask tests updating a task's description
func TestUpdateTask(t *testing.T) {
	setup(t)
	defer os.Remove("tasks.json")

	AddTask("Original task")
	err := UpdateTask(1, "Updated task")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	store, err := repository.LoadTasks()
	if err != nil {
		t.Errorf("Expected no error loading tasks, got %v", err)
	}
	if store.Tasks[0].Description != "Updated task" {
		t.Errorf("Expected description 'Updated task', got %s", store.Tasks[0].Description)
	}
	if store.Tasks[0].UpdatedAt.Before(store.Tasks[0].CreatedAt) {
		t.Error("Expected UpdatedAt to be after CreatedAt")
	}
}

// TestUpdateTaskNotFound tests updating a non-existent task
func TestUpdateTaskNotFound(t *testing.T) {
	setup(t)
	defer os.Remove("tasks.json")

	err := UpdateTask(999, "Non-existent task")
	if err == nil {
		t.Error("Expected error for non-existent task, got none")
	}
}

// TestDeleteTask tests deleting a task
func TestDeleteTask(t *testing.T) {
	setup(t)
	defer os.Remove("tasks.json")

	AddTask("Task to delete")
	err := DeleteTask(1)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	store, err := repository.LoadTasks()
	if err != nil {
		t.Errorf("Expected no error loading tasks, got %v", err)
	}
	if len(store.Tasks) != 0 {
		t.Errorf("Expected 0 tasks after deletion, got %d", len(store.Tasks))
	}
}

// TestDeleteTaskNotFound tests deleting a non-existent task
func TestDeleteTaskNotFound(t *testing.T) {
	setup(t)
	defer os.Remove("tasks.json")

	err := DeleteTask(999)
	if err == nil {
		t.Error("Expected error for non-existent task, got none")
	}
}

// TestMarkTaskInProgress tests marking a task as in-progress
func TestMarkTaskInProgress(t *testing.T) {
	setup(t)
	defer os.Remove("tasks.json")

	AddTask("Task to mark")
	err := MarkTaskInProgress(1)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	store, err := repository.LoadTasks()
	if err != nil {
		t.Errorf("Expected no error loading tasks, got %v", err)
	}
	if store.Tasks[0].Status != model.StatusInProgress {
		t.Errorf("Expected status %s, got %s", model.StatusInProgress, store.Tasks[0].Status)
	}
}

// TestMarkTaskDone tests marking a task as done
func TestMarkTaskDone(t *testing.T) {
	setup(t)
	defer os.Remove("tasks.json")

	AddTask("Task to mark")
	err := MarkTaskDone(1)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	store, err := repository.LoadTasks()
	if err != nil {
		t.Errorf("Expected no error loading tasks, got %v", err)
	}
	if store.Tasks[0].Status != model.StatusDone {
		t.Errorf("Expected status %s, got %s", model.StatusDone, store.Tasks[0].Status)
	}
}

// TestListTasks tests listing all tasks
func TestListTasks(t *testing.T) {
	setup(t)
	defer os.Remove("tasks.json")

	AddTask("Task 1")
	AddTask("Task 2")
	AddTask("Task 3")

	store, err := repository.LoadTasks()
	if err != nil {
		t.Errorf("Expected no error loading tasks, got %v", err)
	}
	if len(store.Tasks) != 3 {
		t.Errorf("Expected 3 tasks, got %d", len(store.Tasks))
	}
}

// TestListTasksByStatus tests listing tasks by status
func TestListTasksByStatus(t *testing.T) {
	setup(t)
	defer os.Remove("tasks.json")

	AddTask("Todo task")
	MarkTaskInProgress(1)
	AddTask("Done task")
	MarkTaskDone(2)

	store, err := repository.LoadTasks()
	if err != nil {
		t.Errorf("Expected no error loading tasks, got %v", err)
	}

	// Test in-progress filter
	for _, task := range store.Tasks {
		if task.Status == model.StatusInProgress && task.ID != 1 {
			t.Errorf("Expected only task ID 1 to be in-progress, found ID %d", task.ID)
		}
	}

	// Test done filter
	for _, task := range store.Tasks {
		if task.Status == model.StatusDone && task.ID != 2 {
			t.Errorf("Expected only task ID 2 to be done, found ID %d", task.ID)
		}
	}
}
