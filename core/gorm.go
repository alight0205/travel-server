package core

import (
	"log"
	"time"
	"travel-server/global"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitGorm() *gorm.DB {

	if global.Config.Mysql.Host == "" {
		global.Log.Warning("mysql config error")
		return nil
	}
	dsn := global.Config.Mysql.Dsn()

	var mysqlLogger logger.Interface

	if global.Config.System.Env == "debug" {
		mysqlLogger = logger.Default.LogMode(logger.Info) // 开发环境显示所有sql
	} else {
		mysqlLogger = logger.Default.LogMode(logger.Error) // 只打印错误的sql
	}
	global.MysqlLog = logger.Default.LogMode(logger.Info)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: mysqlLogger,
	})
	if err != nil {
		global.Log.Errorf("mysql connect error: %s", dsn)
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(global.Config.Mysql.MaxIdleConns)
	sqlDB.SetMaxOpenConns(global.Config.Mysql.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Duration(global.Config.Mysql.ConnMaxLifetime) * time.Hour)

	log.Println("mysql load Init success")
	return db
}
