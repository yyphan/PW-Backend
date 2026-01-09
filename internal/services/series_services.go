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

	// 1. Fetch Series basic info
	err := database.DB.
		Table("series").
		Select("series.bg_url, series.series_slug, series_translations.title, series_translations.description").
		Joins("INNER JOIN series_translations ON series_translations.series_id = series.id").
		Where("series.topic = ? AND series_translations.language_code = ?", topic, lang).
		Order("series_translations.created_at DESC").
		Scan(&response.Series).Error

	if err != nil {
		return response, err
	}

	if len(response.Series) == 0 {
		return response, nil
	}

	// 2. Fetch all post slugs for these series
	var seriesSlugs []string
	for _, s := range response.Series {
		seriesSlugs = append(seriesSlugs, s.SeriesSlug)
	}

	type PostSlugResult struct {
		SeriesSlug string
		PostSlug   string
	}
	var postResults []PostSlugResult

	err = database.DB.
		Table("posts").
		Select("posts.post_slug, series.series_slug").
		Joins("JOIN series ON posts.series_id = series.id").
		Where("series.series_slug IN ?", seriesSlugs).
		Order("posts.idx_in_series ASC").
		Scan(&postResults).Error

	if err != nil {
		return response, fmt.Errorf("error fetching post slugs: %w", err)
	}

	// 3. Map posts to series
	postsMap := make(map[string][]string)
	for _, p := range postResults {
		postsMap[p.SeriesSlug] = append(postsMap[p.SeriesSlug], p.PostSlug)
	}

	for i := range response.Series {
		if slugs, ok := postsMap[response.Series[i].SeriesSlug]; ok {
			response.Series[i].PostSlugs = slugs
		} else {
			response.Series[i].PostSlugs = []string{}
		}
	}

	return response, nil
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
