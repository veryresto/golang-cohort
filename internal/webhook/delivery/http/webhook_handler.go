package webhook

import (
	"net/http"

	dto "online-course/internal/webhook/dto"
	usecase "online-course/internal/webhook/usecase"
	"online-course/pkg/response"

	"github.com/gin-gonic/gin"
)

type WebhookHandler struct {
	usecase usecase.WebhookUsecase
}

func NewWebhookHandler(usecase usecase.WebhookUsecase) *WebhookHandler {
	return &WebhookHandler{usecase}
}

func (handler *WebhookHandler) Route(r *gin.RouterGroup) {
	webhookRoute := r.Group("/api/v1")

	webhookRoute.POST("/webhooks/xendit", handler.Xendit)
}

func (handler *WebhookHandler) Xendit(ctx *gin.Context) {
	var input dto.WebhookRequestBody

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, response.Response(
			http.StatusBadRequest,
			http.StatusText(http.StatusBadRequest),
			err.Error(),
		))
		ctx.Abort()
		return
	}

	err := handler.usecase.UpdatePayment(input.ID)

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
