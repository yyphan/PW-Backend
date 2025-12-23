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

func UpsertPostTranslation(postId uint, req dto.UpsertPostTranslationRequest) error {
	return database.DB.Transaction(func(tx *gorm.DB) error {
		seriesSlug, err := getSeriesSlug(tx, postId)
		if err != nil {
			return err
		}

		post, err := getPost(tx, postId)
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
			return fmt.Errorf("[UpsertPostTranslation]error upserting post translation: %w", result.Error)
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

func getPost(tx *gorm.DB, postId uint) (models.Post, error) {
	var post models.Post
	if result := tx.First(&post, postId); result.Error != nil {
		return models.Post{}, fmt.Errorf("error getting post: %w", result.Error)
	}
	return post, nil
}
