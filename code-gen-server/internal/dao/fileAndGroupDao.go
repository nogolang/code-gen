package dao

import (
	"code-gen/internal/model"
	"code-gen/internal/utils/genUtils"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type FileAndGroupDao struct {
	Logger *zap.Logger
	Db     *gorm.DB
}

func (d *FileAndGroupDao) AddBatch(m []*model.FileAndGroupModel) error {
	tx := d.Db.Create(m)
	if tx.Error != nil {
		return errors.Wrap(tx.Error, "添加文件和组出错")
	}
	return nil
}

func (d *FileAndGroupDao) Add(group *model.FileAndGroupModel) error {
	tx := d.Db.Create(group)
	if tx.Error != nil {
		return errors.Wrap(tx.Error, "添加文件和组出错")
	}
	return nil
}

func (d *FileAndGroupDao) DeleteAllByGroupId(id int) error {
	tx := d.Db.Where("group_id = ?", id).Delete(&model.FileAndGroupModel{})
	if tx.Error != nil {
		return errors.Wrap(tx.Error, "根据组id删除所有关联file出错")
	}
	return nil
}

func (d *FileAndGroupDao) FindAllByGroupId(id int) ([]*model.FileAndGroupModel, error) {
	var obj []*model.FileAndGroupModel
	tx := d.Db.Where("group_id = ?", id).Find(&obj)
	if tx.Error != nil {
		return nil, errors.Wrap(tx.Error, "根据组id查询所有文件数据出错")
	}
	return obj, nil
}

func (d *FileAndGroupDao) UpdateById(id int, m *model.FileAndGroupModel) error {
	//先查询后更新
	old, err := d.FindById(id)
	if err != nil {
		return err
	}
	old.OutDir = genUtils.WindowsPathToLinux(m.OutDir)
	tx := d.Db.Save(old)
	if tx.Error != nil {
		return errors.Wrap(tx.Error, "根据id更新文件组出错")
	}
	return nil
}

func (d *FileAndGroupDao) FindById(id int) (*model.FileAndGroupModel, error) {
	var obj model.FileAndGroupModel
	tx := d.Db.Find(&obj, id)
	if tx.Error != nil {
		return nil, errors.Wrap(tx.Error, "根据id查询文件组出错")
	}
	return &obj, nil
}

func (d *FileAndGroupDao) DeleteById(id int) error {
	tx := d.Db.Delete(&model.FileAndGroupModel{}, id)
	if tx.Error != nil {
		return errors.WithMessage(tx.Error, "删除组和模板文件的中间表出错")
	}
	return nil
}
