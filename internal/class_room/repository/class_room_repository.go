package class_room

import (
	entity "online-course/internal/class_room/entity"
	"online-course/pkg/response"
	"online-course/pkg/utils"

	"gorm.io/gorm"
)

type ClassRoomRepository interface {
	FindAllByUserID(userId int, offset int, limit int) []entity.ClassRoom
	FindOneByUserIdAndProductId(userId int, productId int) (*entity.ClassRoom, *response.Error)
	Create(entity entity.ClassRoom) (*entity.ClassRoom, *response.Error)
}

type classRoomRepository struct {
	db *gorm.DB
}

// Create implements ClassRoomRepository
func (repository *classRoomRepository) Create(entity entity.ClassRoom) (*entity.ClassRoom, *response.Error) {
	if err := repository.db.Create(&entity).Error; err != nil {
		return nil, &response.Error{
			Code: 500,
			Err:  err,
		}
	}

	return &entity, nil
}

// FindAllByUserID implements ClassRoomRepository
func (repository *classRoomRepository) FindAllByUserID(userId int, offset int, limit int) []entity.ClassRoom {
	var classRooms []entity.ClassRoom

	repository.db.Scopes(utils.Paginate(offset, limit)).
		Preload("Product.ProductCategory").
		Where("user_id = ?", userId).
		Find(&classRooms)

	return classRooms
}

// FindOneByUserIdAndProductId implements ClassRoomRepository
func (repository *classRoomRepository) FindOneByUserIdAndProductId(userId int, productId int) (*entity.ClassRoom, *response.Error) {
	var classRoom entity.ClassRoom

	if err := repository.db.Preload("Product.ProductCategory").
		Where("user_id = ?", userId).
		Where("product_id = ?", productId).First(&classRoom).Error; err != nil {
		return nil, &response.Error{
			Code: 500,
			Err:  err,
		}
	}

	return &classRoom, nil
}

func NewClassRoomRepository(db *gorm.DB) ClassRoomRepository {
	return &classRoomRepository{db}
}
