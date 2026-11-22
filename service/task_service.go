package service

import (
	"fmt"
	"slices"
	"sync/atomic"
	"time"

	"github.com/goobermv/task-tracker-CLI/models"
)

var taskIDCounter atomic.Uint64

func AddTask(title, description string, taskList *[]models.Task) error {
	if title == "" {
		return fmt.Errorf("Title is empty, please input a title for the task")
	}
	if description == "" {
		return fmt.Errorf("Description is empty, please input a description for the task")
	}

	task := &models.Task{
		ID:          int(taskIDCounter.Add(1)),
		Title:       title,
		Description: description,
		Completed:   false,
		CreatedAt:   time.Now(),
	}
	fmt.Println("Task created successfully!")

	*taskList = append(*taskList, *task)
	fmt.Println("Task added to task list successfully!")

	return nil
}

func GetTasks(taskList []models.Task) {
	fmt.Println("Task List:")

	for _, task := range taskList {
		fmt.Printf("\nID: %d\nTitle: %s\nDescription: %s\nCompleted: %t\nCreated At: %s\n", task.ID, task.Title, task.Description, task.Completed, task.CreatedAt.Format("2006-01-02 15:04:05"))
	}
}

func GetTaskByID(taskList []models.Task) (*models.Task, error) {
	var taskID int
	fmt.Print("Enter task ID to find: ")
	if _, err := fmt.Scan(&taskID); err != nil {
		return nil, fmt.Errorf("Failed to scan input: %v", err)
	}

	for i := range taskList {
		if taskList[i].ID == taskID {
			fmt.Print("\nTask found:")
			fmt.Printf("\nID: %d\nTitle: %s\nDescription: %s\nCompleted: %t\nCreated At: %s\n", taskList[i].ID, taskList[i].Title, taskList[i].Description, taskList[i].Completed, taskList[i].CreatedAt.Format("2006-01-02 15:04:05"))
			return &taskList[i], nil
		}
	}

	return nil, fmt.Errorf("No task with this ID")
}

func CompleteTask(taskList *[]models.Task) error {
	var taskID int
	fmt.Print("Enter task ID to change status: ")
	if _, err := fmt.Scan(&taskID); err != nil {
		return fmt.Errorf("Failed to scan input: %v", err)
	}

	for i := range *taskList {
		if (*taskList)[i].ID == taskID {
			fmt.Println("\nTask found")
			fmt.Println("Task status being changed")
			(*taskList)[i].Completed = true
			fmt.Println("Task status changed: ")
			fmt.Printf("\nID: %d\nTitle: %s\nDescription: %s\nCompleted: %t\nCreated At: %s\n", (*taskList)[i].ID, (*taskList)[i].Title, (*taskList)[i].Description, (*taskList)[i].Completed, (*taskList)[i].CreatedAt.Format("2006-01-02 15:04:05"))
			return nil
		}
	}

	return fmt.Errorf("No task with this ID")
}

func DeleteTask(taskList *[]models.Task) error {
	var taskID int
	fmt.Print("Enter task ID to delete: ")
	if _, err := fmt.Scan(&taskID); err != nil {
		return fmt.Errorf("Failed to scan input: %v", err)
	}

	for _, task := range *taskList {
		if taskID == task.ID {
			fmt.Print("\nTask found:")
			fmt.Printf("\nID: %d\nTitle: %s\nDescription: %s\nCompleted: %t\nCreated At: %s\n", task.ID, task.Title, task.Description, task.Completed, task.CreatedAt.Format("2006-01-02 15:04:05"))
			*taskList = slices.DeleteFunc(*taskList, func(t models.Task) bool {
				return task.ID == taskID
			})
			fmt.Println("Task deleted successfully")
			return nil
		}
	}

	return fmt.Errorf("No task with this ID")
}
