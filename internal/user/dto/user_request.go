package user

import "time"

type UserRequestBody struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"email"`
	Password string `json:"password" binding:"required"`
}

type UserUpdateRequestBody struct {
	EmailVerifiedAt *time.Time `json:"email_verified_at"`
	Password        *string    `json:"password"`
}
