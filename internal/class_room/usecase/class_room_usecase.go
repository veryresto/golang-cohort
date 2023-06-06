package class_room

import (
	"errors"

	dto "online-course/internal/class_room/dto"
	entity "online-course/internal/class_room/entity"
	repository "online-course/internal/class_room/repository"
	"online-course/pkg/response"

	"gorm.io/gorm"
)

type ClassRoomUsecase interface {
	FindAllByUserID(userId int, offset int, limit int) dto.ClassRoomListResponse
	FindOneByUserIDAndProductID(userId int, productId int) (*dto.ClassRoomResponseBody, *response.Error)
	Create(dto dto.ClassRoomRequestBody) (*entity.ClassRoom, *response.Error)
}

type classRoomUsecase struct {
	repository repository.ClassRoomRepository
}

// Create implements ClassRoomUsecase
func (usecase *classRoomUsecase) Create(classRoomRequestBody dto.ClassRoomRequestBody) (*entity.ClassRoom, *response.Error) {
	// Validasi apakah user id dan product id nya sudah ada
	dataClassRoom, err := usecase.repository.FindOneByUserIdAndProductId(int(classRoomRequestBody.UserID), int(classRoomRequestBody.ProductID))

	if err != nil && !errors.Is(err.Err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	if dataClassRoom != nil {
		return nil, &response.Error{
			Code: 400,
			Err:  errors.New("anda sudah masuk ke dalam class ini"),
		}
	}

	classRoom := entity.ClassRoom{
		UserID:      classRoomRequestBody.UserID,
		ProductID:   &classRoomRequestBody.ProductID,
		CreatedByID: &classRoomRequestBody.UserID,
	}

	data, err := usecase.repository.Create(classRoom)

	if err != nil {
		return nil, err
	}

	return data, nil
}

// FindAllByUserID implements ClassRoomUsecase
func (usecase *classRoomUsecase) FindAllByUserID(userId int, offset int, limit int) dto.ClassRoomListResponse {
	classRooms := usecase.repository.FindAllByUserID(userId, offset, limit)

	classRoomsResp := dto.CreateClassRoomListResponse(classRooms)

	return classRoomsResp
}

// FindOneByUserIDAndProductID implements ClassRoomUsecase
func (usecase *classRoomUsecase) FindOneByUserIDAndProductID(userId int, productId int) (*dto.ClassRoomResponseBody, *response.Error) {
	classRoom, _ := usecase.repository.FindOneByUserIdAndProductId(userId, productId)

	classRoomResp := dto.CreateClassRoomResponse(*classRoom)

	return &classRoomResp, nil
}

func NewClassRoomUseCase(repository repository.ClassRoomRepository) ClassRoomUsecase {
	return &classRoomUsecase{repository}
}
