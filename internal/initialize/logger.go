package initialize

import (
	"go-ecomerce-backend-api/global"
	"go-ecomerce-backend-api/pkg/logger"
)

func InitLogger() {
	global.Logger = logger.NewLogger(global.Config.Logger)
}
