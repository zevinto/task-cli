package repository

import (
	"os"
	"testing"
	"time"

	"github.com/zevinto/task-cli/internal/model"
)

// TestLoadTasksEmptyFile tests loading from a non-existent file
func TestLoadTasksEmptyFile(t *testing.T) {
	// Ensure the test file doesn't exist
	os.Remove(TaskFile)
	store, err := LoadTasks()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if len(store.Tasks) != 0 {
		t.Errorf("Expected empty task list, got %d tasks", len(store.Tasks))
	}
}

// TestSaveAndLoadTasks tests saving and loading tasks
func TestSaveAndLoadTasks(t *testing.T) {
	// Clean up before and after test
	os.Remove(TaskFile)
	defer os.Remove(TaskFile)

	store := model.TaskStore{
		Tasks: []model.Task{
			{
				ID:          1,
				Description: "Test task",
				Status:      model.StatusTodo,
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			},
		},
	}
	err := SaveTasks(store)
	if err != nil {
		t.Errorf("Expected no error saving tasks, got %v", err)
	}

	loadedStore, err := LoadTasks()
	if err != nil {
		t.Errorf("Expected no error loading tasks, got %v", err)
	}
	if len(loadedStore.Tasks) != 1 {
		t.Errorf("Expected 1 task, got %d", len(loadedStore.Tasks))
	}
	if loadedStore.Tasks[0].Description != "Test task" {
		t.Errorf("Expected task description to be 'Test task', got '%s'", loadedStore.Tasks[0].Description)
	}
}
