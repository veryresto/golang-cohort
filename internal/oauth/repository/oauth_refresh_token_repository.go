package oauth

import (
	entity "online-course/internal/oauth/entity"
	response "online-course/pkg/response"

	"gorm.io/gorm"
)

type OauthRefreshTokenRepository interface {
	Create(entity entity.OauthRefreshToken) (*entity.OauthRefreshToken, *response.Error)
	FindOneByToken(token string) (*entity.OauthRefreshToken, *response.Error)
	FindOneByOauthAccessTokenId(oauthAccessTokenId int) (*entity.OauthRefreshToken, *response.Error)
	Delete(entity entity.OauthRefreshToken) *response.Error
}

type oauthRefreshTokenRepository struct {
	db *gorm.DB
}

// FindOneByOauthAccessTokenId implements OauthRefreshTokenRepository
func (repository *oauthRefreshTokenRepository) FindOneByOauthAccessTokenId(oauthAccessTokenId int) (*entity.OauthRefreshToken, *response.Error) {
	var oauthRefreshToken entity.OauthRefreshToken

	if err := repository.db.Where("oauth_access_token_id = ?", oauthAccessTokenId).
		First(&oauthRefreshToken).
		Error; err != nil {
		return nil, &response.Error{
			Code: 500,
			Err:  err,
		}
	}

	return &oauthRefreshToken, nil
}

// Create implements OauthRefreshTokenRepository
func (repository *oauthRefreshTokenRepository) Create(entity entity.OauthRefreshToken) (*entity.OauthRefreshToken, *response.Error) {
	if err := repository.db.Create(&entity).Error; err != nil {
		return nil, &response.Error{
			Code: 500,
			Err:  err,
		}
	}

	return &entity, nil
}

// Delete implements OauthRefreshTokenRepository
func (repository *oauthRefreshTokenRepository) Delete(entity entity.OauthRefreshToken) *response.Error {
	if err := repository.db.Delete(&entity).Error; err != nil {
		return &response.Error{
			Code: 500,
			Err:  err,
		}
	}

	return nil
}

// FindOneByToken implements OauthRefreshTokenRepository
func (repository *oauthRefreshTokenRepository) FindOneByToken(token string) (*entity.OauthRefreshToken, *response.Error) {
	var oauthRefreshToken entity.OauthRefreshToken

	if err := repository.db.Preload("OauthAccessToken").Where("token = ?", token).First(&oauthRefreshToken).Error; err != nil {
		return nil, &response.Error{
			Code: 500,
			Err:  err,
		}
	}

	return &oauthRefreshToken, nil
}

func NewOauthRefreshTokenRepository(db *gorm.DB) OauthRefreshTokenRepository {
	return &oauthRefreshTokenRepository{db}
}
