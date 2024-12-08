package controller

//
//import (
//	"code-gen/internal/model"
//	"code-gen/internal/service"
//	"code-gen/internal/utils/commonRes"
//	"github.com/gin-gonic/gin"
//	"go.uber.org/zap"
//	"net/http"
//	"strconv"
//)
//
//type FileController struct {
//	Logger      *zap.Logger
//	FileService service.FileService
//}
//
//func NewFileController(engine *gin.Engine, FileService service.FileService, Logger *zap.Logger) *FileController {
//	ctl := &FileController{FileService: FileService, Logger: Logger}
//	group := engine.Group("/file")
//	group.GET("/findById/:id", ctl.FindById())
//	group.GET("/deleteById/:id", ctl.DeleteById())
//	group.POST("/updateById/:id", ctl.UpdateById())
//	group.POST("/add", ctl.Add())
//	group.POST("/findAll", ctl.FindAll())
//	group.GET("/findAllNoPagination", ctl.FindAllNoPagination())
//	return ctl
//}
//
//func (receiver *FileController) FindById() gin.HandlerFunc {
//	return func(context *gin.Context) {
//		temp := context.Param("id")
//		id, err := strconv.Atoi(temp)
//		if err != nil || id <= 0 {
//			context.Error(commonRes.ParamInvalidID)
//			return
//		}
//		data, err := receiver.FileService.FindById(id)
//		if err != nil {
//			context.Error(err)
//			return
//		}
//		context.JSON(http.StatusOK, commonRes.OK.WithData(data))
//	}
//}
//func (receiver *FileController) UpdateById() gin.HandlerFunc {
//	return func(context *gin.Context) {
//		temp := context.Param("id")
//		id, err := strconv.Atoi(temp)
//		if err != nil || id <= 0 {
//			context.Error(commonRes.ParamInvalidID)
//			return
//		}
//
//		//获取新的model参数
//		var fileGen model.FileModel
//		err = context.ShouldBindJSON(&fileGen)
//		if err != nil {
//			context.Error(commonRes.FileParamInvalid)
//			return
//		}
//
//		err = receiver.FileService.UpdateById(id, &fileGen)
//		if err != nil {
//			context.Error(err)
//			return
//		}
//		context.JSON(http.StatusOK, commonRes.UpdateOK)
//	}
//}
//
//func (receiver *FileController) DeleteById() gin.HandlerFunc {
//	return func(context *gin.Context) {
//		temp := context.Param("id")
//		id, err := strconv.Atoi(temp)
//		if err != nil || id <= 0 {
//			context.Error(commonRes.ParamInvalidID)
//			return
//		}
//
//		err = receiver.FileService.DeleteById(id)
//		if err != nil {
//			context.Error(err)
//			return
//		}
//		context.JSON(http.StatusOK, commonRes.DeleteOK)
//	}
//}
//
//func (receiver *FileController) Add() gin.HandlerFunc {
//	return func(context *gin.Context) {
//		//获取新的model参数
//		var request model.FileModel
//		err := context.ShouldBindJSON(&request)
//		if err != nil {
//			context.Error(commonRes.FileParamInvalid)
//			return
//		}
//
//		err = receiver.FileService.Add(&request)
//		if err != nil {
//			context.Error(err)
//			return
//		}
//		context.JSON(http.StatusOK, commonRes.AddOK)
//		return
//	}
//}
//
//func (receiver *FileController) FindAll() gin.HandlerFunc {
//	return func(context *gin.Context) {
//		var query model.FileModelQuery
//		err := context.ShouldBindJSON(&query)
//		if err != nil {
//			context.Error(commonRes.ParamInvalid)
//			return
//		}
//
//		data, err := receiver.FileService.FindAll(&query)
//		if err != nil {
//			context.Error(err)
//			return
//		}
//		context.JSON(http.StatusOK, commonRes.OK.WithData(data))
//		return
//	}
//}
//
//func (receiver *FileController) FindAllNoPagination() gin.HandlerFunc {
//	return func(context *gin.Context) {
//		data, err := receiver.FileService.FindAllNoPagination()
//		if err != nil {
//			context.Error(err)
//			return
//		}
//		context.JSON(http.StatusOK, commonRes.OK.WithData(data))
//		return
//	}
//}
