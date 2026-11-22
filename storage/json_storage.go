package storage

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/goobermv/task-tracker-CLI/models"
)

func ToJSON(task models.Task) error {
	data, err := json.MarshalIndent(task, "", " ")
	if err != nil {
		return fmt.Errorf("Failed to marshal into JSON: %v", err)
	}

	err = os.WriteFile("tasks.json", data, 0644)
	if err != nil {
		return fmt.Errorf("Failed to write to file: %v", err)
	}

	return nil
}

func FromJSON(taskList []models.Task) (*[]models.Task, error) {
	data, err := os.ReadFile("tasks.json")
	if err != nil {
		return nil, fmt.Errorf("failed to read file")
	}

	if err := json.Unmarshal(data, &taskList); err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %v", err)
	}

	return &taskList, nil
}

// adding a find by ID JSON function
// adding an update JSON struct function
// adding a delete JSON struct function
