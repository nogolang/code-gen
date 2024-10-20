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

type OrmController struct {
	Logger     *zap.Logger
	OrmService service.OrmSvc
}

func NewOrmController(engine *gin.Engine, OrmService service.OrmSvc, Logger *zap.Logger) *OrmController {
	ctl := &OrmController{OrmService: OrmService, Logger: Logger}
	group := engine.Group("/database")
	group.GET("/findById/:id", ctl.FindById())
	group.GET("/deleteById/:id", ctl.DeleteById())
	group.POST("/updateById/:id", ctl.UpdateById())
	group.POST("/add", ctl.Add())
	group.POST("/findAll", ctl.FindAll())
	group.GET("/findAllNoPagination", ctl.FindAllNoPagination())
	group.GET("/findTablesByDatabaseId/:id", ctl.FindTablesByDatabaseId())
	group.POST("/checkConnect", ctl.CheckConnect())
	return ctl
}
func (receiver *OrmController) FindById() gin.HandlerFunc {
	return func(context *gin.Context) {
		temp := context.Param("id")
		id, err := strconv.Atoi(temp)
		if err != nil || id <= 0 {
			context.Error(commonRes.ParamInvalidID)
			return
		}
		data, err := receiver.OrmService.FindById(id)
		if err != nil {
			context.Error(err)
			return
		}
		context.JSON(http.StatusOK, commonRes.OK.WithData(data))
	}
}

func (receiver *OrmController) DeleteById() gin.HandlerFunc {
	return func(context *gin.Context) {
		temp := context.Param("id")
		id, err := strconv.Atoi(temp)
		if err != nil || id <= 0 {
			context.Error(commonRes.ParamInvalidID)
			return
		}

		err = receiver.OrmService.DeleteById(id)
		if err != nil {
			context.Error(err)
			return
		}
		context.JSON(http.StatusOK, commonRes.DeleteOK)
	}
}

func (receiver *OrmController) UpdateById() gin.HandlerFunc {
	return func(context *gin.Context) {
		temp := context.Param("id")
		id, err := strconv.Atoi(temp)
		if err != nil || id <= 0 {
			context.Error(commonRes.ParamInvalidID)
			return
		}

		//获取新的model参数
		var fileGen model.OrmModel
		err = context.ShouldBindJSON(&fileGen)
		if err != nil {
			context.Error(commonRes.FileParamInvalid)
			return
		}

		err = receiver.OrmService.UpdateById(id, &fileGen)
		if err != nil {
			context.Error(err)
			return
		}
		context.JSON(http.StatusOK, commonRes.UpdateOK)
	}
}

func (receiver *OrmController) Add() gin.HandlerFunc {
	return func(context *gin.Context) {
		//获取新的model参数
		var request model.OrmModel
		err := context.ShouldBindJSON(&request)
		if err != nil {
			context.Error(commonRes.FileParamInvalid)
			return
		}

		err = receiver.OrmService.Add(&request)
		if err != nil {
			context.Error(err)
			return
		}
		context.JSON(http.StatusOK, commonRes.AddOK)
		return
	}

}

func (receiver *OrmController) FindAll() gin.HandlerFunc {
	return func(context *gin.Context) {
		var query model.OrmModelQuery
		err := context.ShouldBindJSON(&query)
		if err != nil {
			context.Error(commonRes.ParamInvalid)
			return
		}

		data, err := receiver.OrmService.FindAll(&query)
		if err != nil {
			context.Error(err)
			return
		}
		context.JSON(http.StatusOK, commonRes.OK.WithData(data))
		return
	}
}

func (receiver *OrmController) CheckConnect() gin.HandlerFunc {
	return func(context *gin.Context) {
		var request model.OrmModel
		err := context.ShouldBindJSON(&request)
		if err != nil {
			context.Error(commonRes.ParamInvalid)
			return
		}
		err = receiver.OrmService.CheckConnect(&request)
		if err != nil {
			context.Error(err)
			return
		}
		context.JSON(http.StatusOK, commonRes.DataBaseConnectOK)
		return

	}
}

func (receiver *OrmController) FindAllNoPagination() gin.HandlerFunc {
	return func(context *gin.Context) {
		data, err := receiver.OrmService.FindAllNoPagination()
		if err != nil {
			context.Error(err)
			return
		}
		context.JSON(http.StatusOK, commonRes.OK.WithData(data))
		return
	}
}

func (receiver *OrmController) FindTablesByDatabaseId() gin.HandlerFunc {
	return func(context *gin.Context) {
		temp := context.Param("id")
		id, err := strconv.Atoi(temp)
		if err != nil || id <= 0 {
			context.Error(commonRes.NeedDatabaseID)
			return
		}
		data, err := receiver.OrmService.FindTablesByDatabaseId(id)
		if err != nil {
			context.Error(err)
			return
		}
		context.JSON(http.StatusOK, commonRes.OK.WithData(data))
	}

}
