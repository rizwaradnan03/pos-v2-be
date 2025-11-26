package infra

import (
	"pos-v2-be/internal/app/repositories"
	"pos-v2-be/internal/initial/intf"

	"gorm.io/gorm"
)

func NewRepository(db *gorm.DB) *intf.Repositories {
	authRepository := repositories.NewAuthRepository(db)

	return &intf.Repositories{
		AuthRepository:                 authRepository,
	}
}
