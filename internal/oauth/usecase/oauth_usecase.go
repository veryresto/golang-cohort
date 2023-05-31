package oauth

import (
	"errors"
	"os"
	"time"

	adminUsecase "online-course/internal/admin/usecase"
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
	adminUsecase                adminUsecase.AdminUsecase
}

// Login implements OauthUsecase
func (usecase *oauthUsecase) Login(dtoLoginRequestBody dto.LoginRequestBody) (*dto.LoginResponse, *response.Error) {
	// Check apakah client_id dan client_secret terdaftar pada database
	oauthClient, err := usecase.oauthClientRepository.FindByClientIDAndClientSecret(dtoLoginRequestBody.ClientID, dtoLoginRequestBody.ClientSecret)

	if err != nil {
		return nil, err
	}

	var user dto.UserResponse

	if oauthClient.Name == "admin" {
		dataAdmin, err := usecase.adminUsecase.FindOneByEmail(dtoLoginRequestBody.Email)

		if err != nil {
			return nil, &response.Error{
				Code: 400,
				Err:  errors.New("username or password is invalid"),
			}
		}

		user.ID = dataAdmin.ID
		user.Email = dataAdmin.Email
		user.Name = dataAdmin.Name
		user.Password = dataAdmin.Password
	} else {
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
	}

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

	if oauthClient.Name == "admin" {
		claims.IsAdmin = true
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
func (usecase *oauthUsecase) Refresh(dtoRefreshToken dto.RefreshTokenRequestBody) (*dto.LoginResponse, *response.Error) {
	// Check oauth refresh token berdasarkan refresh token
	oauthRefreshToken, err := usecase.oauthRefreshTokenRepository.FindOneByToken(dtoRefreshToken.RefreshToken)

	if err != nil {
		return nil, err
	}

	if oauthRefreshToken.ExpiredAt.Before(time.Now()) {
		return nil, &response.Error{
			Code: 400,
			Err:  errors.New("oauth refresh token anda sudah expired"),
		}
	}

	var user dto.UserResponse

	expirationTime := time.Now().Add(24 * 365 * time.Hour)

	if *oauthRefreshToken.OauthAccessToken.OauthClientID == 2 {
		admin, _ := usecase.adminUsecase.FindOneById(int(oauthRefreshToken.UserID))

		user.ID = admin.ID
		user.Name = admin.Name
		user.Email = admin.Email
	} else {
		dataUser, _ := usecase.userUsecase.FindOneById(int(oauthRefreshToken.UserID))

		user.ID = dataUser.ID
		user.Name = dataUser.Name
		user.Email = dataUser.Email
	}

	claims := &dto.ClaimsReponse{
		ID:      user.ID,
		Name:    user.Name,
		Email:   user.Email,
		IsAdmin: false,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	if *oauthRefreshToken.OauthAccessToken.OauthClientID == 2 {
		claims.IsAdmin = true
	}

	jwtKey := []byte(os.Getenv("JWT_SECRET"))

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, errSignedString := token.SignedString(jwtKey)

	if errSignedString != nil {
		return nil, &response.Error{
			Code: 500,
			Err:  errSignedString,
		}
	}

	// Insert ke table oauth access token
	dataOauthAccessToken := entity.OauthAccessToken{
		OauthClientID: oauthRefreshToken.OauthAccessToken.OauthClientID,
		UserID:        oauthRefreshToken.UserID,
		Token:         tokenString,
		Scope:         "*",
		ExpiredAt:     &expirationTime,
	}

	saveOauthAccessToken, err := usecase.oauthAcessTokenRepository.Create(dataOauthAccessToken)

	if err != nil {
		return nil, err
	}

	expirationTimeOauthRefreshToken := time.Now().Add(24 * 366 * time.Hour)

	// Insert ke table oauth refresh token
	dataOauthRefreshToken := entity.OauthRefreshToken{
		OauthAccessTokenID: &saveOauthAccessToken.ID,
		UserID:             oauthRefreshToken.UserID,
		Token:              utils.RandString(128),
		ExpiredAt:          &expirationTimeOauthRefreshToken,
	}

	saveOauthRefreshToken, err := usecase.oauthRefreshTokenRepository.Create(dataOauthRefreshToken)

	if err != nil {
		return nil, err
	}

	// Delete old oauth refresh token
	err = usecase.oauthRefreshTokenRepository.Delete(*oauthRefreshToken)

	if err != nil {
		return nil, err
	}

	// Delete old oauth access token
	usecase.oauthAcessTokenRepository.Delete(*oauthRefreshToken.OauthAccessToken)

	if err != nil {
		return nil, err
	}

	return &dto.LoginResponse{
		AccessToken:  tokenString,
		RefreshToken: saveOauthRefreshToken.Token,
		Type:         "Bearer",
		ExpiredAt:    expirationTime.Format(time.RFC3339),
		Scope:        "*",
	}, nil
}

func NewOauthUsecase(
	oauthClientRepository repository.OauthClientRepository,
	oauthAcessTokenRepository repository.OauthAccessTokenRepository,
	oauthRefreshTokenRepository repository.OauthRefreshTokenRepository,
	userUsecase userUsecase.UserUsecase,
	adminUsecase adminUsecase.AdminUsecase,
) OauthUsecase {
	return &oauthUsecase{
		oauthClientRepository,
		oauthAcessTokenRepository,
		oauthRefreshTokenRepository,
		userUsecase,
		adminUsecase,
	}
}
