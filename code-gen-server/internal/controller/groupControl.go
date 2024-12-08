package controller

import (
	"code-gen/internal/model"
	"code-gen/internal/service"
	"code-gen/internal/utils/commonRes"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

type GroupController struct {
	Logger       *zap.Logger
	GroupService service.GroupSvc
}

func NewGroupController(engine *gin.Engine, GroupService service.GroupSvc, Logger *zap.Logger) *GroupController {
	ctl := &GroupController{GroupService: GroupService, Logger: Logger}
	group := engine.Group("/group")
	group.GET("/findById/:id", ctl.FindById())
	group.POST("/findAllDir", ctl.FindAllDir())
	group.POST("/findAllDirForUpdate", ctl.FindAllDirForUpdate())
	group.GET("/deleteById/:id", ctl.DeleteById())
	group.POST("/updateById/:id", ctl.UpdateById())
	group.POST("/add", ctl.Add())
	group.POST("/findAll", ctl.FindAll())
	group.GET("/findAllNoPagination", ctl.FindAllNoPagination())
	group.DELETE("/deleteFileById/:id", ctl.DeleteFileById())
	group.DELETE("/deleteAllInvalidFile/:id", ctl.DeleteAllInvalidFile())
	return ctl
}

func (receiver *GroupController) FindById() gin.HandlerFunc {
	return func(context *gin.Context) {
		temp := context.Param("id")
		id, err := strconv.Atoi(temp)
		if err != nil || id <= 0 {
			context.Error(commonRes.ParamInvalidID)
			return
		}
		data, err := receiver.GroupService.FindById(id)
		if err != nil {
			context.Error(err)
			return
		}
		context.JSON(http.StatusOK, commonRes.OK.WithData(data))
	}
}

func (receiver *GroupController) DeleteById() gin.HandlerFunc {
	return func(context *gin.Context) {
		temp := context.Param("id")
		id, err := strconv.Atoi(temp)
		if err != nil || id <= 0 {
			context.Error(commonRes.ParamInvalidID)
			return
		}

		err = receiver.GroupService.DeleteById(id)
		if err != nil {
			context.Error(err)
			return
		}
		context.JSON(http.StatusOK, commonRes.DeleteOK)
	}
}

func (receiver *GroupController) UpdateById() gin.HandlerFunc {
	return func(context *gin.Context) {
		temp := context.Param("id")
		id, err := strconv.Atoi(temp)
		if err != nil || id <= 0 {
			context.Error(commonRes.ParamInvalidID)
			return
		}

		//获取新的model参数
		var fileGen model.GroupModel
		err = context.ShouldBindJSON(&fileGen)
		if err != nil {
			context.Error(commonRes.FileParamInvalid)
			return
		}

		err = receiver.GroupService.UpdateById(id, &fileGen)
		if err != nil {
			context.Error(err)
			return
		}
		context.JSON(http.StatusOK, commonRes.UpdateOK)
	}
}

func (receiver *GroupController) Add() gin.HandlerFunc {
	return func(context *gin.Context) {
		//获取新的model参数
		var request model.GroupModel
		err := context.ShouldBindJSON(&request)
		if err != nil {
			context.Error(commonRes.FileParamInvalid)
			return
		}

		err = receiver.GroupService.Add(&request)
		if err != nil {
			context.Error(err)
			return
		}
		context.JSON(http.StatusOK, commonRes.AddOK)
		return
	}

}

func (receiver *GroupController) FindAll() gin.HandlerFunc {
	return func(context *gin.Context) {
		var query model.GroupModelQuery
		err := context.ShouldBindJSON(&query)
		if err != nil {
			context.Error(commonRes.ParamInvalid)
			return
		}

		data, err := receiver.GroupService.FindAll(&query)
		if err != nil {
			context.Error(err)
			return
		}
		context.JSON(http.StatusOK, commonRes.OK.WithData(data))
		return
	}
}

func (receiver *GroupController) FindAllNoPagination() gin.HandlerFunc {
	return func(context *gin.Context) {
		data, err := receiver.GroupService.FindAllNoPagination()
		if err != nil {
			context.Error(err)
			return
		}
		context.JSON(http.StatusOK, commonRes.OK.WithData(data))
		return
	}
}

func (receiver *GroupController) FindAllDir() gin.HandlerFunc {
	return func(context *gin.Context) {
		type Dir struct {
			Path string `json:"path"`
		}
		var dir Dir
		err := context.ShouldBind(&dir)
		if err != nil {
			context.Error(err)
			return
		}
		data, err := receiver.GroupService.FindAllDir(dir.Path)
		if err != nil {
			context.Error(err)
			return
		}
		context.JSON(http.StatusOK, commonRes.OK.WithData(data))
		return
	}
}

func (receiver *GroupController) FindAllDirForUpdate() gin.HandlerFunc {
	return func(context *gin.Context) {
		type Dir struct {
			//分组id
			Id   int    `json:"id"`
			Path string `json:"path"`
		}
		var dir Dir
		err := context.ShouldBind(&dir)
		if err != nil {
			context.Error(err)
			return
		}
		data, err := receiver.GroupService.FindAllDirForUpdate(dir.Path, dir.Id)
		if err != nil {
			context.Error(err)
			return
		}
		context.JSON(http.StatusOK, commonRes.OK.WithData(data))
		return
	}
}

func (receiver *GroupController) DeleteFileById() gin.HandlerFunc {
	return func(context *gin.Context) {
		//不管有没有id都提示删除成功，因为可能还没存到中间表
		//如果id不为0，才去中间表找
		temp := context.Param("id")
		id, err := strconv.Atoi(temp)
		if err != nil || id == 0 {
			context.JSON(http.StatusOK, commonRes.DeleteOK)
			return
		}

		err = receiver.GroupService.DeleteFileById(id)
		if err != nil {
			context.Error(err)
			return
		}
		context.JSON(http.StatusOK, commonRes.DeleteOK)
		return
	}
}

// 删除所有无效的model，也就是本地不存在对应的文件，但是只删除当前group的
func (receiver *GroupController) DeleteAllInvalidFile() gin.HandlerFunc {
	return func(context *gin.Context) {
		temp := context.Param("id")
		id, err := strconv.Atoi(temp)
		if err != nil || id == 0 {
			context.JSON(http.StatusOK, commonRes.DeleteOK)
			return
		}
		//传递groupID
		err = receiver.GroupService.DeleteAllInvalidFile(id)
		if err != nil {
			context.Error(err)
			return
		}

		context.JSON(http.StatusOK, commonRes.OK)
		return
	}
}
