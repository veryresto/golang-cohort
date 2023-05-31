//go:build wireinject
// +build wireinject

package admin

import (
	"github.com/google/wire"
	"gorm.io/gorm"
	handler "online-course/internal/admin/delivery/http"
	repository "online-course/internal/admin/repository"
	usecase "online-course/internal/admin/usecase"
)

func InitializedService(db *gorm.DB) *handler.AdminHandler {
	wire.Build(
		repository.NewAdminRepository,
		usecase.NewAdminUsecase,
		handler.NewAdminHandler,
	)

	return &handler.AdminHandler{}
}
