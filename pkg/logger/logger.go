package logger

import (
	"go-ecomerce-backend-api/pkg/setting"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type LoggerZap struct {
	*zap.Logger
}

func NewLogger(config setting.LoggerSetting) *LoggerZap {
	logLevel := config.LogLevel
	// debug-> info -> warn -> error -> fatal -> panic
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
	case "fatal":
		level = zapcore.FatalLevel
	}
	encoder := getEncoderLog()
	hook := lumberjack.Logger{
		Filename:   config.FileLogName, // path save
		MaxSize:    config.MaxSize,     // megabytes
		MaxBackups: config.MaxBackups,
		MaxAge:     config.MaxAge,   // days
		Compress:   config.Compress, // disable by default
	}
	core := zapcore.NewCore(
		encoder,
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)),
		level)
	//logger := zap.New(core, zap.AddCaller())
	return &LoggerZap{zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))}
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
