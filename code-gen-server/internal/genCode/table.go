package genCode

import (
	"github.com/duke-git/lancet/v2/strutil"
	"gorm.io/gorm"
	"strings"
)

type Table struct {
	DB                      *gorm.DB
	TableName               string  //表名
	TableNameWithBigCamel   string  //大驼峰表名
	TableNameWithSmallCamel string  //小驼峰表名
	TableComment            string  //表的注解
	DataBaseName            string  //表所在的数据库名称
	Fields                  []field //表的字段
}

// Field代表数据库的字段名称和类型
type field struct {
	FieldName               string //原始字段名，从规则上来说应该设计为蛇形命名
	FieldNameWithBigCamel   string //大驼峰字段名,UserInfo
	FieldNameWithSmallCamel string //小驼峰字段名,userInfo
	FieldType               string //字段类型
	FieldComment            string //字段的注解
}

func NewTable(dataBaseName string, tableName string, db *gorm.DB) *Table {
	return &Table{
		TableName:               tableName,
		TableNameWithBigCamel:   strutil.UpperFirst(strutil.CamelCase(tableName)),
		TableNameWithSmallCamel: strutil.CamelCase(tableName),
		DataBaseName:            dataBaseName,
		DB:                      db,
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

			//通过映射获取到映射类型
			fieldTemp.FieldType = mapping[pureType]
		} else {
			//如果没有找到，那么直接通过映射获取到映射类型即可
			fieldTemp.FieldType = mapping[fieldTemp.FieldType]
		}

		//小驼峰，
		fieldTemp.FieldNameWithSmallCamel = strutil.CamelCase(fieldTemp.FieldName)

		//大驼峰
		fieldTemp.FieldNameWithBigCamel = strutil.UpperFirst(strutil.CamelCase(fieldTemp.FieldName))

		myFields = append(myFields, fieldTemp)

	}
	receiver.Fields = myFields
}
