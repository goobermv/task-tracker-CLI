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
		fmt.Printf("\nID: %d\nTitle: %s\nDescription: %s\nCompleted: %t\nCreated At: %s\n", task.ID, task.Title, task.Description, task.Completed, task.CreatedAt.Format("2006-01-02 15:04:05"))
	}
}

func GetTaskByID(id int, taskList []models.Task) (*models.Task, error) {
	var taskID int
	fmt.Scan("Enter task ID to find: ", &taskID)

	for _, task := range taskList {
		if taskID == task.ID {
			fmt.Print("\nTask:")
			fmt.Printf("\nID: %d\nTitle: %s\nDescription: %s\nCompleted: %t\nCreated At: %s\n", task.ID, task.Title, task.Description, task.Completed, task.CreatedAt.Format("2006-01-02 15:04:05"))
			return &task, nil
		}
	}

	return nil, fmt.Errorf("No task with this ID")
}
