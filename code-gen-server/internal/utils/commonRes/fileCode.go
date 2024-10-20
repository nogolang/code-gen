package commonRes

import "net/http"

var (
	FileGenOK = NewResponse(http.StatusOK, "文件生成成功")

	FileParamInvalid = NewResponse(10001, "页面参数错误")
	FileFormatError  = NewResponse(10002, "文件格式错误，正确的是xxx.go.gohtml")

	FileCreateDirError  = NewResponse(10101, "创建文件夹错误")
	FileCreateFileError = NewResponse(10102, "创建文件错误")

	FileCreateTemplateError = NewResponse(10201, "创建模板错误")
	FileTemplateParseError  = NewResponse(10202, "模板解析出错")
	FileMappingError        = NewResponse(10203, "读取映射文件出错")
)

// 数据库相关
var (
	DataBaseInitOK     = NewResponse(http.StatusOK, "数据库初始化成功")
	DataBaseInitError  = NewResponse(10301, "数据库初始化错误")
	DataBaseNotInitYet = NewResponse(10302, "数据库还未初始化")
	DataBaseFailStatus = NewResponse(10303, "请先连接数据库")
)

var (
	FileDataReadError  = NewResponse(10401, "数据文件读取错误")
	FileDataNotFount   = NewResponse(10402, "无文件配置数据，请配置")
	FileDataParseError = NewResponse(10403, "数据文件解析到对象出错，查看是否有破坏性修改")
)

var (
	FileMappingNotFount      = NewResponse(10501, "mapping数据找不到")
	FileOutDirDataNotFount   = NewResponse(10502, "输出目录数据找不到")
	FileTemplateDataNotFount = NewResponse(10502, "模板文件数据找不到")
)

var (
	GroupNotFount = NewResponse(10701, "组数据查询不到")
)
