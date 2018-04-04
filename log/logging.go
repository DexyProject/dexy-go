package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger

func init() {
	logger = zap.NewExample()
}

func Logger() *zap.Logger {
	return logger
}

func Debug(msg string, fields ...zapcore.Field) {
	Logger().Error(msg, fields...)
}

func Info(msg string, fields ...zapcore.Field) {
	Logger().Error(msg, fields...)
}

func Warn(msg string, fields ...zapcore.Field) {
	Logger().Error(msg, fields...)
}

func Error(msg string, fields ...zapcore.Field) {
	Logger().Error(msg, fields...)
}

func Fatal(msg string, fields ...zapcore.Field) {
	Logger().Fatal(msg, fields...)
}
