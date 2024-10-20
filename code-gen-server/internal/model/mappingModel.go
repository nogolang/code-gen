package model

type MappingPathModel struct {
	Id       int    `json:"id" gorm:"primaryKey;autoIncrement;"`
	Describe string `json:"describe" gorm:"column:describe;"`
	Content  string `json:"content" gorm:"column:content;"`
}
type MappingPathModelQuery struct {
	Page     int    `json:"page"`
	Size     int    `json:"size"`
	QueryStr string `json:"queryStr"`
}

type MappingPathModelsAll struct {
	Data  []MappingPathModel `json:"data"`
	Total int64              `json:"total"`
}
