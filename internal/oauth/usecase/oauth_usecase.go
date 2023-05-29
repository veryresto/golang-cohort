package oauth

import (
	"errors"
	"os"
	"time"

	dto "online-course/internal/oauth/dto"
	entity "online-course/internal/oauth/entity"
	repository "online-course/internal/oauth/repository"
	userUsecase "online-course/internal/user/usecase"
	response "online-course/pkg/response"
	"online-course/pkg/utils"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type OauthUsecase interface {
	Login(dtoLoginRequestBody dto.LoginRequestBody) (*dto.LoginResponse, *response.Error)
	Refresh(dto dto.RefreshTokenRequestBody) (*dto.LoginResponse, *response.Error)
}

type oauthUsecase struct {
	oauthClientRepository       repository.OauthClientRepository
	oauthAcessTokenRepository   repository.OauthAccessTokenRepository
	oauthRefreshTokenRepository repository.OauthRefreshTokenRepository
	userUsecase                 userUsecase.UserUsecase
}

// Login implements OauthUsecase
func (usecase *oauthUsecase) Login(dtoLoginRequestBody dto.LoginRequestBody) (*dto.LoginResponse, *response.Error) {
	// Check apakah client_id dan client_secret terdaftar pada database
	oauthClient, err := usecase.oauthClientRepository.FindByClientIDAndClientSecret(dtoLoginRequestBody.ClientID, dtoLoginRequestBody.ClientSecret)

	if err != nil {
		return nil, err
	}

	var user dto.UserResponse

	dataUser, err := usecase.userUsecase.FindByEmail(dtoLoginRequestBody.Email)

	if err != nil {
		return nil, &response.Error{
			Code: 400,
			Err:  errors.New("username or password is invalid"),
		}
	}

	user.ID = dataUser.ID
	user.Email = dataUser.Email
	user.Name = dataUser.Name
	user.Password = dataUser.Password

	jwtKey := []byte(os.Getenv("JWT_SECRET"))

	// Compare password apakah sama atau tidak
	errorBycrpt := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(dtoLoginRequestBody.Password))

	if errorBycrpt != nil {
		return nil, &response.Error{
			Code: 400,
			Err:  errors.New("username or password is invalid"),
		}
	}

	expirationTime := time.Now().Add(24 * 365 * time.Hour)

	claims := &dto.ClaimsReponse{
		ID:      user.ID,
		Name:    user.Name,
		Email:   user.Email,
		IsAdmin: false,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString(jwtKey)

	// Insert data ke table oauth access token
	dataOauthAccessToken := entity.OauthAccessToken{
		OauthClientID: &oauthClient.ID,
		UserID:        user.ID,
		Token:         tokenString,
		Scope:         "*",
		ExpiredAt:     &expirationTime,
	}

	oauthAccessToken, err := usecase.oauthAcessTokenRepository.Create(dataOauthAccessToken)

	if err != nil {
		return nil, err
	}

	expirationTimeOauthAccessToken := time.Now().Add(24 * 366 * time.Hour)

	// Insert ke table oauth refresh token
	dataOauthRefreshToken := entity.OauthRefreshToken{
		OauthAccessTokenID: &oauthAccessToken.ID,
		UserID:             user.ID,
		Token:              utils.RandString(128),
		ExpiredAt:          &expirationTimeOauthAccessToken,
	}

	oauthRefreshToken, err := usecase.oauthRefreshTokenRepository.Create(dataOauthRefreshToken)

	if err != nil {
		return nil, err
	}

	return &dto.LoginResponse{
		AccessToken:  oauthAccessToken.Token,
		RefreshToken: oauthRefreshToken.Token,
		Type:         "Bearer",
		ExpiredAt:    expirationTime.Format(time.RFC3339),
		Scope:        "*",
	}, nil
}

// Refresh implements OauthUsecase
func (*oauthUsecase) Refresh(dto dto.RefreshTokenRequestBody) (*dto.LoginResponse, *response.Error) {
	panic("unimplemented")
}

func NewOauthUsecase(
	oauthClientRepository repository.OauthClientRepository,
	oauthAcessTokenRepository repository.OauthAccessTokenRepository,
	oauthRefreshTokenRepository repository.OauthRefreshTokenRepository,
	userUsecase userUsecase.UserUsecase,
) OauthUsecase {
	return &oauthUsecase{
		oauthClientRepository,
		oauthAcessTokenRepository,
		oauthRefreshTokenRepository,
		userUsecase,
	}
}
