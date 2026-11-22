package service

import (
	"fmt"
	"time"

	"github.com/goobermv/task-tracker-CLI/internal/models"
	"github.com/goobermv/task-tracker-CLI/internal/storage"
)

func AddTask(title, description string) error {
	if title == "" {
		return fmt.Errorf("Title is empty, please input a title for the task")
	}
	if description == "" {
		return fmt.Errorf("Description is empty, please input a description for the task")
	}

	taskList, err := storage.FromJSON()
	if err != nil {
		return err
	}

	var maxID int

	for _, task := range taskList {
		if task.ID > maxID {
			maxID = task.ID
		}
	}

	task := models.Task{
		ID:          maxID + 1,
		Title:       title,
		Description: description,
		Completed:   false,
		CreatedAt:   time.Now(),
	}

	taskList = append(taskList, task)
	if err := storage.ToJSON(taskList); err != nil {
		return err
	}

	fmt.Println("\nTask created successfully!")
	return nil
}

func GetTasks() error {
	taskList, err := storage.FromJSON()
	if err != nil {
		return err
	}

	if len(taskList) == 0 {
		fmt.Println("\nNo tasks found.")
		return nil
	}

	fmt.Println("Task List:")
	for _, task := range taskList {
		fmt.Printf("\nID: %d\nTitle: %s\nDescription: %s\nCompleted: %t\nCreated At: %s\n", task.ID, task.Title, task.Description, task.Completed, task.CreatedAt.Format("2006-01-02 15:04:05"))
	}

	return nil
}

func GetTaskByID(taskID int) error {
	taskList, err := storage.FromJSON()
	if err != nil {
		return err
	}

	for _, task := range taskList {
		if task.ID == taskID {
			fmt.Print("\nTask found:")
			fmt.Printf("\nID: %d\nTitle: %s\nDescription: %s\nCompleted: %t\nCreated At: %s\n", task.ID, task.Title, task.Description, task.Completed, task.CreatedAt.Format("2006-01-02 15:04:05"))
			return nil
		}
	}

	return fmt.Errorf("No task with this ID")
}

func CompleteTask(taskID int) error {
	taskList, err := storage.FromJSON()
	if err != nil {
		return err
	}

	for i, task := range taskList {
		if task.ID != taskID {
			continue
		}

		fmt.Print("\nTask found")
		fmt.Println("Task status being changed")
		taskList[i].Completed = true

		if err := storage.ToJSON(taskList); err != nil {
			return err
		}

		fmt.Println("Task status changed: ")
		fmt.Printf("\nID: %d\nTitle: %s\nDescription: %s\nCompleted: %t\nCreated At: %s\n", taskList[i].ID, taskList[i].Title, taskList[i].Description, taskList[i].Completed, taskList[i].CreatedAt.Format("2006-01-02 15:04:05"))
		return nil
	}

	return fmt.Errorf("No task with this ID")
}

func DeleteTask(taskID int) error {
	taskList, err := storage.FromJSON()
	if err != nil {
		return err
	}

	for i, task := range taskList {
		if task.ID != taskID {
			continue
		}

		fmt.Print("\nTask found:")
		fmt.Printf("\nID: %d\nTitle: %s\nDescription: %s\nCompleted: %t\nCreated At: %s\n", task.ID, task.Title, task.Description, task.Completed, task.CreatedAt.Format("2006-01-02 15:04:05"))

		taskList = append(taskList[:i], taskList[i+1:]...)
		if err := storage.ToJSON(taskList); err != nil {
			return err
		}

		fmt.Println("Task deleted successfully")
		return nil
	}

	return fmt.Errorf("No task with this ID")
}
