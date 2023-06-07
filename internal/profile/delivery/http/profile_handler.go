package profile

import (
	"net/http"
	"strings"

	"online-course/internal/middleware"
	dto "online-course/internal/profile/dto"
	usecase "online-course/internal/profile/usecase"
	dtoUser "online-course/internal/user/dto"
	"online-course/pkg/response"
	"online-course/pkg/utils"

	"github.com/gin-gonic/gin"
)

type ProfileHandler struct {
	usecase usecase.ProfileUsecase
}

func NewProfileHandler(usecase usecase.ProfileUsecase) *ProfileHandler {
	return &ProfileHandler{usecase}
}

func (handler *ProfileHandler) Route(r *gin.RouterGroup) {
	profileRoute := r.Group("/api/v1")

	profileRoute.Use(middleware.AuthJwt)
	{
		profileRoute.GET("/profiles", handler.Profile)
		profileRoute.PATCH("/profiles", handler.Update)
		profileRoute.DELETE("/profiles", handler.Deactive)
		profileRoute.POST("/profiles/logout", handler.Logout)
	}
}

func (handler *ProfileHandler) Profile(ctx *gin.Context) {
	user := utils.GetCurrentUser(ctx)

	data, err := handler.usecase.FindProfile(int(user.ID))

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

func (handler *ProfileHandler) Update(ctx *gin.Context) {
	var input dtoUser.UserUpdateRequestBody

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

	data, err := handler.usecase.Update(int(user.ID), input)

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

func (handler *ProfileHandler) Deactive(ctx *gin.Context) {
	user := utils.GetCurrentUser(ctx)

	err := handler.usecase.Deactive(int(user.ID))

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

func (handler *ProfileHandler) Logout(ctx *gin.Context) {
	var input dto.ProfileRequestLogoutBody

	if err := ctx.ShouldBindHeader(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, response.Response(
			http.StatusBadRequest,
			http.StatusText(http.StatusBadRequest),
			err.Error(),
		))
		ctx.Abort()
		return
	}

	reqToken := input.Authorization // "Bearer tokennya"
	splitToken := strings.Split(reqToken, "Bearer ")

	err := handler.usecase.Logout(splitToken[1])

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
