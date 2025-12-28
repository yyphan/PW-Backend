package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"yyphan-pw/backend/internal/dto"
	"yyphan-pw/backend/internal/services"

	"github.com/gin-gonic/gin"
)

func GetSeriesList(c *gin.Context) {
	var req dto.SeriesListRequest

	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Errorf("[GetSeriesList] Invalid parameters: %w", err),
		})
		return
	}

	data, err := services.GetSeriesList(req.Lang, req.Topic)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Errorf("[GetSeriesList] Error querying DB: %w", err),
		})
		return
	}
	c.JSON(http.StatusOK, data)
}

func PatchSeries(c *gin.Context) {
	seriesIdStr := c.Param("id")
	seriesId, err := strconv.ParseUint(seriesIdStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Errorf("[PatchSeries] invalid series id: %w", err),
		})
		return
	}

	var updates map[string]interface{}
	if err := c.ShouldBindJSON(&updates); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Errorf("[PatchSeries] Invalid parameters: %w", err),
		})
		return
	}

	err = services.PatchSeries(uint(seriesId), updates)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Errorf("[PatchSeries] Error updating series: %w", err),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

func UpsertSeriesTranslation(c *gin.Context) {
	var req dto.UpsertSeriesTranslationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Errorf("[UpsertSeriesTranslation] Invalid parameters: %w", err),
		})
		return
	}

	seriesIdStr := c.Param("id")
	seriesId, err := strconv.ParseUint(seriesIdStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Errorf("[UpsertSeriesTranslation] invalid series id: %w", err),
		})
		return
	}

	err = services.UpsertSeriesTranslation(uint(seriesId), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Errorf("[UpsertSeriesTranslation] Error upserting series translation: %w", err),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}
