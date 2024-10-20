package httpMiddle

import (
	"code-gen/internal/utils/commonRes"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

// MyGinZap 自定义中间件
func MyGinZap(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		//取消caller，因为这是放到中间件调用的，没必要用caller
		newLogger := logger.WithOptions(zap.WithCaller(false))

		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		method := c.Request.Method
		allRequestStr := method + " " + path
		if query != "" {
			allRequestStr = path + "?" + query
		}

		//调用下一个中间件，计算出时间
		//但是这里不这么做，为了更容易观测，这里就不计算时间了，时间由外部计算即可
		c.Next()
		latency := time.Since(start)

		fields := []zapcore.Field{
			zap.Int("status", c.Writer.Status()),
			zap.String("ip", c.ClientIP()),
			//zap.String("user-agent", c.Request.UserAgent()),
			zap.Duration("latency", latency),
		}

		//ginzap这里判断c.Errors，看看我们是否往c.Error里塞入了错误
		//并且判断是否是自定义的错误，如果是自定义的错误，则打印info即可，因为我们已经处理了
		if len(c.Errors) > 0 {
			for _, e := range c.Errors {
				//转成原始的error
				err := e.Err
				//myResponse := commonRes.Response{}

				//判断是不是自定义的错误，打印info即可
				//如果不是，则可能是第三方的错误，比如gorm，打印Error级别
				var response *commonRes.Response
				if errors.As(err, &response) {
					newLogger.Info(allRequestStr, fields...)
				} else {

					//处理第三方的错误
					newLogger.Error(allRequestStr, fields...)
					//同时打印出错误日志,这样是比较好的
					newLogger.Sugar().Errorf("%+v", err)
				}
			}
		} else {
			//如果调用过程中没有产生error，则打印info
			newLogger.Info(allRequestStr, fields...)
		}

		//调用和下一个中间件
		c.Next()

	}
}
