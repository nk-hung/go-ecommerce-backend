package main

import (
	"log"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	// sugar := zap.NewExample().Sugar()
	// sugar.Infof("Hello name :%s, age: %d ", "TipGO", 26)

	// logger
	// logger := zap.NewExample()
	// logger.Info("Hello", zap.String("name", "TipsGo"), zap.Int("age", 26))

	// logger := zap.NewExample()
	// logger.Info("Hello NewExample")
	//
	// // Development
	// logger, _ = zap.NewDevelopment()
	// logger.Info("Hello NewDevelopment")
	//
	// // Production
	// logger, _ = zap.NewProduction()
	// logger.Info("Hello NewProduction")

	// 3.
	encoder := getEncoderLog()
	sync := getWriterSync()
	core := zapcore.NewCore(encoder, sync, zapcore.InfoLevel)

	logger := zap.New(core, zap.AddCaller())

	logger.Info("Info logger", zap.Int("line", 1))
	logger.Error("Info Error", zap.Int("line", 2))
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

func getWriterSync() zapcore.WriteSyncer {
	file, err := os.OpenFile("./log/log.txt", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	// file, _ := os.Create("./log/log.txt")
	syncFile := zapcore.AddSync(file)
	syncConsole := zapcore.AddSync(os.Stderr)

	return zapcore.NewMultiWriteSyncer(syncConsole, syncFile)
}
