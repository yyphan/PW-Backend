package models

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Post struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	SeriesID    uint   `json:"seriesId"`
	IdxInSeries uint   `json:"idxInSeries"`
	PostSlug    string `gorm:"column:post_slug;size:255" json:"postSlug"`

	Series Series `gorm:"foreignKey:SeriesID;references:ID" json:"-"`
}

type PostTranslation struct {
	PostID           uint      `gorm:"primaryKey" json:"-"`
	LanguageCode     string    `gorm:"primaryKey;size:2" json:"languageCode"`
	Title            string    `gorm:"size:255" json:"title"`
	MarkdownFilePath string    `gorm:"size:2048" json:"-"`
	CreatedAt        time.Time `json:"createdAt"`
	UpdatedAt        time.Time `json:"updatedAt"`

	Post Post `gorm:"foreignKey:PostID;references:ID" json:"-"`
}

func InsertPost(tx *gorm.DB, seriesId uint, postSlug string) (*uint, error) {
	postCount, err := CountPostsInSeries(tx, seriesId)
	if err != nil {
		return nil, err
	}

	post := Post{
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

func InsertPostTranslation(tx *gorm.DB, postId uint, lang string, title string, markdownFilePath string) error {
	postTranslation := PostTranslation{
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

func GetPostById(tx *gorm.DB, postId uint) (Post, error) {
	var post Post
	if result := tx.First(&post, postId); result.Error != nil {
		return Post{}, fmt.Errorf("error getting post: %w", result.Error)
	}
	return post, nil
}
