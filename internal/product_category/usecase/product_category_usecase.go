package product_category

import (
	dto "online-course/internal/product_category/dto"
	entity "online-course/internal/product_category/entity"
	repository "online-course/internal/product_category/repository"
	media "online-course/pkg/media/cloudinary"
	"online-course/pkg/response"
)

type ProductCategoryUsecase interface {
	FindAll(offset int, limit int) []entity.ProductCategory
	FindOneById(id int) (*entity.ProductCategory, *response.Error)
	Create(dto dto.ProductCategoryRequestBody) (*entity.ProductCategory, *response.Error)
	Delete(id int) (*entity.ProductCategory, *response.Error)
	Update(id int, dto dto.ProductCategoryRequestBody) (*entity.ProductCategory, *response.Error)
}

type productCategoryUsecase struct {
	repository repository.ProductCategoryRepository
	media      media.Media
}

// Create implements ProductCategoryUsecase
func (usecase *productCategoryUsecase) Create(dto dto.ProductCategoryRequestBody) (*entity.ProductCategory, *response.Error) {
	entity := entity.ProductCategory{
		Name:        dto.Name,
		CreatedByID: dto.CreatedBy,
	}

	if dto.Image != nil {

		// Upload image
		image, err := usecase.media.Upload(*dto.Image)

		if err != nil {
			return nil, &response.Error{
				Code: 500,
				Err:  err.Err,
			}
		}

		if image != nil {
			entity.Image = image
		}
	}

	data, err := usecase.repository.Create(entity)

	if err != nil {
		return nil, &response.Error{
			Code: 500,
			Err:  err.Err,
		}
	}

	return data, nil
}

// Delete implements ProductCategoryUsecase
func (usecase *productCategoryUsecase) Delete(id int) (*entity.ProductCategory, *response.Error) {
	// Cari berdasarkan id
	productCategory, err := usecase.repository.FindOneById(id)

	if err != nil {
		return nil, &response.Error{
			Code: 500,
			Err:  err.Err,
		}
	}

	if err := usecase.repository.Delete(*productCategory); err != nil {
		return nil, &response.Error{
			Code: 500,
			Err:  err.Err,
		}
	}

	return productCategory, nil
}

// FindAll implements ProductCategoryUsecase
func (usecase *productCategoryUsecase) FindAll(offset int, limit int) []entity.ProductCategory {
	return usecase.repository.FindAll(offset, limit)
}

// FindOneById implements ProductCategoryUsecase
func (usecase *productCategoryUsecase) FindOneById(id int) (*entity.ProductCategory, *response.Error) {
	return usecase.repository.FindOneById(id)
}

// Update implements ProductCategoryUsecase
func (usecase *productCategoryUsecase) Update(id int, dto dto.ProductCategoryRequestBody) (*entity.ProductCategory, *response.Error) {
	// Cari data berdasarkan id
	productCategory, err := usecase.repository.FindOneById(id)

	if err != nil {
		return nil, &response.Error{
			Code: 500,
			Err:  err.Err,
		}
	}

	productCategory.Name = dto.Name
	productCategory.UpdatedByID = dto.UpdatedBy

	if dto.Image != nil {
		image, err := usecase.media.Upload(*dto.Image)

		if err != nil {
			return nil, &response.Error{
				Code: 500,
				Err:  err.Err,
			}
		}

		// Check image lama
		if productCategory.Image != nil {
			// Delete image lama
			_, err := usecase.media.Delete(*productCategory.Image)

			if err != nil {
				return nil, &response.Error{
					Code: 500,
					Err:  err.Err,
				}
			}
		}

		if image != nil {
			productCategory.Image = image
		}
	}

	// Update
	data, err := usecase.repository.Update(*productCategory)

	if err != nil {
		return nil, &response.Error{
			Code: 500,
			Err:  err.Err,
		}
	}

	return data, nil
}

func NewProductCategoryUsecase(
	repository repository.ProductCategoryRepository,
	media media.Media,
) ProductCategoryUsecase {
	return &productCategoryUsecase{
		repository,
		media,
	}
}
