package requests

import (
	"errors"
	"strings"
)

func ValidateCreateRequest(dto CreateRequestDTO) error {
	if strings.TrimSpace(dto.Title) == "" {
		return errors.New("title is required")
	}
	if strings.TrimSpace(dto.Description) == "" {
		return errors.New("description is required")
	}
	return nil
}

func ValidateStatus(status string) bool {
	s := strings.ToLower(status)
	return s == "pending" || s == "review" || s == "approved" || s == "rejected"
}
