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

		if req.ExistingSeriesID != nil {
			postId, err := createPost(tx, *req.ExistingSeriesID, req.PostSlug)
			if err != nil {
				return err
			} else {
				seriesSlug, err := getSeriesSlug(tx, *req.ExistingSeriesID)
				if err != nil {
					return err
				}

				markdownFilePath = utils.GetMarkdownRelaPath(req.LanguageCode, seriesSlug, req.PostSlug)

				err = createPostTranslation(tx, *postId, req.LanguageCode, req.Title, markdownFilePath)
				if err != nil {
					return err
				}
			}
		} else { // creates new seriest for this post

		}

		err := utils.WriteFile(markdownFilePath, req.MarkdownContent)
		if err != nil {
			return err
		}

		return nil
	})
}

func createPost(tx *gorm.DB, seriesId uint, postSlug string) (*uint, error) {
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

func createPostTranslation(tx *gorm.DB, postId uint, lang string, title string, markdownFilePath string) error {
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
