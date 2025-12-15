package services

import (
	"yyphan-pw/backend/internal/database"
	"yyphan-pw/backend/internal/dto"
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
