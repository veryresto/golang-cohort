package class_room

import (
	"time"

	productEntity "online-course/internal/product/entity"
	userEntity "online-course/internal/user/entity"

	"gorm.io/gorm"
)

type ClassRoom struct {
	ID          int64                  `json:"id"`
	User        *userEntity.User       `json:"user" gorm:"foreignKey:UserID;references:ID"`
	UserID      int64                  `json:"user_id"`
	Product     *productEntity.Product `json:"product" gorm:"foreignKey:ProductID;references:ID"`
	ProductID   *int64                 `json:"product_id"`
	CreatedBy   *userEntity.User       `json:"-" gorm:"foreignKey:CreatedByID;references:ID"`
	CreatedByID *int64                 `json:"created_by" gorm:"column:created_by"`
	UpdateBy    *userEntity.User       `json:"-" gorm:"foreignKey:CreatedByID;references:ID"`
	UpdatedByID *int64                 `json:"updated_by" gorm:"column:updated_by"`
	CreatedAt   *time.Time             `json:"created_at"`
	UpdatedAt   *time.Time             `json:"updated_at"`
	DeletedAt   gorm.DeletedAt         `json:"deleted_at"`
}