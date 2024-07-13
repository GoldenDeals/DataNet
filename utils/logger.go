package utils

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type LogFactory struct {
	loggerFlush      map[string]func() error
	loggerInitalized bool
}

func InitLoggers() *LogFactory {
	l := new(LogFactory)
	l.loggerFlush = make(map[string]func() error)
	l.loggerInitalized = true
	return l
}

func (l *LogFactory) NewLogger(sys string) *zap.SugaredLogger {
	if !l.loggerInitalized {
		panic("logger is not initialized")
	}
	_, ok := l.loggerFlush[sys]
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
		OutputPaths:      viper.GetStringSlice("log.out"),
		ErrorOutputPaths: viper.GetStringSlice("log.erros"),
	}

	logger, err := cfg.Build()
	if err != nil {
		panic(err)
	}

	logger = logger.Named(sys)

	l.loggerFlush[sys] = logger.Sync

	sugar := logger.Sugar()
	return sugar
}

func (l *LogFactory) SyncLoggers() {
	for _, f := range l.loggerFlush {
		_ = f()
	}
}

func getLevelFromString(s string) zap.AtomicLevel {
	lvl, err := zap.ParseAtomicLevel(s)
	if err != nil {
		panic(err)
	}
	return lvl
}
