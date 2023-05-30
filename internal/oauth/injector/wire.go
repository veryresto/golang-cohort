//go:build wireinject
// +build wireinject

package oauth

import (
	"github.com/google/wire"
	"gorm.io/gorm"
	handler "online-course/internal/oauth/delivery/http"
	oauthAccessTokenRepository "online-course/internal/oauth/repository"
	oauthClientRepository "online-course/internal/oauth/repository"
	oauthRefreshTokenRepository "online-course/internal/oauth/repository"
	usecase "online-course/internal/oauth/usecase"
	userRepository "online-course/internal/user/repository"
	userUsecase "online-course/internal/user/usecase"
)

func InitializedService(db *gorm.DB) *handler.OauthHandler {
	wire.Build(
		userRepository.NewUserRepository,
		userUsecase.NewUserUseCase,
		oauthAccessTokenRepository.NewOauthAccessTokenRepository,
		oauthClientRepository.NewOauthClientRepository,
		oauthRefreshTokenRepository.NewOauthRefreshTokenRepository,
		handler.NewOauthHandler,
		usecase.NewOauthUsecase,
	)

	return &handler.OauthHandler{}
}
