package users

import (
	"errors"

	"reginald_backend/internal/audit"
	"reginald_backend/internal/shared"
	"reginald_backend/pkg/utils"
)

func CreateUser(req CreateUserRequest) (*User, error) {
	if _, exists := GetUserByEmail(req.Email); exists {
		return nil, errors.New("user already exists")
	}

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	user := &User{
		ID:       utils.GenerateID(),
		Name:     req.Name,
		Email:    req.Email,
		Role:     shared.RoleUser,
		Password: hashedPassword,
	}

	SaveUser(user)
	audit.LogAction(user.Email, "REGISTER", "USER", user.ID, 201, 0)
	return user, nil
}

func Authenticate(email, password string) (*User, error) {
	user, exists := GetUserByEmail(email)
	if !exists {
		return nil, errors.New("invalid credentials")
	}

	if !utils.CheckPassword(password, user.Password) {
		return nil, errors.New("invalid credentials")
	}

	return user, nil
}
