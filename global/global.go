package global

import (
	"database/sql"
	"go-ecomerce-backend-api/pkg/logger"
	"go-ecomerce-backend-api/pkg/setting"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	Config setting.Config
	Logger *logger.LoggerZap
	Rdb    *redis.Client
	Mdb    *gorm.DB
	Mdbc   *sql.DB
)

/*
Config
Redis
Mysql
...
*/
