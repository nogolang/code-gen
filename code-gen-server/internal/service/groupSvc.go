package service

import (
	"code-gen/internal/dao"
	"code-gen/internal/model"
	"code-gen/internal/utils/genUtils"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"os"
)

type GroupSvc struct {
	Logger     *zap.Logger
	Dao        *dao.GroupDao
	FileDao    *dao.FileDao
	MappingDao *dao.MappingPathDao
	OrmDao     *dao.OrmDao
	FileGenDao *dao.FileGenDao
}

// 添加组的同时，需要填充fileAndgroup的id信息
func (receiver *GroupSvc) Add(groupModel *model.GroupModel) error {
	//添加组，然后会回显id
	groupModel.GenRootDir = genUtils.WindowsPathToLinux(groupModel.GenRootDir)
	groupModel.SearchRootDir = genUtils.WindowsPathToLinux(groupModel.SearchRootDir)
	err := receiver.Dao.Add(groupModel)
	if err != nil {
		return err
	}

	//添加文件FileModels
	//如果用普通切片[],那么对象里的数值经过for之后是新的对象
	for _, file := range groupModel.FileModels {
		file.GroupId = groupModel.Id
	}
	err = receiver.FileDao.AddBatch(groupModel.FileModels)
	if err != nil {
		return err
	}

	return nil
}

// 查询组的同时，需要填充FileModels
func (receiver *GroupSvc) FindById(id int) (*model.GroupModel, error) {
	allfiles, err := receiver.FileDao.FindAllByGroupId(id)
	if err != nil {
		return nil, err
	}

	//填充一下文件是否存在
	//仅仅是展示给前台
	for _, file := range allfiles {
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

	group.FileModels = allfiles
	return group, nil
}

// 删除组的时候，删除所有的files
func (receiver *GroupSvc) DeleteById(id int) error {
	err := receiver.FileDao.DeleteAllByGroupId(id)
	if err != nil {
		return err
	}

	//把fileGen里的group id置0
	err = receiver.FileGenDao.SetZeroIdWithGroupId(id)
	if err != nil {
		return err
	}

	return receiver.Dao.DeleteById(id)
}

// 更新组的时候，也需要更新files的信息，因为files的更新是在组里的
func (receiver *GroupSvc) UpdateById(id int, m *model.GroupModel) error {
	for _, file := range m.FileModels {
		file.GroupId = m.Id
		err := receiver.FileDao.UpdateById(file.Id, file)
		if err != nil {
			return err
		}
	}
	m.GenRootDir = genUtils.WindowsPathToLinux(m.GenRootDir)
	m.SearchRootDir = genUtils.WindowsPathToLinux(m.SearchRootDir)
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
	//var newData []model.GroupModel
	//for _, d := range data {
	//	fullModel, _ := receiver.FindById(d.Id)
	//	newData = append(newData, *fullModel)
	//}

	modelAll.Data = data
	modelAll.Total = total
	return &modelAll, nil
}

func (receiver *GroupSvc) FindAllNoPagination() ([]model.GroupModel, error) {
	return receiver.Dao.FindAllNoPagination()
}

// 遍历出所有的文件,并返回一个空的model表单让前端填充
// 这是新增的时候要的
func (receiver *GroupSvc) FindAllDir(path string) ([]*model.FileModel, error) {
	newPath := genUtils.WindowsPathToLinux(path)
	files, err := genUtils.RecursionFiles(newPath)
	if err != nil {
		return nil, err
	}
	var models []*model.FileModel
	for _, fileName := range files {
		var m model.FileModel
		m.TemplatePath = fileName
		m.TemplatePath = genUtils.WindowsPathToLinux(m.TemplatePath)
		m.TemplatePathIsExist = true
		models = append(models, &m)
	}
	return models, err
}

// 如果是更新的时候,我搜索，那么是需要携带groupId的
// 先去查询所有的groupId对应的files，然后查询路径下所有的file
func (receiver *GroupSvc) FindAllDirForUpdate(path string, id int) ([]*model.FileModel, error) {
	newPath := genUtils.WindowsPathToLinux(path)
	var models []*model.FileModel
	//我在更新，此时会自动查找，当数据库的内容在本地没有存在，那么就显示本地文件不存在即可
	//  到时候我们手动在前端剔除掉，或者一键剔除掉
	allfilesDataBase, err := receiver.FileDao.FindAllByGroupId(id)
	for _, fileDatabase := range allfilesDataBase {
		//判断模板文件是否存在
		_, err = genUtils.ReadFile(fileDatabase.TemplatePath)
		if errors.Is(err, os.ErrNotExist) {
			fileDatabase.TemplatePathIsExist = false
		} else {
			fileDatabase.TemplatePathIsExist = true
		}
	}
	//放到新数组里，返回给前台
	models = append(models, allfilesDataBase...)

	//但是如果本地存在，而数据库不存在，那么就新增一个model让前端填充
	// 通过path去找，当前groupId有没有对应的path
	rawfiles, err := genUtils.RecursionFiles(newPath)
	if err != nil {
		return nil, err
	}
	for _, rawFile := range rawfiles {
		//如果不存在，则新增一个新的model
		isExist := isPathExist(allfilesDataBase, rawFile)
		if !isExist {
			var m model.FileModel
			m.TemplatePath = rawFile
			m.TemplatePath = genUtils.WindowsPathToLinux(m.TemplatePath)
			m.TemplatePathIsExist = true
			models = append(models, &m)
		}
	}

	return models, err
}

func (receiver *GroupSvc) DeleteFileById(id int) error {
	return receiver.FileDao.DeleteById(id)
}

func (receiver *GroupSvc) DeleteAllInvalidFile(groupId int) error {
	//查询当前分组下的无效文件
	allfilesDataBase, err := receiver.FileDao.FindAllByGroupId(groupId)
	if err != nil {
		return err
	}
	for _, fileDatabase := range allfilesDataBase {
		//判断模板文件是否存在
		_, err = genUtils.ReadFile(fileDatabase.TemplatePath)
		if errors.Is(err, os.ErrNotExist) {
			fileDatabase.TemplatePathIsExist = false
		} else {
			fileDatabase.TemplatePathIsExist = true
		}
	}

	//填充id
	var ids []int
	for _, file := range allfilesDataBase {
		//如果不存在，则删除这些
		//但是前提是我们TemplatePathIsExist都正确处理了
		//特别是查询的时候，就要去校验
		if !file.TemplatePathIsExist {
			ids = append(ids, file.Id)
		}
	}
	return receiver.FileDao.DeleteAllInvalidFile(ids)
}

func isPathExist(allfilesDataBase []*model.FileModel, path string) bool {
	for _, file := range allfilesDataBase {
		if file.TemplatePath == path {
			return true
		}
	}
	return false
}
