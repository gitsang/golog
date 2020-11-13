package main

import (
	log "github.com/gitsang/golog"
	"go.uber.org/zap"
)

func logWithoutInit() {
	log.Info("test without init")
}

func logWithInit() {
	log.InitLogger(zap.InfoLevel, "main.log")
	log.Info("test with init")
}

func main() {
	logWithoutInit()
	logWithInit()
}
