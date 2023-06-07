package admin

import (
	entity "online-course/internal/admin/entity"
	"online-course/pkg/response"
	"online-course/pkg/utils"

	"gorm.io/gorm"
)

type AdminRepository interface {
	FindAll(offset int, limit int) []entity.Admin
	FindOneById(id int) (*entity.Admin, *response.Error)
	FindOneByEmail(email string) (*entity.Admin, *response.Error)
	Create(entity entity.Admin) (*entity.Admin, *response.Error)
	Update(entity entity.Admin) (*entity.Admin, *response.Error)
	Delete(entity entity.Admin) *response.Error
	TotalCountAdmin() int64
}

type adminRepository struct {
	db *gorm.DB
}

// TotalCountAdmin implements AdminRepository
func (repository *adminRepository) TotalCountAdmin() int64 {
	var admin entity.Admin

	var totalAdmin int64

	repository.db.Model(&admin).Count(&totalAdmin)

	return totalAdmin
}

// Create implements AdminRepository
func (repository *adminRepository) Create(entity entity.Admin) (*entity.Admin, *response.Error) {
	if err := repository.db.Create(&entity).Error; err != nil {
		return nil, &response.Error{
			Code: 500,
			Err:  err,
		}
	}

	return &entity, nil
}

// Delete implements AdminRepository
func (repository *adminRepository) Delete(entity entity.Admin) *response.Error {
	if err := repository.db.Delete(&entity).Error; err != nil {
		return &response.Error{
			Code: 500,
			Err:  err,
		}
	}

	return nil
}

// FindAll implements AdminRepository
func (repository *adminRepository) FindAll(offset int, limit int) []entity.Admin {
	var admin []entity.Admin

	repository.db.Scopes(utils.Paginate(offset, limit)).Find(&admin)

	return admin
}

// FindOneByEmail implements AdminRepository
func (repository *adminRepository) FindOneByEmail(email string) (*entity.Admin, *response.Error) {
	var admin entity.Admin

	if err := repository.db.Where("email = ?", email).First(&admin).Error; err != nil {
		return nil, &response.Error{
			Code: 500,
			Err:  err,
		}
	}

	return &admin, nil
}

// FindOneById implements AdminRepository
func (repository *adminRepository) FindOneById(id int) (*entity.Admin, *response.Error) {
	var admin entity.Admin

	if err := repository.db.First(&admin, id).Error; err != nil {
		return nil, &response.Error{
			Code: 500,
			Err:  err,
		}
	}

	return &admin, nil
}

// Update implements AdminRepository
func (repository *adminRepository) Update(entity entity.Admin) (*entity.Admin, *response.Error) {
	if err := repository.db.Save(&entity).Error; err != nil {
		return nil, &response.Error{
			Code: 500,
			Err:  err,
		}
	}

	return &entity, nil
}

func NewAdminRepository(db *gorm.DB) AdminRepository {
	return &adminRepository{db}
}
