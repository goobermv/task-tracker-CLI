package service

import (
	"errors"
	"testing"

	"github.com/goobermv/task-tracker-CLI/internal/models"
)

func TestAddTaskValidation(t *testing.T) {
	tests := []struct {
		name        string
		title       string
		description string
		wantErr     error
	}{
		{
			name:        "empty title",
			title:       "",
			description: "Learn Go",
			wantErr:     models.ErrNoTitle,
		},
		{
			name:        "empty description",
			title:       "Learn Go",
			description: "",
			wantErr:     models.ErrNoDescription,
		},
		{
			name:        "empty title and description",
			title:       "",
			description: "",
			wantErr:     models.ErrNoTitle,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := AddTask(tt.title, tt.description)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("AddTask() error = %v, want %v", err, tt.wantErr)
			}

		})
	}
}
