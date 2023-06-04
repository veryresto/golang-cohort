package cart

import (
	"net/http"
	"strconv"

	dto "online-course/internal/cart/dto"
	usecase "online-course/internal/cart/usecase"
	"online-course/internal/middleware"
	"online-course/pkg/response"
	"online-course/pkg/utils"

	"github.com/gin-gonic/gin"
)

type CartHandler struct {
	usecase usecase.CartUsecase
}

func NewCartHandler(usecase usecase.CartUsecase) *CartHandler {
	return &CartHandler{usecase}
}

func (handler *CartHandler) Route(r *gin.RouterGroup) {
	cartRoute := r.Group("/api/v1")

	cartRoute.Use(middleware.AuthJwt)
	{
		cartRoute.GET("/carts", handler.FindByUserId)
		cartRoute.POST("/carts", handler.Create)
		cartRoute.PATCH("/carts/:id", handler.Update)
		cartRoute.DELETE("/carts/:id", handler.Delete)
	}
}

func (handler *CartHandler) FindByUserId(ctx *gin.Context) {
	offset, _ := strconv.Atoi(ctx.Query("offset"))
	limit, _ := strconv.Atoi(ctx.Query("limit"))
	user := utils.GetCurrentUser(ctx)

	data := handler.usecase.FindByUserId(int(user.ID), offset, limit)

	ctx.JSON(http.StatusOK, response.Response(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		data,
	))
}

func (handler *CartHandler) Create(ctx *gin.Context) {
	var input dto.CartRequestBody

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
	input.CreatedBy = user.ID

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

func (handler *CartHandler) Delete(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	user := utils.GetCurrentUser(ctx)

	err := handler.usecase.Delete(id, int(user.ID))

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
		http.StatusText(http.StatusOK),
	))
}

func (handler *CartHandler) Update(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	var input dto.CartRequestUpdateBody

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

	input.UserID = &user.ID

	data, err := handler.usecase.Update(id, input)

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
