package service

import (
	"fmt"
	"sync/atomic"
	"time"

	"github.com/goobermv/task-tracker-CLI/models"
)

var taskIDCounter atomic.Uint64

func AddTask(title, description string, taskList []models.Task) error {
	if title == "" {
		return fmt.Errorf("Title is empty, please input a title for the task...")
	}
	if description == "" {
		return fmt.Errorf("Description is empty, please input a description for the task...")
	}

	task := &models.Task{
		ID:          int(taskIDCounter.Add(1)),
		Title:       title,
		Description: description,
		Completed:   false,
		CreatedAt:   time.Now(),
	}
	fmt.Println("Task created successfully!")

	taskList = append(taskList, *task)
	fmt.Println("Task added to task list successfully!")

	return nil
}

func GetTasks(taskList []models.Task) {
	fmt.Println("Task List:")

	for _, task := range taskList {
		fmt.Printf("\nID: %d\nTitle: %s\nDescription: %s\nCompleted: %b\nCreated At: %t\n", task.ID, task.Title, task.Description, task.Completed, task.CreatedAt)
	}
}
