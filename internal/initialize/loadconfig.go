package initialize

import (
	"fmt"
	"go-ecomerce-backend-api/global"

	viper2 "github.com/spf13/viper"
)

func LoadConfig() {
	viper := viper2.New()
	viper.AddConfigPath("./config/") // path to config
	viper.SetConfigName("local")     // ten file
	viper.SetConfigType("yaml")

	// read configuration
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	// read server configuration
	fmt.Printf("Server host: %v \n", viper.GetString("server.host"))
	fmt.Printf("Server port: %d \n", viper.GetInt("server.port"))

	// configure structure
	if err := viper.Unmarshal(&global.Config); err != nil {
		fmt.Printf("Viper unmarshal err %v \n", err)
	}
}
