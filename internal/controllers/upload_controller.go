package controllers

import (
	"net/http"
	"yyphan-pw/backend/internal/services"

	"github.com/gin-gonic/gin"
)

func UploadImages(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "[UploadImages] Invalid form"})
		return
	}

	files := form.File["files"]
	var urls []string

	for _, f := range files {
		url, err := services.UploadImage(f)
		if err != nil {
			continue
		}
		urls = append(urls, url)
	}

	c.JSON(http.StatusOK, gin.H{"urls": urls})
}
