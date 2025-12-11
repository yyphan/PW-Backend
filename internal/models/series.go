package models

type Series struct {
	ID               uint   `gorm:"primaryKey" json:"id"`
	BackgroundImgURL string `gorm:"column:bg_url;size:2048" json:"backgroundImgUrl"`
	Topic            string `gorm:"index:idx_series_topic;column:topic;size:255" json:"topic"`
	SeriesSlug       string `gorm:"column:series_slug;size:255" json:"seriesSlug"`
}

type SeriesTranslation struct {
	SeriesID     uint   `gorm:"primaryKey" json:"-"`
	LanguageCode string `gorm:"primaryKey;size:2" json:"languageCode"`
	Title        string `gorm:"size:255" json:"title"`

	Series Series `gorm:"foreignKey:SeriesID;references:ID" json:"-"`
}
