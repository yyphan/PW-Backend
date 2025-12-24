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

	err := database.DB.
		Table("series").
		Select("series.bg_url, series.series_slug, series_translations.title, series_translations.description").
		Joins("INNER JOIN series_translations ON series_translations.series_id = series.id").
		Where("series.topic = ? AND series_translations.language_code = ?", topic, lang).
		Order("series_translations.created_at DESC").
		Scan(&response).Error

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

	return updateSeries(id, cleanUpdates)
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

func insertSeries(tx *gorm.DB, dto dto.NewSeriesRequest, lang string) (*uint, error) {
	series := models.Series{
		BackgroundImgURL: dto.BackgroundImgURL,
		SeriesSlug:       dto.SeriesSlug,
		Topic:            dto.Topic,
	}

	if result := tx.Create(&series); result.Error != nil {
		return nil, fmt.Errorf("error inserting into series: %w", result.Error)
	}

	seriesTranslation := models.SeriesTranslation{
		SeriesID:     series.ID, // successful insert above will fill ID
		LanguageCode: lang,
		Title:        dto.Title,
		Description:  dto.Description,
	}

	if result := tx.Create(&seriesTranslation); result.Error != nil {
		return nil, fmt.Errorf("error inserting into series_translations: %w", result.Error)
	}

	return &series.ID, nil
}

func countPostsInSeries(tx *gorm.DB, seriesId uint) (int64, error) {
	var existingPostsCount int64
	err := tx.Model(&models.Post{}).
		Where("series_id = ?", seriesId).
		Count(&existingPostsCount).Error
	if err != nil {
		return 0, fmt.Errorf("error counting posts in series: %w", err)
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

func updateSeries(id uint, updates map[string]interface{}) error {
	result := database.DB.Model(&models.Series{}).
		Where("id = ?", id).
		Updates(updates)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
