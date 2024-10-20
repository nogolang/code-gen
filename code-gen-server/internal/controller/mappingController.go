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

type MappingController struct {
	Logger     *zap.Logger
	MappingSvc service.MappingSvc
}

func NewMappingController(engine *gin.Engine, MappingSvc service.MappingSvc, Logger *zap.Logger) *MappingController {
	ctl := &MappingController{MappingSvc: MappingSvc, Logger: Logger}
	group := engine.Group("/mapping")
	group.GET("/findById/:id", ctl.FindById())
	group.GET("/deleteById/:id", ctl.DeleteById())
	group.POST("/updateById/:id", ctl.UpdateById())
	group.POST("/add", ctl.Add())
	group.POST("/findAll", ctl.FindAll())
	group.GET("/findAllNoPagination", ctl.FindAllNoPagination())
	return ctl
}

func (receiver *MappingController) FindAllNoPagination() gin.HandlerFunc {
	return func(context *gin.Context) {
		data, err := receiver.MappingSvc.FindAllNoPagination()
		if err != nil {
			context.Error(err)
			return
		}
		context.JSON(http.StatusOK, commonRes.OK.WithData(data))
		return
	}
}

func (receiver *MappingController) FindById() gin.HandlerFunc {
	return func(context *gin.Context) {
		temp := context.Param("id")
		id, err := strconv.Atoi(temp)
		if err != nil || id <= 0 {
			context.Error(commonRes.ParamInvalidID)
			return
		}
		data, err := receiver.MappingSvc.FindById(id)
		if err != nil {
			context.Error(err)
			return
		}
		context.JSON(http.StatusOK, commonRes.OK.WithData(data))
	}
}

func (receiver *MappingController) DeleteById() gin.HandlerFunc {
	return func(context *gin.Context) {
		temp := context.Param("id")
		id, err := strconv.Atoi(temp)
		if err != nil || id <= 0 {
			context.Error(commonRes.ParamInvalidID)
			return
		}

		err = receiver.MappingSvc.DeleteById(id)
		if err != nil {
			context.Error(err)
			return
		}
		context.JSON(http.StatusOK, commonRes.DeleteOK)
	}
}

func (receiver *MappingController) UpdateById() gin.HandlerFunc {
	return func(context *gin.Context) {
		temp := context.Param("id")
		id, err := strconv.Atoi(temp)
		if err != nil || id <= 0 {
			context.Error(commonRes.ParamInvalidID)
			return
		}

		//获取新的model参数
		var fileGen model.MappingPathModel
		err = context.ShouldBindJSON(&fileGen)
		if err != nil {
			context.Error(commonRes.FileParamInvalid)
			return
		}

		err = receiver.MappingSvc.UpdateById(id, &fileGen)
		if err != nil {
			context.Error(err)
			return
		}
		context.JSON(http.StatusOK, commonRes.UpdateOK)
	}
}

func (receiver *MappingController) Add() gin.HandlerFunc {
	return func(context *gin.Context) {
		//获取新的model参数
		var request model.MappingPathModel
		err := context.ShouldBindJSON(&request)
		if err != nil {
			context.Error(commonRes.FileParamInvalid)
			return
		}

		err = receiver.MappingSvc.Add(&request)
		if err != nil {
			context.Error(err)
			return
		}
		context.JSON(http.StatusOK, commonRes.AddOK)
		return
	}

}

func (receiver *MappingController) FindAll() gin.HandlerFunc {
	return func(context *gin.Context) {
		var query model.MappingPathModelQuery
		err := context.ShouldBindJSON(&query)
		if err != nil {
			context.Error(commonRes.ParamInvalid)
			return
		}

		data, err := receiver.MappingSvc.FindAll(&query)
		if err != nil {
			context.Error(err)
			return
		}
		context.JSON(http.StatusOK, commonRes.OK.WithData(data))
		return
	}
}
