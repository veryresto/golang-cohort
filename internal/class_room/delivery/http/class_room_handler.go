package class_room

import (
	"net/http"
	"strconv"

	usecase "online-course/internal/class_room/usecase"
	"online-course/internal/middleware"
	"online-course/pkg/response"
	"online-course/pkg/utils"

	"github.com/gin-gonic/gin"
)

type ClassRoomHandler struct {
	usecase usecase.ClassRoomUsecase
}

func NewClassRoomHandler(usecase usecase.ClassRoomUsecase) *ClassRoomHandler {
	return &ClassRoomHandler{usecase}
}

func (handler *ClassRoomHandler) Route(r *gin.RouterGroup) {
	classRoomRoute := r.Group("/api/v1")

	classRoomRoute.Use(middleware.AuthJwt)
	{
		classRoomRoute.GET("/class_rooms", handler.FindAllByUserId)
	}
}

func (handler *ClassRoomHandler) FindAllByUserId(ctx *gin.Context) {
	offset, _ := strconv.Atoi(ctx.Query("offset"))
	limit, _ := strconv.Atoi(ctx.Query("limit"))

	user := utils.GetCurrentUser(ctx)

	data := handler.usecase.FindAllByUserID(int(user.ID), offset, limit)

	ctx.JSON(http.StatusOK, response.Response(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		data,
	))

}
