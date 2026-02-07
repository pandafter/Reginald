package requests

import (
	"log"
	"reginald_backend/pkg/database"
)

func SaveRequest(req *Request) error {
	query := `INSERT INTO requests (id, title, description, status, created_by, created_at) 
	          VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := database.DB.Exec(query, req.ID, req.Title, req.Description, req.Status, req.CreatedBy, req.CreatedAt)
	if err != nil {
		log.Printf("Error saving request: %v", err)
		return err
	}
	return nil
}

func GetAllRequests() ([]Request, error) {
	query := `SELECT id, title, description, status, created_by, COALESCE(handled_by, ''), created_at FROM requests ORDER BY created_at DESC`
	rows, err := database.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []Request
	for rows.Next() {
		var r Request
		err := rows.Scan(&r.ID, &r.Title, &r.Description, &r.Status, &r.CreatedBy, &r.HandledBy, &r.CreatedAt)
		if err != nil {
			return nil, err
		}
		list = append(list, r)
	}
	return list, nil
}

func GetRequestsByUserID(userID string) ([]Request, error) {
	query := `SELECT id, title, description, status, created_by, COALESCE(handled_by, ''), created_at FROM requests WHERE created_by = $1 ORDER BY created_at DESC`
	rows, err := database.DB.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []Request
	for rows.Next() {
		var r Request
		err := rows.Scan(&r.ID, &r.Title, &r.Description, &r.Status, &r.CreatedBy, &r.HandledBy, &r.CreatedAt)
		if err != nil {
			return nil, err
		}
		list = append(list, r)
	}
	return list, nil
}

func UpdateRequestStatus(id, status, handledBy string) error {
	query := `UPDATE requests SET status = $1, handled_by = $2 WHERE id = $3`
	_, err := database.DB.Exec(query, status, handledBy, id)
	return err
}

func DeleteRequest(id string) error {
	query := `DELETE FROM requests WHERE id = $1`
	_, err := database.DB.Exec(query, id)
	if err != nil {
		log.Printf("Error deleting request: %v", err)
	}
	return err
}
