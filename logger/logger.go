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
	config.EncoderConfig = encodedConfig

	log, err = config.Build(zap.AddCallerSkip(1))

	// log, err = zap.NewProduction(zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}
}

func Info(messager string, fields ...zap.Field) {
	log.Info(messager, fields...)
}

func Debug(messager string, fields ...zap.Field) {
	log.Debug(messager, fields...)
}

func Error(messager string, fields ...zap.Field) {
	log.Error(messager, fields...)
}
