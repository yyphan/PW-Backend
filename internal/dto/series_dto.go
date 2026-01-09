package dto

type SeriesListRequest struct {
	Topic string `form:"topic" binding:"required,oneof=techie reader"`
	Lang  string `form:"lang,default=en" binding:"oneof=en cn"`
}

type SeriesListResponse struct {
	Topic  string          `json:"topic"`
	Lang   string          `json:"lang"`
	Series []SeriesCardDto `json:"series"`
}

type SeriesCardDto struct {
	BgUrl       string   `json:"bgUrl"`
	SeriesSlug  string   `json:"seriesSlug"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	PostSlugs   []string `json:"postSlugs"`
}

type UpsertSeriesTranslationRequest struct {
	LanguageCode string `json:"languageCode"`
	Title        string `json:"title"`
	Description  string `json:"description"`
}
