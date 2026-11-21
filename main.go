package main

import (
	"fmt"

	"github.com/goobermv/task-tracker-CLI/models"
	"github.com/goobermv/task-tracker-CLI/service"
)

func main() {
	fmt.Println("---Task Tracker CLI---")

	var taskList []models.Task

	for {
		fmt.Println("\nChoose an option:")
		fmt.Println("1 - Add a new task")
		fmt.Println("2 - Print all tasks")
		fmt.Println("3 - Find task by ID")
		fmt.Println("4 - Mark a task as completed")
		fmt.Println("5 - Delete a task")
		fmt.Println("0 - Exit")

		var i int
		fmt.Println("Enter a number: ")
		fmt.Scan(&i)
		if i == 0 {
			fmt.Print("Exiting program..")
			return
		}

		switch i {
		case 1:
			var title, description string
			fmt.Println("\nEnter a title: ")
			fmt.Scan(&title)
			fmt.Println("Enter a description: ")
			fmt.Scan(&description)

			service.AddTask(title, description, &taskList)

		case 2:
			service.GetTasks(taskList)
		case 3:
			service.GetTaskByID(taskList)
		case 4:
			service.CompleteTask(&taskList)
		case 5:
			service.DeleteTask(&taskList)
		}
	}
}
