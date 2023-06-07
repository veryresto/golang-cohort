package product

import (
	entity "online-course/internal/product/entity"
	"online-course/pkg/response"
	"online-course/pkg/utils"

	"gorm.io/gorm"
)

type ProductRepository interface {
	FindAll(offset int, limit int) []entity.Product
	FindOneById(id int) (*entity.Product, *response.Error)
	Create(entity entity.Product) (*entity.Product, *response.Error)
	Update(entity entity.Product) (*entity.Product, *response.Error)
	Delete(entity entity.Product) *response.Error
	TotalCountProduct() int64
}

type productRepository struct {
	db *gorm.DB
}

// TotalCountProduct implements ProductRepository
func (repository *productRepository) TotalCountProduct() int64 {
	var product entity.Product

	var totalProduct int64

	repository.db.Model(&product).Count(&totalProduct)

	return totalProduct
}

// Create implements ProductRepository
func (repository *productRepository) Create(entity entity.Product) (*entity.Product, *response.Error) {
	if err := repository.db.Create(&entity).Error; err != nil {
		return nil, &response.Error{
			Code: 500,
			Err:  err,
		}
	}

	return &entity, nil
}

// Delete implements ProductRepository
func (repository *productRepository) Delete(entity entity.Product) *response.Error {
	if err := repository.db.Delete(&entity).Error; err != nil {
		return &response.Error{
			Code: 500,
			Err:  err,
		}
	}

	return nil
}

// FindAll implements ProductRepository
func (repository *productRepository) FindAll(offset int, limit int) []entity.Product {
	var products []entity.Product

	repository.db.Scopes(utils.Paginate(offset, limit)).Find(&products)

	return products
}

// FindOneById implements ProductRepository
func (repository *productRepository) FindOneById(id int) (*entity.Product, *response.Error) {
	var product entity.Product

	if err := repository.db.First(&product, id).Error; err != nil {
		return nil, &response.Error{
			Code: 500,
			Err:  err,
		}
	}

	return &product, nil
}

// Update implements ProductRepository
func (repository *productRepository) Update(entity entity.Product) (*entity.Product, *response.Error) {
	if err := repository.db.Save(&entity).Error; err != nil {
		return nil, &response.Error{
			Code: 500,
			Err:  err,
		}
	}

	return &entity, nil
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db}
}
