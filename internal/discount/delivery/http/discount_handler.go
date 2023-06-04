package discount

import (
	"net/http"
	"strconv"

	dto "online-course/internal/discount/dto"
	usecase "online-course/internal/discount/usecase"
	"online-course/internal/middleware"
	"online-course/pkg/response"
	"online-course/pkg/utils"

	"github.com/gin-gonic/gin"
)

type DiscountHandler struct {
	usecase usecase.DiscountUsecase
}

func NewDiscountHandler(usecase usecase.DiscountUsecase) *DiscountHandler {
	return &DiscountHandler{usecase}
}

func (handler *DiscountHandler) Route(r *gin.RouterGroup) {
	discountRouter := r.Group("/api/v1")

	discountRouter.GET("/discounts", handler.FindAll)
	discountRouter.GET("/discounts/:id", handler.FindById)

	discountRouter.Use(middleware.AuthJwt, middleware.AuthAdmin)
	{
		discountRouter.POST("/discounts", handler.Create)
		discountRouter.PATCH("/discounts/:id", handler.Update)
		discountRouter.DELETE("/discounts/:id", handler.Delete)
	}
}

func (handler *DiscountHandler) FindAll(ctx *gin.Context) {
	offset, _ := strconv.Atoi(ctx.Query("offset"))
	limit, _ := strconv.Atoi(ctx.Query("limit"))

	data := handler.usecase.FindAll(offset, limit)

	ctx.JSON(http.StatusOK, response.Response(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		data,
	))
}

func (handler *DiscountHandler) FindById(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	data, err := handler.usecase.FindOneById(id)

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

func (handler *DiscountHandler) Create(ctx *gin.Context) {
	var input dto.DiscountRequestBody

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, response.Response(
			http.StatusBadRequest,
			http.StatusText(http.StatusBadRequest),
			err.Error(),
		))
		ctx.Abort()
		return
	}

	admin := utils.GetCurrentUser(ctx)

	input.CreatedBy = &admin.ID

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

func (handler *DiscountHandler) Update(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	var input dto.DiscountRequestBody

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, response.Response(
			http.StatusBadRequest,
			http.StatusText(http.StatusBadRequest),
			err.Error(),
		))
		ctx.Abort()
		return
	}

	admin := utils.GetCurrentUser(ctx)

	input.UpdatedBy = &admin.ID

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

func (handler *DiscountHandler) Delete(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	err := handler.usecase.Delete(id)

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
