package storage

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/goobermv/task-tracker-CLI/internal/models"
)

func ToJSON(tasks []models.Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal into JSON: %w", err)
	}

	err = os.WriteFile("data/tasks.json", data, 0644)
	if err != nil {
		return fmt.Errorf("Failed to write to file: %w", err)
	}

	return nil
}

func FromJSON() ([]models.Task, error) {
	data, err := os.ReadFile("data/tasks.json")
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return []models.Task{}, nil
		}
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	if len(bytes.TrimSpace(data)) == 0 {
		return []models.Task{}, nil
	}

	var taskList []models.Task

	if err := json.Unmarshal(data, &taskList); err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %w", err)
	}

	return taskList, nil
}
