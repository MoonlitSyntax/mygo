package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mygo/config"
	"mygo/internal/db"
	"mygo/pkg/logger"
)

func main() {
	config.LoadSetting()
	AppSetting := config.AppSetting
	logger.InitLogger()
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local",
		AppSetting.Database.MySQL.UserName,
		AppSetting.Database.MySQL.Password,
		AppSetting.Database.MySQL.Address,
		AppSetting.Database.MySQL.DBName,
		AppSetting.Database.MySQL.Charset,
		AppSetting.Database.MySQL.ParseTime,
	)
	db.MysqlDatabaseInit(dsn)
	gin.SetMode(AppSetting.Server.RunMode)
	logger.Logger.Infof("Starting service %s", AppSetting.Server.HttpPort)

}
