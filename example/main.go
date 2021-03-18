package main

import (
	log "github.com/gitsang/golog"
	"go.uber.org/zap"
	"time"
)

func main() {
	defer log.Sync()

	// GetLogLevel: curl http://localhost:9090/loglevel
	// SetLogLevel: curl -XPUT --data '{"level":"error"}' http://localhost:9090/loglevel
	// LogLevel: "debug" "info" "warn" "error" "dpanic" "panic" "fatal"
	log.InitLogger(
		log.WithLogLevel("debug"),
		log.WithEncoderType("console"),
		log.WithEnableHttp(true),
		log.WithHttpPort(9090),
	)

	for {
		log.Debug("debug message", zap.Any("k", "v"), zap.Any("k2", "v2"))
		log.Info("info message", zap.Any("k", "v"), zap.Any("k2", "v2"))
		log.Warn("warn message", zap.Any("k", "v"), zap.Any("k2", "v2"))
		log.Error("error message", zap.Any("k", "v"), zap.Any("k2", "v2"))

		log.DebugS("debug message", "k1", "v1", "k2", true)
		log.InfoS("info message", "k1", "v1", "k2", true)
		log.WarnS("warn message", "k1", "v1", "k2", true)
		log.ErrorS("error message", "k1", "v1", "k2", true)

		time.Sleep(10 * time.Second)
	}
}
