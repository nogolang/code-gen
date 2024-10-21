package model

type FileModelQuery struct {
	Page     int    `json:"page"`
	Size     int    `json:"size"`
	QueryStr string `json:"queryStr"`
}
type FileModesAll struct {
	Data  []*FileModel `json:"data"`
	Total int          `json:"total"`
}

type OutDirModel struct {
	Id          int    `json:"id" gorm:"primaryKey;autoIncrement;"`
	OutDir      string `json:"outDir" gorm:"column:out_dir"`
	IsExist     bool   `json:"isExist" gorm:"column:is_exist"`
	FileModelId int    `json:"fileModelId" gorm:"column:fileModel_id"`
}

type FileModelAllRequest struct {
	FileModelAll []FileModel `json:"fileModelAll"`
}

type FileModel struct {
	Id int `json:"id" gorm:"primaryKey;autoIncrement;"`
	//描述
	Describe string `json:"describe" gorm:"column:describe;"`

	//名称的后缀，比如叫_controller
	NameSuffix string `json:"nameSuffix" gorm:"column:name_suffix"`

	//文件的后后缀，比如.go
	FileSuffix string `json:"fileSuffix" gorm:"column:file_suffix"`

	//生成的文件名称是否是驼峰形式
	IsCamelCase int `json:"isCamelCase" gorm:"column:is_camel_case;"`

	//所属mapping
	MappingId int `json:"mappingId" gorm:"column:mapping_id"`

	TemplatePath        string `json:"templatePath" gorm:"column:template_path"`
	TemplatePathIsExist bool   `json:"templatePathIsExist" gorm:"column:template_path_isExist;default:1"`
}
