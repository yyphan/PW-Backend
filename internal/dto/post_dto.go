package dto

import (
	"mime/multipart"
)

type GetPostRequest struct {
	LanguageCode string `form:"languageCode" binding:"required"`
	SeriesSlug   string `form:"seriesSlug" binding:"required"`
	PostSlug     string `form:"postSlug" binding:"required"`
}

type GetPostResponse struct {
	Title           string `json:"title"`
	UpdatedAt       string `json:"updatedAt"`
	MarkdownContent string `json:"markdownContent"`
}

// Metadata and file are sent as multipart form data
type CreatePostRequest struct {
	Data         string                `form:"data" binding:"required"`
	MarkdownFile *multipart.FileHeader `form:"markdownFile" binding:"required"`
}

// The normalized data parsed from the request body
type CreatePostData struct {
	PostSlug         string            `json:"postSlug"`
	LanguageCode     string            `json:"languageCode" binding:"required"`
	Title            string            `json:"title" binding:"required"`
	MarkdownContent  string            `json:"markdownContent"`
	ExistingSeriesID *uint             `json:"seriesId"`
	NewSeries        *NewSeriesRequest `json:"newSeries"`
}

type NewSeriesRequest struct {
	BackgroundImgURL string `json:"backgroundImgUrl"`
	Topic            string `json:"topic"`
	SeriesSlug       string `json:"seriesSlug"`
	Title            string `json:"title"`
	Description      string `json:"description"`
}

// Metadata and file are sent as multipart form data
type UpsertPostTranslationRequest struct {
	Data         string                `form:"data" binding:"required"`
	MarkdownFile *multipart.FileHeader `form:"markdownFile" binding:"required"`
}

// The normalized data parsed from the request body
type UpsertPostTranslationData struct {
	LanguageCode    string `json:"languageCode" binding:"required"`
	Title           string `json:"title" binding:"required"`
	MarkdownContent string `json:"markdownContent"`
}
