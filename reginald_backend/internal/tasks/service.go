// LOGICA DE NEGOCIO PARA TAREAS
package tasks

import (
	"time"

	"reginald_backend/internal/audit"
	"reginald_backend/pkg/utils"
)

func CreateTask(req CreateTaskRequest) (*Task, error) {
	task := &Task{
		ID:         utils.GenerateID(),
		Title:      req.Title,
		Status:     "pending",
		CreatedBy:  req.CreatedBy,
		AssignedTo: req.AssignedTo,
		CreatedAt:  time.Now(),
	}

	err := SaveTask(task)
	if err != nil {
		return nil, err
	}

	audit.LogAction(task.CreatedBy, "CREATE", "TASK", task.ID, 201, 0)
	return task, nil
}

func ListTasks(userID string) ([]Task, error) {
	if userID == "" {
		return GetAllTasks()
	}
	return GetTasksByUserID(userID)
}

func UpdateTaskByID(id, title, status, userEmail string) error {
	if title == "" {
		task, err := GetTaskByID(id)
		if err != nil {
			return err
		}
		title = task.Title
	}

	if status != "" && !IsValidStatus(status) {
		return ErrInvalidStatus
	}

	err := UpdateTask(id, title, status)
	if err == nil {
		audit.LogAction(userEmail, "UPDATE", "TASK", id, 200, 0)
	}
	return err
}

func ChangeTaskStatus(taskID, status, userEmail string) error {
	if !IsValidStatus(status) {
		return ErrInvalidStatus
	}

	err := UpdateTask(taskID, "", status)
	if err == nil {
		audit.LogAction(userEmail, "STATUS_CHANGE", "TASK", taskID+" ("+status+")", 200, 0)
	}
	return err
}

func DeleteTaskByID(id, userEmail string) error {
	err := DeleteTask(id)
	if err == nil {
		audit.LogAction(userEmail, "DELETE", "TASK", id, 204, 0)
	}
	return err
}
