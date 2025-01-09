package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" // 导入 MySQL 驱动
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"mygo/config"
	"mygo/internal/model"
	"mygo/pkg/logger"
)

var DB *gorm.DB

func MysqlDatabaseInit(dsn string) {
	// 提取数据库名称
	dbName := config.AppSetting.Database.MySQL.DBName
	baseDSN := fmt.Sprintf("%s:%s@tcp(%s)/",
		config.AppSetting.Database.MySQL.UserName,
		config.AppSetting.Database.MySQL.Password,
		config.AppSetting.Database.MySQL.Address,
	)

	// 连接到 MySQL，不指定具体的数据库
	sqlDB, err := sql.Open("mysql", baseDSN)
	if err != nil {
		logger.Logger.Error("MySQL连接失败: ", err)
		return
	}
	defer sqlDB.Close()

	// 检查并创建数据库
	_, err = sqlDB.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci", dbName))
	if err != nil {
		logger.Logger.Error("创建数据库失败: ", err)
		return
	}
	logger.Logger.Info("数据库已确认存在或创建成功")

	// 使用完整的 DSN 连接数据库
	dsnWithDB := fmt.Sprintf("%s%s?charset=utf8mb4&parseTime=True&loc=Local", baseDSN, dbName)
	db, err := gorm.Open(mysql.Open(dsnWithDB), &gorm.Config{})
	if err != nil {
		logger.Logger.Error("MySQL数据库初始化失败: ", err)
		return
	}

	// 获取底层 *sql.DB
	sqlDB2, err := db.DB()
	if err != nil {
		logger.Logger.Error("获取底层数据库连接失败: ", err)
		return
	}

	// 设置连接池参数
	sqlDB2.SetMaxIdleConns(config.AppSetting.Database.MySQL.MaxIdleConns)
	sqlDB2.SetMaxOpenConns(config.AppSetting.Database.MySQL.MaxOpenConns)

	// 保存全局 DB 实例
	DB = db

	// 自动迁移表
	err = db.AutoMigrate(&model.Article{})
	if err != nil {
		logger.Logger.Error("数据库表迁移失败: ", err)
		return
	}

	logger.Logger.Info("数据库初始化完成并迁移表成功")
}
