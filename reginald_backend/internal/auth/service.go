package auth

import (
	"errors"
	"reginald_backend/internal/users"
	"reginald_backend/pkg/utils"
)

func Login(email, password string) (string, *users.User, error) {
	user, err := users.Authenticate(email, password)
	if err != nil {
		return "", nil, errors.New("invalid credentials")
	}

	token := utils.GenerateJWT(user.ID, user.Name, user.Email, user.Role)
	return token, user, nil
}
