package dashboard

import (
	adminUsecase "online-course/internal/admin/usecase"
	dto "online-course/internal/dashboard/dto"
	orderUsecase "online-course/internal/order/usecase"
	productUsecase "online-course/internal/product/usecase"
	userUsecase "online-course/internal/user/usecase"
)

type DashboardUsecase interface {
	GetDashboard() dto.DashboardResponseBody
}

type dashboardUsecase struct {
	userUsecase    userUsecase.UserUsecase
	adminUsecase   adminUsecase.AdminUsecase
	productUsecase productUsecase.ProductUsecase
	orderUsecase   orderUsecase.OrderUsecase
}

// GetDashboard implements DashboardUsecase
func (usecase dashboardUsecase) GetDashboard() dto.DashboardResponseBody {
	var dataDashboard dto.DashboardResponseBody

	dataDashboard.TotalAdmin = usecase.adminUsecase.TotalCountAdmin()
	dataDashboard.TotalUser = usecase.userUsecase.TotalCountUser()
	dataDashboard.TotalProduct = usecase.productUsecase.TotalCountProduct()
	dataDashboard.TotalOrder = usecase.orderUsecase.TotalCountOrder()

	return dataDashboard
}

func NewDashboardUsecase(
	userUsecase userUsecase.UserUsecase,
	adminUsecase adminUsecase.AdminUsecase,
	productUsecase productUsecase.ProductUsecase,
	orderUsecase orderUsecase.OrderUsecase,
) DashboardUsecase {
	return &dashboardUsecase{
		userUsecase,
		adminUsecase,
		productUsecase,
		orderUsecase,
	}
}
