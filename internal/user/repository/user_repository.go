package user

import (
	entity "online-course/internal/user/entity"
	response "online-course/pkg/response"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindOneById(id int) (*entity.User, *response.Error)
	FindByEmail(email string) (*entity.User, *response.Error)
	Create(entity entity.User) (*entity.User, *response.Error)
	FindOneByCodeVerified(codeVerified string) (*entity.User, *response.Error)
	Update(entity entity.User) (*entity.User, *response.Error)
}

type userRepository struct {
	db *gorm.DB
}

// FindOneById implements UserRepository
func (repository *userRepository) FindOneById(id int) (*entity.User, *response.Error) {
	var user entity.User

	if err := repository.db.First(&user, id).Error; err != nil {
		return nil, &response.Error{
			Code: 500,
			Err:  err,
		}
	}

	return &user, nil
}

// Update implements UserRepository
func (repository *userRepository) Update(entity entity.User) (*entity.User, *response.Error) {
	if err := repository.db.Save(&entity).Error; err != nil {
		return nil, &response.Error{
			Code: 500,
			Err:  err,
		}
	}

	return &entity, nil
}

// FindOneByCodeVerified implements UserRepository
func (repository *userRepository) FindOneByCodeVerified(codeVerified string) (*entity.User, *response.Error) {
	var user entity.User

	if err := repository.db.Where("code_verified = ?", codeVerified).First(&user).Error; err != nil {
		return nil, &response.Error{
			Code: 500,
			Err:  err,
		}
	}

	return &user, nil
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
