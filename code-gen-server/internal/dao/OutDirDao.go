package dao

import (
	"code-gen/internal/model"
	"code-gen/internal/utils/commonRes"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type OutDir struct {
	Db     *gorm.DB
	Logger *zap.Logger
}

func (receiver *OutDir) Add(m *model.OutDirModel) error {
	tx := receiver.Db.Create(m)
	if tx.Error != nil {
		return errors.Wrap(tx.Error, "创建ourDir出错")
	}
	return nil
}

func (receiver *OutDir) FindById(id int) (*model.OutDirModel, error) {
	var obj model.OutDirModel
	tx := receiver.Db.Find(&obj, id)
	if tx.Error != nil {
		return nil, errors.Wrap(tx.Error, "查询出错")
	}
	if tx.RowsAffected == 0 {
		return nil, commonRes.FileOutDirDataNotFount
	}
	return &obj, nil
}
func (receiver *OutDir) FindByFileId(id int) (*model.OutDirModel, error) {
	var obj model.OutDirModel
	tx := receiver.Db.Where("fileModel_id = ?", id).Find(&obj)
	if tx.Error != nil {
		return nil, errors.Wrap(tx.Error, "查询出错")
	}
	return &obj, nil
}

func (receiver *OutDir) UpdateById(id int, newModel *model.OutDirModel) error {
	newModel.Id = id
	tx := receiver.Db.Save(newModel)
	if tx.Error != nil {
		return errors.Wrap(tx.Error, "保存出错")
	}
	return nil
}

func (receiver *OutDir) DeleteByFileId(id int) error {
	tx := receiver.Db.Where("fileModel_id = ?", id).Delete(&model.OutDirModel{})
	if tx.Error != nil {
		return errors.Wrap(tx.Error, "删除出错")
	}
	return nil
}
