package utils

import (
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

type LoggerGroup func(...*zerolog.Event) *zerolog.Event

func NewLoggerGroup(name string) LoggerGroup {
	return func(e ...*zerolog.Event) *zerolog.Event {
		var er *zerolog.Event
		if len(e) < 1 {
			er = log.Debug()
		} else {
			er = e[0]
		}

		return er.Str("group", name).Str("call", fileWithLineNum())
	}
}

func InitLogger() {
	zerolog.SetGlobalLevel(getLogLevel(viper.GetString("loglevel")))

	if !viper.GetBool("log_json") {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})
	}
}

func getLogLevel(name string) zerolog.Level {
	return map[string]zerolog.Level{
		"trace":   zerolog.TraceLevel,
		"debug":   zerolog.DebugLevel,
		"info":    zerolog.DebugLevel,
		"warning": zerolog.WarnLevel,
		"error":   zerolog.ErrorLevel,
		"disable": zerolog.Disabled,
	}[strings.ToLower(name)]
}

func fileWithLineNum() string {
	pc, filename, line, _ := runtime.Caller(2)

	realFN := strings.Split(filename, "/")
	realPC := strings.Split(runtime.FuncForPC(pc).Name(), "/")

	return fmt.Sprintf("%s/%s:%d", realFN[len(realFN)-1], realPC[len(realPC)-1], line)
}
