package model

import "time"

// Task represents a single task with all required properties
type Task struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

// TaskStore manages the collections of tasks
type TaskStore struct {
	Tasks []Task `json:"tasks"`
}

// Status constants
const (
	StatusTodo       = "todo"
	StatusInProgress = "in-progress"
	StatusDone       = "done"
)
