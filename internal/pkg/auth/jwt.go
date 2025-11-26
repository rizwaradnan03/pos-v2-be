package auth

import (
	"fmt"
	"os"
	"pos-v2-be/internal/app/dtos"
	"pos-v2-be/internal/enums"
	"pos-v2-be/internal/pkg/validate"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

func JwtEncode(data dtos.TokenDto) (string, error) {
	expTime := time.Now().Add(time.Hour * 24).Unix()

	if data.Exp != nil {
		expTime = int64(*data.Exp)
	}

	claims := jwt.MapClaims{
		"id":    data.ID,
		"email": data.Email,
		"role":  data.Role,
		"exp":   expTime,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func JwtDecode(tokenString string) (dtos.TokenDto, error) {
	claims := jwt.MapClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil || !token.Valid {
		return dtos.TokenDto{}, fmt.Errorf("failed to decode jwt: %w", err)
	}

	returnedPayload := dtos.TokenDto{}

	if Id, ok := claims["id"].(string); ok {
		uid, err := uuid.Parse(Id)

		if err != nil {
			return dtos.TokenDto{}, err
		}
		returnedPayload.ID = uid
	}

	if Email, ok := claims["email"].(string); ok {
		returnedPayload.Email = Email
	}

	if exp, ok := claims["exp"].(int64); ok {
		returnedPayload.Exp = &exp
	}

	if roleStr, ok := claims["role"].(string); ok {
		returnedPayload.Role = enums.RoleType(roleStr)
	}

	return returnedPayload, nil
}

func JwtVerify(tokenString string) (bool, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		return false, err
	}

	return token.Valid, nil
}

func ExtractJwtAndDecode(c *gin.Context) (*dtos.TokenDto, error) {
	authHeader := c.GetHeader("Authorization")

	if authHeader == "" {
		c.Abort()

		return nil, validate.NewError("Gagal Verifikasi Token")
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		c.Abort()

		return nil, validate.NewError("Gagal Verifikasi Token")
	}

	token := parts[1]
	result, err := JwtVerify(token)
	if !result || err != nil {
		c.Abort()

		return nil, validate.NewError("Gagal Verifikasi Token")
	}

	decodedToken, err := JwtDecode(token)
	if err != nil {
		c.Abort()

		return nil, validate.NewError("Gagal Verifikasi Token")
	}

	return &decodedToken, nil
}
