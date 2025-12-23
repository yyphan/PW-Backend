package services

import (
	"fmt"
	"yyphan-pw/backend/internal/database"
	"yyphan-pw/backend/internal/dto"
	"yyphan-pw/backend/internal/models"
	"yyphan-pw/backend/internal/utils"

	"gorm.io/gorm"
)

// Also creates series if not exists
func CreatePost(req dto.CreatePostRequest) error {
	return database.DB.Transaction(func(tx *gorm.DB) error {
		var markdownFilePath string

		targetSeriesId := req.ExistingSeriesID
		if targetSeriesId == nil { // creates new series for this post

			if req.NewSeries == nil {
				return fmt.Errorf("error [CreatePost]: CreatePostRequest must provide either ExistingSeriesID or NewSeries")
			}

			var err error
			targetSeriesId, err = insertSeries(tx, *req.NewSeries, req.LanguageCode)
			if err != nil {
				return err
			}
		}

		postId, err := insertPost(tx, *targetSeriesId, req.PostSlug)
		if err != nil {
			return err
		} else {
			seriesSlug, err := getSeriesSlug(tx, *targetSeriesId)
			if err != nil {
				return err
			}

			markdownFilePath = utils.GetMarkdownRelaPath(req.LanguageCode, seriesSlug, req.PostSlug)

			err = insertPostTranslation(tx, *postId, req.LanguageCode, req.Title, markdownFilePath)
			if err != nil {
				return err
			}
		}

		err = utils.WriteFile(markdownFilePath, req.MarkdownContent)
		if err != nil {
			return err
		}

		return nil
	})
}

func insertPost(tx *gorm.DB, seriesId uint, postSlug string) (*uint, error) {
	postCount, err := countPostsInSeries(tx, seriesId)
	if err != nil {
		return nil, err
	}

	post := models.Post{
		SeriesID:    seriesId,
		PostSlug:    postSlug,
		IdxInSeries: uint(postCount),
	}

	result := tx.Create(&post)
	if result.Error != nil {
		return nil, fmt.Errorf("error inserting into posts: %w", result.Error)
	}

	return &post.ID, nil
}

func insertPostTranslation(tx *gorm.DB, postId uint, lang string, title string, markdownFilePath string) error {
	postTranslation := models.PostTranslation{
		PostID:           postId,
		LanguageCode:     lang,
		Title:            title,
		MarkdownFilePath: markdownFilePath,
	}

	err := tx.Create(&postTranslation).Error
	if err != nil {
		return fmt.Errorf("error inserting into post_translations: %w", err)
	}

	return nil
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
