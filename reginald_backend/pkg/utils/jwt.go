package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func getSecret() []byte {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		return []byte("reginald-secret-fallback-change-me")
	}
	return []byte(secret)
}

func GenerateJWT(id, name, email, role string) string {
	claims := jwt.MapClaims{
		"id":    id,
		"name":  name,
		"email": email,
		"role":  role,
		"exp":   time.Now().Add(time.Hour * 24).Unix(), // Increased to 24h as per cookie
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, _ := token.SignedString(getSecret())
	return signed
}

func ValidateJWT(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return getSecret(), nil
	})

	if err != nil {
		return nil, err
	}

	return token.Claims.(jwt.MapClaims), nil
}
