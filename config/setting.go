package config

import (
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"log"
	"os"
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
			MinIdleConns int
		}
		Redis struct {
			Address  string
			Password string
			DB       int
			PoolSize int
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

	if err := viper.Unmarshal(AppSetting); err != nil {
		log.Fatalf("无法解析配置: %v", err)
	}

	AppSetting.App.JWT.Secret = os.Getenv("JWT_SECRET")
	AppSetting.Database.MySQL.Password = os.Getenv("MYSQL_PASSWORD")
	AppSetting.Database.MySQL.Address = os.Getenv("MYSQL_ADDR")
	AppSetting.Database.Redis.Password = os.Getenv("REDIS_PASSWORD")
	AppSetting.Database.Redis.Address = os.Getenv("REDIS_ADDR")

	log.Println("配置文件加载完成")
}
