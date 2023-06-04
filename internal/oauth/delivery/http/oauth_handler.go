package oauth

import (
	"net/http"

	dto "online-course/internal/oauth/dto"
	usecase "online-course/internal/oauth/usecase"
	response "online-course/pkg/response"

	"github.com/gin-gonic/gin"
)

type OauthHandler struct {
	usecase usecase.OauthUsecase
}

func NewOauthHandler(usecase usecase.OauthUsecase) *OauthHandler {
	return &OauthHandler{usecase}
}

func (handler *OauthHandler) Route(r *gin.RouterGroup) {
	oauthRouter := r.Group("/api/v1")

	oauthRouter.POST("/oauth", handler.Login)
	oauthRouter.POST("/oauth/refresh_tokens", handler.Refresh)
}

func (handler *OauthHandler) Login(ctx *gin.Context) {
	var input dto.LoginRequestBody

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, response.Response(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), err.Error()))
		ctx.Abort()
		return
	}

	// Memanggil fungsi dari login
	data, err := handler.usecase.Login(input)

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
		data,
	))
}

func (handler *OauthHandler) Refresh(ctx *gin.Context) {
	var input dto.RefreshTokenRequestBody

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, response.Response(
			http.StatusBadRequest,
			http.StatusText(http.StatusBadRequest),
			err.Error(),
		))
		ctx.Abort()
		return
	}

	data, err := handler.usecase.Refresh(input)

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
		http.StatusOK,
		http.StatusText(http.StatusOK),
		data,
	))
}
