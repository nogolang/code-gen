package model

type FileGenModel struct {
	Id         int    `json:"id" gorm:"primaryKey;autoIncrement;"`
	Describe   string `json:"describe" gorm:"column:describe;"`
	DataBaseId int    `json:"dataBaseId" gorm:"column:database_id;"`
	TableNames string `json:"tableNames" gorm:"column:table_name;"`
	GroupId    int    `json:"groupId" gorm:"column:group_id;"`

	//下面几个冗余字段，只是显示给前台
	GroupDescribe string   `json:"groupDescribe" gorm:"-"`
	DatabaseName  string   `json:"databaseName" gorm:"-"`
	TableNamesArr []string `json:"tableNamesArr" gorm:"-"`
}
type FileGenModelQuery struct {
	Page     int    `json:"page"`
	Size     int    `json:"size"`
	QueryStr string `json:"queryStr"`
}

type FileGenModelRequest struct {
	Describe      string   `json:"describe"`
	DataBaseId    int      `json:"dataBaseId" `
	TableNamesArr []string `json:"tableNamesArr"`
	GroupId       int      `json:"groupId"`
}

type FileGenModelsAll struct {
	Data  []*FileGenModel `json:"data"`
	Total int64           `json:"total"`
}
