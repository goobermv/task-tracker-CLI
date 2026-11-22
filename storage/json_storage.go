package storage

import (
	"encoding/json"
	"os"

	"github.com/goobermv/task-tracker-CLI/models"
)

func ToJSON(task models.Task) error {
	data, err := json.MarshalIndent(task, "", " ")
	if err != nil {
		return err
	}

	err = os.WriteFile("tasks.json", data, 0644)
	if err != nil {
		return err
	}

	return nil
}
