package controllers

import (
	"net/http"
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
