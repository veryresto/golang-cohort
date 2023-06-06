package user

import (
	entity "online-course/internal/user/entity"
	response "online-course/pkg/response"
	"online-course/pkg/utils"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindAll(offset int, limit int) []entity.User
	FindOneById(id int) (*entity.User, *response.Error)
	FindByEmail(email string) (*entity.User, *response.Error)
	Create(entity entity.User) (*entity.User, *response.Error)
	FindOneByCodeVerified(codeVerified string) (*entity.User, *response.Error)
	Update(entity entity.User) (*entity.User, *response.Error)
	Delete(entity entity.User) *response.Error
}

type userRepository struct {
	db *gorm.DB
}

// FindAll implements UserRepository
func (repository *userRepository) FindAll(offset int, limit int) []entity.User {
	var users []entity.User

	repository.db.Scopes(utils.Paginate(offset, limit)).Find(&users)

	return users
}

// Delete implements UserRepository
func (repository *userRepository) Delete(entity entity.User) *response.Error {
	if err := repository.db.Delete(&entity).Error; err != nil {
		return &response.Error{
			Code: 500,
			Err:  err,
		}
	}

	return nil
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
