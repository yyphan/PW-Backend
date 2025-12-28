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

const ImgRootFolderKey = "LOCAL_IMAGE_ROOT"

func UploadImage(file *multipart.FileHeader) (string, error) {
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	ext := filepath.Ext(file.Filename)
	newFileName := uuid.New().String() + ext

	imgRootRela := os.Getenv(ImgRootFolderKey)
	absRoot, err := filepath.Abs(imgRootRela)
	if err != nil {
		return "", err
	}

	now := time.Now()
	dstFolder := filepath.Join(absRoot, fmt.Sprintf("%d", now.Year()), fmt.Sprintf("%02d", now.Month()))

	if err := os.MkdirAll(dstFolder, 0755); err != nil {
		return "", err
	}

	dstPath := filepath.Join(dstFolder, newFileName)
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
