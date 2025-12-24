package services

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"

	"github.com/google/uuid"
)

func UploadImage(file *multipart.FileHeader) (string, error) {
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	ext := filepath.Ext(file.Filename)
	newFileName := uuid.New().String() + ext

	now := time.Now()
	relDir := filepath.Join("public", "images", fmt.Sprintf("%d", now.Year()), fmt.Sprintf("%02d", now.Month()))

	if err := os.MkdirAll(relDir, 0755); err != nil {
		return "", err
	}

	dstPath := filepath.Join(relDir, newFileName)
	dst, err := os.Create(dstPath)
	if err != nil {
		return "", err
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return "", err
	}
	publicURL := fmt.Sprintf("/images/%d/%02d/%s", now.Year(), now.Month(), newFileName)

	return publicURL, nil
}
