package serviceinterface

import (
	"pos-v2-be/internal/app/dtos"
	"pos-v2-be/internal/app/responses"
	"pos-v2-be/internal/schema"
)

type AuthServiceInterface interface {
	SignIn(input dtos.AuthSignInDto) (*responses.AuthSignInResponse, error)
	SignUp(input dtos.AuthSignInUpDto) (*schema.User, error)
	RefreshAccessToken(input dtos.RefreshAccessToken) (*responses.AuthSignInResponse, error)
}
