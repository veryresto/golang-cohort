//go:build wireinject
// +build wireinject

package dashboard

import (
	"github.com/google/wire"
	"gorm.io/gorm"
	adminRepository "online-course/internal/admin/repository"
	adminUsecase "online-course/internal/admin/usecase"
	cartRepository "online-course/internal/cart/repository"
	cartUsecase "online-course/internal/cart/usecase"
	handler "online-course/internal/dashboard/delivery/http"
	usecase "online-course/internal/dashboard/usecase"
	discountRepository "online-course/internal/discount/repository"
	discountUsecase "online-course/internal/discount/usecase"
	orderRepository "online-course/internal/order/repository"
	orderUsecase "online-course/internal/order/usecase"
	orderDetailRepository "online-course/internal/order_detail/repository"
	orderDetailUsecase "online-course/internal/order_detail/usecase"
	paymentUsecase "online-course/internal/payment/usecase"
	productRepository "online-course/internal/product/repository"
	productUsecase "online-course/internal/product/usecase"
	userRepository "online-course/internal/user/repository"
	userUsecase "online-course/internal/user/usecase"
	media "online-course/pkg/media/cloudinary"
)

func InitializedService(db *gorm.DB) *handler.DashboardHandler {
	wire.Build(
		handler.NewDashboardHandler,
		usecase.NewDashboardUsecase,
		cartRepository.NewCartRepository,
		cartUsecase.NewCartUsecase,
		discountRepository.NewDiscountRepository,
		discountUsecase.NewDiscountUseCase,
		orderRepository.NewOrderRepository,
		orderUsecase.NewOrderUseCase,
		orderDetailRepository.NewOrderDetailRepository,
		orderDetailUsecase.NewOrderDetailUsecase,
		paymentUsecase.NewPaymentUsecase,
		productRepository.NewProductRepository,
		productUsecase.NewProductUsecase,
		userRepository.NewUserRepository,
		userUsecase.NewUserUseCase,
		media.NewMedia,
		adminRepository.NewAdminRepository,
		adminUsecase.NewAdminUsecase,
	)

	return &handler.DashboardHandler{}
}
