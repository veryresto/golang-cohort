package admin

import (
	dto "online-course/internal/admin/dto"
	entity "online-course/internal/admin/entity"
	repository "online-course/internal/admin/repository"
	"online-course/pkg/response"

	"golang.org/x/crypto/bcrypt"
)

type AdminUsecase interface {
	FindAll(offset int, limit int) []entity.Admin
	FindOneById(id int) (*entity.Admin, *response.Error)
	FindOneByEmail(email string) (*entity.Admin, *response.Error)
	Create(dto dto.AdminRequestBody) (*entity.Admin, *response.Error)
	Update(id int, dto dto.AdminRequestBody) (*entity.Admin, *response.Error)
	Delete(id int) *response.Error
}

type adminUsecase struct {
	repository repository.AdminRepository
}

// Create implements AdminUsecase
func (usecase *adminUsecase) Create(dto dto.AdminRequestBody) (*entity.Admin, *response.Error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(*dto.Password), bcrypt.DefaultCost)

	if err != nil {
		return nil, &response.Error{
			Code: 500,
			Err:  err,
		}
	}

	dataAdmin := entity.Admin{
		Email:    dto.Email,
		Name:     dto.Name,
		Password: string(hashedPassword),
	}

	admin, errCreateAdmin := usecase.repository.Create(dataAdmin)

	if errCreateAdmin != nil {
		return nil, errCreateAdmin
	}

	return admin, nil
}

// Delete implements AdminUsecase
func (usecase *adminUsecase) Delete(id int) *response.Error {
	// Search admin berdasarkan id
	admin, err := usecase.repository.FindOneById(id)

	if err != nil {
		return err
	}

	if err := usecase.repository.Delete(*admin); err != nil {
		return err
	}

	return nil
}

// FindAll implements AdminUsecase
func (usecase *adminUsecase) FindAll(offset int, limit int) []entity.Admin {
	return usecase.repository.FindAll(offset, limit)
}

// FindOneByEmail implements AdminUsecase
func (usecase *adminUsecase) FindOneByEmail(email string) (*entity.Admin, *response.Error) {
	return usecase.repository.FindOneByEmail(email)
}

// FindOneById implements AdminUsecase
func (usecase *adminUsecase) FindOneById(id int) (*entity.Admin, *response.Error) {
	return usecase.repository.FindOneById(id)
}

// Update implements AdminUsecase
func (usecase *adminUsecase) Update(id int, dto dto.AdminRequestBody) (*entity.Admin, *response.Error) {
	// Search admin berdasarkan id
	admin, err := usecase.repository.FindOneById(id)

	if err != nil {
		return nil, err
	}

	admin.Name = dto.Name

	// Validasi jika email dari admin tidak sama maka akan di update
	if admin.Email != dto.Email {
		admin.Email = dto.Email
	}

	if dto.Password != nil {
		hashedPassword, errHashedPassword := bcrypt.GenerateFromPassword([]byte(*dto.Password), bcrypt.DefaultCost)

		if errHashedPassword != nil {
			return nil, &response.Error{
				Code: 500,
				Err:  errHashedPassword,
			}
		}

		admin.Password = string(hashedPassword)
	}

	updateAdmin, err := usecase.repository.Update(*admin)

	if err != nil {
		return nil, err
	}

	return updateAdmin, nil
}

func NewAdminUsecase(repository repository.AdminRepository) AdminUsecase {
	return &adminUsecase{repository}
}
