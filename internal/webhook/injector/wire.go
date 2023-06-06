//go:build wireinject
// +build wireinject

package webhook

import (
	"github.com/google/wire"
	"gorm.io/gorm"
	cartRepository "online-course/internal/cart/repository"
	cartUsecase "online-course/internal/cart/usecase"
	classRoomRepository "online-course/internal/class_room/repository"
	classRoomUsecase "online-course/internal/class_room/usecase"
	discountRepository "online-course/internal/discount/repository"
	discountUsecase "online-course/internal/discount/usecase"
	orderRepository "online-course/internal/order/repository"
	orderUsecase "online-course/internal/order/usecase"
	orderDetailRepository "online-course/internal/order_detail/repository"
	orderDetailUsecase "online-course/internal/order_detail/usecase"
	paymentUsecase "online-course/internal/payment/usecase"
	productRepository "online-course/internal/product/repository"
	productUsecase "online-course/internal/product/usecase"
	handler "online-course/internal/webhook/delivery/http"
	usecase "online-course/internal/webhook/usecase"
	media "online-course/pkg/media/cloudinary"
)

func InitializedService(db *gorm.DB) *handler.WebhookHandler {
	wire.Build(
		handler.NewWebhookHandler,
		usecase.NewWebhookUsecase,
		classRoomRepository.NewClassRoomRepository,
		classRoomUsecase.NewClassRoomUseCase,
		orderRepository.NewOrderRepository,
		orderUsecase.NewOrderUseCase,
		orderDetailRepository.NewOrderDetailRepository,
		orderDetailUsecase.NewOrderDetailUsecase,
		cartRepository.NewCartRepository,
		cartUsecase.NewCartUsecase,
		discountRepository.NewDiscountRepository,
		discountUsecase.NewDiscountUseCase,
		paymentUsecase.NewPaymentUsecase,
		productRepository.NewProductRepository,
		productUsecase.NewProductUsecase,
		media.NewMedia,
	)

	return &handler.WebhookHandler{}
}
