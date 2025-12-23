package controllers

import (
	"net/http"
	"strconv"
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

func UpsertPostTranslation(c *gin.Context) {
	var req dto.UpsertPostTranslationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "[UpsertPostTranslation] invalid reqeust body: " + err.Error(),
		})
		return
	}

	postIdStr := c.Param("id")
	postId, err := strconv.ParseUint(postIdStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "[UpsertPostTranslation] invalid post id: " + err.Error(),
		})
		return
	}

	err = services.UpsertPostTranslation(uint(postId), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "[UpsertPostTranslation] error upserting post translation: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success"})
}
