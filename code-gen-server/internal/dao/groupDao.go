package dao

import (
	"code-gen/internal/model"
	"code-gen/internal/utils/commonRes"
	"code-gen/internal/utils/gormUtils"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type GroupDao struct {
	Logger *zap.Logger
	Db     *gorm.DB
}

func (d *GroupDao) Add(m *model.GroupModel) error {
	tx := d.Db.Create(m)
	if tx.Error != nil {
		return errors.Wrap(tx.Error, "添加组出错")
	}
	return nil
}

func (d *GroupDao) FindAll(query *model.GroupModelQuery) ([]*model.GroupModel, int64, error) {

	conditionDb := d.Db
	if query.QueryStr != "" {
		conditionDb = d.Db.Where("model.describe like ?", "%"+query.QueryStr+"%").
			Or("model.table_name like ?", "%"+query.QueryStr+"%")
	}
	conditionDb = conditionDb.Table("group_model as model")
	//获取所有内容
	var total int64
	conditionDb.Count(&total)
	var files []*model.GroupModel
	tx := conditionDb.
		Scopes(gormUtils.Pagination(query.Page, query.Size)).
		Find(&files)
	if tx.Error != nil {
		return nil, 0, errors.Wrap(tx.Error, "查询所有组数据出错")
	}
	return files, total, nil
}

func (d *GroupDao) UpdateById(id int, m *model.GroupModel) error {
	m.Id = id
	tx := d.Db.Save(m)
	if tx.Error != nil {
		return errors.Wrap(tx.Error, "更新组出错")
	}
	return nil
}

func (d *GroupDao) DeleteById(id int) error {
	tx := d.Db.Delete(&model.GroupModel{}, id)
	if tx.Error != nil {
		return errors.Wrap(tx.Error, "删除组出错")
	}
	return nil
}

func (d *GroupDao) FindById(id int) (*model.GroupModel, error) {
	var obj model.GroupModel
	tx := d.Db.Find(&obj, id)
	if tx.Error != nil {
		return nil, errors.Wrap(tx.Error, "根据id查询组出错")
	}
	if tx.RowsAffected == 0 {
		return nil, commonRes.GroupNotFount
	}
	return &obj, nil
}

func (d *GroupDao) FindAllNoPagination() ([]model.GroupModel, error) {
	var files []model.GroupModel
	tx := d.Db.Model(&model.GroupModel{}).
		Find(&files)
	if tx.Error != nil {
		return nil, errors.Wrap(tx.Error, "过滤组出错")
	}
	return files, nil
}
