package verification_email

import (
	"net/http"

	dto "online-course/internal/verification_email/dto"
	usecase "online-course/internal/verification_email/usecase"
	"online-course/pkg/response"

	"github.com/gin-gonic/gin"
)

type VerificationEmailHandler struct {
	usecase usecase.VerificationEmailUsecase
}

func NewVerificationEmailHandler(usecase usecase.VerificationEmailUsecase) *VerificationEmailHandler {
	return &VerificationEmailHandler{usecase}
}

func (handler *VerificationEmailHandler) Route(r *gin.RouterGroup) {
	verificationEmailRouter := r.Group("/api/v1")

	verificationEmailRouter.POST("/verification_emails", handler.VerificationEmail)
}

func (handler *VerificationEmailHandler) VerificationEmail(ctx *gin.Context) {
	// Validate input
	var input dto.VerificationEmailRequestBody

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, response.Response(
			http.StatusBadRequest,
			http.StatusText(http.StatusBadRequest),
			err.Error(),
		))
		ctx.Abort()
		return
	}

	err := handler.usecase.VerificationCode(input)

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
