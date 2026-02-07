package users

import (
	"log"
	"reginald_backend/pkg/database"
)

func SaveUser(user *User) {
	query := `INSERT INTO users (id, name, email, password, role) 
	          VALUES ($1, $2, $3, $4, $5)
	          ON CONFLICT (email) DO UPDATE 
	          SET name = EXCLUDED.name, password = EXCLUDED.password, role = EXCLUDED.role`
	_, err := database.DB.Exec(query, user.ID, user.Name, user.Email, user.Password, user.Role)
	if err != nil {
		log.Printf("Error saving user: %v", err)
	}
}

func GetUserByEmail(email string) (*User, bool) {
	query := `SELECT id, name, email, password, role FROM users WHERE email = $1`
	row := database.DB.QueryRow(query, email)

	var user User
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Role)
	if err != nil {
		return nil, false
	}
	return &user, true
}

func GetAllUsers() []*User {
	query := `SELECT id, name, email, password, role FROM users`
	rows, err := database.DB.Query(query)
	if err != nil {
		log.Printf("Error getting all users: %v", err)
		return nil
	}
	defer rows.Close()

	var list []*User
	for rows.Next() {
		var u User
		if err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.Password, &u.Role); err != nil {
			continue
		}
		list = append(list, &u)
	}
	return list
}
