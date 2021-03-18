package log

import (
	"fmt"
	"go.uber.org/zap"
)

func Sync() {
	err := logger.Sync()
	if err != nil {
		fmt.Println(err.Error())
	}
}

func Debug(msg string, fields ...zap.Field) {
	logger.Debug(msg, fields...)
}

func Info(msg string, fields ...zap.Field) {
	logger.Info(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	logger.Warn(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	logger.Error(msg, fields...)
}

func SyncS() {
	err := sugarLogger.Sync()
	if err != nil {
		fmt.Println(err.Error())
	}
}

func DebugS(args ...interface{}) {
	sugarLogger.Debug(args)
}

func InfoS(args ...interface{}) {
	sugarLogger.Info(args)
}

func WarnS(args ...interface{}) {
	sugarLogger.Warn(args)
}

func ErrorS(args ...interface{}) {
	sugarLogger.Error(args)
}
