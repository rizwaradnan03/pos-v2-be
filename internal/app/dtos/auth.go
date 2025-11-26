package dtos

import (
	"pos-v2-be/internal/enums"
	"pos-v2-be/internal/pkg/translate"
	"pos-v2-be/internal/pkg/validate"

	"github.com/google/uuid"
)

type AuthSignInDto struct {
	Username    string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type AuthSignInUpDto struct {
	Username    string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type RequestForgotPasswordDto struct {
	Username string `json:"username" validate:"required"`
}

type SendForgotPasswordDto struct {
	Username    string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type TokenDto struct {
	ID    uuid.UUID      `json:"id"`
	Username string         `json:"username"`
	Role  enums.RoleType `json:"role"`
	Exp   *int64         `json:"exp"`
}

type RefreshAccessToken struct {
	AccessToken string `json:"access_token"`
}

type OauthGoogleDto struct {
	Username string `json:"username"`
}

func ValidateRequestForgotPassword(dto RequestForgotPasswordDto) error {
	err := validate.Validator.Struct(dto)
	if err != nil {
		return translate.TranslateError(err)
	}
	return nil
}

func ValidateSendForgotPassword(dto SendForgotPasswordDto) error {
	err := validate.Validator.Struct(dto)
	if err != nil {
		return translate.TranslateError(err)
	}
	return nil
}

func ValidateAuthSignIn(dto AuthSignInDto) error {
	err := validate.Validator.Struct(dto)
	if err != nil {
		return translate.TranslateError(err)
	}
	return nil
}

func ValidateAuthSignUp(dto AuthSignInUpDto) error {
	err := validate.Validator.Struct(dto)
	if err != nil {
		return translate.TranslateError(err)
	}
	return nil
}
