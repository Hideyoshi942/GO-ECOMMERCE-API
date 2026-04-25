package initialize

import (
	"bytes"
	"fmt"
	"go-ecomerce-backend-api/global"
	"os"
	"strings"

	"github.com/subosito/gotenv"
	viper2 "github.com/spf13/viper"
)

func LoadConfig() {
	viper := viper2.New()

	if err := loadEnvFiles(".env", ".evn"); err != nil {
		panic(fmt.Errorf("fatal error env file: %w", err))
	}

	viper.SetConfigType("yaml")
	configBytes, err := os.ReadFile("./config/local.yaml")
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	expandedConfig := os.ExpandEnv(string(configBytes))

	// Base config from yaml is required.
	if err := viper.ReadConfig(bytes.NewBufferString(expandedConfig)); err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AllowEmptyEnv(true)
	viper.AutomaticEnv()
	bindEnvs(viper)

	fmt.Printf("Server mode: %s\n", viper.GetString("server.mode"))
	fmt.Printf("Server port: %d\n", viper.GetInt("server.port"))

	if err := viper.Unmarshal(&global.Config); err != nil {
		fmt.Printf("viper unmarshal err %v\n", err)
	}
}

func loadEnvFiles(filePaths ...string) error {
	for _, filePath := range filePaths {
		if _, err := os.Stat(filePath); err == nil {
			if err := gotenv.Load(filePath); err != nil {
				return fmt.Errorf("%s: %w", filePath, err)
			}
		}
	}

	return nil
}

func bindEnvs(viper *viper2.Viper) {
	_ = viper.BindEnv("server.port", "SERVER_PORT")
	_ = viper.BindEnv("server.mode", "SERVER_MODE")

	_ = viper.BindEnv("mysql.host", "MYSQL_HOST")
	_ = viper.BindEnv("mysql.port", "MYSQL_PORT")
	_ = viper.BindEnv("mysql.username", "MYSQL_USERNAME")
	_ = viper.BindEnv("mysql.password", "MYSQL_PASSWORD")
	_ = viper.BindEnv("mysql.dbname", "MYSQL_DBNAME")
	_ = viper.BindEnv("mysql.maxIdleConns", "MYSQL_MAX_IDLE_CONNS")
	_ = viper.BindEnv("mysql.maxOpenConns", "MYSQL_MAX_OPEN_CONNS")
	_ = viper.BindEnv("mysql.maxLifetime", "MYSQL_MAX_LIFETIME")
	_ = viper.BindEnv("mysql.maxIdleTime", "MYSQL_MAX_IDLE_TIME")
	_ = viper.BindEnv("mysql.maxRetries", "MYSQL_MAX_RETRIES")
	_ = viper.BindEnv("mysql.retryInterval", "MYSQL_RETRY_INTERVAL")
	_ = viper.BindEnv("mysql.readTimeout", "MYSQL_READ_TIMEOUT")
	_ = viper.BindEnv("mysql.writeTimeout", "MYSQL_WRITE_TIMEOUT")
	_ = viper.BindEnv("mysql.timeout", "MYSQL_TIMEOUT")
	_ = viper.BindEnv("mysql.idleTimeout", "MYSQL_IDLE_TIMEOUT")

	_ = viper.BindEnv("logger.log_level", "LOGGER_LOG_LEVEL")
	_ = viper.BindEnv("logger.file_log_name", "LOGGER_FILE_LOG_NAME")
	_ = viper.BindEnv("logger.max_size", "LOGGER_MAX_SIZE")
	_ = viper.BindEnv("logger.max_backups", "LOGGER_MAX_BACKUPS")
	_ = viper.BindEnv("logger.max_age", "LOGGER_MAX_AGE")
	_ = viper.BindEnv("logger.compress", "LOGGER_COMPRESS")

	_ = viper.BindEnv("redis.host", "REDIS_HOST")
	_ = viper.BindEnv("redis.port", "REDIS_PORT")
	_ = viper.BindEnv("redis.password", "REDIS_PASSWORD")
	_ = viper.BindEnv("redis.db", "REDIS_DB")
	_ = viper.BindEnv("redis.pool_size", "REDIS_POOL_SIZE")
}
