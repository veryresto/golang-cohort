//go:build wireinject
// +build wireinject

package order

import (
	"github.com/google/wire"
	"gorm.io/gorm"
	cartRepository "online-course/internal/cart/repository"
	cartUsecase "online-course/internal/cart/usecase"
	discoutRepository "online-course/internal/discount/repository"
	discountUsecase "online-course/internal/discount/usecase"
	handler "online-course/internal/order/delivery/http"
	repository "online-course/internal/order/repository"
	usecase "online-course/internal/order/usecase"
	orderDetailRepository "online-course/internal/order_detail/repository"
	orderDetailUsecase "online-course/internal/order_detail/usecase"
	paymentUsecase "online-course/internal/payment/usecase"
	productRepository "online-course/internal/product/repository"
	productUsecase "online-course/internal/product/usecase"
	media "online-course/pkg/media/cloudinary"
)

func InitializedService(db *gorm.DB) *handler.OrderHandler {
	wire.Build(
		handler.NewOrderHandler,
		usecase.NewOrderUseCase,
		repository.NewOrderRepository,
		productRepository.NewProductRepository,
		productUsecase.NewProductUsecase,
		orderDetailRepository.NewOrderDetailRepository,
		orderDetailUsecase.NewOrderDetailUsecase,
		paymentUsecase.NewPaymentUsecase,
		cartUsecase.NewCartUsecase,
		cartRepository.NewCartRepository,
		discoutRepository.NewDiscountRepository,
		discountUsecase.NewDiscountUseCase,
		media.NewMedia,
	)

	return &handler.OrderHandler{}
}
