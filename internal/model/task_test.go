package model

import (
	"testing"
	"time"
)

// TestTaskStruct validates the Task struct initialization
func TestTaskStruct(t *testing.T) {
	task := Task{
		ID:          1,
		Description: "Test task",
		Status:      StatusDone,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	if task.ID != 1 {
		t.Errorf("Expected ID to be 1, got %d", task.ID)
	}
	if task.Description != "Test task" {
		t.Errorf("Expected Description to be \"Test task\", got \"%s\"", task.Description)
	}
	if task.Status != StatusDone {
		t.Errorf("Expected Status to be %s, got %s", StatusDone, task.Status)
	}
	if task.CreatedAt.IsZero() {
		t.Error("Expected CreatedAt to be set, got zero time")
	}
	if task.UpdatedAt.IsZero() {
		t.Error("Expected UpdatedAt to be set, got zero time")
	}
}

// TestTaskStore validates the TaskStore struct
func TestTaskStore(t *testing.T) {
	store := TaskStore{
		Tasks: []Task{
			{ID: 1, Description: "Task 1", Status: StatusTodo},
			{ID: 2, Description: "Task 2", Status: StatusInProgress},
		},
	}

	if len(store.Tasks) != 2 {
		t.Errorf("Expected 2 tasks in store, got %d", len(store.Tasks))
	}
	if store.Tasks[0].ID != 1 {
		t.Errorf("Expected first task ID to be 1, got %d", store.Tasks[0].ID)
	}
}
