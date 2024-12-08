package conf

import (
	"code-gen/internal/controller"
	"code-gen/internal/utils"
	"code-gen/internal/utils/httpMiddle"
	"context"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type HttpServer struct {
	Logger *zap.Logger
	Engine *gin.Engine

	//我们写了什么controller,需要在这里进行加入
	//并且编写对应的provider
	//FileControl    *controller.FileController
	GroupControl   *controller.GroupController
	OrmControl     *controller.OrmController
	MappingControl *controller.MappingController
	FileGenControl *controller.FileGenController
}

func loadMiddleware(logger *zap.Logger, engine *gin.Engine) {
	//跨域中间件，如果用了网关，在网关里设置了，这里就不需要了
	engine.Use(cors.New(cors.Config{
		AllowAllOrigins:  true, // 允许所有来源
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: true,
	}))

	//引入ginzap中间件，gin的日志打印由zap来完成
	//主要是因为生产环境下，需要打印错误日志到文件中
	//engine.Use(ginzap.Ginzap(logger, time.DateTime, false))

	//引入我们自定义的中间件
	engine.Use(httpMiddle.MyGinZap(logger))

	//统一错误处理中间件,所有的错误向外抛出，由gin处理
	engine.Use(httpMiddle.HandlerError())

	//未知错误恢复错误中间件
	engine.Use(gin.CustomRecovery(httpMiddle.DefaultHandleRecovery))
}

func NewGin(logger *zap.Logger) *gin.Engine {
	engine := gin.New()

	//加载中间件
	loadMiddleware(logger, engine)

	return engine
}

type HttpServerConfig struct {
	Port string `json:"port"`
}

func (receiver *HttpServer) RunServer() {
	var config HttpServerConfig
	err := viper.UnmarshalKey("http", &config)
	if err != nil {
		receiver.Logger.Fatal("读取http配置失败", zap.Error(err))
		return
	}

	//注册全局验证器，里面会有自定义的一些校验
	err = utils.InitGlobalValidator()
	if err != nil {
		receiver.Logger.Sugar().Fatal(err)
		return
	}

	//启动gin
	receiver.ginServer(config)
}
func (receiver *HttpServer) ginServer(config HttpServerConfig) {
	httpServer := &http.Server{
		Addr:         ":" + config.Port,
		Handler:      receiver.Engine,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	//协程启动服务器
	go func() {
		err := httpServer.ListenAndServe()
		if err != nil {
			receiver.Logger.Error("服务已关闭：", zap.Error(err))
		}
	}()

	//创建信号，返回一个channel
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)

	//程序关闭了，则协程有值，执行到这里
	<-quit

	//Shutdown之后不会再接收新的请求
	//原生http.Shutdown 会通知ListenAndServe，这里不会
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if err := httpServer.Shutdown(ctx); err != nil {
		receiver.Logger.Error("关闭服务错误", zap.Error(err))
	}

	receiver.Logger.Info("释放资源中....")

	//阻塞
	<-ctx.Done()
	receiver.Logger.Info("退出了....")
}
