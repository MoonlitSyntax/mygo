package util

import (
	"fmt"
	"mygo/config"
	"mygo/internal/db"
	"mygo/pkg/logger"
)

func InitAll() {
	config.LoadSetting()
	logger.InitLogger()
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local",
		config.AppSetting.Database.MySQL.UserName,
		config.AppSetting.Database.MySQL.Password,
		config.AppSetting.Database.MySQL.Address,
		config.AppSetting.Database.MySQL.DBName,
		config.AppSetting.Database.MySQL.Charset,
		config.AppSetting.Database.MySQL.ParseTime,
	)
	db.MysqlDatabaseInit(dsn)

	logger.Logger.Info("系统初始化完成")
}
