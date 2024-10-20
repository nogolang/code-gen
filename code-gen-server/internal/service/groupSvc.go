package service

import (
	"code-gen/internal/dao"
	"code-gen/internal/model"
	"code-gen/internal/utils/genUtils"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"os"
	"strings"
)

type GroupSvc struct {
	Logger          *zap.Logger
	Dao             *dao.GroupDao
	FileDao         *dao.FileDao
	OutDirDao       *dao.OutDir
	MappingDao      *dao.MappingPathDao
	OrmDao          *dao.OrmDao
	FileAndGroupDao *dao.FileAndGroupDao
}

// 添加组的同时，需要填充fileAndgroup的id信息
func (receiver *GroupSvc) Add(groupModel *model.GroupModel) error {
	//添加组，然后会回显id
	groupModel.RootDir = genUtils.WindowsPathToLinux(groupModel.RootDir)
	err := receiver.Dao.Add(groupModel)
	if err != nil {
		return err
	}

	//添加文件和组关联表id，这里用的是指针切片[]* 所以可以直接改
	//如果用普通切片[],那么对象里的数值经过for之后是新的对象
	for _, fileAndGroup := range groupModel.FileAndGroups {
		fileAndGroup.GroupId = groupModel.Id
	}
	err = receiver.FileAndGroupDao.AddBatch(groupModel.FileAndGroups)
	if err != nil {
		return err
	}

	return nil
}

// 查询组的同时，需要填充allFileAndGroup
func (receiver *GroupSvc) FindById(id int) (*model.GroupModel, error) {
	allFileAndGroup, err := receiver.FileAndGroupDao.FindAllByGroupId(id)
	if err != nil {
		return nil, err
	}

	//处理一下fileInfo和templateName
	//仅仅是展示给前台
	for _, afg := range allFileAndGroup {
		file, _ := receiver.FileDao.FindById(afg.FileId)
		//原始fileInfo
		afg.FileInfo = file

		//再填充一下TemplateName
		index := strings.LastIndex(file.TemplatePath, "/")
		afg.TemplateName = file.TemplatePath[index+1:]

		//判断模板文件释放存在
		_, err = genUtils.ReadFile(file.TemplatePath)
		if errors.Is(err, os.ErrNotExist) {
			file.TemplatePathIsExist = false
		} else {
			file.TemplatePathIsExist = true
		}

	}

	group, err := receiver.Dao.FindById(id)
	if err != nil {
		return nil, err
	}

	group.FileAndGroups = allFileAndGroup
	return group, nil
}

// 删除组的时候，删除allFileAndGroup
func (receiver *GroupSvc) DeleteById(id int) error {
	err := receiver.FileAndGroupDao.DeleteAllByGroupId(id)
	if err != nil {
		return err
	}
	return receiver.Dao.DeleteById(id)
}

// 更新组的时候，同时更新文件目录信息
func (receiver *GroupSvc) UpdateById(id int, m *model.GroupModel) error {
	for _, fileAndGroup := range m.FileAndGroups {
		//根据FileAndGroup自己的id更新
		err := receiver.FileAndGroupDao.UpdateById(fileAndGroup.Id, fileAndGroup)
		if err != nil {
			return err
		}
	}
	m.RootDir = genUtils.WindowsPathToLinux(m.RootDir)
	return receiver.Dao.UpdateById(id, m)
}

func (receiver *GroupSvc) FindAll(query *model.GroupModelQuery) (*model.GroupModelsAll, error) {
	var modelAll model.GroupModelsAll
	data, total, err := receiver.Dao.FindAll(query)
	if err != nil {
		return nil, err
	}

	//如果数量等于0，那么不返回错误，则是返回空数组
	if total == 0 {
		return &modelAll, nil
	}

	//填充一下files的信息，可能会有用
	var newData []model.GroupModel
	for _, d := range data {
		fullModel, _ := receiver.FindById(d.Id)
		newData = append(newData, *fullModel)
	}

	modelAll.Data = newData
	modelAll.Total = total
	return &modelAll, nil
}

func (receiver *GroupSvc) FindAllNoPagination() ([]model.GroupModel, error) {
	return receiver.Dao.FindAllNoPagination()
}
