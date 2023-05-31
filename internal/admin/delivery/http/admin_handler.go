package admin

import (
	"net/http"
	"strconv"

	dto "online-course/internal/admin/dto"
	usecase "online-course/internal/admin/usecase"
	"online-course/pkg/response"

	"github.com/gin-gonic/gin"
)

type AdminHandler struct {
	usecase usecase.AdminUsecase
}

func NewAdminHandler(usecase usecase.AdminUsecase) *AdminHandler {
	return &AdminHandler{usecase}
}

func (handler *AdminHandler) Route(r *gin.RouterGroup) {
	adminRouter := r.Group("/api/v1")

	adminRouter.GET("/admins", handler.FindAll)
	adminRouter.POST("/admins", handler.Create)
	adminRouter.GET("/admins/:id", handler.FindById)
	adminRouter.PATCH("/admins/:id", handler.Update)
	adminRouter.DELETE("/admins/:id", handler.Delete)
}

func (handler *AdminHandler) Create(ctx *gin.Context) {
	var input dto.AdminRequestBody

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, response.Response(
			http.StatusBadRequest,
			http.StatusText(http.StatusBadRequest),
			err.Error(),
		))
		return
	}

	// Create data
	_, err := handler.usecase.Create(input)

	if err != nil {
		ctx.JSON(int(err.Code), response.Response(
			int(err.Code),
			http.StatusText(int(err.Code)),
			err.Err.Error(),
		))
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusCreated, response.Response(
		http.StatusCreated,
		http.StatusText(http.StatusCreated),
		"created",
	))
}

func (handler *AdminHandler) Update(ctx *gin.Context) {
	// Kita perlu mendapatkan params id /api/v1/admins/:param
	id, _ := strconv.Atoi(ctx.Param("id"))

	var input dto.AdminRequestBody

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, response.Response(
			http.StatusBadRequest,
			http.StatusText(http.StatusBadRequest),
			err.Error(),
		))
		ctx.Abort()
		return
	}

	// Update data
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

func (handler *AdminHandler) FindAll(ctx *gin.Context) {
	// api/v1/admins?offset=1&limit=5 query
	offset, _ := strconv.Atoi(ctx.Query("offset"))
	limit, _ := strconv.Atoi(ctx.Query("limit"))

	data := handler.usecase.FindAll(offset, limit)

	ctx.JSON(http.StatusOK, response.Response(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		data,
	))
}

func (handler *AdminHandler) FindById(ctx *gin.Context) {
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

func (handler *AdminHandler) Delete(ctx *gin.Context) {
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
		"ok",
	))
}
