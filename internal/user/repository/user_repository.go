package user

import (
	"gorm.io/gorm"
	entity "online-course/internal/user/entity"
	response "online-course/pkg/response"
)

type UserRepository interface {
	FindByEmail(email string) (*entity.User, *response.Error)
	Create(entity entity.User) (*entity.User, *response.Error)
}

type userRepository struct {
	db *gorm.DB
}

// Create implements UserRepository
func (repository *userRepository) Create(entity entity.User) (*entity.User, *response.Error) {
	if err := repository.db.Create(&entity).Error; err != nil {
		return nil, &response.Error{
			Code: 500,
			Err:  err,
		}
	}

	return &entity, nil
}

// FindByEmail implements UserRepository
func (repository *userRepository) FindByEmail(email string) (*entity.User, *response.Error) {
	var user entity.User

	if err := repository.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, &response.Error{
			Code: 500,
			Err:  err,
		}
	}

	return &user, nil
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}
