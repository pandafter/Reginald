package requests

import (
	"reginald_backend/internal/audit"
	"reginald_backend/pkg/utils"
	"time"
)

func CreateRequest(reqDTO CreateRequestDTO) (*Request, error) {
	req := &Request{
		ID:          utils.GenerateID(),
		Title:       reqDTO.Title,
		Description: reqDTO.Description,
		Status:      "pending",
		CreatedBy:   reqDTO.CreatedBy,
		CreatedAt:   time.Now(),
	}

	err := SaveRequest(req)
	if err != nil {
		return nil, err
	}

	audit.LogAction(req.CreatedBy, "CREATE", "REQUEST", req.ID, 201, 0)
	return req, nil
}

func ListRequests(userEmail, role string) ([]Request, error) {
	if role == "admin" {
		return GetAllRequests()
	}
	return GetRequestsByUserID(userEmail)
}

func ProcessRequest(id, status, adminEmail string) error {
	if !IsValidStatus(status) {
		return ErrInvalidStatus
	}
	err := UpdateRequestStatus(id, status, adminEmail)
	if err == nil {
		audit.LogAction(adminEmail, "PROCESS", "REQUEST", id+" ("+status+")", 200, 0)
	}
	return err
}
