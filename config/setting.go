package config

import (
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"log"
	"strings"
)

type Setting struct {
	Server struct {
		RunMode      string
		HttpPort     int
		ReadTimeout  int
		WriteTimeout int
	}
	App struct {
		DefaultPageSize int
		MaxPageSize     int
		Logger          struct {
			LogFilePath       string
			LogFileMaxSize    int
			LogFileMaxAge     int
			LogFileMaxBackups int
			LogLevel          string
		}
		JWT struct {
			Secret     string
			Issuer     string
			ExpireTime int
		}
		ContextTimeout        int
		DefaultContextTimeout int
	}
	Database struct {
		MySQL struct {
			UserName     string
			Password     string
			Address      string
			DBName       string
			TablePrefix  string
			Charset      string
			ParseTime    bool
			MaxIdleConns int
			MaxOpenConns int
		}
		Redis struct {
			Address      string
			Password     string
			DB           int
			PoolSize     int
			MaxIdleConns int
			MaxOpenConns int
		}
	}
}

var AppSetting = &Setting{}

func LoadSetting() {
	if err := godotenv.Load("./config/.env"); err != nil {
		log.Println("未找到 .env 文件，尝试使用系统环境变量")
	}
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config/")

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("无法读取配置文件: %v", err)
	}

	AppSetting = &Setting{}

	err := viper.BindEnv("Database.MySQL.Password", "MYSQL_PASSWORD")
	if err != nil {
		return
	}
	err = viper.BindEnv("Database.MySQL.Address", "MYSQL_ADDR")
	if err != nil {
		return
	}
	err = viper.BindEnv("Database.Redis.Password", "REDIS_PASSWORD")
	if err != nil {
		return
	}
	err = viper.BindEnv("Database.Redis.Address", "REDIS_ADDR")
	if err != nil {
		return
	}
	err = viper.BindEnv("Database.Redis.Address", "JWT_SECRET")
	if err != nil {
		return
	}
	if err := viper.Unmarshal(AppSetting); err != nil {
		log.Fatalf("无法解析配置: %v", err)
	}

	log.Println("配置文件加载完成")
}
