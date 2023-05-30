//go:build wireinject
// +build wireinject

package forgot_password

import (
	"github.com/google/wire"
	"gorm.io/gorm"
	handler "online-course/internal/forgot_password/delivery/http"
	repository "online-course/internal/forgot_password/repository"
	usecase "online-course/internal/forgot_password/usecase"
	userRepository "online-course/internal/user/repository"
	userUsecase "online-course/internal/user/usecase"
	mail "online-course/pkg/mail/sendgrid"
)

func InitializedService(db *gorm.DB) *handler.ForgotPasswordHandler {
	wire.Build(
		handler.NewForgotPasswordHandler,
		repository.NewForgotPasswordRepository,
		usecase.NewForgotPasswordUsecase,
		userRepository.NewUserRepository,
		userUsecase.NewUserUseCase,
		mail.NewMail,
	)

	return &handler.ForgotPasswordHandler{}
}
