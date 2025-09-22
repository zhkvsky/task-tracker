package task

import (
	"encoding/json"
	"os"
	"time"
)

type TaskStatus string

const (
	TaskStatusTodo       TaskStatus = "todo"
	TaskStatusInProgress TaskStatus = "in-progress"
	TaskStatusDone       TaskStatus = "done"
)

type Task struct {
	ID          int        `json:"id"`
	Description string     `json:"description"`
	Status      TaskStatus `json:"status"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

func LoadTask(filename string) ([]Task, error) {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return []Task{}, nil
	}

	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	tasks := []Task{}
	if err := json.Unmarshal(data, &tasks); err != nil {
		return nil, err
	}

	return tasks, nil
}

func SaveTasks(filename string, tasks []Task) error {
	data, err := json.Marshal(tasks)
	if err != nil {
		return err
	}

	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		return err
	}

	return nil
}

func AddTask(filename string, description string) (Task, error) {
	tasks, err := LoadTask(filename)
	if err != nil {
		return Task{}, err
	}

	newID := len(tasks) + 1
	newTask := Task{
		ID:          newID,
		Description: description,
		Status:      TaskStatusTodo,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	tasks = append(tasks, newTask)

	err = SaveTasks(filename, tasks)
	if err != nil {
		return newTask, err
	}

	return newTask, err
}
