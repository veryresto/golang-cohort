package order

import (
	"net/http"
	"strconv"

	"online-course/internal/middleware"
	dto "online-course/internal/order/dto"
	usecase "online-course/internal/order/usecase"
	"online-course/pkg/response"
	"online-course/pkg/utils"

	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	usecase usecase.OrderUsecase
}

func NewOrderHandler(usecase usecase.OrderUsecase) *OrderHandler {
	return &OrderHandler{usecase}
}

func (handler *OrderHandler) Route(r *gin.RouterGroup) {
	orderRoute := r.Group("/api/v1")

	orderRoute.Use(middleware.AuthJwt)
	{
		orderRoute.POST("/orders", handler.Create)
		orderRoute.GET("/orders", handler.FindAllByUserId)
		orderRoute.GET("/orders/:id", handler.FindById)
	}
}

func (handler *OrderHandler) FindAllByUserId(ctx *gin.Context) {
	offset, _ := strconv.Atoi(ctx.Query("offset"))
	limit, _ := strconv.Atoi(ctx.Query("limit"))

	user := utils.GetCurrentUser(ctx)

	data := handler.usecase.FindAllByUserId(int(user.ID), offset, limit)

	ctx.JSON(http.StatusOK, response.Response(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		data,
	))
}

func (handler *OrderHandler) FindById(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	user := utils.GetCurrentUser(ctx)

	data, err := handler.usecase.FindOneById(id, int(user.ID))

	if err != nil {
		ctx.JSON(int(err.Code), response.Response(
			int(err.Code),
			http.StatusText(int(err.Code)),
			err.Err.Error(),
		))
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, response.Response(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		data,
	))
}

func (handler *OrderHandler) Create(ctx *gin.Context) {
	var input dto.OrderRequestBody

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, response.Response(
			http.StatusBadRequest,
			http.StatusText(http.StatusBadRequest),
			err.Error(),
		))
		ctx.Abort()
		return
	}

	user := utils.GetCurrentUser(ctx)

	input.UserID = user.ID
	input.Email = user.Email

	data, err := handler.usecase.Create(input)

	if err != nil {
		ctx.JSON(int(err.Code), response.Response(
			int(err.Code),
			http.StatusText(int(err.Code)),
			err.Err.Error(),
		))
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, response.Response(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		data,
	))
}
