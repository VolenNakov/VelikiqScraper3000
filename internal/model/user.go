package model

type User struct {
	ID         int64  `json:"id"`
	Email      string `json:"email"`
	CreatedAt  string `json:"created_at"`
	IsVerified bool   `json:"is_verified"`
}

type RegisterRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type RegisterResponse struct {
	ID    *int64 `json:"id"`
	Email string `json:"email"`
}

type LoginRequest = RegisterRequest

type LoginResponse struct {
	ID    int64  `json:"id"`
	Email string `json:"email"`
	Token string `json:"token"`
}
