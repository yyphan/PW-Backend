package controllers

import (
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
			"error": "[GetSeriesList] Invalid parameters: " + err.Error(),
		})
		return
	}

	data, err := services.GetSeriesList(req.Lang, req.Topic)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "[GetSeriesList] Error querying DB " + err.Error(),
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
			"error": "[PatchSeries] invalid series id: " + err.Error(),
		})
		return
	}

	var updates map[string]interface{}
	if err := c.ShouldBindJSON(&updates); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "[PatchSeries] Invalid parameters: " + err.Error(),
		})
		return
	}

	err = services.PatchSeries(uint(seriesId), updates)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "[PatchSeries] Error updating series " + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}
