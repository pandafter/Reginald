package users

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Role     string `json:"role"` // admin | user
	Email    string `json:"email"`
	Password string `json:"-"` // Ocultamos el password en las respuestas JSON
}

type CreateUserRequest struct {
	Name     string `json:"name"`
	Role     string `json:"role"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
