//go:build wireinject
// +build wireinject

package verification_email

import (
	"github.com/google/wire"
	"gorm.io/gorm"
	userRepository "online-course/internal/user/repository"
	userUsecase "online-course/internal/user/usecase"
	handler "online-course/internal/verification_email/delivery/http"
	usecase "online-course/internal/verification_email/usecase"
)

func InitializedService(db *gorm.DB) *handler.VerificationEmailHandler {
	wire.Build(
		handler.NewVerificationEmailHandler,
		usecase.NewVerificationEmailUsecase,
		userRepository.NewUserRepository,
		userUsecase.NewUserUseCase,
	)

	return &handler.VerificationEmailHandler{}
}
