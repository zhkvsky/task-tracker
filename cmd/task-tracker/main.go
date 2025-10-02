package main

import (
	"fmt"
	"github.com/zhkvsky/task-tracker/internal/task"
	"os"
	"strconv"
	"strings"
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
		tasks, err := task.ListTasks("tasks.json")
		if err != nil {
			fmt.Println("Error listing tasks:", err)
			return
		}

		if len(tasks) == 0 {
			fmt.Println("No tasks found.")
			return
		}

		for _, t := range tasks {
			fmt.Printf("[%d] %s (%s)\n", t.ID, t.Description, t.Status)
		}

	case "update":
		if len(os.Args) < 4 {
			fmt.Println("Please provide a id and description for the task")
			return
		}

		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Invalid id:", os.Args[2])
			return
		}

		description := strings.Join(os.Args[3:], " ")

		updatedTask, err := task.UpdateTask("tasks.json", id, description)
		if err != nil {
			fmt.Println("Error updating task:", err)
			return
		}

		fmt.Printf("Task updated successfully: (ID: %d)\n", updatedTask.ID)

	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Please provide an id of the task")
			return
		}

		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Invalid id:", os.Args[2])
			return
		}

		deletedTask, err := task.DeleteTask("tasks.json", id)
		if err != nil {
			fmt.Println("Error deleting task:", err)
			return
		}

		fmt.Printf("Task deleted successfully: (ID: %d, %s)\n", deletedTask.ID, deletedTask.Description)

	case "mark-in-progress":
		fmt.Println("Marking task an in progress:")

	case "mark-done":
		fmt.Println("Marking task as done:")

	default:
		fmt.Println("Unknown command")
	}
}
