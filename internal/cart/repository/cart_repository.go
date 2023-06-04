package cart

import (
	entity "online-course/internal/cart/entity"
	"online-course/pkg/response"
	"online-course/pkg/utils"

	"gorm.io/gorm"
)

type CartRepository interface {
	FindByUserId(userId int, offset int, limit int) []entity.Cart
	FindOneById(id int) (*entity.Cart, *response.Error)
	Create(entity entity.Cart) (*entity.Cart, *response.Error)
	Update(entity entity.Cart) (*entity.Cart, *response.Error)
	Delete(entity entity.Cart) *response.Error
	DeleteByUserId(userId int) *response.Error
}

type cartRepository struct {
	db *gorm.DB
}

// Create implements CartRepository
func (repository *cartRepository) Create(entity entity.Cart) (*entity.Cart, *response.Error) {
	if err := repository.db.Create(&entity).Error; err != nil {
		return nil, &response.Error{
			Code: 500,
			Err:  err,
		}
	}

	return &entity, nil
}

// Delete implements CartRepository
func (repository *cartRepository) Delete(entity entity.Cart) *response.Error {
	if err := repository.db.Delete(&entity).Error; err != nil {
		return &response.Error{
			Code: 500,
			Err:  err,
		}
	}

	return nil
}

// DeleteOneByUserId implements CartRepository
func (repository *cartRepository) DeleteByUserId(userId int) *response.Error {
	var cart entity.Cart

	if err := repository.db.Where("user_id = ?", userId).Delete(&cart).Error; err != nil {
		return &response.Error{
			Code: 500,
			Err:  err,
		}
	}

	return nil
}

// FindByUserId implements CartRepository
func (repository *cartRepository) FindByUserId(userId int, offset int, limit int) []entity.Cart {
	var carts []entity.Cart

	repository.db.Scopes(utils.Paginate(offset, limit)).
		Preload("User").Preload("Product").
		Where("user_id = ?", userId).
		Find(&carts)

	return carts
}

// FindOneById implements CartRepository
func (repository *cartRepository) FindOneById(id int) (*entity.Cart, *response.Error) {
	var cart entity.Cart

	if err := repository.db.Preload("User").Preload("Product").Find(&cart, id).Error; err != nil {
		return nil, &response.Error{
			Code: 500,
			Err:  err,
		}
	}

	return &cart, nil
}

// Update implements CartRepository
func (repository *cartRepository) Update(entity entity.Cart) (*entity.Cart, *response.Error) {
	if err := repository.db.Save(&entity).Error; err != nil {
		return nil, &response.Error{
			Code: 500,
			Err:  err,
		}
	}

	return &entity, nil
}

func NewCartRepository(db *gorm.DB) CartRepository {
	return &cartRepository{db}
}
