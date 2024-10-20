package service

import (
	"code-gen/internal/dao"
	"code-gen/internal/model"
	"code-gen/internal/utils/genUtils"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"os"
)

type FileService struct {
	Logger     *zap.Logger
	Dao        *dao.FileDao
	OutDirDao  *dao.OutDir
	MappingDao *dao.MappingPathDao
}

// 根据id查询的时候，其它字段也要填充进来
func (receiver *FileService) FindById(id int) (*model.FileModel, error) {
	oldModel, err := receiver.Dao.FindById(id)
	if err != nil {
		return nil, err
	}
	return oldModel, nil
}

func (receiver *FileService) UpdateById(id int, request *model.FileModel) error {

	//fileModel数据，先查询后修改
	old, err := receiver.FindById(id)
	if err != nil {
		return err
	}
	old.Describe = request.Describe
	old.IsCamelCase = request.IsCamelCase
	old.NameSuffix = request.NameSuffix
	old.FileSuffix = request.FileSuffix
	old.MappingId = request.MappingId
	request.TemplatePath = genUtils.WindowsPathToLinux(request.TemplatePath)
	err = receiver.Dao.UpdateById(id, old)
	if err != nil {
		return err
	}

	return nil
}

func (receiver *FileService) DeleteById(id int) error {
	return receiver.Dao.DeleteById(id)
}

func (receiver *FileService) Add(request *model.FileModel) error {
	request.TemplatePath = genUtils.WindowsPathToLinux(request.TemplatePath)
	return receiver.Dao.Add(request)
}

func (receiver *FileService) FindAll(query *model.FileModelQuery) (*model.FileModesAll, error) {
	var modelAll model.FileModesAll
	data, total, err := receiver.Dao.FindAll(query)
	if err != nil {
		return nil, err
	}

	//如果数量等于0，那么不返回错误，则是返回空数组
	if total == 0 {
		return &modelAll, nil
	}

	//查询每个模板文件是否真的存在
	for _, d := range data {
		_, err = genUtils.ReadFile(d.TemplatePath)
		if errors.Is(err, os.ErrNotExist) {
			d.TemplatePathIsExist = false
		} else {
			d.TemplatePathIsExist = true
		}

	}

	modelAll.Data = data
	modelAll.Total = total

	return &modelAll, nil
}

func (receiver *FileService) FindAllNoPagination() ([]model.FileModel, error) {
	return receiver.Dao.FindAllNoPagination()
}
