package webhook

import (
	"errors"
	"fmt"
	"os"
	"strings"

	classRoomDto "online-course/internal/class_room/dto"
	classRoomUsecase "online-course/internal/class_room/usecase"
	orderDto "online-course/internal/order/dto"
	orderUsecase "online-course/internal/order/usecase"
	"online-course/pkg/response"

	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/invoice"
)

type WebhookUsecase interface {
	UpdatePayment(id string) *response.Error
}

type webHookUsecase struct {
	orderUsecase     orderUsecase.OrderUsecase
	classRoomUsecase classRoomUsecase.ClassRoomUsecase
}

// UpdatePayment implements WebhookUsecase
func (usecase *webHookUsecase) UpdatePayment(id string) *response.Error {
	// Kita akan memeriksa data dari xendit
	params := invoice.GetParams{
		ID: id,
	}

	dataXendit, err := invoice.Get(&params)

	if err != nil {
		return &response.Error{
			Code: 500,
			Err:  err,
		}
	}

	if dataXendit == nil {
		return &response.Error{
			Code: 400,
			Err:  errors.New("order is not found"),
		}
	}

	dataOrder, errOrder := usecase.orderUsecase.FindOneByExternalId(dataXendit.ExternalID)

	if err != nil {
		return errOrder
	}

	if dataOrder == nil {
		return &response.Error{
			Code: 400,
			Err:  errors.New("order is not found"),
		}
	}

	if dataOrder.Status == "settled" {
		return &response.Error{
			Code: 400,
			Err:  errors.New("payment has been already processed"),
		}
	}

	if dataOrder.Status != "paid" {
		if dataXendit.Status == "PAID" || dataXendit.Status == "SETTLED" {
			// add to class room
			for _, orderDetail := range dataOrder.OrderDetails {
				dataClassRoom := classRoomDto.ClassRoomRequestBody{
					UserID:    *dataOrder.UserID,
					ProductID: *orderDetail.ProductID,
				}

				_, err := usecase.classRoomUsecase.Create(dataClassRoom)

				if err != nil {
					fmt.Println(err)
				}
			}
		}

		// Mengirimkan notif
	}

	// Update data order
	order := orderDto.OrderRequestBody{
		Status: strings.ToLower(dataXendit.Status),
	}

	usecase.orderUsecase.Update(int(dataOrder.ID), order)

	return nil
}

func NewWebhookUsecase(
	orderUsecase orderUsecase.OrderUsecase,
	classRoomUsecase classRoomUsecase.ClassRoomUsecase,
) WebhookUsecase {
	xendit.Opt.SecretKey = os.Getenv("XENDIT_APIKEY")

	return &webHookUsecase{
		orderUsecase,
		classRoomUsecase,
	}
}
