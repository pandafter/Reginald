package users

import (
	"log"
	"os"

	"reginald_backend/internal/shared"
	"reginald_backend/pkg/utils"
)

func BootstrapAdmin() {
	email := os.Getenv("ADMIN_EMAIL")
	name := os.Getenv("ADMIN_NAME")
	password := os.Getenv("ADMIN_PASSWORD")

	if email == "" || password == "" {
		log.Fatal("admin env vars missing")
	}

	if _, exists := GetUserByEmail(email); exists {
		log.Println("Admin already exists")
		return
	}

	hash, _ := utils.HashPassword(password)

	admin := &User{
		ID:       "admin-0001",
		Name:     name,
		Email:    email,
		Role:     shared.RoleAdmin,
		Password: hash,
	}

	SaveUser(admin)
	log.Println("Admin user created")
}
