package initialize

import (
	"fmt"
	"go-ecomerce-backend-api/global"
	"go-ecomerce-backend-api/internal/po"
	"time"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func checkErrorPanic(err error, errString string) {
	if err != nil {
		global.Logger.Error(errString, zap.Error(err))
		panic(err)
	}
}

func InitMysql() {
	m := global.Config.MySql
	dsn := "%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	var s = fmt.Sprintf(dsn, m.Username, m.Password, m.Host, m.Port, m.DBName)
	db, err := gorm.Open(mysql.Open(s), &gorm.Config{
		SkipDefaultTransaction: false,
	})
	checkErrorPanic(err, "Failed to connect to MySQL database")
	global.Logger.Info("Connected to MySQL database")
	global.Mdb = db

	// set Pool
	SetPool()

	// migrate
	MigrateTables()
}

func SetPool() {
	m := global.Config.MySql

	sqlDb, err := global.Mdb.DB()
	if err != nil {
		fmt.Printf("MySql error: %v\n", err)
	}
	sqlDb.SetConnMaxIdleTime(time.Duration(m.MaxIdleTime))
	sqlDb.SetMaxIdleConns(m.MaxIdleConns)
	sqlDb.SetMaxOpenConns(m.MaxOpenConns)
	sqlDb.SetConnMaxLifetime(time.Duration(m.MaxLifetime))
}

func MigrateTables() {
	err := global.Mdb.AutoMigrate(
		&po.User{},
		&po.Role{},
	)

	if err != nil {
		fmt.Printf("MySql error: %v\n", err)
	}
}
