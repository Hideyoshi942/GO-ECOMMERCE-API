package viper

import (
	"fmt"

	viper2 "github.com/spf13/viper"
)

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

}
