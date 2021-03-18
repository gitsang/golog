package log

import "go.uber.org/zap"

var (
	logger      *zap.Logger
	sugarLogger *zap.SugaredLogger
	atomicLevel zap.AtomicLevel
)

func init() {
	InitLogger()
}
