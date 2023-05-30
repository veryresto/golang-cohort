package verification_email

import (
	"time"

	userDto "online-course/internal/user/dto"
	userUsecase "online-course/internal/user/usecase"
	dto "online-course/internal/verification_email/dto"
	response "online-course/pkg/response"
)

type VerificationEmailUsecase interface {
	VerificationCode(dto dto.VerificationEmailRequestBody) *response.Error
}

type verificationEmailUseCase struct {
	userUsecase userUsecase.UserUsecase
}

// VerificationCode implements VerificationEmailUsecase
func (usecase *verificationEmailUseCase) VerificationCode(dto dto.VerificationEmailRequestBody) *response.Error {
	user, err := usecase.userUsecase.FindOneByCodeVerified(dto.CodeVerified)

	if err != nil {
		return err
	}

	timeNow := time.Now()

	dataUpdateUser := userDto.UserUpdateRequestBody{
		EmailVerifiedAt: &timeNow,
	}

	_, err = usecase.userUsecase.Update(int(user.ID), dataUpdateUser)

	if err != nil {
		return err
	}

	return nil
}

func NewVerificationEmailUsecase(userUsecase userUsecase.UserUsecase) VerificationEmailUsecase {
	return &verificationEmailUseCase{userUsecase}
}
