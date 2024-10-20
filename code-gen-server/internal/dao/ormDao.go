package dao

import (
	"code-gen/internal/model"
	"code-gen/internal/utils/commonRes"
	"code-gen/internal/utils/gormUtils"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type OrmDao struct {
	Logger *zap.Logger
	Db     *gorm.DB
}

func (d *OrmDao) Add(m *model.OrmModel) error {
	tx := d.Db.Create(m)
	if tx.Error != nil {
		return errors.Wrap(tx.Error, "添加组出错")
	}
	return nil
}

func (d *OrmDao) FindAll(query *model.OrmModelQuery) ([]model.OrmModel, int64, error) {

	//拼接查询条件
	conditionDb := d.Db
	if query.QueryStr != "" {
		conditionDb = d.Db.Where("model.describe like ?", "%"+query.QueryStr+"%").
			Or("model.port like ?", "%"+query.QueryStr+"%").
			Or("model.username like ?", "%"+query.QueryStr+"%").
			Or("model.host like ?", "%"+query.QueryStr+"%")
	}

	conditionDb = conditionDb.Table("orm_model as model")

	//获取所有内容
	var total int64
	conditionDb.Count(&total)
	var files []model.OrmModel
	tx := conditionDb.
		Scopes(gormUtils.Pagination(query.Page, query.Size)).
		Find(&files)
	if tx.Error != nil {
		return nil, 0, errors.Wrap(tx.Error, "查询所有组数据出错")
	}
	return files, total, nil
}

func (d *OrmDao) UpdateById(id int, m *model.OrmModel) error {
	m.Id = id
	tx := d.Db.Save(m)
	if tx.Error != nil {
		return errors.Wrap(tx.Error, "更新组出错")
	}
	return nil
}

func (d *OrmDao) DeleteById(id int) error {
	tx := d.Db.Delete(&model.OrmModel{}, id)
	if tx.Error != nil {
		return errors.Wrap(tx.Error, "删除组出错")
	}
	return nil
}

func (d *OrmDao) FindById(id int) (*model.OrmModel, error) {
	var obj model.OrmModel
	tx := d.Db.Find(&obj, id)
	if tx.Error != nil {
		return nil, errors.Wrap(tx.Error, "根据id查询orm出错")
	}
	if tx.RowsAffected == 0 {
		return nil, commonRes.GroupNotFount
	}
	return &obj, nil
}

func (d *OrmDao) FindAllNoPagination() ([]model.OrmModel, error) {
	var files []model.OrmModel
	tx := d.Db.Table("orm_model as model").
		Find(&files)
	if tx.Error != nil {
		return nil, errors.Wrap(tx.Error, "过滤数据库出错")
	}
	return files, nil
}

func (d *OrmDao) FindTables(db *gorm.DB) ([]string, error) {
	type table struct {
		TablesInMyTest []string
	}
	sql := "show tables;"
	var t table
	tx := db.Raw(sql).Scan(&t.TablesInMyTest)
	if tx.Error != nil {
		return nil, errors.Wrap(tx.Error, "查询所有表出错")
	}
	return t.TablesInMyTest, nil
}
