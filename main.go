package main

import (
	"log"
	"travel-server/core"
	"travel-server/flags"
	"travel-server/global"
	"travel-server/routers"

	_ "travel-server/docs" // swagger docs
)

// @title travel-server API文档
// @version 2.0
// @description travel-server API文档
// @host 127.0.0.1:8000
// @BasePath /
func main() {
	// 读取配置文件
	core.InitConf()
	// 连接数据库
	global.DB = core.InitGorm()
	// 初始化OSS
	global.AliOSS = core.InitOSS()

	// 初始化命令
	option := flags.Parse()
	if option.Run() {
		return
	}
	// 初始化路由
	router := routers.InitRouter()
	addr := global.Config.System.Addr()
	log.Printf("travel-server start at: %s", addr)
	router.Run(addr)
}
