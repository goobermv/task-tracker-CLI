package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/goobermv/task-tracker-CLI/internal/service"
)

func main() {
	fmt.Println("---Task Tracker CLI---")

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\nChoose an option:")
		fmt.Println("1 - Add a new task")
		fmt.Println("2 - Print all tasks")
		fmt.Println("3 - Find task by ID")
		fmt.Println("4 - Mark a task as completed")
		fmt.Println("5 - Delete a task")
		fmt.Println("0 - Exit")

		fmt.Print("\nEnter a number: ")

		if !scanner.Scan() {
			fmt.Println("\nInput closed. Exiting...")
			return
		}

		choice, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Please enter a valid number.")
			continue
		}

		if choice == 0 {
			fmt.Println("Exiting program...")
			return
		}

		switch choice {
		case 1:
			fmt.Print("\nEnter a title: ")
			if !scanner.Scan() {
				return
			}
			title := scanner.Text()

			fmt.Print("Enter a description: ")
			if !scanner.Scan() {
				return
			}
			description := scanner.Text()

			if err := service.AddTask(title, description); err != nil {
				fmt.Println("Error:", err)
			}

		case 2:
			if err := service.GetTasks(); err != nil {
				fmt.Println("Error:", err)
			}

		case 3:
			fmt.Print("\nEnter task ID to find: ")
			if !scanner.Scan() {
				return
			}
			taskID, err := strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println("Error: ", err)
			}

			if err := service.GetTaskByID(taskID); err != nil {
				fmt.Println("Error:", err)
			}

		case 4:
			fmt.Print("\nEnter task ID of completed task: ")
			if !scanner.Scan() {
				return
			}
			taskID, err := strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println("Error: ", err)
			}

			if err := service.CompleteTask(taskID); err != nil {
				fmt.Println("Error:", err)
			}

		case 5:
			fmt.Print("\nEnter task ID of task to be deleted: ")
			if !scanner.Scan() {
				return
			}
			taskID, err := strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println("Error: ", err)
			}

			if err := service.DeleteTask(taskID); err != nil {
				fmt.Println("Error:", err)
			}

		default:
			fmt.Println("Unknown option. Try again.")
		}
	}
}
