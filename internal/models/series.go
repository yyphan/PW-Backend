package models

import (
	"fmt"
	"time"
	"yyphan-pw/backend/internal/database"
	"yyphan-pw/backend/internal/dto"

	"gorm.io/gorm"
)

type Series struct {
	ID               uint   `gorm:"primaryKey" json:"id"`
	BackgroundImgURL string `gorm:"column:bg_url;size:2048" json:"backgroundImgUrl"`
	Topic            string `gorm:"index:idx_series_topic;column:topic;size:255" json:"topic"`
	SeriesSlug       string `gorm:"column:series_slug;size:255" json:"seriesSlug"`
}

type SeriesTranslation struct {
	SeriesID     uint      `gorm:"primaryKey" json:"-"`
	LanguageCode string    `gorm:"primaryKey;size:2" json:"languageCode"`
	Title        string    `gorm:"size:255" json:"title"`
	Description  string    `gorm:"size:255" json:"description"`
	CreatedAt    time.Time `json:"createdAt"`

	Series Series `gorm:"foreignKey:SeriesID;references:ID" json:"-"`
}

func InsertSeries(tx *gorm.DB, dto dto.NewSeriesRequest, lang string) (*uint, error) {
	series := Series{
		BackgroundImgURL: dto.BackgroundImgURL,
		SeriesSlug:       dto.SeriesSlug,
		Topic:            dto.Topic,
	}

	if result := tx.Create(&series); result.Error != nil {
		return nil, fmt.Errorf("error inserting into series: %w", result.Error)
	}

	seriesTranslation := SeriesTranslation{
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

func CountPostsInSeries(tx *gorm.DB, seriesId uint) (int64, error) {
	var existingPostsCount int64
	err := tx.Model(&Post{}).
		Where("series_id = ?", seriesId).
		Count(&existingPostsCount).Error
	if err != nil {
		return 0, fmt.Errorf("error counting posts in series: %w", err)
	}

	return existingPostsCount, nil
}

func GetSeriesSlug(tx *gorm.DB, seriesID uint) (string, error) {
	var slug string

	err := tx.Model(&Series{}).
		Select("series_slug").
		Where("id = ?", seriesID).
		Take(&slug).Error

	if err != nil {
		return "", fmt.Errorf("error reading series_slug: %w", err)
	}

	return slug, nil
}

func UpdateSeries(id uint, updates map[string]interface{}) error {
	result := database.DB.Model(&Series{}).
		Where("id = ?", id).
		Updates(updates)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
