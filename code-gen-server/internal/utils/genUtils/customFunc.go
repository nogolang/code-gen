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
				strings.Contains(str, "DeleteAt") ||
				strings.Contains(str, "deleted_at") ||
				strings.Contains(str, "deletedAt") {
				return true
			}
			return false
		},
		"isGormCreateAt": func(str string) bool {
			if strings.Contains(str, "CreatedAt") ||
				strings.Contains(str, "CreateAt") ||
				strings.Contains(str, "created_at") ||
				strings.Contains(str, "create_at") {
				return true
			}
			return false
		},
		"isGormUpdateAt": func(str string) bool {
			if strings.Contains(str, "UpdatedAt") ||
				strings.Contains(str, "UpdateAt") ||
				strings.Contains(str, "updated_at") ||
				strings.Contains(str, "update_at") {
				return true
			}
			return false
		},
		"isJsNumberType": func(str string) bool {
			//判断类型是string还是number，因为vue里对于数值类型需要用v-model.number
			if strings.Contains(str, "bigint") ||
				strings.Contains(str, "int") ||
				strings.Contains(str, "float") ||
				strings.Contains(str, "double") ||
				strings.Contains(str, "tinyint") {
				return true
			}
			return false
		},
		//添加大括号，比如在proto里要生成/{id}，那么模板里就是{{{IdName}}}
		//但这会语法错误，所以需要我们用函数添加大括号
		//{{addBrace $table.IdName}}，这样最终就会变成 {id}
		"addBrace": func(str string) string {
			return "{" + str + "}"
		},
	}
)
