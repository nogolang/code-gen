package commonRes

import "net/http"

var (
	DataBaseConnectOK   = NewResponse(http.StatusOK, "连接成功")
	DataBaseConnectFail = NewResponse(12001, "连接失败")
	NeedDatabaseID      = NewResponse(12003, "需要先指定数据库配置")
)
