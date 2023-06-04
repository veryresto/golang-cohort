package order_detail

import (
	"fmt"

	entity "online-course/internal/order_detail/entity"
	repository "online-course/internal/order_detail/repository"
)

type OrderDetailUsecase interface {
	Create(entity entity.OrderDetail)
}

type orderDetailUsecase struct {
	repository repository.OrderDetailRepository
}

// Create implements OrderDetailUsecase
func (usecase *orderDetailUsecase) Create(entity entity.OrderDetail) {
	_, err := usecase.repository.Create(entity)

	if err != nil {
		fmt.Println(err)
	}
}

func NewOrderDetailUsecase(repository repository.OrderDetailRepository) OrderDetailUsecase {
	return &orderDetailUsecase{repository}
}
