package file

import (
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
)

func UploadFile(file *multipart.FileHeader, uploadPath string) (*string, error) {
	idFile := uuid.New()
	fileName := idFile.String() + filepath.Ext(file.Filename)
	fullPath := filepath.Join(uploadPath, fileName)

	fullPath = strings.ReplaceAll(fullPath, "\\", "/")

	src, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer src.Close()

	dst, err := os.Create(fullPath)
	if err != nil {
		return nil, err
	}
	defer dst.Close()

	if _, err := io.Copy(dst, src); err != nil {
		return nil, err
	}

	return &fullPath, nil
}

func UploadBytesFile(data []byte, ext string, uploadPath string) (*string, error) {
	if len(data) == 0 {
		// return nil, ("data file kosong")
	}

	if ext == "" {
		ext = "bin"
	}

	idFile := uuid.New()
	fileName := idFile.String() + "." + ext
	fullPath := filepath.Join(uploadPath, fileName)

	fullPath = strings.ReplaceAll(fullPath, "\\", "/")

	err := os.WriteFile(fullPath, data, 0644)
	if err != nil {
		return nil, err
	}

	return &fullPath, nil
}
