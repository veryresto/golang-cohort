package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	dto "online-course/internal/user/dto"
	entity "online-course/internal/user/entity"
	repository "online-course/internal/user/repository"
	response "online-course/pkg/response"
	utils "online-course/pkg/utils"
)

type UserUsecase interface {
	FindByEmail(email string) (*entity.User, *response.Error)
	Create(dto dto.UserRequestBody) (*entity.User, *response.Error)
}

type userUsecase struct {
	repository repository.UserRepository
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
