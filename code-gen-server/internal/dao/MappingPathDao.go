package dao

import (
	"code-gen/internal/model"
	"code-gen/internal/utils/commonRes"
	"code-gen/internal/utils/gormUtils"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type MappingPathDao struct {
	Db     *gorm.DB
	Logger *zap.Logger
}

func (receiver *MappingPathDao) Add(m *model.MappingPathModel) error {
	tx := receiver.Db.Create(m)
	if tx.Error != nil {
		return errors.Wrap(tx.Error, "创建mapping出错")
	}
	return nil
}

func (receiver *MappingPathDao) FindById(id int) (*model.MappingPathModel, error) {
	var obj model.MappingPathModel
	tx := receiver.Db.Find(&obj, id)
	if tx.Error != nil {
		return nil, errors.Wrap(tx.Error, "查询出错")
	}
	if tx.RowsAffected == 0 {
		return nil, commonRes.FileMappingNotFount
	}
	return &obj, nil
}

func (receiver *MappingPathDao) UpdateById(id int, newModel *model.MappingPathModel) error {
	newModel.Id = id
	tx := receiver.Db.Save(newModel)
	if tx.Error != nil {
		return errors.Wrap(tx.Error, "保存出错")
	}
	return nil
}

func (receiver *MappingPathDao) DeleteById(id int) error {
	tx := receiver.Db.Delete(&model.MappingPathModel{}, id)
	if tx.Error != nil {
		return errors.Wrap(tx.Error, "删除mapping出错")
	}
	return nil
}

func (receiver *MappingPathDao) FindAll(query *model.MappingPathModelQuery) ([]model.MappingPathModel, int64, error) {
	//拼接查询条件
	conditionDb := receiver.Db
	if query.QueryStr != "" {
		conditionDb = receiver.Db.Where("model.describe like ?", "%"+query.QueryStr+"%")
	}
	conditionDb = conditionDb.Table("mapping_path_model as model")
	//获取所有内容
	var total int64
	conditionDb.Count(&total)
	var files []model.MappingPathModel
	tx := conditionDb.
		Scopes(gormUtils.Pagination(query.Page, query.Size)).
		Find(&files)
	if tx.Error != nil {
		return nil, 0, errors.Wrap(tx.Error, "查询所有组数据出错")
	}
	return files, total, nil
}

func (receiver *MappingPathDao) FindAllNoPagination() ([]model.MappingPathModel, error) {
	var files []model.MappingPathModel
	tx := receiver.Db.Model(&model.MappingPathModel{}).
		Find(&files)
	if tx.Error != nil {
		return nil, errors.Wrap(tx.Error, "过滤mapping出错")
	}
	return files, nil
}
