package dto

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

type CreatePostRequest struct {
	PostSlug         string            `json:"postSlug"`
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
