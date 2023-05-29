package main

import (
	userRepository "online-course/internal/user/repository"
	userUsecase "online-course/internal/user/usecase"
	mysql "online-course/pkg/db/mysql"

	"github.com/gin-gonic/gin"

	registerHandler "online-course/internal/register/delivery/http"
	registerUsecase "online-course/internal/register/usecase"

	oauthAccessTokenRepository "online-course/internal/oauth/repository"
	oauthClientRepository "online-course/internal/oauth/repository"
	oauthRefreshTokenRepository "online-course/internal/oauth/repository"

	oauthHandler "online-course/internal/oauth/delivery/http"
	oauthUsecase "online-course/internal/oauth/usecase"
)

func main() {
	r := gin.Default()

	db := mysql.DB()

	userRepository := userRepository.NewUserRepository(db)
	userUsecase := userUsecase.NewUserUseCase(userRepository)
	registerUsecase := registerUsecase.NewRegisterUseCase(userUsecase)
	registerHandler.NewRegisterHandler(registerUsecase).Route(&r.RouterGroup)

	oauthClientRepository := oauthClientRepository.NewOauthClientRepository(db)
	oauthAccessTokenRepository := oauthAccessTokenRepository.NewOauthAccessTokenRepository(db)
	oauthRefreshTokenRepository := oauthRefreshTokenRepository.NewOauthRefreshTokenRepository(db)
	oauthUsecase := oauthUsecase.NewOauthUsecase(oauthClientRepository, oauthAccessTokenRepository, oauthRefreshTokenRepository, userUsecase)
	oauthHandler.NewOauthHandler(oauthUsecase).Route(&r.RouterGroup)

	r.Run()
}
