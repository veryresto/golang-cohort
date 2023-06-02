package discount

import (
	dto "online-course/internal/discount/dto"
	entity "online-course/internal/discount/entity"
	repository "online-course/internal/discount/repository"
	"online-course/pkg/response"
)

type DiscountUsecase interface {
	FindAll(offset int, limit int) []entity.Discount
	FindOneById(id int) (*entity.Discount, *response.Error)
	FindOneByCode(code string) (*entity.Discount, *response.Error)
	Create(dto dto.DiscountRequestBody) (*entity.Discount, *response.Error)
	Update(id int, dto dto.DiscountRequestBody) (*entity.Discount, *response.Error)
	Delete(id int) *response.Error
	UpdateRemainingQuantity(id int, quantity int, operator string) (*entity.Discount, *response.Error)
}

type discountUsecase struct {
	repository repository.DiscountRepository
}

// Create implements DiscountUsecase
func (usecase *discountUsecase) Create(dto dto.DiscountRequestBody) (*entity.Discount, *response.Error) {
	discount := &entity.Discount{
		Name:              dto.Name,
		Code:              dto.Code,
		Quantity:          dto.Quantity,
		RemainingQuantity: dto.Quantity,
		Type:              dto.Type,
		Value:             dto.Value,
		CreatedByID:       dto.CreatedBy,
	}

	if dto.StartDate != nil {
		discount.StartDate = dto.StartDate
	}

	if dto.EndDate != nil {
		discount.EndDate = dto.EndDate
	}

	data, err := usecase.repository.Create(*discount)

	if err != nil {
		return nil, err
	}

	return data, nil
}

// Delete implements DiscountUsecase
func (usecase *discountUsecase) Delete(id int) *response.Error {
	discount, err := usecase.repository.FindOneById(id)

	if err != nil {
		return err
	}

	err = usecase.repository.Delete(*discount)

	if err != nil {
		return err
	}

	return nil
}

// FindAll implements DiscountUsecase
func (usecase *discountUsecase) FindAll(offset int, limit int) []entity.Discount {
	return usecase.repository.FindAll(offset, limit)
}

// FindOneByCode implements DiscountUsecase
func (usecase *discountUsecase) FindOneByCode(code string) (*entity.Discount, *response.Error) {
	return usecase.repository.FindOneByCode(code)
}

// FindOneById implements DiscountUsecase
func (usecase *discountUsecase) FindOneById(id int) (*entity.Discount, *response.Error) {
	return usecase.repository.FindOneById(id)
}

// Update implements DiscountUsecase
func (usecase *discountUsecase) Update(id int, dto dto.DiscountRequestBody) (*entity.Discount, *response.Error) {
	discount, err := usecase.repository.FindOneById(id)

	if err != nil {
		return nil, err
	}

	discount.Name = dto.Name
	discount.Code = dto.Code
	discount.Quantity = dto.Quantity
	discount.RemainingQuantity = dto.Quantity
	discount.Type = dto.Type // Fix/Rebate atau Percent
	discount.Value = dto.Value

	if dto.StartDate != nil {
		discount.StartDate = dto.StartDate
	}

	if dto.EndDate != nil {
		discount.EndDate = dto.EndDate
	}

	data, err := usecase.repository.Update(*discount)

	if err != nil {
		return nil, err
	}

	return data, err
}

// UpdateRemainingQuantity implements DiscountUsecase
func (*discountUsecase) UpdateRemainingQuantity(id int, quantity int, operator string) (*entity.Discount, *response.Error) {
	panic("unimplemented")
}

func NewDiscountUseCase(repository repository.DiscountRepository) DiscountUsecase {
	return &discountUsecase{repository}
}
