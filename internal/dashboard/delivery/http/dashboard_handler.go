package dashboard

import (
	"net/http"

	usecase "online-course/internal/dashboard/usecase"
	"online-course/internal/middleware"
	"online-course/pkg/response"

	"github.com/gin-gonic/gin"
)

type DashboardHandler struct {
	usecase usecase.DashboardUsecase
}

func NewDashboardHandler(usecase usecase.DashboardUsecase) *DashboardHandler {
	return &DashboardHandler{usecase}
}

func (handler *DashboardHandler) Route(r *gin.RouterGroup) {
	dashboardRoute := r.Group("/api/v1")

	dashboardRoute.Use(middleware.AuthJwt, middleware.AuthAdmin)
	{
		dashboardRoute.GET("/dashboards", handler.GetDataDashboard)
	}
}

func (handler *DashboardHandler) GetDataDashboard(ctx *gin.Context) {
	data := handler.usecase.GetDashboard()

	ctx.JSON(http.StatusOK, response.Response(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		data,
	))
}
