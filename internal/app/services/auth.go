package services

import (
	"fmt"
	"pos-v2-be/internal/app/dtos"
	serviceinterface "pos-v2-be/internal/app/interfaces/service_interface"
	"pos-v2-be/internal/app/responses"
	"pos-v2-be/internal/initial/intf"
	"pos-v2-be/internal/pkg/auth"
	"pos-v2-be/internal/pkg/number"
	"pos-v2-be/internal/pkg/validate"
	"pos-v2-be/internal/schema"
	"time"
)

type authService struct {
	repo *intf.Repositories
}

func NewAuthService(r *intf.Repositories) serviceinterface.AuthServiceInterface {
	return &authService{repo: r}
}

func (s *authService) SignIn(input dtos.AuthSignInDto) (*responses.AuthSignInResponse, error) {
	err := dtos.ValidateAuthSignIn(input)
	if err != nil {
		return nil, err
	}

	sign, err := s.repo.AuthRepository.SignIn(input)
	if err != nil || sign == nil {
		return nil, validate.NewError("Email / Password Salah!")
	}

	return sign, nil
}

func (s *authService) SignUp(input dtos.AuthSignInUpDto) (*schema.User, error) {
	return s.repo.AuthRepository.SignUp(input)
}

func (s *authService) RefreshAccessToken(input dtos.RefreshAccessToken) (*responses.AuthSignInResponse, error) {
	decoded, err := auth.JwtDecode(input.AccessToken)
	if err != nil {
		return nil, validate.NewError("Gagal decode access token!")
	}

	fmt.Println("email se : ", decoded.Username)

	account, err := s.repo.AuthRepository.FindOneByEmail(decoded.Username)
	if err != nil {
		return nil, validate.NewError("Email tidak ditemukan!")
	}

	if account.RefreshToken == nil {
		return nil, validate.NewError("Refresh token tidak ditemukan untuk akun ini!")
	}

	decodedRefreshToken, err := auth.JwtDecode(*account.RefreshToken)
	if err != nil {
		return nil, validate.NewError("Gagal decode access token!")
	}

	if *decodedRefreshToken.Exp < number.TimeToNumber(time.Now()) {
		return nil, validate.NewError("Sesi telah habis!")
	}

	return s.SignIn(dtos.AuthSignInDto{
		Username:    account.Username,
		Password: *account.Password,
	})
}
