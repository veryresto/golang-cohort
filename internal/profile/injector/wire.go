//go:build wireinject
// +build wireinject

package profile

import (
	"github.com/google/wire"
	"gorm.io/gorm"
	adminRepository "online-course/internal/admin/repository"
	adminUsecase "online-course/internal/admin/usecase"
	oauthAccessTokenRepository "online-course/internal/oauth/repository"
	oauthClientRepository "online-course/internal/oauth/repository"
	oauthRefreshTokenRepository "online-course/internal/oauth/repository"
	oauthUsecase "online-course/internal/oauth/usecase"
	handler "online-course/internal/profile/delivery/http"
	usecase "online-course/internal/profile/usecase"
	userRepository "online-course/internal/user/repository"
	userUsecase "online-course/internal/user/usecase"
)

func InitializedService(db *gorm.DB) *handler.ProfileHandler {
	wire.Build(
		adminRepository.NewAdminRepository,
		adminUsecase.NewAdminUsecase,
		oauthClientRepository.NewOauthClientRepository,
		oauthAccessTokenRepository.NewOauthAccessTokenRepository,
		oauthRefreshTokenRepository.NewOauthRefreshTokenRepository,
		oauthUsecase.NewOauthUsecase,
		usecase.NewProfileUsecase,
		userRepository.NewUserRepository,
		userUsecase.NewUserUseCase,
		handler.NewProfileHandler,
	)

	return &handler.ProfileHandler{}
}
