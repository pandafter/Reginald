// Acceso a datos
package tasks

import (
	"database/sql"
	"log"
	"reginald_backend/pkg/database"
)

func SaveTask(task *Task) error {
	query := `INSERT INTO tasks (id, title, status, created_by, assigned_to) 
	          VALUES ($1, $2, $3, $4, $5)`
	_, err := database.DB.Exec(query, task.ID, task.Title, task.Status, task.CreatedBy, task.AssignedTo)
	if err != nil {
		log.Printf("Error saving task: %v", err)
		return err
	}
	return nil
}

func GetTaskByID(id string) (*Task, error) {
	query := `SELECT id, title, status, created_by, COALESCE(assigned_to, ''), created_at FROM tasks WHERE id = $1`
	row := database.DB.QueryRow(query, id)

	var task Task
	err := row.Scan(&task.ID, &task.Title, &task.Status, &task.CreatedBy, &task.AssignedTo, &task.CreatedAt)
	if err != nil {
		return nil, ErrTaskNotFound
	}
	return &task, nil
}

func GetAllTasks() ([]Task, error) {
	query := `SELECT id, title, status, created_by, COALESCE(assigned_to, ''), created_at FROM tasks ORDER BY created_at DESC`
	rows, err := database.DB.Query(query)
	if err != nil {
		log.Printf("Error querying all tasks: %v", err)
		return nil, err
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var t Task
		err := rows.Scan(&t.ID, &t.Title, &t.Status, &t.CreatedBy, &t.AssignedTo, &t.CreatedAt)
		if err != nil {
			log.Printf("Error scanning task row: %v", err)
			return nil, err
		}
		tasks = append(tasks, t)
	}
	return tasks, nil
}

func GetTasksByUserID(userID string) ([]Task, error) {
	// Filter by assigned_to so each user sees their tasks
	query := `SELECT id, title, status, created_by, COALESCE(assigned_to, ''), created_at FROM tasks WHERE assigned_to = $1 OR created_by = $1 ORDER BY created_at DESC`
	rows, err := database.DB.Query(query, userID)
	if err != nil {
		log.Printf("Error querying tasks for user %s: %v", userID, err)
		return nil, err
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var t Task
		err := rows.Scan(&t.ID, &t.Title, &t.Status, &t.CreatedBy, &t.AssignedTo, &t.CreatedAt)
		if err != nil {
			log.Printf("Error scanning task row for user: %v", err)
			return nil, err
		}
		tasks = append(tasks, t)
	}
	return tasks, nil
}

func UpdateTask(id, title, status string) error {
	query := `UPDATE tasks SET title = $1, status = $2 WHERE id = $3`
	result, err := database.DB.Exec(query, title, status, id)
	if err != nil {
		return err
	}
	rows, _ := result.RowsAffected()
	if rows == 0 {
		return ErrTaskNotFound
	}
	return nil
}

func DeleteTask(id string) error {
	query := `DELETE FROM tasks WHERE id = $1`
	result, err := database.DB.Exec(query, id)
	if err != nil {
		return err
	}
	rows, _ := result.RowsAffected()
	if rows == 0 {
		return ErrTaskNotFound
	}
	return nil
}

// Ensure sql import is used if needed for NullString in future
var _ = sql.NullString{}
