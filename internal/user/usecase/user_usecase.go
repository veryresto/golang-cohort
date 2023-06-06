package user

import (
	"errors"

	dto "online-course/internal/user/dto"
	entity "online-course/internal/user/entity"
	repository "online-course/internal/user/repository"
	response "online-course/pkg/response"
	utils "online-course/pkg/utils"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserUsecase interface {
	FindAll(offset int, limit int) []entity.User
	FindByEmail(email string) (*entity.User, *response.Error)
	FindOneById(id int) (*entity.User, *response.Error)
	Create(dto dto.UserRequestBody) (*entity.User, *response.Error)
	FindOneByCodeVerified(codeVerified string) (*entity.User, *response.Error)
	Update(id int, dto dto.UserUpdateRequestBody) (*entity.User, *response.Error)
	Delete(id int) *response.Error
}

type userUsecase struct {
	repository repository.UserRepository
}

// FindAll implements UserUsecase
func (usecase *userUsecase) FindAll(offset int, limit int) []entity.User {
	return usecase.repository.FindAll(offset, limit)
}

// Delete implements UserUsecase
func (usecase *userUsecase) Delete(id int) *response.Error {
	user, err := usecase.repository.FindOneById(id)

	if err != nil {
		return err
	}

	err = usecase.repository.Delete(*user)

	if err != nil {
		return err
	}

	return nil
}

// FindOneById implements UserUsecase
func (usecase *userUsecase) FindOneById(id int) (*entity.User, *response.Error) {
	return usecase.repository.FindOneById(id)
}

// Update implements UserUsecase
func (usecase *userUsecase) Update(id int, dto dto.UserUpdateRequestBody) (*entity.User, *response.Error) {
	user, err := usecase.repository.FindOneById(id)

	user.Name = dto.Name

	if user.Email != dto.Email {
		user.Email = dto.Email
	}

	if err != nil {
		return nil, err
	}

	if dto.EmailVerifiedAt != nil {
		user.EmailVerifiedAt = dto.EmailVerifiedAt
	}

	if dto.Password != nil {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(*dto.Password), bcrypt.DefaultCost)

		if err != nil {
			return nil, &response.Error{
				Code: 500,
				Err:  err,
			}
		}

		user.Password = string(hashedPassword)
	}

	if dto.UpdatedBy != nil {
		user.UpdatedByID = dto.UpdatedBy
	}

	updateUser, err := usecase.repository.Update(*user)

	if err != nil {
		return nil, err
	}

	return updateUser, nil
}

// FindOneByCodeVerified implements UserUsecase
func (usecase *userUsecase) FindOneByCodeVerified(codeVerified string) (*entity.User, *response.Error) {
	return usecase.repository.FindOneByCodeVerified(codeVerified)
}

// Create implements UserUsecase
func (usecase *userUsecase) Create(dto dto.UserRequestBody) (*entity.User, *response.Error) {
	// Cari berdasarkan email
	checkUser, err := usecase.repository.FindByEmail(dto.Email)

	if err != nil && !errors.Is(err.Err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	if checkUser != nil {
		return nil, &response.Error{
			Code: 409,
			Err:  errors.New("email sudah terdaftar"),
		}
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(dto.Password), bcrypt.DefaultCost)

	// if err != nil {
	// 	return nil, &response.Error{
	// 		Code: 500,
	// 		Err:  err.Error(),
	// 	}
	// }

	user := entity.User{
		Name:         dto.Name,
		Email:        dto.Email,
		Password:     string(hashedPassword),
		CodeVerified: utils.RandString(32),
	}

	if dto.CreatedBy != nil {
		user.CreatedByID = dto.CreatedBy
	}

	dataUser, err := usecase.repository.Create(user)

	if err != nil {
		return nil, err
	}

	return dataUser, nil

}

// FindByEmail implements UserUsecase
func (usecase *userUsecase) FindByEmail(email string) (*entity.User, *response.Error) {
	return usecase.repository.FindByEmail(email)
}

func NewUserUseCase(repository repository.UserRepository) UserUsecase {
	return &userUsecase{repository}
}
