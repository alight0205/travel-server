package global

import (
	"travel-server/config"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	Config   *config.Config
	DB       *gorm.DB
	Redis    *redis.Client
	AliOSS   *oss.Client
	Log      *logrus.Logger
	MysqlLog logger.Interface

	Adapter  *gormadapter.Adapter
	Enforcer *casbin.Enforcer
)
