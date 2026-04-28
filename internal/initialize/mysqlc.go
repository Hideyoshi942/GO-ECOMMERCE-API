package initialize

import (
	"database/sql"
	"fmt"
	"go-ecomerce-backend-api/global"
	"go-ecomerce-backend-api/internal/po"
	"time"

	"go.uber.org/zap"
	"gorm.io/gen"
)

func checkErrorPanicC(err error, errString string) {
	if err != nil {
		global.Logger.Error(errString, zap.Error(err))
		panic(err)
	}
}

func InitMysqlC() {
	m := global.Config.MySql
	dsn := "%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	var s = fmt.Sprintf(dsn, m.Username, m.Password, m.Host, m.Port, m.DBName)
	db, err := sql.Open("mysql", s)
	checkErrorPanicC(err, "Failed to connect to MySQL database")
	global.Logger.Info("Connected to MySQL database")
	global.Mdbc = db

	// set Pool
	SetPoolC()

	// gen
	genTableDAOC()

	// migrate
	//MigrateTablesC()
}

func SetPoolC() {
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

func genTableDAOC() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "./internal/models",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
	})

	g.UseDB(global.Mdb)
	//g.GenerateAllTable()
	//g.GenerateModel("users", gen.FieldIgnore("password"))
	g.Execute()
}

func MigrateTablesC() {
	err := global.Mdb.AutoMigrate(
		&po.User{},
		&po.Role{},
	)

	if err != nil {
		fmt.Printf("MySql error: %v\n", err)
	}
}
