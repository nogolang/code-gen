package genCode

import (
	"github.com/duke-git/lancet/v2/strutil"
	"gorm.io/gorm"
	"strings"
)

type Table struct {
	DB                                 *gorm.DB
	TableName                          string //表名
	IdName                             string //当前表的id的名称和类型，id必须在第1个字段
	IdType                             string
	IdNameWithSmallCamel               string  //小驼峰形式的id名称
	IdNameWithBigCamel                 string  //大驼峰形式的id名称
	TableNameWithBigCamel              string  //大驼峰表名
	TableNameWithSmallCamel            string  //小驼峰表名
	TableComment                       string  //表的注解
	DataBaseName                       string  //表所在的数据库名称
	DataBaseNameWithNoPrefix           string  //去除前缀的数据库名称
	DataBaseNameWithNoPrefixSmallCamel string  //去除前缀的数据库名称,小驼峰
	DataBaseNameWithNoPrefixBigCamel   string  //去除前缀的数据库名称,大驼峰
	Fields                             []field //表的字段
}

// Field代表数据库的字段名称和类型
type field struct {
	FieldName               string //原始字段名，从规则上来说应该设计为蛇形命名
	FieldNameWithBigCamel   string //大驼峰字段名,UserInfo
	FieldNameWithSmallCamel string //小驼峰字段名,userInfo
	FieldType               string //字段类型

	//原始数据库类型，去掉了后面的括号的，比如varchar(255)，变成varchar，方便判断
	RawFieldType string
	FieldComment string //字段的注解
}

func NewTable(dataBaseName string, dataBaseNamePrefix string, tableName string, db *gorm.DB) *Table {
	return &Table{
		TableName:                          tableName,
		TableNameWithBigCamel:              strutil.UpperFirst(strutil.CamelCase(tableName)),
		TableNameWithSmallCamel:            strutil.CamelCase(tableName),
		DataBaseName:                       dataBaseName,
		DataBaseNameWithNoPrefix:           strings.ReplaceAll(dataBaseName, dataBaseNamePrefix, ""),
		DataBaseNameWithNoPrefixSmallCamel: strutil.CamelCase(strings.ReplaceAll(dataBaseName, dataBaseNamePrefix, "")),
		DataBaseNameWithNoPrefixBigCamel:   strutil.UpperFirst(strutil.CamelCase(strings.ReplaceAll(dataBaseName, dataBaseNamePrefix, ""))),
		DB:                                 db,
	}
}

func (receiver *Table) GetTable(resolveMap map[string]string) *Table {
	//获取表的注解
	receiver.fillTableComment()

	//获取字段的信息
	receiver.fillFields(resolveMap)
	return receiver
}

// 单独获取表的注解
func (receiver *Table) fillTableComment() {

	//获取表的注解
	tableSQL := `SELECT table_comment
	FROM information_schema.tables 
	WHERE table_schema=? AND table_name = ?;`

	var temp string
	receiver.DB.Raw(tableSQL, receiver.DataBaseName, receiver.TableName).Find(&temp)

	//把"表"这个后缀去掉，比如用户表，改为用户
	cutStr, _ := strings.CutSuffix(temp, "表")
	receiver.TableComment = cutStr
}

// 填充字段信息，根据mapping
func (receiver *Table) fillFields(mapping map[string]string) {
	var myFields []field

	//获取字段的信息，这里必须要ORDER BY ordinal_position，不然取出来是无序的
	fieldSQL := `
		SELECT column_name as FieldName,column_type as FieldType,column_comment as FieldComment
		FROM information_schema.columns
		WHERE table_schema = ? AND table_name = ? ORDER BY ordinal_position;
	`
	var fieldTemps []field
	receiver.DB.Raw(fieldSQL, receiver.DataBaseName, receiver.TableName).Find(&fieldTemps)

	for _, fieldTemp := range fieldTemps {
		//根据mapping改变字段类型，但是返回的type带有小括号
		//比如varchar(255),我们只需要255，此时我们小括号的位置，然后从这个位置往后截取掉
		cutIndex := strings.Index(fieldTemp.FieldType, "(")

		//如果找到了(，那就去掉，去掉后返回回去
		if cutIndex != -1 {
			pureType := fieldTemp.FieldType[:cutIndex]

			//赋值原始类型，外部则可以判断
			fieldTemp.RawFieldType = pureType

			//通过映射获取到映射类型
			fieldTemp.FieldType = mapping[pureType]
		} else {
			//如果没有找到，那么直接通过映射获取到映射类型即可
			fieldTemp.RawFieldType = fieldTemp.FieldType
			fieldTemp.FieldType = mapping[fieldTemp.FieldType]
		}

		//小驼峰，
		fieldTemp.FieldNameWithSmallCamel = strutil.CamelCase(fieldTemp.FieldName)

		//大驼峰
		fieldTemp.FieldNameWithBigCamel = strutil.UpperFirst(strutil.CamelCase(fieldTemp.FieldName))

		myFields = append(myFields, fieldTemp)

	}

	//赋值id的类型和字符串，确保id在第1个字段
	receiver.IdName = myFields[0].FieldName
	receiver.IdNameWithSmallCamel = strutil.CamelCase(myFields[0].FieldName)
	receiver.IdNameWithBigCamel = strutil.UpperFirst(strutil.CamelCase(myFields[0].FieldName))
	receiver.IdType = myFields[0].FieldType

	receiver.Fields = myFields
}
