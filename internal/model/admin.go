package model

type VerifyRequest struct {
	UserID *int `json:"user_id" validate:"required"`
}

type VerifyResponse struct {
	UserID     int  `json:"user_id"`
	IsVerified bool `json:"is_verified"`
}
