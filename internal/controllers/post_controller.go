package controllers

import (
	"net/http"
	"yyphan-pw/backend/internal/dto"
	"yyphan-pw/backend/internal/services"

	"github.com/gin-gonic/gin"
)

func CreatePost(c *gin.Context) {
	var req dto.CreatePostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "[CreatePost] invalid reqeust body: " + err.Error(),
		})
	}

	err := services.CreatePost(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "[CreatePost] error creating post: " + err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{"status": "success"})
}
