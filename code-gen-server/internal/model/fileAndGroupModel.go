package model

type FileAndGroupModel struct {
	Id           int    `json:"id" gorm:"primaryKey;autoIncrement;"`
	FileId       int    `json:"fileId" gorm:"column:file_id"`
	GroupId      int    `json:"groupId" gorm:"column:group_id"`
	OutDir       string `json:"outDir" gorm:"column:out_dir"`
	TemplateName string `json:"templateName" gorm:"-"`

	//展示给前台的
	FileInfo *FileModel `json:"fileInfo" gorm:"-"`
}
