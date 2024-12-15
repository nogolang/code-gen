package conf

import (
	"code-gen/configs"
	"code-gen/internal/controller"
	"code-gen/internal/utils"
	"code-gen/internal/utils/httpMiddle"
	"context"
	"embed"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"io/fs"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

type HttpServer struct {
	Logger    *zap.Logger
	Engine    *gin.Engine
	AllConfig *configs.AllConfig

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

//go:embed all:dist
var files embed.FS

func NewGin(logger *zap.Logger) *gin.Engine {
	engine := gin.New()

	//要对assets也放行，需要用sub获取到子目录
	assets, _ := fs.Sub(files, "dist/assets")
	engine.StaticFS("/assets", http.FS(assets))

	//处理html，html需要设置Content-Type为text/html
	// 实际上如果直接使用static是可以的，gin会自动添加Content-Type
	//但是我们使用的是文件系统，需要从里面读取文件
	engine.GET("/", func(c *gin.Context) {
		data, _ := files.ReadFile("dist/index.html")
		c.Header("Content-Type", "text/html")
		c.Header("Accept", "text/html")
		c.Data(http.StatusOK, "text/html", data)
	})

	//解决刷新404的问题
	//NoRoute也就是如果访问的路由不存在，则用下面的路由处理
	engine.NoRoute(func(c *gin.Context) {
		//如果它是请求一个html页面，则直接返回index.html页面
		//然后交给我们前端的路由去处理
		accept := c.Request.Header.Get("Accept")
		flag := strings.Contains(accept, "text/html")
		if flag {
			content, err := files.ReadFile("dist/index.html")
			if err != nil {
				c.Writer.WriteHeader(404)
				c.Writer.WriteString("Not Found")
				return
			}
			c.Writer.WriteHeader(200)
			c.Writer.Header().Add("Accept", "text/html")
			c.Writer.Write(content)
			c.Writer.Flush()
		}
	})

	//加载中间件
	loadMiddleware(logger, engine)

	return engine
}

type HttpServerConfig struct {
	Port string `json:"port"`
}

func (receiver *HttpServer) RunServer() {

	//注册全局验证器，里面会有自定义的一些校验
	err := utils.InitGlobalValidator()
	if err != nil {
		receiver.Logger.Sugar().Fatal(err)
		return
	}

	//启动gin
	receiver.ginServer()
}
func (receiver *HttpServer) ginServer() {
	httpServer := &http.Server{
		Addr:         ":" + fmt.Sprintf("%d", receiver.AllConfig.Server.HttpPort),
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

	receiver.Logger.Sugar().Infof("启动端口==%d", receiver.AllConfig.Server.HttpPort)

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

	receiver.Logger.Info(receiver.AllConfig.Server.ServerName + "  释放资源中....")
	receiver.Logger.Info(receiver.AllConfig.Server.ServerName + "  退出了....")

	//阻塞
	<-ctx.Done()
}
