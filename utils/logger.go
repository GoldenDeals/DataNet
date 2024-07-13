package utils

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	loggerFlush      map[string]func() error
	loggerInitalized = false
)

func InitLoggers() {
	loggerFlush = make(map[string]func() error)
	loggerInitalized = true
}

func NewLogger(sys string) *zap.SugaredLogger {
	if !loggerInitalized {
		panic("logger is not initialized")
	}
	_, ok := loggerFlush[sys]
	if ok {
		panic("logger sys exiting")
	}
	cfg := zap.Config{
		Level:    getLevelFromString(viper.GetString("log.level")),
		Encoding: viper.GetString("log.encoding"),
		EncoderConfig: zapcore.EncoderConfig{
			NameKey:     "sys",
			MessageKey:  "message",
			LevelKey:    "lvl",
			TimeKey:     "time",
			EncodeLevel: zapcore.LowercaseColorLevelEncoder,
			EncodeTime:  zapcore.EpochTimeEncoder,
		},
		InitialFields:    map[string]interface{}{"sys": sys},
		OutputPaths:      viper.GetStringSlice("log.out"),
		ErrorOutputPaths: viper.GetStringSlice("log.erros"),
	}

	logger, err := cfg.Build()
	if err != nil {
		panic(err)
	}

	logger = logger.Named(sys)

	loggerFlush[sys] = logger.Sync

	sugar := logger.Sugar()
	return sugar
}

func SyncLoggers() {
	for _, f := range loggerFlush {
		f()
	}
}

func getLevelFromString(s string) zap.AtomicLevel {
	lvl, err := zap.ParseAtomicLevel(s)
	if err != nil {
		panic(err)
	}
	return lvl
}
