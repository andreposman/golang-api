package api

import (
	"github.com/andreposman/golang-api/internal/config"
	"github.com/andreposman/golang-api/internal/location/controller"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func InitAPI() {
	router := gin.Default()
	envConfig := config.GetConfig()
	if envConfig.ServerPort == "" {
		envConfig.ServerPort = "8080"
	}

	log.Infof("ENV VARS: HISTORY_SERVER_LISTEN_ADDR: %s  |  FALLBACK_HISTORY_SERVER_LISTEN_ADDR: %s, ", envConfig.ServerPort,envConfig.FallbackServerPort)

	api := router.Group("/location")
		api.GET("/:order_id", controller.Get)
		api.PUT("/:order_id", controller.Put)
		api.DELETE(":order_id", controller.Delete)

	err := router.Run(":" + envConfig.ServerPort)
	if err != nil {
		router.Run(":" + envConfig.FallbackServerPort)
	}

}
