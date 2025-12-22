package services

import (
	"fmt"
	"yyphan-pw/backend/internal/database"
	"yyphan-pw/backend/internal/dto"
	"yyphan-pw/backend/internal/models"

	"gorm.io/gorm"
)

func GetSeriesList(lang string, topic string) (dto.SeriesListResponse, error) {
	var response dto.SeriesListResponse

	err := database.DB.
		Table("series").
		Select("series.bg_url, series.series_slug, series_translations.title, series_translations.description").
		Joins("INNER JOIN series_translations ON series_translations.series_id = series.id").
		Where("series.topic = ? AND series_translations.language_code = ?", topic, lang).
		Order("series_translations.created_at DESC").
		Scan(&response).Error

	return response, err
}

func createSeries(tx *gorm.DB, backgroundImgUrl string, topic string, seriesSlug string) (uint, error) {
	return 0, nil
}

func createSeriesTranslation(tx *gorm.DB, seriesId uint, lang string, title string, description string) error {
	return nil
}

func countPostsInSeries(tx *gorm.DB, seriesId uint) (int64, error) {
	var existingPostsCount int64
	err := tx.Model(&models.Post{}).
		Where("series_id = ?", seriesId).
		Count(&existingPostsCount).Error
	if err != nil {
		return 0, fmt.Errorf("error couting posts in series: %w", err)
	}

	return existingPostsCount, nil
}

func getSeriesSlug(tx *gorm.DB, seriesID uint) (string, error) {
	var slug string

	err := tx.Model(&models.Series{}).
		Select("series_slug").
		Where("id = ?", seriesID).
		Take(&slug).Error

	if err != nil {
		return "", fmt.Errorf("error reading series_slug: %w", err)
	}

	return slug, nil
}
