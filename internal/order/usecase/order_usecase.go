package order

import (
	"errors"
	"strconv"
	"time"

	cartUsecase "online-course/internal/cart/usecase"
	discountEntity "online-course/internal/discount/entity"
	discountUsecase "online-course/internal/discount/usecase"
	dto "online-course/internal/order/dto"
	entity "online-course/internal/order/entity"
	repository "online-course/internal/order/repository"
	orderDetailEntity "online-course/internal/order_detail/entity"
	orderDetailUsecase "online-course/internal/order_detail/usecase"
	paymentDto "online-course/internal/payment/dto"
	paymentUsecase "online-course/internal/payment/usecase"
	productEntity "online-course/internal/product/entity"
	productUsecase "online-course/internal/product/usecase"
	"online-course/pkg/response"

	"github.com/google/uuid"
)

type OrderUsecase interface {
	FindAllByUserId(userId int, offset int, limit int) []entity.Order
	FindOneById(id int, userId int) (*entity.Order, *response.Error)
	FindOneByExternalId(externalId string) (*entity.Order, *response.Error)
	Create(dto dto.OrderRequestBody) (*entity.Order, *response.Error)
	Update(id int, dto dto.OrderRequestBody) (*entity.Order, *response.Error)
}

type orderUsecase struct {
	repository         repository.OrderRepository
	cartUsecase        cartUsecase.CartUsecase
	discountUsecase    discountUsecase.DiscountUsecase
	productUsecase     productUsecase.ProductUsecase
	orderDetailUsecase orderDetailUsecase.OrderDetailUsecase
	paymentUsecase     paymentUsecase.PaymentUsecase
}

