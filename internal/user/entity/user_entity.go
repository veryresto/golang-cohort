package user

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID              int64      `json:"id"`
	Name            string     `json:"name"`
	Email           string     `json:"email"`
	Password        string     `json:"-"`
	CodeVerified    string     `json:"-"`
	EmailVerifiedAt *time.Time `json:"email_verified_at"`
	// CreatedByID     *int64         `json:"created_by" gorm:"coloumn:created_by"`
	// CreatedBy       string         `json:"-"`
	// UpdatedByID     *int64         `json:"updated_by" gorm:"coloumn:updated_by"`
	// UpdateddBy      string         `json:"-"`
	CreatedAt *time.Time     `json:"created_at"`
	UpdatedAt *time.Time     `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
