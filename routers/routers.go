package routers

import (
	"travel-server/global"
	"travel-server/middleware"
	"travel-server/module/article"
	"travel-server/module/auth"
	"travel-server/module/common"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	gs "github.com/swaggo/gin-swagger"
)

type RouterGroup struct {
	*gin.RouterGroup
}

func InitRouter() *gin.Engine {
	gin.SetMode(global.Config.System.Env)
	router := gin.Default()
	router.Use(middleware.Cors())

	api := router.Group("/api") // 正常路由
	router.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))

	auth.InitRouter(api)
	common.InitRouter(api)
	article.InitRouter(api)
	return router
}
