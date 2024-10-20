package dao

import (
	"code-gen/internal/model"
	"code-gen/internal/utils/commonRes"
	"code-gen/internal/utils/gormUtils"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type FileDao struct {
	Logger *zap.Logger
	Db     *gorm.DB
}

func (f *FileDao) FindById(id int) (*model.FileModel, error) {
	var obj model.FileModel
	tx := f.Db.Find(&obj, id)
	if tx.Error != nil {
		return nil, errors.Wrap(tx.Error, "文件配置查询出错")
	}
	if tx.RowsAffected == 0 {
		return nil, commonRes.FileDataNotFount
	}
	return &obj, nil
}

func (f *FileDao) UpdateById(id int, newModel *model.FileModel) error {
	newModel.Id = id
	tx := f.Db.Save(newModel)
	if tx.Error != nil {
		return errors.Wrap(tx.Error, "文件配置保存出错")
	}
	return nil
}

func (f *FileDao) DeleteById(id int) error {
	tx := f.Db.Delete(&model.FileModel{}, id)
	if tx.Error != nil {
		return errors.Wrap(tx.Error, "文件配置删除出错")
	}
	return nil
}

func (f *FileDao) Add(fileModel *model.FileModel) error {
	tx := f.Db.Create(fileModel)
	if tx.Error != nil {
		return errors.Wrap(tx.Error, "文件配置保存出错")
	}
	return nil
}

func (f *FileDao) FindAll(query *model.FileModelQuery) ([]*model.FileModel, int, error) {
	//拼接查询条件
	conditionDb := f.Db
	if query.QueryStr != "" {
		conditionDb = f.Db.Where("model.describe like ?", "%"+query.QueryStr+"%").
			Or("model.name_suffix like ?", "%"+query.QueryStr+"%").
			Or("model.file_suffix like ?", "%"+query.QueryStr+"%")
	}
	conditionDb = conditionDb.Table("file_model as model")

	//获取所有内容
	var total int64
	conditionDb.Count(&total)
	var files []*model.FileModel
	tx := conditionDb.
		Scopes(gormUtils.Pagination(query.Page, query.Size)).
		Find(&files)
	if tx.Error != nil {
		return nil, 0, errors.Wrap(tx.Error, "文件配置根据条件查询出错")
	}
	return files, int(total), nil
}

func (f *FileDao) FindAllNoPagination() ([]model.FileModel, error) {
	var files []model.FileModel
	tx := f.Db.
		Find(&files)
	if tx.Error != nil {
		return nil, errors.Wrap(tx.Error, "查询所有files出错")
	}
	return files, nil
}
