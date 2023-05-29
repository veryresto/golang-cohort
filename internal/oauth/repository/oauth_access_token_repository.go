package oauth

import (
	"gorm.io/gorm"
	entity "online-course.faerul.com/internal/oauth/entity"
	response "online-course.faerul.com/pkg/response"
)

type OauthAccessTokenRepository interface {
	Create(entity entity.OauthAccessToken) (*entity.OauthAccessToken, *response.Error)
	Delete(id int) *response.Error
}

type oauthAccessTokenRepository struct {
	db *gorm.DB
}

// Create implements OauthAccessTokenRepository
func (repository *oauthAccessTokenRepository) Create(entity entity.OauthAccessToken) (*entity.OauthAccessToken, *response.Error) {
	if err := repository.db.Create(&entity).Error; err != nil {
		return nil, &response.Error{
			Code: 500,
			Err:  err,
		}
	}

	return &entity, nil
}

// Delete implements OauthAccessTokenRepository
func (repository *oauthAccessTokenRepository) Delete(id int) *response.Error {
	var oauthAccessToken entity.OauthAccessToken

	if err := repository.db.Delete(&oauthAccessToken, id).Error; err != nil {
		return &response.Error{
			Code: 500,
			Err:  err,
		}
	}

	return nil
}

func NewOauthAccessTokenRepository(db *gorm.DB) OauthAccessTokenRepository {
	return &oauthAccessTokenRepository{db}
}
