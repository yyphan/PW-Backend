package models

func GetAllDBModels() []interface{} {
	return []interface{}{
		&Post{},
		&Series{},
		&PostTranslation{},
		&SeriesTranslation{},
	}
}
