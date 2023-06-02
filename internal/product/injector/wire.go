//go:build wireinject
// +build wireinject

package product

import (
	"github.com/google/wire"
	"gorm.io/gorm"
	handler "online-course/internal/product/delivery/http"
	repository "online-course/internal/product/repository"
	usecase "online-course/internal/product/usecase"
	media "online-course/pkg/media/cloudinary"
)

func InitializedService(db *gorm.DB) *handler.ProductHandler {
	wire.Build(
		handler.NewProductHandler,
		repository.NewProductRepository,
		usecase.NewProductUsecase,
		media.NewMedia,
	)

	return &handler.ProductHandler{}
}
