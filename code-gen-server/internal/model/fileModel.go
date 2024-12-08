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

type FileModelAllRequest struct {
	FileModelAll []FileModel `json:"fileModelAll"`
}

type FileModel struct {
	Id int `json:"id" gorm:"primaryKey;autoIncrement;"`

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
	GroupId             int    `json:"groupId" gorm:"column:group_id"`
	GenPath             string `json:"genPath" gorm:"column:gen_path"`
}