// Create implements OrderUsecase
func (usecase *orderUsecase) Create(dto dto.OrderRequestBody) (*entity.Order, *response.Error) {
	price := 0
	totalPrice := 0
	description := ""
	var products []productEntity.Product

	order := &entity.Order{
		UserID: &dto.UserID,
		Status: "pending",
	}

	var dataDiscount *discountEntity.Discount

	// Cari data keranjang berdasarkan user id
	carts := usecase.cartUsecase.FindByUserId(int(dto.UserID), 0, 9999)

	if len(carts) == 0 {
		// Jika kosong kita akan melakukan pemeriksaan product idnya apakah dikirim oleh client
		if dto.ProductID == nil {
			return nil, &response.Error{
				Code: 400,
				Err:  errors.New("cart anda kosong atau anda belum memasukkan product id"),
			}
		}
	}

	// Check data discount
	if dto.DiscountCode != nil {
		discount, err := usecase.discountUsecase.FindOneByCode(*dto.DiscountCode)

		if err != nil {
			return nil, &response.Error{
				Code: 400,
				Err:  errors.New("code sudah expired"),
			}
		}

		if discount.RemainingQuantity == 0 {
			return nil, &response.Error{
				Code: 400,
				Err:  errors.New("code sudah expired"),
			}
		}

		if discount.StartDate != nil && discount.EndDate != nil {
			if discount.StartDate.After(time.Now()) || discount.EndDate.Before(time.Now()) {
				return nil, &response.Error{
					Code: 400,
					Err:  errors.New("code sudah expired"),
				}
			}
		} else if discount.StartDate != nil {
			if discount.StartDate.After(time.Now()) {
				return nil, &response.Error{
					Code: 400,
					Err:  errors.New("code sudah expired"),
				}
			}
		} else if discount.EndDate != nil {
			if discount.EndDate.Before(time.Now()) {
				return nil, &response.Error{
					Code: 400,
					Err:  errors.New("code sudah expired"),
				}
			}
		}

		dataDiscount = discount
	}

	if len(carts) > 0 {
		// Jika menggunakan cart
		for _, cart := range carts {
			product, err := usecase.productUsecase.FindOneById(int(*cart.ProductID))

			if err != nil {
				return nil, &response.Error{
					Code: 500,
					Err:  err.Err,
				}
			}

			products = append(products, *product)
		}
	} else if dto.ProductID != nil {
		product, err := usecase.productUsecase.FindOneById(int(*dto.ProductID))

		if err != nil {
			return nil, &response.Error{
				Code: 500,
				Err:  err.Err,
			}
		}

		products = append(products, *product)
	}

	// Kalkulasi price serta membuat description untuk xendit
	for index, product := range products {
		price += int(product.Price)

		i := strconv.Itoa(index + 1)

		description = i + ". Product : " + product.Title + "<br/>"
	}

	totalPrice = price

	// Check apakah terdapat data discount
	if dataDiscount != nil {
		// Hitung logic discount
		if dataDiscount.Type == "rebate" {
			totalPrice = price - int(dataDiscount.Value)
		} else if dataDiscount.Type == "percent" {
			totalPrice = price - (price / 100 * int(dataDiscount.Value))
		}

		order.DiscountID = &dataDiscount.ID
	}

	order.Price = int64(price)
	order.TotalPrice = int64(totalPrice)
	order.CreatedByID = &dto.UserID

	externalId := uuid.New().String()

	order.ExternalID = externalId

	// Insert ke table order
	data, err := usecase.repository.Create(*order)

	if err != nil {
		return nil, err
	}

	// Insert ke table order details
	for _, product := range products {
		orderDetail := &orderDetailEntity.OrderDetail{
			ProductID:   &product.ID,
			Price:       product.Price,
			CreatedByID: order.UserID,
			OrderID:     data.ID,
		}

		usecase.orderDetailUsecase.Create(*orderDetail)
	}

	// Hit payment xendit
	dataPayment := paymentDto.PaymentRequestBody{
		ExternalID:  externalId,
		Amount:      int64(totalPrice),
		PayerEmail:  dto.Email,
		Description: description,
	}

	payment, err := usecase.paymentUsecase.Create(dataPayment)

	if err != nil {
		return nil, &response.Error{
			Code: 500,
			Err:  err.Err,
		}
	}

	data.CheckoutLink = payment.InvoiceURL

	usecase.repository.Update(*data)

	// Update remaining quantity discount
	if dto.DiscountCode != nil {
		_, err := usecase.discountUsecase.UpdateRemainingQuantity(int(dataDiscount.ID), 1, "-")

		if err != nil {
			return nil, &response.Error{
				Code: 500,
				Err:  err.Err,
			}
		}
	}

	// Delete carts
	err = usecase.cartUsecase.DeleteByUserId(int(dto.UserID))

	if err != nil {
		return nil, &response.Error{
			Code: 500,
			Err:  err.Err,
		}
	}

	return data, nil

}

// FindAllByUserId implements OrderUsecase
func (usecase *orderUsecase) FindAllByUserId(userId int, offset int, limit int) []entity.Order {
	return usecase.repository.FindAllByUserId(userId, offset, limit)
}

// FindOneByExternalId implements OrderUsecase
func (usecase *orderUsecase) FindOneByExternalId(externalId string) (*entity.Order, *response.Error) {
	return usecase.repository.FindOneByExternalId(externalId)
}

// FindOneById implements OrderUsecase
func (usecase *orderUsecase) FindOneById(id int, userId int) (*entity.Order, *response.Error) {
	return usecase.repository.FindOneById(id)
}

// Update implements OrderUsecase
func (*orderUsecase) Update(id int, dto dto.OrderRequestBody) (*entity.Order, *response.Error) {
	panic("unimplemented")
}

func NewOrderUseCase(
	repository repository.OrderRepository,
	cartUsecase cartUsecase.CartUsecase,
	discountUsecase discountUsecase.DiscountUsecase,
	productUsecase productUsecase.ProductUsecase,
	orderDetailUsecase orderDetailUsecase.OrderDetailUsecase,
	paymentUsecase paymentUsecase.PaymentUsecase,
) OrderUsecase {
	return &orderUsecase{
		repository,
		cartUsecase,
		discountUsecase,
		productUsecase,
		orderDetailUsecase,
		paymentUsecase,
	}
}
