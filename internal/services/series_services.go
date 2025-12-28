package services

import (
	"fmt"
	"yyphan-pw/backend/internal/database"
	"yyphan-pw/backend/internal/dto"
	"yyphan-pw/backend/internal/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func GetSeriesList(lang string, topic string) (dto.SeriesListResponse, error) {
	var response dto.SeriesListResponse

	response.Lang = lang
	response.Topic = topic

	err := database.DB.
		Table("series").
		Select("series.bg_url, series.series_slug, series_translations.title, series_translations.description").
		Joins("INNER JOIN series_translations ON series_translations.series_id = series.id").
		Where("series.topic = ? AND series_translations.language_code = ?", topic, lang).
		Order("series_translations.created_at DESC").
		Scan(&response.Series).Error

	return response, err
}

func PatchSeries(id uint, input map[string]interface{}) error {
	allowedFields := map[string]string{
		"backgroundImgUrl": "bg_url",
		"topic":            "topic",
		"seriesSlug":       "series_slug",
	}

	cleanUpdates := make(map[string]interface{})

	for jsonKey, dbCol := range allowedFields {
		if val, exists := input[jsonKey]; exists {
			cleanUpdates[dbCol] = val
		}
	}

	if len(cleanUpdates) == 0 {
		return nil
	}

	return models.UpdateSeries(id, cleanUpdates)
}

func UpsertSeriesTranslation(seriesId uint, req dto.UpsertSeriesTranslationRequest) error {
	return database.DB.Transaction(func(tx *gorm.DB) error {
		seriesTranslation := models.SeriesTranslation{
			SeriesID:     seriesId,
			LanguageCode: req.LanguageCode,
			Title:        req.Title,
			Description:  req.Description,
		}

		result := tx.Clauses(clause.OnConflict{
			UpdateAll: true,
		}).Create(&seriesTranslation)

		if result.Error != nil {
			return fmt.Errorf("[UpsertSeriesTranslation] error upserting series translation: %w", result.Error)
		}

		return nil
	})
}
