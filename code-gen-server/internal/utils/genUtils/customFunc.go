package genUtils

import (
	"github.com/duke-git/lancet/v2/strutil"
	"html/template"
	"strings"
)

// CustomFunc 自定义函数
var (
	CustomFunc = template.FuncMap{
		//转换到驼峰后，首字母小写
		"lowerFirstCamel": func(str string) string {
			camelStr := strutil.CamelCase(str)
			first := camelStr[:1]
			remain := camelStr[1:]
			first = strings.ToLower(first)
			return first + remain
		},
		"isGormDeleteAt": func(str string) bool {
			//如果是deleteAt字段，那么类型变为gorm.DeletedAt
			if strings.Contains(str, "DeletedAt") ||
				strings.Contains(str, "deleted_at") ||
				strings.Contains(str, "deleteAt") {
				return true
			}
			return false
		},
	}
)
