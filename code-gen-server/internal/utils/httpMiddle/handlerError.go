package httpMiddle

import (
	"code-gen/internal/utils/commonRes"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"net/http"
)

// DefaultHandleRecovery
// 未知的错误
// gin的customerRecovery会把error传递过来
func DefaultHandleRecovery(c *gin.Context, error any) {
	//这里使用全局的logger，无伤大雅
	zap.L().Error("系统出现panic", zap.Any("error", error))
	c.JSON(http.StatusInternalServerError, gin.H{
		"code":    http.StatusInternalServerError,
		"message": "发生未知错误,请联系管理员提交bug的信息",
	})
	return
}

func HandlerError() gin.HandlerFunc {
	return func(context *gin.Context) {

		//让请求和其它中间件走
		context.Next()

		//这里只当我们使用ctx.Error(err)把错误设置进去了，这里才会有值
		for _, e := range context.Errors {
			//把gin.err返回成error
			err := e.Err

			//把error转换到我们的错误，因为我们的Response实现了error接口的
			//如果是我们的错误，则返回我们的信息给前台
			var myResponse *commonRes.Response
			if errors.As(err, &myResponse) {
				//这里没有传递data，因为错误是没有data的
				context.JSON(http.StatusOK, gin.H{
					"code":    myResponse.Code,
					"message": myResponse.Message,
					"reason":  myResponse.Reason,
				})
			} else {
				//如果不是我们返回的错误，比如gorm的语句错误
				//我们不打印错误，打印放到ginzap自定义的中间件里
				context.JSON(http.StatusInternalServerError, gin.H{
					"code":    500,
					"message": "发生未知错误,请联系管理员提交bug的信息",
				})
			}

			//直接return,不返回可能有多个错误给前台
			return
		}

	}
}
