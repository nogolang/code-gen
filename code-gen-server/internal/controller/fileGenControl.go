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

type FileGenController struct {
	Logger         *zap.Logger
	FileGenService service.FileGenSvc
}

func NewFileGenController(engine *gin.Engine, FileGenService service.FileGenSvc, Logger *zap.Logger) *FileGenController {
	ctl := &FileGenController{FileGenService: FileGenService, Logger: Logger}
	group := engine.Group("/fileGen")
	group.GET("/findById/:id", ctl.FindById())
	group.GET("/deleteById/:id", ctl.DeleteById())
	group.POST("/updateById/:id", ctl.UpdateById())
	group.POST("/add", ctl.Add())
	group.POST("/findAll", ctl.FindAll())
	group.POST("/genFiles", ctl.GenFiles())
	return ctl
}

func (receiver *FileGenController) FindById() gin.HandlerFunc {
	return func(context *gin.Context) {
		temp := context.Param("id")
		id, err := strconv.Atoi(temp)
		if err != nil || id <= 0 {
			context.Error(commonRes.ParamInvalidID)
			return
		}
		data, err := receiver.FileGenService.FindById(id)
		if err != nil {
			context.Error(err)
			return
		}
		context.JSON(http.StatusOK, commonRes.OK.WithData(data))
	}
}

func (receiver *FileGenController) DeleteById() gin.HandlerFunc {
	return func(context *gin.Context) {
		temp := context.Param("id")
		id, err := strconv.Atoi(temp)
		if err != nil || id <= 0 {
			context.Error(commonRes.ParamInvalidID)
			return
		}

		err = receiver.FileGenService.DeleteById(id)
		if err != nil {
			context.Error(err)
			return
		}
		context.JSON(http.StatusOK, commonRes.DeleteOK)
	}
}

func (receiver *FileGenController) UpdateById() gin.HandlerFunc {
	return func(context *gin.Context) {
		temp := context.Param("id")
		id, err := strconv.Atoi(temp)
		if err != nil || id <= 0 {
			context.Error(commonRes.ParamInvalidID)
			return
		}

		//获取新的model参数
		var fileGen model.FileGenModelRequest
		err = context.ShouldBindJSON(&fileGen)
		if err != nil {
			context.Error(commonRes.FileParamInvalid)
			return
		}

		err = receiver.FileGenService.UpdateById(id, &fileGen)
		if err != nil {
			context.Error(err)
			return
		}
		context.JSON(http.StatusOK, commonRes.UpdateOK)
	}
}

func (receiver *FileGenController) Add() gin.HandlerFunc {
	return func(context *gin.Context) {
		//获取新的model参数
		var request model.FileGenModelRequest
		err := context.ShouldBindJSON(&request)
		if err != nil {
			context.Error(commonRes.FileParamInvalid)
			return
		}

		err = receiver.FileGenService.Add(&request)
		if err != nil {
			context.Error(err)
			return
		}
		context.JSON(http.StatusOK, commonRes.AddOK)
		return
	}

}

func (receiver *FileGenController) FindAll() gin.HandlerFunc {
	return func(context *gin.Context) {
		var query model.FileGenModelQuery
		err := context.ShouldBindJSON(&query)
		if err != nil {
			context.Error(commonRes.ParamInvalid)
			return
		}

		data, err := receiver.FileGenService.FindAll(&query)
		if err != nil {
			context.Error(err)
			return
		}
		context.JSON(http.StatusOK, commonRes.OK.WithData(data))
		return
	}
}

func (receiver *FileGenController) GenFiles() gin.HandlerFunc {
	return func(context *gin.Context) {
		var Ids []int
		err := context.ShouldBind(&Ids)
		if err != nil {
			context.Error(commonRes.ParamInvalid)
			return
		}
		err = receiver.FileGenService.GenFiles(Ids)
		if err != nil {
			context.Error(err)
			return
		}
		context.JSON(http.StatusOK, commonRes.FileGenOK)
	}
}
