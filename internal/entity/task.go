package entity

import "time"

// Task struct represents the 'task' table
type Task struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	Completed   bool      `json:"completed"`
}

type TaskRequest struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	Completed   bool      `json:"completed"`
}

// Deadline struct represents the 'deadline' table
type Deadline struct {
	ID     int       `json:"id"`
	TaskID int       `json:"task_id"`
	DueAt  time.Time `json:"due_at"`
}

// SubTask struct represents the 'sub_task' table
type SubTask struct {
	ID        int    `json:"id"`
	TaskID    int    `json:"task_id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

// TaskProgress struct represents the 'task_progress' table
type TaskProgress struct {
	TaskID     int `json:"task_id"`
	Percentage int `json:"percentage"`
}

// TaskStatus struct represents the 'task_status' table
type TaskStatus struct {
	ID     int    `json:"id"`
	Status string `json:"status"`
}
