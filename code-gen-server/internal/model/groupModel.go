package model

type GroupModelQuery struct {
	Page     int    `json:"page"`
	Size     int    `json:"size"`
	QueryStr string `json:"queryStr"`
}

type GroupModelsAll struct {
	Data  []*GroupModel `json:"data"`
	Total int64         `json:"total"`
}

type GroupModel struct {
	Id            int    `json:"id" gorm:"primaryKey;autoIncrement;"`
	Describe      string `json:"describe" gorm:"column:describe;"`
	GenRootDir    string `json:"rootDir" gorm:"column:gen_root_dir"`
	SearchRootDir string `json:"searchRootDir" gorm:"column:search_root_dir"`

	//返回给前台的文件信息，同时也会用来接收
	FileModels []*FileModel `json:"fileModels" gorm:"-"`
}
