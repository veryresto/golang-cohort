package register

import (
	registerDto "online-course/internal/register/dto"
	userDto "online-course/internal/user/dto"
	userUsecase "online-course/internal/user/usecase"
	mail "online-course/pkg/mail/sendgrid"
	response "online-course/pkg/response"
)

type RegisterUsecase interface {
	Register(dto userDto.UserRequestBody) *response.Error
}

type registerUsecase struct {
	userUsecase userUsecase.UserUsecase
	mail        mail.Mail
}

// Register implements RegisterUsecase
func (usecase *registerUsecase) Register(dto userDto.UserRequestBody) *response.Error {
	user, err := usecase.userUsecase.Create(dto)

	if err != nil {
		return err
	}

	// Kirim email melalui sendgrid
	data := registerDto.EmailVerification{
		SUBJECT:           "Verifikasi Akun",
		EMAIL:             dto.Email,
		VERIFICATION_CODE: user.CodeVerified,
	}

	go usecase.mail.SendVerification(dto.Email, data)

	return nil
}

func NewRegisterUseCase(userUsecase userUsecase.UserUsecase, mail mail.Mail) RegisterUsecase {
	return &registerUsecase{userUsecase, mail}
}
