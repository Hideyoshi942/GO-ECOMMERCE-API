package initialize

import (
	"fmt"
	"go-ecomerce-backend-api/global"
	"log"

	"go.uber.org/zap"
)

func Run() {

	// load configuration
	LoadConfig()
	m := global.Config.MySql
	fmt.Println("Loading configuration..." + m.Host)
	InitLogger()
	global.Logger.Info("Configuration log successfully", zap.String("oke", "success"))
	InitMysql()
	InitRedis()

	r := InitRouter()

	if err := r.Run(":8082"); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
