package model

type User struct {
	ID         int64  `json:"id"`
	Username   string `json:"username"`
	CreatedAt  string `json:"created_at"`
	IsVerified bool   `json:"is_verified"`
}

type RegisterRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required,min=6"`
}

type RegisterResponse struct {
	ID       *int64 `json:"id"`
	Username string `json:"username"`
}

type LoginRequest = RegisterRequest

type LoginResponse struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Token    string `json:"token"`
}
