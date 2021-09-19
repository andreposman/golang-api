package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Environment struct {
	ServerPort string
	BackupServerPort string
}


func GetConfig() Environment {
	var configuration Environment
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(".")
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