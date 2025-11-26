package dtos

import (
	"pos-v2-be/internal/enums"
	"pos-v2-be/internal/pkg/translate"
	"pos-v2-be/internal/pkg/validate"

	"github.com/google/uuid"
)

type AuthSignInDto struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type AuthSignInUpDto struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type RequestForgotPasswordDto struct {
	Email string `json:"email" validate:"required"`
}

type SendForgotPasswordDto struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type TokenDto struct {
	ID    uuid.UUID      `json:"id"`
	Email string         `json:"email"`
	Role  enums.RoleType `json:"role"`
	Exp   *int64         `json:"exp"`
}

type RefreshAccessToken struct {
	AccessToken string `json:"access_token"`
}

type OauthGoogleDto struct {
	Email string `json:"email"`
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
