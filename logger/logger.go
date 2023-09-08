package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	log *zap.Logger
	err error
)

func init() {
	config := zap.NewProductionConfig()
	encodedConfig := zap.NewProductionEncoderConfig()
	encodedConfig.TimeKey = "timestamp"
	encodedConfig.MessageKey = "message"
	encodedConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encodedConfig.StacktraceKey = ""
	config.EncoderConfig = encodedConfig

	log, err = config.Build(zap.AddCallerSkip(1))
	// log, err = zap.NewProduction(zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}
}

func Info(message string, fields ...zap.Field) {
	log.Info(message, fields...)
}

func Debug(message string, fields ...zap.Field) {
	log.Debug(message, fields...)
}

func Error(message string, fields ...zap.Field) {
	log.Error(message, fields...)
}
