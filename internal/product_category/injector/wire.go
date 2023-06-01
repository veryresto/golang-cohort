//go:build wireinject
// +build wireinject

package product_category

import (
	"github.com/google/wire"
	"gorm.io/gorm"
	handler "online-course/internal/product_category/delivery/http"
	repository "online-course/internal/product_category/repository"
	usecase "online-course/internal/product_category/usecase"
	media "online-course/pkg/media/cloudinary"
)

func InitializedService(db *gorm.DB) *handler.ProductCategoryHandler {
	wire.Build(
		handler.NewProductCategoryHandler,
		repository.NewProductCategoryRepository,
		usecase.NewProductCategoryUsecase,
		media.NewMedia,
	)

	return &handler.ProductCategoryHandler{}
}
