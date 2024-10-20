package dao

import (
	"code-gen/internal/model"
	"code-gen/internal/utils/commonRes"
	"code-gen/internal/utils/gormUtils"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type FileGenDao struct {
	Db     *gorm.DB
	Logger *zap.Logger
}

func (receiver *FileGenDao) Add(m *model.FileGenModel) error {
	tx := receiver.Db.Create(m)
	if tx.Error != nil {
		return errors.Wrap(tx.Error, "创建生成配置出错")
	}
	return nil
}

func (receiver *FileGenDao) FindById(id int) (*model.FileGenModel, error) {
	var obj model.FileGenModel
	tx := receiver.Db.Find(&obj, id)
	if tx.Error != nil {
		return nil, errors.Wrap(tx.Error, "查询生成配置出错")
	}
	if tx.RowsAffected == 0 {
		return nil, commonRes.FileMappingNotFount
	}
	return &obj, nil
}

func (receiver *FileGenDao) UpdateById(newModel *model.FileGenModel) error {
	tx := receiver.Db.Save(newModel)
	if tx.Error != nil {
		return errors.Wrap(tx.Error, "保存生成配置出错")
	}
	return nil
}

func (receiver *FileGenDao) DeleteById(id int) error {
	tx := receiver.Db.Delete(&model.FileGenModel{}, id)
	if tx.Error != nil {
		return errors.Wrap(tx.Error, "删除生成配置出错")
	}
	return nil
}

func (receiver *FileGenDao) FindAll(query *model.FileGenModelQuery) ([]*model.FileGenModel, int64, error) {
	//拼接查询条件
	conditionDb := receiver.Db
	if query.QueryStr != "" {
		conditionDb = receiver.Db.Where("model.describe like ?", "%"+query.QueryStr+"%").
			Or("model.table_name like ?", "%"+query.QueryStr+"%")
	}
	conditionDb = conditionDb.Table("file_gen_model as model")
	//获取所有内容
	var total int64
	conditionDb.Count(&total)
	var files []*model.FileGenModel
	tx := conditionDb.
		Scopes(gormUtils.Pagination(query.Page, query.Size)).
		Find(&files)
	if tx.Error != nil {
		return nil, 0, errors.Wrap(tx.Error, "查询所有生成配置数据出错")
	}
	return files, total, nil
}
