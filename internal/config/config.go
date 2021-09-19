package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Enviroment struct {
	ServerPort string
	BackupServerPort string
}


func GetConfig() Enviroment {
	var configuration Enviroment
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./internal/config/")
	err := viper.ReadInConfig()
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}

	err2 := viper.Unmarshal(&configuration)
	if err2 != nil {
		fmt.Printf("Unable to decode into struct, %v", err2)
	}

	configuration.ServerPort = viper.GetString("HISTORY_SERVER_LISTEN_ADDR")
	configuration.BackupServerPort = viper.GetString("BACKUP_HISTORY_SERVER_LISTEN_ADDR")

	return configuration
}