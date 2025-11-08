package models

type Series struct {
	ID               uint   `gorm:"primaryKey" json:"id"`
	BackgroundImgURL string `gorm:"column:bg_url;size:2048" json:"backgroundImgUrl"`
}

type SeriesTranslation struct {
	SeriesID     uint   `gorm:"primaryKey" json:"-"`
	LanguageCode string `gorm:"primaryKey;size:2" json:"languageCode"`
	Title        string `gorm:"size:255" json:"title"`

	Series Series `gorm:"foreignKey:SeriesID;references:ID" json:"-"`
}
