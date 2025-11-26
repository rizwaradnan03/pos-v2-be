package infra

import (
	"pos-v2-be/internal/app/http/controllers"
	"pos-v2-be/internal/initial/intf"
)

func NewController(service *intf.Services) *intf.Controllers {
	authController := controllers.NewAuthController(service.AuthService)

	return &intf.Controllers{
		AuthController:                 *authController,
	}
}
