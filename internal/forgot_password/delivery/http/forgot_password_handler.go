package forgot_password

import (
	"net/http"

	dto "online-course/internal/forgot_password/dto"
	usecase "online-course/internal/forgot_password/usecase"
	"online-course/pkg/response"

	"github.com/gin-gonic/gin"
)

type ForgotPasswordHandler struct {
	usecase usecase.ForgotPasswordUsecase
}

func NewForgotPasswordHandler(usecase usecase.ForgotPasswordUsecase) *ForgotPasswordHandler {
	return &ForgotPasswordHandler{usecase}
}

func (handler *ForgotPasswordHandler) Route(r *gin.RouterGroup) {
	forgotPasswordRouter := r.Group("/api/v1")

	forgotPasswordRouter.POST("/forgot_passwords", handler.Create)
	forgotPasswordRouter.PUT("/forgot_passwords", handler.Update)
}

func (handler *ForgotPasswordHandler) Create(ctx *gin.Context) {
	var input dto.ForgotPasswordRequestBody

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, response.Response(
			http.StatusBadRequest,
			http.StatusText(http.StatusBadRequest),
			err.Error(),
		))
		ctx.Abort()
		return
	}

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

	ctx.JSON(http.StatusOK, response.Response(http.StatusOK, http.StatusText(http.StatusOK), "Success, please check your email."))
}

func (handler *ForgotPasswordHandler) Update(ctx *gin.Context) {
	var input dto.ForgotPasswordUpdateRequestBody

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, response.Response(
			http.StatusBadRequest,
			http.StatusText(http.StatusBadRequest),
			err.Error(),
		))
		ctx.Abort()
		return
	}

	_, err := handler.usecase.Update(input)

	if err != nil {
		ctx.JSON(int(err.Code),
			response.Response(int(err.Code),
				http.StatusText(int(err.Code)),
				err.Err.Error(),
			))
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, response.Response(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		"Success change your password.",
	))
}
