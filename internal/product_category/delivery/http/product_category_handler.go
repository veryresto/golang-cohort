package product_category

import (
	"net/http"
	"strconv"

	"online-course/internal/middleware"
	dto "online-course/internal/product_category/dto"
	usecase "online-course/internal/product_category/usecase"
	"online-course/pkg/response"
	"online-course/pkg/utils"

	"github.com/gin-gonic/gin"
)

type ProductCategoryHandler struct {
	usecase usecase.ProductCategoryUsecase
}

func NewProductCategoryHandler(usecase usecase.ProductCategoryUsecase) *ProductCategoryHandler {
	return &ProductCategoryHandler{usecase}
}

func (handler *ProductCategoryHandler) Route(r *gin.RouterGroup) {
	productCategoryRouter := r.Group("/api/v1")

	productCategoryRouter.Use(middleware.AuthJwt, middleware.AuthAdmin)
	{
		productCategoryRouter.GET("/product_categories", handler.FindAll)
		productCategoryRouter.GET("/product_categories/:id", handler.FindById)
		productCategoryRouter.POST("/product_categories", handler.Create)
		productCategoryRouter.PATCH("/product_categories/:id", handler.Update)
		productCategoryRouter.DELETE("/product_categories/:id", handler.Delete)
	}
}

func (handler *ProductCategoryHandler) FindAll(ctx *gin.Context) {
	offset, _ := strconv.Atoi(ctx.Query("offset"))
	limit, _ := strconv.Atoi(ctx.Query("limit"))

	data := handler.usecase.FindAll(offset, limit)

	ctx.JSON(http.StatusOK, response.Response(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		data,
	))
}

func (handler *ProductCategoryHandler) FindById(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	data, err := handler.usecase.FindOneById(id)

	if err != nil {
		ctx.JSON(int(err.Code), response.Response(
			int(err.Code),
			http.StatusText(int(err.Code)),
			err,
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

func (handler *ProductCategoryHandler) Create(ctx *gin.Context) {
	// Validate input
	var input dto.ProductCategoryRequestBody

	if err := ctx.ShouldBind(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, response.Response(
			http.StatusBadRequest,
			http.StatusText(http.StatusBadRequest),
			err,
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
			err,
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

func (handler *ProductCategoryHandler) Update(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	var input dto.ProductCategoryRequestBody

	if err := ctx.ShouldBind(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, response.Response(
			http.StatusBadRequest,
			http.StatusText(http.StatusBadRequest),
			err,
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
			err,
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

func (handler *ProductCategoryHandler) Delete(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	data, err := handler.usecase.Delete(id)

	if err != nil {
		ctx.JSON(int(err.Code), response.Response(
			int(err.Code),
			http.StatusText(int(err.Code)),
			err,
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
