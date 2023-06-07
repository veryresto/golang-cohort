package profile

import (
	userEntity "online-course/internal/user/entity"
)

type ProfileResponseBody struct {
	ID         int64  `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	IsVerified bool   `json:"is_verified"`
}

func CreateProfileResponse(user userEntity.User) ProfileResponseBody {
	isVerified := false

	if user.EmailVerifiedAt != nil {
		isVerified = true
	}

	return ProfileResponseBody{
		ID:         user.ID,
		Name:       user.Name,
		Email:      user.Email,
		IsVerified: isVerified,
	}
}
