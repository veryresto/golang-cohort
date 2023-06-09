// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package dashboard

import (
	"gorm.io/gorm"
	"online-course/internal/admin/repository"
	admin2 "online-course/internal/admin/usecase"
	"online-course/internal/cart/repository"
	cart2 "online-course/internal/cart/usecase"
	"online-course/internal/dashboard/delivery/http"
	dashboard2 "online-course/internal/dashboard/usecase"
	"online-course/internal/discount/repository"
	discount2 "online-course/internal/discount/usecase"
	"online-course/internal/order/repository"
	order2 "online-course/internal/order/usecase"
	"online-course/internal/order_detail/repository"
	order_detail2 "online-course/internal/order_detail/usecase"
	"online-course/internal/payment/usecase"
	"online-course/internal/product/repository"
	product2 "online-course/internal/product/usecase"
	"online-course/internal/user/repository"
	user2 "online-course/internal/user/usecase"
	"online-course/pkg/media/cloudinary"
)

// Injectors from wire.go:

func InitializedService(db *gorm.DB) *dashboard.DashboardHandler {
	userRepository := user.NewUserRepository(db)
	userUsecase := user2.NewUserUseCase(userRepository)
	adminRepository := admin.NewAdminRepository(db)
	adminUsecase := admin2.NewAdminUsecase(adminRepository)
	productRepository := product.NewProductRepository(db)
	mediaMedia := media.NewMedia()
	productUsecase := product2.NewProductUsecase(productRepository, mediaMedia)
	orderRepository := order.NewOrderRepository(db)
	cartRepository := cart.NewCartRepository(db)
	cartUsecase := cart2.NewCartUsecase(cartRepository)
	discountRepository := discount.NewDiscountRepository(db)
	discountUsecase := discount2.NewDiscountUseCase(discountRepository)
	orderDetailRepository := order_detail.NewOrderDetailRepository(db)
	orderDetailUsecase := order_detail2.NewOrderDetailUsecase(orderDetailRepository)
	paymentUsecase := payment.NewPaymentUsecase()
	orderUsecase := order2.NewOrderUseCase(orderRepository, cartUsecase, discountUsecase, productUsecase, orderDetailUsecase, paymentUsecase)
	dashboardUsecase := dashboard2.NewDashboardUsecase(userUsecase, adminUsecase, productUsecase, orderUsecase)
	dashboardHandler := dashboard.NewDashboardHandler(dashboardUsecase)
	return dashboardHandler
}
