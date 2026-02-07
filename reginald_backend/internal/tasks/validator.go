// Validaciones
package tasks

import "errors"

func ValidateCreateTask(req CreateTaskRequest) error {
	if req.Title == "" {
		return errors.New("title is required")
	}
	if req.CreatedBy == "" {
		return errors.New("created_by is required")
	}
	return nil
}

func IsValidStatus(status string) bool {
	validStatuses := map[string]bool{
		"pending":     true,
		"in-progress": true,
		"done":        true,
	}

	return validStatuses[status]
}
