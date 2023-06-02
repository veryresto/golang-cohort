package product

import (
	dto "online-course/internal/product/dto"
	entity "online-course/internal/product/entity"
	repository "online-course/internal/product/repository"
	media "online-course/pkg/media/cloudinary"
	"online-course/pkg/response"
)

type ProductUsecase interface {
	FindAll(offset int, limit int) []entity.Product
	FindOneById(id int) (*entity.Product, *response.Error)
	Create(dto dto.ProductRequestBody) (*entity.Product, *response.Error)
	Update(id int, dto dto.ProductRequestBody) (*entity.Product, *response.Error)
	Delete(id int) (*entity.Product, *response.Error)
}

type productUsecase struct {
	repositoy repository.ProductRepository
	media     media.Media
}

// Create implements ProductUsecase
func (usecase *productUsecase) Create(dto dto.ProductRequestBody) (*entity.Product, *response.Error) {
	product := &entity.Product{
		ProductCategoryID: &dto.ProductCategoryID,
		Title:             dto.Title,
		Description:       dto.Description,
		IsHighlighted:     dto.IsHighlighted,
		Price:             int64(dto.Price),
		CreatedByID:       dto.CreatedBy,
	}

	// Upload image
	if dto.Image != nil {
		image, err := usecase.media.Upload(*dto.Image)

		if err != nil {
			return nil, err
		}

		if image != nil {
			product.Image = image
		}
	}

	// Upload video
	if dto.Video != nil {
		video, err := usecase.media.Upload(*dto.Video)

		if err != nil {
			return nil, err
		}

		if video != nil {
			product.Video = video
		}
	}

	data, err := usecase.repositoy.Create(*product)

	if err != nil {
		return nil, err
	}

	return data, nil
}

// Delete implements ProductUsecase
func (usecase *productUsecase) Delete(id int) (*entity.Product, *response.Error) {
	product, err := usecase.repositoy.FindOneById(id)

	if err != nil {
		return nil, err
	}

	err = usecase.repositoy.Delete(*product)

	if err != nil {
		return nil, err
	}

	return product, nil
}

// FindAll implements ProductUsecase
func (usecase *productUsecase) FindAll(offset int, limit int) []entity.Product {
	return usecase.repositoy.FindAll(offset, limit)
}

// FindOneById implements ProductUsecase
func (usecase *productUsecase) FindOneById(id int) (*entity.Product, *response.Error) {
	return usecase.repositoy.FindOneById(id)
}

// Update implements ProductUsecase
func (usecase *productUsecase) Update(id int, dto dto.ProductRequestBody) (*entity.Product, *response.Error) {
	// Cari product berdasarkan id
	product, err := usecase.repositoy.FindOneById(id)

	if err != nil {
		return nil, err
	}

	product.ProductCategoryID = &dto.ProductCategoryID
	product.Title = dto.Title
	product.Description = dto.Description
	product.IsHighlighted = dto.IsHighlighted
	product.Price = int64(dto.Price)
	product.UpdatedByID = dto.UpdatedBy

	// Upload image
	if dto.Image != nil {
		image, err := usecase.media.Upload(*dto.Image)

		if err != nil {
			return nil, err
		}

		// Check image yang lama
		if product.Image != nil {
			// Delete
			_, err := usecase.media.Delete(*product.Image)

			if err != nil {
				return nil, err
			}
		}

		if image != nil {
			product.Image = image
		}
	}

	// Upload video
	if dto.Video != nil {
		video, err := usecase.media.Upload(*dto.Video)

		if err != nil {
			return nil, err
		}

		// Check video yang lama
		if product.Video != nil {
			_, err := usecase.media.Delete(*product.Video)

			if err != nil {
				return nil, err
			}
		}

		if video != nil {
			product.Video = video
		}
	}

	data, err := usecase.repositoy.Update(*product)

	if err != nil {
		return nil, err
	}

	return data, nil
}

func NewProductUsecase(
	repositoy repository.ProductRepository,
	media media.Media,
) ProductUsecase {
	return &productUsecase{
		repositoy,
		media,
	}
}
