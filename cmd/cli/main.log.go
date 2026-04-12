package main

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	// sugar use in
	//sugar := zap.NewExample().Sugar()
	//sugar.Infof("failed to fetch URL",
	//	// Structured context as loosely typed key-value pairs.
	//	"url", "http://example.com",
	//	"attempt", 3,
	//	"backoff", time.Second,
	//) // like fmt.Printf()

	// logger.go use in
	//logger.go := zap.NewExample()
	//logger.go.Info("Hello", zap.String("name", "John"), zap.Int("age", 25))
	//
	//logger.go.Info("Hello NewExample")
	//
	//logger.go, _ = zap.NewDevelopment()
	//logger.go.Info("Hello NewDevelopment")
	//
	//logger.go, _ = zap.NewProduction()
	//logger.go.Info("Hello NewProduction")

	//
	encoder := getEncoderLog()
	sync := getWriterSync()
	core := zapcore.NewCore(encoder, sync, zapcore.InfoLevel)
	logger := zap.New(core, zap.AddCaller())

	logger.Info("Info log", zap.Int("line", 1))
	logger.Error("Error log", zap.Int("line", 1))

}

// format log
func getEncoderLog() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()

	// convert timestamp -> localdatetime
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	// ts -> time
	encoderConfig.TimeKey = "time"

	// from info INFO
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	// "caller": "cli/main.log.go:24"
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder

	return zapcore.NewJSONEncoder(encoderConfig)
}

// write log into file
func getWriterSync() zapcore.WriteSyncer {
	file, _ := os.OpenFile("./log/log.txt", os.O_CREATE|os.O_WRONLY, os.ModePerm)
	syncFile := zapcore.AddSync(file)
	syncConsole := zapcore.AddSync(os.Stderr)

	return zapcore.NewMultiWriteSyncer(syncFile, syncConsole)
}
