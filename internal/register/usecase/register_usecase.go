package register

import (
	userDto "online-course/internal/user/dto"
	userUsecase "online-course/internal/user/usecase"
	response "online-course/pkg/response"
)

type RegisterUsecase interface {
	Register(dto userDto.UserRequestBody) *response.Error
}

type registerUsecase struct {
	userUsecase userUsecase.UserUsecase
}

// Register implements RegisterUsecase
func (usecase *registerUsecase) Register(dto userDto.UserRequestBody) *response.Error {
	_, err := usecase.userUsecase.Create(dto)

	if err != nil {
		return err
	}

	// Kirim email melalui sendgrid

	return nil
}

func NewRegisterUseCase(userUsecase userUsecase.UserUsecase) RegisterUsecase {
	return &registerUsecase{userUsecase}
}
