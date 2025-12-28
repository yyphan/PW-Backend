package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"strconv"
	"yyphan-pw/backend/internal/dto"
	"yyphan-pw/backend/internal/services"

	"github.com/gin-gonic/gin"
)

func GetPost(c *gin.Context) {
	var req dto.GetPostRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Errorf("[GetPost] invalid reqeust body: %w", err).Error(),
		})
		return
	}

	post, err := services.GetPost(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Errorf("[GetPost] error getting post: %w", err).Error(),
		})
		return
	}

	c.JSON(http.StatusOK, post)
}

func CreatePost(c *gin.Context) {
	fileHeader, err := c.FormFile("markdownFile")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "[CreatePost] Markdown file is required"})
		return
	}

	markdownContent, err := parseMarkdownFile(fileHeader)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Errorf("[CreatePost] error parsing markdown file: %w", err).Error()})
		return
	}

	var req dto.CreatePostRequest
	req.MarkdownContent = markdownContent

	metaDataStr := c.PostForm("data")
	if metaDataStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "[CreatePost] Metadata is required"})
		return
	}

	if err := json.Unmarshal([]byte(metaDataStr), &req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Errorf("[CreatePost] invalid reqeust body: %w", err).Error(),
		})
	}

	if req.PostSlug == "" {
		req.PostSlug = "index"
	}

	err = services.CreatePost(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Errorf("[CreatePost] error creating post: %w", err).Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

func UpsertPostTranslation(c *gin.Context) {
	var req dto.UpsertPostTranslationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Errorf("[UpsertPostTranslation] invalid reqeust body: %w", err).Error(),
		})
		return
	}

	postIdStr := c.Param("id")
	postId, err := strconv.ParseUint(postIdStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Errorf("[UpsertPostTranslation] invalid post id: %w", err).Error(),
		})
		return
	}

	err = services.UpsertPostTranslation(uint(postId), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Errorf("[UpsertPostTranslation] error upserting post translation: %w", err).Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

func parseMarkdownFile(fileHeader *multipart.FileHeader) (string, error) {
	file, err := fileHeader.Open()
	if err != nil {
		return "", fmt.Errorf("Failed to open file: %w", err)
	}
	defer file.Close()

	contentBytes, err := io.ReadAll(file)
	if err != nil {
		return "", fmt.Errorf("Failed to read file content: %w", err)
	}

	return string(contentBytes), nil
}
