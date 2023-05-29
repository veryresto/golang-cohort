package oauth

import (
	"gorm.io/gorm"
	entity "online-course.faerul.com/internal/oauth/entity"
	response "online-course.faerul.com/pkg/response"
)

type OauthClientRepository interface {
	FindByClientIDAndClientSecret(clientID string, clientSecret string) (*entity.OauthClient, *response.Error)
}

type oauthClientRepository struct {
	db *gorm.DB
}

// FindByClientIDAndClientSecret implements OauthClientRepository
func (repository *oauthClientRepository) FindByClientIDAndClientSecret(clientID string, clientSecret string) (*entity.OauthClient, *response.Error) {
	var oauthClient entity.OauthClient

	if err := repository.db.Where("client_id = ?", clientID).Where("client_secret = ?", clientSecret).First(&oauthClient).Error; err != nil {
		return nil, &response.Error{
			Code: 500,
			Err:  err,
		}
	}

	return &oauthClient, nil
}

func NewOauthClientRepository(db *gorm.DB) OauthClientRepository {
	return &oauthClientRepository{db}
}
