package genCode

import (
	"code-gen/internal/utils/commonRes"
	"code-gen/internal/utils/genUtils"
	"github.com/Masterminds/sprig/v3"
	"github.com/duke-git/lancet/v2/strutil"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"os"
	"strings"
	"text/template"
)

type FileGen struct {
	CustomFunc template.FuncMap
	logger     *zap.Logger

	//db对象
	DB *gorm.DB

	//数据库名
	DataBaseName string

	//表名
	TableName []string

	//模板路径
	TemplatePath string

	//映射文件内容
	MappingStr string

	//文件的后缀，比如叫_controller
	NameSuffix string

	//文件的后缀，比如.go
	FileSuffix string

	//最终输出目录
	FinalOutDir string

	//生成的文件名称是否是驼峰形式
	IsCamelCase bool
}

func NewFileGen(customFunc template.FuncMap, logger *zap.Logger, DB *gorm.DB, dataBaseName string, tableName []string, templatePath string, mappingStr string, nameSuffix string, fileSuffix string, finalOutDir string, isCamelCase bool) *FileGen {
	return &FileGen{CustomFunc: customFunc, logger: logger, DataBaseName: dataBaseName, DB: DB, TableName: tableName, TemplatePath: templatePath, MappingStr: mappingStr, NameSuffix: nameSuffix, FileSuffix: fileSuffix, FinalOutDir: finalOutDir, IsCamelCase: isCamelCase}
}

// GenFile 生成文件
func (receiver *FileGen) GenFile() error {
	//根据多个表生成
	for _, tbName := range receiver.TableName {
		//创建table对象
		table := NewTable(receiver.DataBaseName, tbName, receiver.DB)

		//把mapping json转换到map
		//读取文件的映射字符串
		mapping, err := genUtils.JsonToMap(receiver.MappingStr)
		if err != nil {
			return err
		}

		//获取到table数据
		tableData := table.GetTable(mapping)

		//根据table解析模板
		err = receiver.parseTemplateFile(tableData, tbName)
		if err != nil {
			return err
		}
	}
	return nil
}

// 解析模板文件
// 传递数据
func (receiver *FileGen) parseTemplateFile(data interface{}, tbName string) error {
	//先删除之前的目录
	//err := os.RemoveAll(receiver.FinalOutDir)
	//if err != nil {
	//	return errors.Wrap(err,"删除出错")
	//}

	//如果目录没有，则还需要创建目录
	_, err := os.Stat(receiver.FinalOutDir)
	if os.IsNotExist(err) {
		err = os.MkdirAll(receiver.FinalOutDir, os.ModePerm)
		if err != nil {
			return commonRes.FileCreateDirError.WithReason(err.Error())
		}
	}

	//获取到模板文件名称，用于template.New
	lastIndex := strings.LastIndex(receiver.TemplatePath, "/")
	fileNameHasSuffix := receiver.TemplatePath[lastIndex+1:]

	files, err := template.New(fileNameHasSuffix).
		Funcs(receiver.CustomFunc).
		Funcs(sprig.FuncMap()).
		ParseFiles(receiver.TemplatePath)
	if err != nil {
		return commonRes.FileCreateTemplateError.WithReason(err.Error())
	}

	//比如生成的目录是/gen/out/httpApi/，文件名称是userHttp.go
	outFile := receiver.FinalOutDir + "/" + tbName + receiver.NameSuffix + receiver.FileSuffix

	//看看外部指定的文件名称是否是驼峰形式
	//因为有些人想生成的文件名称是这样的user_http 蛇形命名
	//有些是userHttp 驼峰
	if receiver.IsCamelCase {
		tbName = strutil.CamelCase(tbName)
	}

	//创建文件
	outIO, err := os.OpenFile(outFile, os.O_CREATE, 0666)
	if err != nil {
		return commonRes.FileCreateFileError.WithReason(err.Error())
	}

	//生成到文件里
	err = files.Execute(outIO, data)
	if err != nil {
		//把模板文件出错的信息返回
		return commonRes.FileTemplateParseError.WithReason(err.Error())
	}

	//生成后关闭流，不然下次生成就替换不了了
	outIO.Close()

	return nil
}
