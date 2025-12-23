package dto

type SeriesListRequest struct {
	Topic string `form:"topic" binding:"required,oneof=techie reader"`
	Lang  string `form:"lang,default=en" binding:"oneof=en cn"`
}

type SeriesListResponse struct {
	BackgroundImgURL string `json:"backgroundImgUrl"`
	SeriesSlug       string `json:"seriesSlug"`
	Title            string `json:"title"`
	Description      string `json:"description"`
}
