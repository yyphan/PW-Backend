package models

import "time"

type Post struct {
	ID          uint  `gorm:"primaryKey" json:"id"`
	SeriesID    *uint `json:"seriesId,omitempty"`
	IdxInSeries *uint `json:"idxInSeries,omitempty"`

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
