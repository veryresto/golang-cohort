package middleware

import (
	"net/http"
	"os"
	"strings"

	dto "online-course/internal/oauth/dto"
	"online-course/pkg/response"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type Header struct {
	Authorization string `header:"authorization" binding:"required"`
}

func AuthJwt(ctx *gin.Context) {
	var input Header

	if err := ctx.ShouldBindHeader(&input); err != nil {
		ctx.JSON(http.StatusUnauthorized, response.Response(
			http.StatusUnauthorized,
			http.StatusText(http.StatusUnauthorized),
			"Unauthorized",
		))
		ctx.Abort()
		return
	}

	reqToken := input.Authorization // "Bearer tokennya"
	splitToken := strings.Split(reqToken, "Bearer ")

	if len(splitToken) != 2 {
		ctx.JSON(http.StatusUnauthorized, response.Response(
			http.StatusUnauthorized,
			http.StatusText(http.StatusUnauthorized),
			"Unauthorized",
		))
		ctx.Abort()
		return
	}

	reqToken = splitToken[1]
	claims := &dto.MapClaimsResponse{}

	token, err := jwt.ParseWithClaims(reqToken, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, response.Response(
			http.StatusUnauthorized,
			http.StatusText(http.StatusUnauthorized),
			"Unauthorized",
		))
		ctx.Abort()
		return
	}

	if !token.Valid {
		ctx.JSON(http.StatusUnauthorized, response.Response(
			http.StatusUnauthorized,
			http.StatusText(http.StatusUnauthorized),
			"Unauthorized",
		))
		ctx.Abort()
		return
	}

	claims = token.Claims.(*dto.MapClaimsResponse)

	ctx.Set("user", claims)
	ctx.Next()
}
