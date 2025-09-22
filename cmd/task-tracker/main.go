package main

import (
	"fmt"
	"github.com/zhkvsky/task-tracker/internal/task"
	"os"
)

func main() {
	fmt.Println("Task tracker 1.0")

	if len(os.Args) < 2 {
		fmt.Println("Please provide a task name. Type \"Help\" for more information.")
		return
	}

	command := os.Args[1]

	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Please provide a description for the task")
			return
		}
		description := os.Args[2]
		newTask, err := task.AddTask("tasks.json", description)
		if err != nil {
			fmt.Println("Error adding task:", err)
			return
		}
		fmt.Printf("Task added successfully: (ID: %d)\n", newTask.ID)
	case "list":
		fmt.Println("Listing all tasks:")

	case "update":
		fmt.Println("Updating task:")

	case "delete":
		fmt.Println("Deleting task:")

	case "mark-in-progress":
		fmt.Println("Marking task an in progress:")

	case "mark-done":
		fmt.Println("Marking task as done:")

	default:
		fmt.Println("Unknown command")
	}
}
