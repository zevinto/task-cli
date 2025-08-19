package repository

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/zevinto/task-cli/internal/model"
)

// TaskFile is the name of the JSON file to store tasks
const TaskFile = "tasks.json"

// LoadTasks reads task from the JSON file
func LoadTasks() (model.TaskStore, error) {
	var store model.TaskStore

	if _, err := os.Stat(TaskFile); os.IsNotExist(err) {
		return model.TaskStore{Tasks: []model.Task{}}, nil
	}

	data, err := os.ReadFile(TaskFile)
	if err != nil {
		return store, fmt.Errorf("failed to read task file: %w", err)
	}

	err = json.Unmarshal(data, &store)
	if err != nil {
		return store, fmt.Errorf("failed to parse task file: %w", err)
	}
	return store, nil
}

// SaveTasks writes tasks to the JSON file
func SaveTasks(store model.TaskStore) error {
	data, err := json.MarshalIndent(store, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal tasks: %w", err)
	}
	err = os.WriteFile(TaskFile, data, os.ModePerm)
	if err != nil {
		return fmt.Errorf("failed to save tasks: %w", err)
	}
	return nil
}
