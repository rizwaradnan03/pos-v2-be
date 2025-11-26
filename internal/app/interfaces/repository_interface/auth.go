package repositoryinterface

import (
	"pos-v2-be/internal/app/dtos"
	"pos-v2-be/internal/app/responses"
	"pos-v2-be/internal/schema"
)

type AuthRepositoryInterface interface {
	IsEmailExist(email string) (bool, error)
	SignIn(input dtos.AuthSignInDto) (*responses.AuthSignInResponse, error)
	SignUp(input dtos.AuthSignInUpDto) (*schema.User, error)
	FindOneByEmail(email string) (*schema.User, error)

}
