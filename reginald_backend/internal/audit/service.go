package audit

import (
	"log"
	"reginald_backend/pkg/database"
	"reginald_backend/pkg/utils"
	"time"
)

type AuditLog struct {
	ID         string    `json:"id"`
	UserEmail  string    `json:"user_email"`
	Action     string    `json:"action"`
	Resource   string    `json:"resource"`
	ResourceID string    `json:"resource_id"`
	StatusCode int       `json:"status_code"`
	DurationMS int       `json:"duration_ms"`
	Timestamp  time.Time `json:"timestamp"`
}

func LogAction(userEmail, action, resource, resourceID string, statusCode, duration int) {
	entry := AuditLog{
		ID:         utils.GenerateID(),
		UserEmail:  userEmail,
		Action:     action,
		Resource:   resource,
		ResourceID: resourceID,
		StatusCode: statusCode,
		DurationMS: duration,
		Timestamp:  time.Now(),
	}

	query := `INSERT INTO audit_logs (id, user_email, action, resource, resource_id, status_code, duration_ms, timestamp) 
	          VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`

	_, err := database.DB.Exec(query, entry.ID, entry.UserEmail, entry.Action, entry.Resource, entry.ResourceID, entry.StatusCode, entry.DurationMS, entry.Timestamp)
	if err != nil {
		log.Printf("Failed to save audit log: %v", err)
	}
}

func GetRecentLogs(limit int) ([]AuditLog, error) {
	query := `SELECT id, user_email, action, resource, COALESCE(resource_id, ''), status_code, duration_ms, timestamp 
	          FROM audit_logs ORDER BY timestamp DESC LIMIT $1`

	rows, err := database.DB.Query(query, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var logs []AuditLog
	for rows.Next() {
		var l AuditLog
		err := rows.Scan(&l.ID, &l.UserEmail, &l.Action, &l.Resource, &l.ResourceID, &l.StatusCode, &l.DurationMS, &l.Timestamp)
		if err != nil {
			return nil, err
		}
		logs = append(logs, l)
	}
	return logs, nil
}

func GetSecurityAlerts(limit int) ([]AuditLog, error) {
	query := `SELECT id, user_email, action, resource, COALESCE(resource_id, ''), status_code, duration_ms, timestamp 
	          FROM audit_logs WHERE action = 'SECURITY_ALERT' ORDER BY timestamp DESC LIMIT $1`

	rows, err := database.DB.Query(query, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var logs []AuditLog
	for rows.Next() {
		var l AuditLog
		err := rows.Scan(&l.ID, &l.UserEmail, &l.Action, &l.Resource, &l.ResourceID, &l.StatusCode, &l.DurationMS, &l.Timestamp)
		if err != nil {
			return nil, err
		}
		logs = append(logs, l)
	}
	return logs, nil
}
