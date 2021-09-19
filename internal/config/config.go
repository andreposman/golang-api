package config

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Environment struct {
	ServerPort         string
	FallbackServerPort string
}


func GetConfig() Environment {
	var configuration Environment
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Error("Fatal error config file: %w \n", err)
	}

	err2 := viper.Unmarshal(&configuration)
	if err2 != nil {
		fmt.Printf("Unable to decode into struct, %v", err2)
	}

	configuration.ServerPort = viper.GetString("HISTORY_SERVER_LISTEN_ADDR")
	configuration.FallbackServerPort = viper.GetString("FALLBACK_HISTORY_SERVER_LISTEN_ADDR")

	return configuration
}