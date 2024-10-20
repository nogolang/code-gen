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
	Logger          *zap.Logger
	Dao             *dao.FileGenDao
	OrmDao          *dao.OrmDao
	GroupDao        *dao.GroupDao
	FileDao         *dao.FileDao
	FileAndGroupDao *dao.FileAndGroupDao
	MappingDao      *dao.MappingPathDao
}

func (receiver *FileGenSvc) Add(m *model.FileGenModelRequest) error {
	var fileModel model.FileGenModel
	err := copier.Copy(&fileModel, m)
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
		dbModel, _ := receiver.OrmDao.FindById(d.DataBaseId)
		d.DatabaseName = dbModel.Describe + "/" + dbModel.DataBaseName
		d.TableNamesArr = strings.Split(d.TableNames, ",")
	}

	//获取到组的describe
	for _, d := range data {
		groupModel, _ := receiver.GroupDao.FindById(d.GroupId)
		d.GroupDescribe = groupModel.Describe
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

		//获取组ID获取到多个文件和组的对象
		fileGroupsModels, _ := receiver.FileAndGroupDao.FindAllByGroupId(genFileModel.GroupId)

		//根据组id获取组对象
		groupModel, _ := receiver.GroupDao.FindById(genFileModel.GroupId)

		//根据db去查询指定的数据库，以及要生成的表
		for _, fileGroup := range fileGroupsModels {
			//查询到文件对象
			fileModel, _ := receiver.FileDao.FindById(fileGroup.FileId)

			//查询到mapping对象的内容
			mappingModel, _ := receiver.MappingDao.FindById(fileModel.MappingId)

			//最终生成目录，是组里的rootDir加上我们中间表的相对路径
			finalOutDir := groupModel.RootDir + fileGroup.OutDir

			gen := genCode.NewFileGen(
				genUtils.CustomFunc,
				receiver.Logger,
				connect,
				dbModel.DataBaseName,
				tables,
				fileModel.TemplatePath,
				mappingModel.Content,
				fileModel.NameSuffix,
				fileModel.FileSuffix,
				finalOutDir,
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
