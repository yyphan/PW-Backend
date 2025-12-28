package services

import (
	"fmt"
	"yyphan-pw/backend/internal/database"
	"yyphan-pw/backend/internal/dto"
	"yyphan-pw/backend/internal/models"
	"yyphan-pw/backend/internal/utils"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// Also creates series if not exists
func CreatePost(req dto.CreatePostRequest) error {
	return database.DB.Transaction(func(tx *gorm.DB) error {
		var markdownFilePath string

		targetSeriesId := req.ExistingSeriesID
		if targetSeriesId == nil { // creates new series for this post

			if req.NewSeries == nil {
				return fmt.Errorf("error: [CreatePost] CreatePostRequest must provide either ExistingSeriesID or NewSeries")
			}

			var err error
			targetSeriesId, err = models.InsertSeries(tx, *req.NewSeries, req.LanguageCode)
			if err != nil {
				return err
			}
		}

		postId, err := models.InsertPost(tx, *targetSeriesId, req.PostSlug)
		if err != nil {
			return err
		} else {
			seriesSlug, err := models.GetSeriesSlug(tx, *targetSeriesId)
			if err != nil {
				return err
			}

			markdownFilePath = utils.GetMarkdownRelaPath(req.LanguageCode, seriesSlug, req.PostSlug)

			err = models.InsertPostTranslation(tx, *postId, req.LanguageCode, req.Title, markdownFilePath)
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

func UpsertPostTranslation(postId uint, req dto.UpsertPostTranslationRequest) error {
	return database.DB.Transaction(func(tx *gorm.DB) error {
		seriesSlug, err := models.GetSeriesSlug(tx, postId)
		if err != nil {
			return err
		}

		post, err := models.GetPostById(tx, postId)
		if err != nil {
			return err
		}

		markdownFilePath := utils.GetMarkdownRelaPath(req.LanguageCode, seriesSlug, post.PostSlug)

		newPostTranslation := models.PostTranslation{
			PostID:           postId,
			LanguageCode:     req.LanguageCode,
			Title:            req.Title,
			MarkdownFilePath: markdownFilePath,
		}

		result := tx.Clauses(clause.OnConflict{
			UpdateAll: true,
		}).Create(&newPostTranslation)

		if result.Error != nil {
			return fmt.Errorf("[UpsertPostTranslation] error upserting post translation: %w", result.Error)
		}

		err = utils.WriteFile(markdownFilePath, req.MarkdownContent)
		if err != nil {
			return err
		}

		return nil
	})
}
