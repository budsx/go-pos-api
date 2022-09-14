package helpers

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.Logger

func init() {

	config := zap.NewProductionConfig()

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.StacktraceKey = ""

	config.EncoderConfig = encoderConfig

	var err error
	log, err = config.Build(zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}
}

func Info(message string, fields ...zapcore.Field) {
	log.Info(message, fields...)
}

func Debug(message string, fields ...zapcore.Field) {
	log.Debug(message, fields...)
}

func Error(message string, fields ...zapcore.Field) {
	log.Error(message, fields...)
}

func Fatal(message string, fields ...zapcore.Field) {
	log.Fatal(message, fields...)
}
