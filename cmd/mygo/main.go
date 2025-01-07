package main

import (
	"github.com/gin-gonic/gin"
	"mygo/config"
	"mygo/pkg/logger"
)

func main() {
	config.LoadSetting()
	AppSetting := config.AppSetting
	logger.InitLogger()

	gin.SetMode(AppSetting.Server.RunMode)
	logger.Logger.Infof("Starting service %s", AppSetting.Server.HttpPort)
}
