package services

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

const ImgRootFolderKey = "LOCAL_IMAGE_ROOT"

func UploadImage(file *multipart.FileHeader) (string, error) {
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	ext := filepath.Ext(file.Filename)
	rawName := strings.TrimSuffix(filepath.Base(file.Filename), ext)

	cleanName := sanitizeFilename(rawName)
	timestamp := time.Now().Format("20060102")
	newFileName := fmt.Sprintf("%s-%s%s", timestamp, cleanName, ext)

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
	publicURL := fmt.Sprintf("![%s](/images/%d/%02d/%s)", rawName, now.Year(), now.Month(), newFileName)

	return publicURL, nil
}

func sanitizeFilename(name string) string {
	name = strings.ReplaceAll(name, " ", "-")
	reg := regexp.MustCompile(`[^A-Za-z0-9-_]+`)
	name = reg.ReplaceAllString(name, "")

	if len(name) > 50 {
		name = name[:50]
	}

	if name == "" {
		name = "image"
	}

	return name
}
