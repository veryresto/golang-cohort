//go:build wireinject
// +build wireinject

package cart

import (
	"github.com/google/wire"
	"gorm.io/gorm"
	handler "online-course/internal/cart/delivery/http"
	repository "online-course/internal/cart/repository"
	usecase "online-course/internal/cart/usecase"
)

func InitializedService(db *gorm.DB) *handler.CartHandler {
	wire.Build(
		handler.NewCartHandler,
		repository.NewCartRepository,
		usecase.NewCartUsecase,
	)

	return &handler.CartHandler{}
}
