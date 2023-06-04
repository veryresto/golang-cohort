package payment

import (
	"os"

	dto "online-course/internal/payment/dto"
	"online-course/pkg/response"

	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/invoice"
)

type PaymentUsecase interface {
	Create(dto dto.PaymentRequestBody) (*xendit.Invoice, *response.Error)
}

type paymentUsecase struct {
}

// Create implements PaymentUsecase
func (usecase *paymentUsecase) Create(dto dto.PaymentRequestBody) (*xendit.Invoice, *response.Error) {
	data := invoice.CreateParams{
		ExternalID:  dto.ExternalID,
		Amount:      float64(dto.Amount),
		Description: dto.Description,
		PayerEmail:  dto.PayerEmail,
		Customer: xendit.InvoiceCustomer{
			Email: dto.PayerEmail,
		},
		CustomerNotificationPreference: xendit.InvoiceCustomerNotificationPreference{
			InvoiceCreated:  []string{"email"},
			InvoiceReminder: []string{"email"},
			InvoicePaid:     []string{"email"},
			InvoiceExpired:  []string{"email"},
		},
		InvoiceDuration:    86400,
		SuccessRedirectURL: os.Getenv("XENDIT_SUCCESS_URL"),
		FailureRedirectURL: os.Getenv("XENDIT_SUCCESS_URL"),
	}

	resp, err := invoice.Create(&data)

	if err != nil {
		return nil, &response.Error{
			Code: 500,
			Err:  err,
		}
	}

	return resp, nil
}

func NewPaymentUsecase() PaymentUsecase {
	// Xendit
	xendit.Opt.SecretKey = os.Getenv("XENDIT_APIKEY")
	return &paymentUsecase{}
}
