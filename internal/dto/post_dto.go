package dto

type CreatePostRequest struct {
	PostSlug         string            `json:"postSlug" binding:"required"`
	LanguageCode     string            `json:"languageCode" binding:"required"`
	Title            string            `json:"title" binding:"required"`
	MarkdownContent  string            `json:"markdownContent" binding:"required"`
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

type UpsertPostTranslationRequest struct {
	LanguageCode    string `json:"languageCode" binding:"required"`
	Title           string `json:"title" binding:"required"`
	MarkdownContent string `json:"markdownContent" binding:"required"`
}
