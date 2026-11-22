package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/goobermv/task-tracker-CLI/service"
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
			if err := service.GetTaskByID(); err != nil {
				fmt.Println("Error:", err)
			}

		case 4:
			if err := service.CompleteTask(); err != nil {
				fmt.Println("Error:", err)
			}

		case 5:
			if err := service.DeleteTask(); err != nil {
				fmt.Println("Error:", err)
			}

		default:
			fmt.Println("Unknown option. Try again.")
		}
	}
}
