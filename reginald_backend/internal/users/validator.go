package users

import (
	"errors"
	"regexp"
)

func ValidateRegistration(req CreateUserRequest) error {
	if req.Email == "" {
		return errors.New("email is required")
	}
	if len(req.Password) < 8 {
		return ErrInvalidPassword
	}

	// Email regex básica
	re := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	if !re.MatchString(req.Email) {
		return errors.New("invalid email format")
	}

	return nil
}

func ValidateLogin(req LoginRequest) error {
	if req.Email == "" || req.Password == "" {
		return ErrInvalidCredentials
	}
	return nil
}

func ValidateCreateUser(req CreateUserRequest) error {
	if req.Name == "" {
		return errors.New("name is required")
	}
	if req.Email == "" {
		return errors.New("email is required")
	}
	if req.Password == "" {
		return errors.New("password is required")
	}
	return nil
}
