package tests

import (
	"reginald_backend/internal/users"
	"testing"
)

func TestRegisterAndLogin(t *testing.T) {
	// Note: This test requires a running database to pass.
	req := users.CreateUserRequest{
		Name:     "Test User",
		Email:    "test@example.com",
		Password: "password123",
		Role:     "user",
	}

	// Registrar
	user, err := users.CreateUser(req)
	if err != nil {
		t.Skip("Skipping test: database not connected or user already exists")
		return
	}

	if user.Email != req.Email {
		t.Errorf("expected email %s, got %s", req.Email, user.Email)
	}

	// Login
	loggedUser, err := users.Authenticate(req.Email, req.Password)
	if err != nil {
		t.Fatalf("login failed: %v", err)
	}

	if loggedUser.ID != user.ID {
		t.Errorf("expected user ID %s, got %s", user.ID, loggedUser.ID)
	}
}

func TestRegisterDuplicateEmail(t *testing.T) {
	req := users.CreateUserRequest{
		Name:     "Test User",
		Email:    "duplicate@example.com",
		Password: "password123",
		Role:     "user",
	}

	_, _ = users.CreateUser(req)
	_, err := users.CreateUser(req)

	if err == nil || err.Error() != "user already exists" {
		t.Errorf("expected error 'user already exists', got %v", err)
	}
}
