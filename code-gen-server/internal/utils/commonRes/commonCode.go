package commonRes

import "net/http"

var (
	OK             = NewResponse(http.StatusOK, "成功")
	UpdateOK       = NewResponse(http.StatusOK, "更新成功")
	DeleteOK       = NewResponse(http.StatusOK, "删除成功")
	AddOK          = NewResponse(http.StatusOK, "添加成功")
	ParamInvalid   = NewResponse(http.StatusBadRequest, "参数错误")
	ParamInvalidID = NewResponse(http.StatusBadRequest, "无效的id")
)
