//go:build wireinject
// +build wireinject

package register

import (
	"github.com/google/wire"
	"gorm.io/gorm"
	handler "online-course/internal/register/delivery/http"
	usecase "online-course/internal/register/usecase"
	userRepository "online-course/internal/user/repository"
	userUsecase "online-course/internal/user/usecase"
	mail "online-course/pkg/mail/sendgrid"
)

func InitializedService(db *gorm.DB) *handler.RegisterHandler {
	wire.Build(
		handler.NewRegisterHandler,
		usecase.NewRegisterUseCase,
		userRepository.NewUserRepository,
		userUsecase.NewUserUseCase,
		mail.NewMail,
	)

	return &handler.RegisterHandler{}
}
