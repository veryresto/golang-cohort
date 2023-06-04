package register

import (
	"net/http"

	registerUseCase "online-course/internal/register/usecase"
	userDto "online-course/internal/user/dto"
	response "online-course/pkg/response"

	"github.com/gin-gonic/gin"
)

type RegisterHandler struct {
	registerUseCase registerUseCase.RegisterUsecase
}

func NewRegisterHandler(registerUseCase registerUseCase.RegisterUsecase) *RegisterHandler {
	return &RegisterHandler{registerUseCase}
}

func (handler *RegisterHandler) Route(r *gin.RouterGroup) {
	r.POST("/api/v1/registers", handler.Register)
}

func (handler *RegisterHandler) Register(ctx *gin.Context) {
	// Validate input
	var registerRequestInput userDto.UserRequestBody

	// Validasi dari body json yang dikirim oleh client
	if err := ctx.ShouldBindJSON(&registerRequestInput); err != nil {
		ctx.JSON(http.StatusBadRequest, response.Response(
			http.StatusBadRequest,
			http.StatusText(http.StatusBadRequest),
			err.Error(),
		))
		ctx.Abort()
		return
	}

	err := handler.registerUseCase.Register(registerRequestInput)

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
		"Success, please check your email",
	))
}
