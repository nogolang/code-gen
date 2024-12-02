package model

type OrmModel struct {
	Id           int    `json:"id" gorm:"primaryKey;autoIncrement;"`
	Describe     string `json:"describe"`
	Host         string `json:"host"`
	Port         string `json:"port"`
	DataBaseName string `json:"dataBaseName"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	Prefix       string `json:"prefix"`
}
type OrmModelQuery struct {
	Page     int    `json:"page"`
	Size     int    `json:"size"`
	QueryStr string `json:"queryStr"`
}

type OrmModelsAll struct {
	Data  []OrmModel `json:"data"`
	Total int64      `json:"total"`
}
