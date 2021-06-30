package main

import (
	log "github.com/gitsang/golog"
	"go.uber.org/zap"
)

var (
	s = "string"
	b = false
	i = 100

)

func sugarFun() {
	log.DebugS("sugar debug message", "k1", "v1", "k2", true)
	log.InfoS("sugar info message", "k1", "v1", "k2", true)
	log.WarnS("sugar warn message", "k1", "v1", "k2", true)
	log.ErrorS("sugar error message", "k1", "v1", "k2", true)
}

func errorFun() {
	log.Error("error message",
		zap.String("s", s),
		zap.Bool("b", b),
		zap.Int("i", i))
}

func warnFun() {
	log.Warn("warn message",
		zap.String("s", s),
		zap.Bool("b", b),
		zap.Int("i", i))
}

func infoFun() {
	log.Info("info message",
		zap.String("s", s),
		zap.Bool("b", b),
		zap.Int("i", i))
}

func debugFun() {
	log.Debug("debug message",
		zap.String("s", s),
		zap.Bool("b", b),
		zap.Int("i", i))
}

func main() {
	defer log.Sync()

	// GetLogLevel: curl http://localhost:9090/loglevel
	// SetLogLevel: curl -XPUT --data '{"level":"error"}' http://localhost:9090/loglevel
	// LogLevel: "debug" "info" "warn" "error"
	log.InitLogger(
		log.WithLogLevel(log.LevelDebug),
		log.WithEncoderType(log.EncoderTypeConsole),
		log.WithDisplayFuncEnable(true),
		log.WithEnableHttp(true),
		log.WithHttpPort(9090),
	)

	debugFun()
	infoFun()
	warnFun()
	errorFun()
}
