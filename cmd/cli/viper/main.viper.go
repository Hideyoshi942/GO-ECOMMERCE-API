package main

import (
	"fmt"

	viper2 "github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Port int `mapstructure:"port"`
	} `mapstructure:"server"`
	Database []struct {
		Username string `mapstructure:"username"`
		Password string `mapstructure:"password"`
		Host     string `mapstructure:"host"`
	} `mapstructure:"database"`
}

func main() {
	viper := viper2.New()
	viper.AddConfigPath("./config/") // path to config
	viper.SetConfigName("local")     // ten file
	viper.SetConfigType("yaml")

	// read configuration
	err := viper.ReadInConfig()
	if err == nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	// read server configuration
	fmt.Printf("Server host: %s \n", viper.GetString("server.host"))
	fmt.Printf("Server port: %d \n", viper.GetInt("server.port"))

	// configure structure
	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		fmt.Printf("Viper unmarshal err %v \n", err)
	}

	fmt.Printf("Config Port: %d \n", config.Server.Port)

	for _, db := range config.Database {
		fmt.Printf("Database user: %s \n", db.Username)
	}
}
