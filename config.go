package log

import (
	"os"
	"strings"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	EncoderTypeConsole = "console"
	EncoderTypeJson    = "json"
)

type EncoderConfig struct {
	EncoderType      string
	MessageKey       string
	LevelKey         string
	TimeKey          string
	NameKey          string
	CallerKey        string
	FunctionKey      string
	StacktraceKey    string
	LineEnding       string
	EncodeLevel      zapcore.LevelEncoder
	EncodeTime       zapcore.TimeEncoder
	EncodeDuration   zapcore.DurationEncoder
	EncodeCaller     zapcore.CallerEncoder
	EncodeName       zapcore.NameEncoder
	ConsoleSeparator string
}

type LogFileConfig struct {
	Filename   string
	MaxSize    int
	MaxBackups int
	MaxAge     int
	Compress   bool
}

type LogConfig struct {
	LogLevel   zapcore.Level
	EnableHttp bool
	HttpPort   int
	EncoderConfig
	LogFileConfig
}

type LogConfigOption func(config *LogConfig)

func defaultLogConfig() *LogConfig {
	return &LogConfig{
		LogLevel:   zapcore.InfoLevel,
		EnableHttp: false,
		EncoderConfig: EncoderConfig{
			EncoderType:      EncoderTypeConsole,
			MessageKey:       "msg",
			LevelKey:         "level",
			TimeKey:          "ts",
			NameKey:          "logger",
			CallerKey:        "caller",
			FunctionKey:      zapcore.OmitKey,
			StacktraceKey:    "trace",
			LineEnding:       zapcore.DefaultLineEnding,
			EncodeLevel:      zapcore.CapitalLevelEncoder,
			EncodeTime:       zapcore.ISO8601TimeEncoder,
			EncodeDuration:   zapcore.MillisDurationEncoder,
			EncodeCaller:     zapcore.ShortCallerEncoder,
			EncodeName:       zapcore.FullNameEncoder,
			ConsoleSeparator: "",
		},
		LogFileConfig: LogFileConfig{
			Filename:   "",
			MaxSize:    1,
			MaxBackups: 5,
			MaxAge:     30,
			Compress:   true,
		},
	}
}

func WithLogLevel(level string) LogConfigOption {
	return func(config *LogConfig) {
		config.LogLevel = StringToLogLevel(level)
	}
}

func WithEnableHttp(enable bool) LogConfigOption {
	return func(config *LogConfig) {
		config.EnableHttp = enable
	}
}

func WithHttpPort(port int) LogConfigOption {
	return func(config *LogConfig) {
		config.HttpPort = port
	}
}

func WithEncoderType(t string) LogConfigOption {
	return func(config *LogConfig) {
		config.EncoderConfig.EncoderType = strings.ToLower(t)
	}
}

func WithLogFile(file string) LogConfigOption {
	return func(config *LogConfig) {
		config.LogFileConfig.Filename = file
	}
}

func WithLogFileCompress(compress bool) LogConfigOption {
	return func(config *LogConfig) {
		config.LogFileConfig.Compress = compress
	}
}

func WithDisplayFuncEnable(enable bool) LogConfigOption {
	return func(config *LogConfig) {
		if enable {
			config.FunctionKey = "func"
		} else {
			config.FunctionKey = zapcore.OmitKey
		}
	}
}

func InitLogger(opts ...LogConfigOption) {
	conf := defaultLogConfig()
	for _, apply := range opts {
		apply(conf)
	}

	encoderConfig := zapcore.EncoderConfig{
		MessageKey:     conf.MessageKey,
		LevelKey:       conf.LevelKey,
		TimeKey:        conf.TimeKey,
		NameKey:        conf.NameKey,
		CallerKey:      conf.CallerKey,
		FunctionKey:    conf.FunctionKey,
		StacktraceKey:  conf.StacktraceKey,
		LineEnding:     conf.LineEnding,
		EncodeLevel:    conf.EncodeLevel,
		EncodeTime:     conf.EncodeTime,
		EncodeDuration: conf.EncodeDuration,
		EncodeCaller:   conf.EncodeCaller,
		EncodeName:     conf.EncodeName,
	}

	var encoder zapcore.Encoder
	if conf.EncoderType == EncoderTypeConsole {
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	} else if conf.EncoderType == EncoderTypeJson {
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	} else {
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	}

	writeSyncer := zapcore.NewMultiWriteSyncer(
		zapcore.AddSync(&lumberjack.Logger{
			Filename:   conf.Filename,
			MaxSize:    conf.MaxAge,
			MaxBackups: conf.MaxBackups,
			MaxAge:     conf.MaxAge,
			Compress:   conf.Compress,
		}),
		zapcore.AddSync(os.Stdout),
	)

	atomicLevel = zap.NewAtomicLevelAt(conf.LogLevel)

	core := zapcore.NewCore(encoder, writeSyncer, atomicLevel)

	logger = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	sugarLogger = logger.Sugar()

	if conf.EnableHttp && conf.HttpPort != 0 {
		StartLogLevelHttpHandle(conf.HttpPort)
	}
}
