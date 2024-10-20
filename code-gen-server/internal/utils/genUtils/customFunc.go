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
	}
)
