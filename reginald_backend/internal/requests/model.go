package requests

import (
	"time"
)

type Request struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      string    `json:"status"`     // pending, review, approved, rejected
	CreatedBy   string    `json:"created_by"` // user email
	HandledBy   string    `json:"handled_by"` // admin email
	CreatedAt   time.Time `json:"created_at"`
}

type CreateRequestDTO struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	CreatedBy   string `json:"created_by"`
}

func IsValidStatus(status string) bool {
	switch status {
	case "pending", "review", "approved", "rejected":
		return true
	}
	return false
}
