package service

import (
	"code-gen/internal/dao"
	"code-gen/internal/genCode"
	"code-gen/internal/model"
	"code-gen/internal/utils/genUtils"
	"code-gen/internal/utils/gormUtils"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"strings"
)

type FileGenSvc struct {
	Logger     *zap.Logger
	Dao        *dao.FileGenDao
	OrmDao     *dao.OrmDao
	GroupDao   *dao.GroupDao
	FileDao    *dao.FileDao
	MappingDao *dao.MappingPathDao
}

func (receiver *FileGenSvc) Add(m *model.FileGenModelRequest) error {
	var fileModel model.FileGenModel
	err := copier.Copy(&fileModel, m)
	//处理表名，扁平化
	fileModel.TableNames = strings.Join(m.TableNamesArr, ",")
	if err != nil {
		return errors.Wrap(err, "新增'生成配置'数据复制出错")
	}
	return receiver.Dao.Add(&fileModel)
}

func (receiver *FileGenSvc) FindById(id int) (*model.FileGenModel, error) {
	m, _ := receiver.Dao.FindById(id)
	m.TableNamesArr = strings.Split(m.TableNames, ",")
	return m, nil
}

func (receiver *FileGenSvc) DeleteById(id int) error {
	return receiver.Dao.DeleteById(id)
}

func (receiver *FileGenSvc) UpdateById(id int, m *model.FileGenModelRequest) error {
	//先查询后更新
	old, err := receiver.FindById(id)
	if err != nil {
		return err
	}

	//处理表名，扁平化
	str := strings.Join(m.TableNamesArr, ",")

	old.TableNames = str
	old.DataBaseId = m.DataBaseId
	old.Describe = m.Describe
	old.GroupId = m.GroupId

	//old里面已经有id数据的
	return receiver.Dao.UpdateById(old)
}

func (receiver *FileGenSvc) FindAll(query *model.FileGenModelQuery) (*model.FileGenModelsAll, error) {
	var modelAll model.FileGenModelsAll
	data, total, err := receiver.Dao.FindAll(query)
	if err != nil {
		return nil, err
	}

	//如果数量等于0，那么不返回错误，则是返回空数组
	if total == 0 {
		return &modelAll, nil
	}

	//获取到数据库名
	for _, d := range data {
		dbModel, err := receiver.OrmDao.FindById(d.DataBaseId)
		//如果没有对应的配置源，则下面会出错，如果我们删除了配置源
		//那么应该给本表对应的datasourceID置空
		if err != nil {
			return nil, err
		}
		d.DatabaseName = dbModel.Describe + "/" + dbModel.DataBaseName
		d.TableNamesArr = strings.Split(d.TableNames, ",")
	}
	
	modelAll.Data = data
	modelAll.Total = total
	return &modelAll, nil
}

func (receiver *FileGenSvc) GenFiles(ids []int) error {
	for _, genFileId := range ids {
		//根据genFileId查询数据库，获取db对象
		genFileModel, _ := receiver.Dao.FindById(genFileId)
		dbModel, _ := receiver.OrmDao.FindById(genFileModel.DataBaseId)
		connect, err := gormUtils.GetConnect(dbModel)
		if err != nil {
			return err
		}

		//把数据库的扁平化表转换为数组
		tables := strings.Split(genFileModel.TableNames, ",")

		//通过组ID获取到多个文件和组的对象
		allFiles, _ := receiver.FileDao.FindAllByGroupId(genFileModel.GroupId)

		//根据组id获取组对象
		groupModel, _ := receiver.GroupDao.FindById(genFileModel.GroupId)

		//根据db去查询指定的数据库，以及要生成的表
		for _, fileModel := range allFiles {
			//查询到mapping对象的内容
			mappingModel, _ := receiver.MappingDao.FindById(fileModel.MappingId)

			//最终生成目录，是组里的rootDir加上中间表的相对路径
			//这个中间件表是组和模板文件的中间表
			finalOutDir := groupModel.GenRootDir + fileModel.GenPath

			gen := genCode.NewFileGen(
				genUtils.CustomFunc,
				receiver.Logger,
				connect,
				dbModel,
				tables,
				fileModel.TemplatePath,
				mappingModel.Content,
				fileModel.NameSuffix,
				fileModel.FileSuffix,
				finalOutDir,
				groupModel.GenRootDir,
				fileModel.IsCamelCase,
			)

			//生成文件
			err := gen.GenFile()
			if err != nil {
				return err
			}
		}
	}
	return nil
}
