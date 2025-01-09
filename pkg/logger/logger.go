package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"mygo/config"
)

var Logger *zap.SugaredLogger

// InitLogger 初始化 SugaredLogger
func InitLogger() {
	// 配置日志文件切分
	writeSyncer := zapcore.AddSync(&lumberjack.Logger{
		Filename:   config.AppSetting.App.Logger.LogFilePath,
		MaxSize:    config.AppSetting.App.Logger.LogFileMaxSize,
		MaxAge:     config.AppSetting.App.Logger.LogFileMaxAge,
		MaxBackups: config.AppSetting.App.Logger.LogFileMaxBackups,
		Compress:   true,
	})

	// 配置日志格式
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder // 时间格式
	encoderConfig.LevelKey = "level"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.CallerKey = "caller"

	// 配置日志级别
	logLevel := parseLogLevel(config.AppSetting.App.Logger.LogLevel)

	// 创建核心日志器
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig), // JSON 格式日志
		writeSyncer,                           // 输出到文件
		logLevel,                              // 日志级别过滤
	)

	// 创建 Logger
	logger := zap.New(core, zap.AddCaller(),
		zap.AddStacktrace(zapcore.ErrorLevel),
	)

	Logger = logger.Sugar()
}

func parseLogLevel(level string) zapcore.Level {
	switch level {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.ErrorLevel
	case "fatal":
		return zapcore.FatalLevel
	default:
		return zapcore.InfoLevel
	}
}
