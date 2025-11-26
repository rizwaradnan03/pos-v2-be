package str

import (
	"encoding/base64"
	"pos-v2-be/internal/pkg/auth"
	"pos-v2-be/internal/pkg/validate"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func ComparePassword(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	return err == nil
}

func StringToUUID(str string) (uuid.UUID, error) {
	uid, err := uuid.Parse(str)
	if err != nil {
		return uuid.Nil, err
	}

	return uid, nil
}

func UUIDToString(id uuid.UUID) string {
	return id.String()
}

func GetIdFromToken(c *gin.Context) (*uuid.UUID, error) {
	decodedToken, err := auth.ExtractJwtAndDecode(c)
	if err != nil {
		return nil, err
	}

	return &decodedToken.ID, nil
}

func GetIpFromHeader(c *gin.Context) string {
	ip := c.ClientIP()

	return ip
}

func DecodeBase64ToBytes(base64Str string) ([]byte, string, error) {
	if base64Str == "" {
		return nil, "", validate.NewError("base64 string kosong")
	}

	parts := strings.Split(base64Str, ",")
	if len(parts) < 2 {
		return nil, "", validate.NewError("format base64 tidak valid")
	}

	dataPart := parts[1]

	mimeType := strings.Split(parts[0], ";")[0]
	ext := ""
	if strings.HasPrefix(mimeType, "data:image/") {
		ext = strings.TrimPrefix(mimeType, "data:image/")
	} else if strings.HasPrefix(mimeType, "data:application/") {
		ext = strings.TrimPrefix(mimeType, "data:application/")
	} else {
		ext = "bin"
	}

	data, err := base64.StdEncoding.DecodeString(dataPart)
	if err != nil {
		return nil, "", err
	}

	return data, ext, nil
}
