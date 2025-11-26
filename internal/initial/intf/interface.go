package intf

import (
	"pos-v2-be/internal/app/http/controllers"
	repositoryinterface "pos-v2-be/internal/app/interfaces/repository_interface"
	serviceinterface "pos-v2-be/internal/app/interfaces/service_interface"
)

type Controllers struct {
	AuthController                 controllers.AuthController
}

type Services struct {
	AuthService                 serviceinterface.AuthServiceInterface

}

type Repositories struct {
	AuthRepository                 repositoryinterface.AuthRepositoryInterface
}
