package forgot_password

import (
	"errors"
	"time"

	dto "online-course/internal/forgot_password/dto"
	entity "online-course/internal/forgot_password/entity"
	repository "online-course/internal/forgot_password/repository"
	userDto "online-course/internal/user/dto"
	userUsecase "online-course/internal/user/usecase"
	mail "online-course/pkg/mail/sendgrid"
	"online-course/pkg/response"
	"online-course/pkg/utils"
)

type ForgotPasswordUsecase interface {
	Create(dto dto.ForgotPasswordRequestBody) (*entity.ForgotPassword, *response.Error)
	Update(dto dto.ForgotPasswordUpdateRequestBody) (*entity.ForgotPassword, *response.Error)
}

type forgotPasswordUsecase struct {
	repository  repository.ForgotPasswordRepository
	userUsecase userUsecase.UserUsecase
	mail        mail.Mail
}

// Create implements ForgotPasswordUsecase
func (usecase *forgotPasswordUsecase) Create(dtoForgotPassword dto.ForgotPasswordRequestBody) (*entity.ForgotPassword, *response.Error) {
	// Check email
	user, err := usecase.userUsecase.FindByEmail(dtoForgotPassword.Email)

	if err != nil {
		return nil, err
	}

	dateTime := time.Now().Add(24 * 1 * time.Hour)

	forgotPassword := entity.ForgotPassword{
		UserID:    &user.ID,
		Valid:     true,
		Code:      utils.RandString(32),
		ExpiredAt: &dateTime,
	}

	saveForgotPasswordData, err := usecase.repository.Create(forgotPassword)

	// Send email
	dataEmailForgotPassword := dto.ForgotPasswordEmailRequestBody{
		SUBJECT: "Kode Forgot Password",
		EMAIL:   user.Email,
		CODE:    forgotPassword.Code,
	}

	go usecase.mail.SendForgotPassword(user.Email, dataEmailForgotPassword)

	if err != nil {
		return nil, err
	}

	return saveForgotPasswordData, nil
}

// Update implements ForgotPasswordUsecase
func (usecase *forgotPasswordUsecase) Update(dto dto.ForgotPasswordUpdateRequestBody) (*entity.ForgotPassword, *response.Error) {
	// Check code
	code, err := usecase.repository.FindOneByCode(dto.Code)

	if err != nil || !code.Valid {
		return nil, &response.Error{
			Code: 400,
			Err:  errors.New("code is invalid"),
		}
	}

	// Search user
	user, err := usecase.userUsecase.FindOneById(int(*code.UserID))

	if err != nil {
		return nil, err
	}

	dataUser := userDto.UserUpdateRequestBody{
		Password: &dto.Password,
	}

	_, err = usecase.userUsecase.Update(int(user.ID), dataUser)

	if err != nil {
		return nil, err
	}

	code.Valid = false

	usecase.repository.Update(*code)

	return code, nil
}

func NewForgotPasswordUsecase(
	repository repository.ForgotPasswordRepository,
	userUsecase userUsecase.UserUsecase,
	mail mail.Mail,
) ForgotPasswordUsecase {
	return &forgotPasswordUsecase{
		repository,
		userUsecase,
		mail,
	}
}
