// Estructuras de Task
package tasks

import "time"

type Task struct {
	ID         string    `json:"id"`
	Title      string    `json:"title"`
	Status     string    `json:"status"`
	CreatedBy  string    `json:"created_by"`
	AssignedTo string    `json:"assigned_to"`
	CreatedAt  time.Time `json:"created_at"`
}

type CreateTaskRequest struct {
	Title      string `json:"title"`
	CreatedBy  string `json:"created_by"`  // Admin who creates it
	AssignedTo string `json:"assigned_to"` // User who will perform it
}

type UpdateStatusRequest struct {
	ID     string `json:"id"`
	Status string `json:"status"`
}
