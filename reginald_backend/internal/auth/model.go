package auth

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserInfo struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
	Role  string `json:"role"`
}

type LoginResponse struct {
	Token string   `json:"token"`
	User  UserInfo `json:"user"`
}
