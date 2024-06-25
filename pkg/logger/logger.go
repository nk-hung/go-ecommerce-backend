package logger

import (
	"os"

	"github.com/nk-hung/go-ecommerce-backend-api/pkg/setting"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type LoggerZap struct {
	*zap.Logger
}

func NewLogger(config setting.LoggerSetting) *LoggerZap {
	logLevel := config.Log_level

	var level zapcore.Level
	switch logLevel {
	case "debug":
		level = zapcore.DebugLevel
	case "info":
		level = zapcore.InfoLevel
	case "warn":
		level = zapcore.WarnLevel
	case "error":
		level = zapcore.ErrorLevel
	default:
		level = zapcore.InfoLevel
	}

	encoder := getEncoderLog()
	hook := lumberjack.Logger{
		Filename:   config.File_log_name,
		MaxSize:    config.Max_size, // megabytes
		MaxBackups: config.Max_backups,
		MaxAge:     config.Max_age, // days
		Compress:   config.Compress,
	}
	core := zapcore.NewCore(
		encoder,
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)),
		level,
	)

	// logger := zap.New(core, zap.AddCaller())
	return &LoggerZap{zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))}
}

// Format log a message
func getEncoderLog() zapcore.Encoder {
	encodeConfig := zap.NewProductionEncoderConfig()

	// Ex: "level":"info","ts":1717833583.0767105,"caller":"cli/main.log.go:22","msg":"Hello NewProduction"
	// timestamp to datetime
	encodeConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	// ts => Time
	encodeConfig.TimeKey = "Time"
	// level: info => INFO
	encodeConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	// caller
	encodeConfig.EncodeCaller = zapcore.ShortCallerEncoder //

	return zapcore.NewJSONEncoder(encodeConfig)
}
