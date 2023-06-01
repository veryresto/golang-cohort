package middleware

import (
	"net/http"

	"online-course/pkg/response"
	"online-course/pkg/utils"

	"github.com/gin-gonic/gin"
)

func AuthAdmin(ctx *gin.Context) {
	admin := utils.GetCurrentUser(ctx)

	if !admin.IsAdmin {
		ctx.JSON(http.StatusUnauthorized, response.Response(
			http.StatusUnauthorized,
			http.StatusText(http.StatusUnauthorized),
			"Unauthorized",
		))
		ctx.Abort()
		return
	}

	ctx.Next()
}
