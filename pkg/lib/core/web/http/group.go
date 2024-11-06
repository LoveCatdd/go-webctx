package http

import (
	"fmt"

	"github.com/LoveCatdd/util/pkg/lib/core/log"
	"github.com/LoveCatdd/webctx/pkg/lib/core/web/middleware"
	"github.com/LoveCatdd/webctx/pkg/lib/core/web/server"
	"github.com/gin-gonic/gin"
)

type AppEngine struct {
	engine         *gin.Engine      // gin 引擎
	root           *gin.RouterGroup // gin 根路由组
	port           string           // 端口
	rootServerName string           // 根路由路径
}

var appEngine *AppEngine

func NewAppEngine(engine *gin.Engine) {
	appEngine = new(AppEngine)
	appEngine.engine = engine
	appEngine.port = server.AppConf.Server.Port
	appEngine.rootServerName = server.AppConf.Server.Name
	appEngine.root = engine.Group(fmt.Sprintf("%v", appEngine.rootServerName), middleware.TraceMiddleware()) // 注册全局trace log
}

func RootRouterGroup() *gin.RouterGroup {
	return appEngine.root
}

func Run() {
	if err := appEngine.engine.Run(fmt.Sprintf(":%v", appEngine.port)); err != nil {
		log.Fatal("Error starting server: %v", err)
	}
}
