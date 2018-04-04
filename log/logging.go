package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger

func init() {
	log := zap.NewExample()
	//if err != nil {
	//	panic(err)
	//}

	logger = log
	//zap.NewJSONEncoder(zap.TimeFormatter(TimestampField)),
	//zap.Fields(zap.Int("pid", os.Getpid()),
	//	zap.String("exe", path.Base(os.Args[0]))),
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
