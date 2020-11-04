package log

import (
	"fmt"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var sugarLogger *zap.SugaredLogger

func InitLogger(level zapcore.Level, path string) {
	writeSyncer := getLogWriter(path)
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writeSyncer, level)

	logger := zap.New(core, zap.AddCaller())
	sugarLogger = logger.Sugar()
}

func Sync() {
	err := sugarLogger.Sync()
	if err != nil {
		fmt.Println(err.Error())
	}
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogWriter(path string) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   path,
		MaxSize:    1,
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   false,
	}
	return zapcore.NewMultiWriteSyncer(
		zapcore.AddSync(lumberJackLogger),
		zapcore.AddSync(os.Stdout),
	)
}

func Debug(args ...interface{}) {
	sugarLogger.Debug(args)
}

func Info(args ...interface{}) {
	sugarLogger.Info(args)
}

func Warn(args ...interface{}) {
	sugarLogger.Warn(args)
}

func Error(args ...interface{}) {
	sugarLogger.Error(args)
}
