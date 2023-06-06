//go:build wireinject
// +build wireinject

package class_room

import (
	"github.com/google/wire"
	"gorm.io/gorm"
	handler "online-course/internal/class_room/delivery/http"
	repository "online-course/internal/class_room/repository"
	usecase "online-course/internal/class_room/usecase"
)

func InitializedService(db *gorm.DB) *handler.ClassRoomHandler {
	wire.Build(
		handler.NewClassRoomHandler,
		repository.NewClassRoomRepository,
		usecase.NewClassRoomUseCase,
	)

	return &handler.ClassRoomHandler{}
}
