package infra

import (
	"pos-v2-be/internal/app/services"
	"pos-v2-be/internal/initial/intf"

	"gorm.io/gorm"
)

func NewService(r *intf.Repositories, db *gorm.DB) *intf.Services {
	authService := services.NewAuthService(r)


	return &intf.Services{
		AuthService:                 authService,
	}
}
