package controllers

import (
	"fmt"
	"net/http"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func UploadImage(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"[UploadImage] error": "No file uploaded"})
		return
	}

	ext := filepath.Ext(file.Filename)
	newFileName := uuid.New().String() + ext

	now := time.Now()
	dirPath := fmt.Sprintf("public/images/%d/%02d", now.Year(), now.Month())

	dst := filepath.Join(dirPath, newFileName)
	if err := c.SaveUploadedFile(file, dst); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"[UploadImage] error": "Failed to save file"})
		return
	}

	publicURL := fmt.Sprintf("/images/%d/%02d/%s", now.Year(), now.Month(), newFileName)

	c.JSON(http.StatusOK, gin.H{
		"url": publicURL,
	})
}
